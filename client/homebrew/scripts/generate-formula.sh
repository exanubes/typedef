#!/bin/bash
set -e

# Usage: ./generate-formula.sh <version> <checksums_file> <output_file>
# Example: ./generate-formula.sh 0.0.15 checksums.txt typedef.rb

if [ "$#" -ne 3 ]; then
  echo "Usage: $0 <version> <checksums_file> <output_file>"
  echo "Example: $0 0.0.15 checksums.txt typedef.rb"
  exit 1
fi

VERSION="$1"
CHECKSUMS_FILE="$2"
OUTPUT_FILE="$3"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TEMPLATE_FILE="${SCRIPT_DIR}/../Formula/typedef.rb.template"

# Validate template file exists
if [ ! -f "$TEMPLATE_FILE" ]; then
  echo "Error: Template file not found at $TEMPLATE_FILE"
  exit 1
fi

# Validate checksums file exists
if [ ! -f "$CHECKSUMS_FILE" ]; then
  echo "Error: Checksums file not found at $CHECKSUMS_FILE"
  exit 1
fi

# Extract SHA256 checksums for all platforms
SHA256_DARWIN_ARM64=$(grep "typedef-cli-darwin-arm64.tar.gz" "$CHECKSUMS_FILE" | awk '{print $1}')
SHA256_DARWIN_AMD64=$(grep "typedef-cli-darwin-amd64.tar.gz" "$CHECKSUMS_FILE" | awk '{print $1}')
SHA256_LINUX_ARM64=$(grep "typedef-cli-linux-arm64.tar.gz" "$CHECKSUMS_FILE" | awk '{print $1}')
SHA256_LINUX_AMD64=$(grep "typedef-cli-linux-amd64.tar.gz" "$CHECKSUMS_FILE" | awk '{print $1}')

# Validate all checksums were extracted
if [ -z "$SHA256_DARWIN_ARM64" ] || [ -z "$SHA256_DARWIN_AMD64" ] || \
   [ -z "$SHA256_LINUX_ARM64" ] || [ -z "$SHA256_LINUX_AMD64" ]; then
  echo "Error: Could not extract all required checksums from $CHECKSUMS_FILE"
  echo "Darwin ARM64: ${SHA256_DARWIN_ARM64:-<missing>}"
  echo "Darwin AMD64: ${SHA256_DARWIN_AMD64:-<missing>}"
  echo "Linux ARM64:  ${SHA256_LINUX_ARM64:-<missing>}"
  echo "Linux AMD64:  ${SHA256_LINUX_AMD64:-<missing>}"
  exit 1
fi

# Generate formula from template by replacing placeholders
sed -e "s/{{VERSION}}/$VERSION/g" \
    -e "s/{{SHA256_DARWIN_ARM64}}/$SHA256_DARWIN_ARM64/g" \
    -e "s/{{SHA256_DARWIN_AMD64}}/$SHA256_DARWIN_AMD64/g" \
    -e "s/{{SHA256_LINUX_ARM64}}/$SHA256_LINUX_ARM64/g" \
    -e "s/{{SHA256_LINUX_AMD64}}/$SHA256_LINUX_AMD64/g" \
    "$TEMPLATE_FILE" > "$OUTPUT_FILE"

echo "Generated formula at $OUTPUT_FILE"
echo "  Version: $VERSION"
echo "  Darwin ARM64 SHA256: $SHA256_DARWIN_ARM64"
echo "  Darwin AMD64 SHA256: $SHA256_DARWIN_AMD64"
echo "  Linux ARM64 SHA256:  $SHA256_LINUX_ARM64"
echo "  Linux AMD64 SHA256:  $SHA256_LINUX_AMD64"
