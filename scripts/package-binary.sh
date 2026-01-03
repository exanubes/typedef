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

# Extract binary name from path
BINARY_NAME=$(basename "${BINARY_PATH}")

# Derived values
ARCHIVE_NAME="${BINARY_NAME}-${OS}-${ARCH}.tar.gz"
TEMP_DIR="temp-${BINARY_NAME}-${OS}-${ARCH}"

echo "Packaging ${BINARY_NAME} for ${OS}-${ARCH}..."

# Create temporary directory
mkdir -p "${TEMP_DIR}"

# Copy binary (preserve original name)
cp "${BINARY_PATH}" "${TEMP_DIR}/${BINARY_NAME}"
chmod +x "${TEMP_DIR}/${BINARY_NAME}"

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
