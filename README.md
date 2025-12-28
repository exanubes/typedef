 [![CI](https://github.com/exanubes/typedef/actions/workflows/ci.yml/badge.svg)](https://github.com/exanubes/typedef/actions/workflows/ci.yml)
 [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/exanubes/typedef/blob/master/LICENSE)
 [![Latest Release](https://img.shields.io/github/v/release/exanubes/typedef)](https://github.com/exanubes/typedef/releases)

# typedef

Convert JSON to type definitions in multiple programming languages.

`typedef` is a code generator that analyzes JSON input and produces strongly-typed definitions for your target language. Whether you're working with API responses, configuration files, or any JSON data, typedef helps you quickly generate type-safe code.

## Features

- **Multiple Output Formats**: Generate Go structs, TypeScript interfaces, Zod schemas, or JSDoc typedefs
- **Flexible Input**: Provide JSON via command flag, stdin pipe, or clipboard
- **Smart Output**: Send generated types to clipboard, stdout, or file
- **Intelligent Type Deduplication**: Automatically identifies and reuses structurally identical types
- **Clipboard Integration**: Seamless clipboard support with automatic fallback on macOS and Linux
- **Web Interface**: Also available as a browser-based tool at [https://exanubes.github.io/typedef](https://exanubes.github.io/typedef)

## Quick Start

```bash
# Build the tool
go build -o typedef main.go

# Generate Go types from JSON
./typedef --format go --input '{"name": "John", "age": 30, "active": true}'

# Output (copied to clipboard by default):
# type Root struct {
#     Active bool   
#     Age    int    
#     Name   string 
# }
```

## Installation

### Option 1: Download Pre-built Binary

Download the latest release for your platform from [GitHub Releases](https://github.com/exanubes/typedef/releases):

- **macOS (Intel)**: `typedef-darwin-amd64`
- **macOS (Apple Silicon)**: `typedef-darwin-arm64`
- **Linux (Intel/AMD)**: `typedef-linux-amd64`
- **Linux (ARM)**: `typedef-linux-arm64`

Make the binary executable and verify it works:

```bash
chmod +x typedef
./typedef --format go --input '{"test": 1}'
```

### Option 2: Build from Source

**Prerequisites:** Go 1.24.4 or higher

```bash
# Clone the repository
git clone https://github.com/exanubes/typedef.git
cd typedef

# Build the CLI tool
go build -o typedef main.go

# Or use the Makefile
make build-cli
```

**Verify the build:**

```bash
./typedef --format go --input '{"test": 1}'
```

## Usage

### Input Methods

typedef accepts JSON input in three ways (checked in this order):

1. **`--input` flag**: Explicit JSON string
   ```bash
   ./typedef --format go --input '{"name": "Alice", "age": 25}'
   ```

2. **stdin** (piped input):
   ```bash
   echo '{"name": "Bob", "score": 95}' | ./typedef --format typescript
   ```

3. **Clipboard** (automatic): If no input is provided, typedef reads from your clipboard
   ```bash
   # Copy JSON to clipboard, then run:
   ./typedef --format zod
   ```

### Output Targets

Control where generated types are sent using the `--target` flag:

- **`clipboard`** (default): Copies output to clipboard for easy pasting
  ```bash
  ./typedef --format go --input '{"data": true}'
  # Output copied to clipboard automatically
  ```

- **`cli`**: Prints to stdout
  ```bash
  ./typedef --format typescript --target cli --input '{"id": 1}'
  ```

- **`file`**: Saves to a file (use `--output-path` to specify location, or it will fallback to default location) 
  ```bash
  ./typedef --format go --input '{"user": "john"}' --target file --output-path ./types.go
  ```

### Supported Formats

| Format | Flag Values | Output |
|--------|-------------|--------|
| Go | `go`, `golang` | Go structs |
| TypeScript | `typescript`, `ts` | TypeScript interfaces |
| Zod | `zod`, `ts-zod` | Zod schema definitions |
| JSDoc | `jsdoc` | JSDoc typedef comments |

### Command Flags Reference

| Flag | Shorthand | Description | Default |
|------|-----------|-------------|---------|
| `--format` | `-f` | Output format (go/typescript/zod/jsdoc) | Required |
| `--input` | `-i` | JSON input string | Uses stdin or clipboard |
| `--target` | `-t` | Output target (clipboard/cli/file) | `clipboard` |
| `--output-path` | `-o` | File path when using `--target file` | `./typedef.txt` |

## Clipboard Support Setup

typedef can read JSON from your clipboard and copy generated types back to it automatically. Support varies by platform:

| Platform | Requirements               | Setup Needed |
|----------|----------------------------|--------------|
| macOS    | pbcopy, pbpaste (built-in) | None         |
| Linux    | wl-clipboard, xclip        | Yes          |

### macOS

No setup required! typedef uses the built-in `pbcopy` and `pbpaste` utilities that come with macOS.

### Linux

#### wl-clipboard 

Install the `wl-clipboard` package:

```bash
# Debian/Ubuntu
sudo apt-get install wl-clipboard

# Fedora
sudo dnf install wl-clipboard

# Arch Linux
sudo pacman -S wl-clipboard
```

Verify installation:
```bash
which wl-copy wl-paste
```

#### xclip 

Install the `xclip` package:

```bash
# Debian/Ubuntu
sudo apt-get install xclip

# Fedora
sudo dnf install xclip

# Arch Linux
sudo pacman -S xclip
```

Verify installation:
```bash
which xclip
```

### Automatic Fallback

If supported clipboard tools aren't detected, typedef will automatically fall back to stdout output without errors:

```
INFO: required dependencies for clipboard not detected. Falling back to cli target
```

The tool continues to work normally, printing output to your terminal instead. You can also explicitly use `--target cli` or `--target file` to bypass clipboard functionality.

## Examples

### Basic Usage

```bash
# Generate Go struct
typedef --format go --input '{"username": "alice", "email": "alice@example.com"}'

# Generate TypeScript interface
typedef --format typescript --input '{"id": 1, "title": "Hello", "published": true}'

# Generate Zod schema
typedef --format zod --input '{"age": 30, "name": "Bob"}'

# Generate JSDoc typedef
typedef --format jsdoc --input '{"count": 5, "items": ["a", "b"]}'
```

### Piped Input

```bash
# From echo
echo '{"status": "success", "data": {"id": 123}}' | typedef --format go

# From file
cat config.json | typedef --format typescript

# From API response with curl
curl -s https://api.example.com/user/1 | typedef --format go
```

### Clipboard Workflow

```bash
# 1. Copy JSON from anywhere (browser, editor, etc.)
# 2. Run typedef
typedef --format typescript

# 3. Generated TypeScript interface is now in your clipboard
# 4. Paste directly into your code editor
```

### File Output

```bash
# Save Go types to file
./typedef --format go --input '{"name": "test"}' --target file --output-path ./models/user.go

# Process API response and save
curl -s https://api.example.com/data | ./typedef --format typescript --target file --output-path ./types/api.ts
```

### Real-World Scenarios

**Converting API Response to TypeScript:**
```bash
curl -s https://jsonplaceholder.typicode.com/users/1 | ./typedef --format typescript
```

**Generating Go Structs from Config:**
```bash
cat settings.json | ./typedef --format go --target file --output-path ./config/types.go
```

**Creating Zod Schemas for Form Validation:**
```bash
./typedef --format zod --input '{"email": "user@example.com", "age": 25, "subscribe": true}'
```

**Using with jq for Complex JSON:**
```bash
# Extract specific data with jq, then generate types
curl -s https://api.github.com/repos/golang/go | jq '.owner' | ./typedef --format typescript
```

## Troubleshooting

### Clipboard Issues

**Problem**: "Clipboard support disabled" message on Linux

**Solution**: Install the required clipboard tool for your display server:
- Wayland: `sudo apt-get install wl-clipboard`
- X11: `sudo apt-get install xclip`

**Problem**: Clipboard not working after installation

**Solution**: Verify the tools are in your PATH:
```bash
which wl-copy wl-paste  # Wayland
which xclip              # X11
```

### Invalid JSON Errors

**Problem**: "Failed to parse JSON" or similar errors

**Solution**: Validate your JSON before passing to typedef:
```bash
# Using jq to validate
echo '{"test": 1}' | jq .

# Using online validators
# https://jsonlint.com/
```

### Platform-Specific Issues

**Linux: Permission denied**
```bash
# Make the binary executable
chmod +x typedef
```

**macOS: "cannot be opened because the developer cannot be verified"**
```bash
# Allow the binary in System Preferences > Security & Privacy
# Or build from source as shown in Installation section
```

## Web Interface

Prefer a browser-based tool? typedef is also available as a web application at:

**[https://exanubes.github.io/typedef](https://exanubes.github.io/typedef)**

The web version offers the same type generation capabilities with a graphical interface, perfect for quick conversions without installing the CLI tool.

## License

MIT License - see LICENSE file for details
