#!/bin/bash
set -e

VERSION="${1:-dev}"
INSTALL_TEST=false

if [ "$2" = "--install" ]; then
  INSTALL_TEST=true
fi

# Check if local builds exist
echo "Checking for local builds..."
for path in darwin/amd64 darwin/arm64 linux/amd64 linux/arm64; do
  if [ ! -f "dist/cli/$path/typedef-cli" ]; then
    echo "Error: Missing build at dist/cli/$path/typedef-cli"
    echo "Run 'make build-cli' to build all platforms first"
    exit 1
  fi
done

# Create temp directory for testing
TEST_DIR=$(mktemp -d)
echo "Using test directory: $TEST_DIR"

# Package binaries into tar.gz (similar to release process)
echo "Packaging binaries..."
for os_arch in darwin-amd64 darwin-arm64 linux-amd64 linux-arm64; do
  os=$(echo $os_arch | cut -d- -f1)
  arch=$(echo $os_arch | cut -d- -f2)

  tar -czf "$TEST_DIR/typedef-cli-$os_arch.tar.gz" \
    -C "dist/cli/$os/$arch" typedef-cli
done

# Generate checksums
echo "Generating checksums..."
cd "$TEST_DIR"
shasum -a 256 typedef-cli-*.tar.gz > checksums.txt
cd - > /dev/null

echo "Checksums generated:"
cat "$TEST_DIR/checksums.txt"

# Generate formula
echo "Generating Homebrew formula..."
./client/homebrew/scripts/generate-formula.sh \
  "$VERSION" \
  "$TEST_DIR/checksums.txt" \
  "$TEST_DIR/typedef.rb"

echo "Formula generated at: $TEST_DIR/typedef.rb"

# Validate Ruby syntax
echo "Validating formula syntax..."
ruby -c "$TEST_DIR/typedef.rb"

# Show formula
echo "Generated formula:"
cat "$TEST_DIR/typedef.rb"

# Test installation if requested
if [ "$INSTALL_TEST" = true ]; then
  echo ""
  echo "Testing installation..."
  echo "Note: This will install typedef-test locally (test binary)"
  echo "Proceeding with installation test..."

  # Create temporary tap for testing
  TAP_NAME="exanubes/typedef-test"
  echo "Creating temporary tap: $TAP_NAME"

  # Remove tap if it exists (from previous failed run)
  brew untap "$TAP_NAME" 2>/dev/null || true

  # Create new tap without git repo (--no-git flag)
  brew tap-new "$TAP_NAME" --no-git
  TAP_PATH="$(brew --repository)/Library/Taps/exanubes/homebrew-typedef-test"

  # Copy formula with local file URLs and rename binary to typedef-test
  echo "Installing formula into tap..."
  sed -e "s|https://github.com/exanubes/typedef/releases/download/v$VERSION/|file://$TEST_DIR/|g" \
      -e 's|bin.install "typedef-cli" => "typedef"|bin.install "typedef-cli" => "typedef-test"|g' \
    "$TEST_DIR/typedef.rb" > "$TAP_PATH/Formula/typedef.rb"

  # Install from tap
  echo "Installing typedef-test from tap..."
  brew install "$TAP_NAME/typedef"

  # Test installed binary
  echo "Testing installed binary..."
  typedef-test version
  typedef-test --format go --input '{"test": 1}' --target cli

  # Cleanup
  echo "Cleaning up tap..."
  brew uninstall typedef
  brew untap "$TAP_NAME"

  echo "Installation test completed successfully"
fi

echo ""
echo "Testing complete! Formula location: $TEST_DIR/typedef.rb"
echo "To clean up test files, run:"
echo "trash $TEST_DIR"
