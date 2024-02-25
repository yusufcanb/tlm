#!/bin/bash

# Function to perform builds
build() {
  os=$1
  app_name=$2
  version=$3
  arch=$4

  # Determine output filename with optional .exe extension
  output_name="${app_name}_${version}_${os}_${arch}"
  if [[ "$os" == "windows" ]]; then
    output_name="${output_name}.exe"
  fi

  echo "Building for $os/$arch (version: $version) -> $output_name"
  CGO_ENABLED=0 GOOS=$os GOARCH=$arch go build -o "dist/${version}/${output_name}" main.go
}

# Operating systems to target
targets=("darwin" "linux" "windows")
app_name="tlm"

if [ $# -eq 0 ]; then
  echo "Error: Please provide a version number as a command-line argument."
  exit 1
fi

version=$1

# Clear old build artifacts
rm -rf dist

# Create the output directory
mkdir dist

# Build for each OS
for os in "${targets[@]}"; do
  build $os $app_name $version "arm64"
  build $os $app_name $version "amd64"
done

echo "Done!"