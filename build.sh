#!/bin/bash

PROJECT_ROOT=$(pwd)
GO_PROGRAM="$PROJECT_ROOT/pleh/go"
OUTPUT_DIR="$PROJECT_ROOT/bin"

echo "Making build directories..."
mkdir -p "$PROJECT_ROOT/bin"
mkdir -p "$OUTPUT_DIR/windows" "$OUTPUT_DIR/darwin_amd64" "$OUTPUT_DIR/darwin_arm64" "$OUTPUT_DIR/linux"

echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o "$OUTPUT_DIR/windows/pleh.exe" "$GO_PROGRAM"

echo "Building for Intel Mac..."
GOOS=darwin GOARCH=amd64 go build -o "$OUTPUT_DIR/darwin_amd64/pleh" "$GO_PROGRAM"

echo "Building for Apple Silicon..."
GOOS=darwin GOARCH=arm64 go build -o "$OUTPUT_DIR/darwin_arm64/pleh" "$GO_PROGRAM"

echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o "$OUTPUT_DIR/linux/pleh" "$GO_PROGRAM"

echo "Build completed. Binaries are located in the '$OUTPUT_DIR' directory."

