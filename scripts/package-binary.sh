#!/bin/bash
set -e

# Usage: ./scripts/package-binary.sh <binary_path> <os> <arch> <output_dir>
# Example: ./scripts/package-binary.sh ./artifacts/linux/amd64/cli linux amd64 ./release

if [ "$#" -ne 4 ]; then
  echo "Usage: $0 <binary_path> <os> <arch> <output_dir>"
  exit 1
fi

BINARY_PATH="$1"
OS="$2"
ARCH="$3"
OUTPUT_DIR="$4"

# Derived values
ARCHIVE_NAME="typedef-${OS}-${ARCH}.tar.gz"
TEMP_DIR="temp-${OS}-${ARCH}"

echo "Packaging ${OS}-${ARCH}..."

# Create temporary directory
mkdir -p "${TEMP_DIR}"

# Copy and rename binary
cp "${BINARY_PATH}" "${TEMP_DIR}/typedef"
chmod +x "${TEMP_DIR}/typedef"

# Copy documentation (if exists)
if [ -f "README.md" ]; then
  cp README.md "${TEMP_DIR}/"
fi

if [ -f "LICENSE" ]; then
  cp LICENSE "${TEMP_DIR}/"
fi

# Create output directory if needed
mkdir -p "${OUTPUT_DIR}"

# Create tar.gz archive
tar -czf "${OUTPUT_DIR}/${ARCHIVE_NAME}" -C "${TEMP_DIR}" .

# Cleanup
rm -rf "${TEMP_DIR}"

echo "Created ${OUTPUT_DIR}/${ARCHIVE_NAME}"
