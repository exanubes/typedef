#!/bin/bash
set -e

VERSION="${1:-dev}"
PACKAGES_DIR="${2:-.}"
INSTALL_TEST=false

if [ "$3" = "--install" ]; then
  INSTALL_TEST=true
fi

# Validate inputs
if [ -z "$VERSION" ]; then
  echo "Error: VERSION argument is required"
  echo "Usage: $0 <version> <packages_dir> [--install]"
  exit 1
fi

# Check if brew is available when installation test is requested
if [ "$INSTALL_TEST" = true ] && ! command -v brew &> /dev/null; then
  echo "Error: Homebrew (brew) is not installed or not in PATH"
  exit 1
fi

# Validate packages directory exists
if [ ! -d "$PACKAGES_DIR" ]; then
  echo "Error: Packages directory not found: $PACKAGES_DIR"
  exit 1
fi

# Validate checksums file exists
CHECKSUMS_FILE="$PACKAGES_DIR/checksums.txt"
if [ ! -f "$CHECKSUMS_FILE" ]; then
  echo "Error: Checksums file not found at $CHECKSUMS_FILE"
  exit 1
fi

# Validate all required packages exist
echo "Validating packaged binaries..."
for os_arch in darwin-amd64 darwin-arm64 linux-amd64 linux-arm64; do
  PACKAGE="$PACKAGES_DIR/typedef-cli-$os_arch.tar.gz"
  if [ ! -f "$PACKAGE" ]; then
    echo "Error: Missing package: $PACKAGE"
    exit 1
  fi
done

echo "All packages validated successfully"
echo "Checksums file:"
cat "$CHECKSUMS_FILE"

# Create temp directory for formula
TEST_DIR=$(mktemp -d)
echo "Using test directory: $TEST_DIR"

# Generate formula using existing script
echo "Generating Homebrew formula..."
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
"$SCRIPT_DIR/generate-formula.sh" \
  "$VERSION" \
  "$CHECKSUMS_FILE" \
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
  echo "Checking for existing tap..."
  brew untap "$TAP_NAME" 2>&1 | grep -v "Error: No available tap" || true

  # Create new tap without git repo
  echo "Creating tap directory structure..."
  brew tap-new "$TAP_NAME" --no-git
  TAP_PATH="$(brew --repository)/Library/Taps/exanubes/homebrew-typedef-test"

  # Get absolute path to packages directory for file:// URLs
  PACKAGES_ABS_PATH="$(cd "$PACKAGES_DIR" && pwd)"

  # Copy formula with local file URLs and rename binary to typedef-test
  echo "Installing formula into tap..."
  sed -e "s|https://github.com/exanubes/typedef/releases/download/v$VERSION/|file://$PACKAGES_ABS_PATH/|g" \
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
