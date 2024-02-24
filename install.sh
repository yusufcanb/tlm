#!/bin/bash

# OS and Architecture Detection
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  os="linux"
elif [[ "$OSTYPE" == "darwin"* ]]; then
  os="darwin"
else
  echo "Unsupported operating system. Only Linux and macOS are currently supported."
  exit 1
fi

if [[ "$(uname -m)" == "x86_64" ]]; then
  arch="amd64"
elif [[ "$(uname -m)" == "aarch64" || "$(uname -m)" == "arm64" ]]; then
  arch="arm64"
else
  echo "Unsupported architecture. tlm requires a 64-bit system (x86_64 or arm64)."
  exit 1
fi

# Download URL Construction
version=$(cat VERSION)
base_url="https://github.com/yusufcanb/tlm/releases/download"
download_url="${base_url}/${version}/tlm_${version}_${os}_${arch}"

# Docker check
if ! command -v docker &>/dev/null; then
  echo "Docker not found. Please install Docker from https://www.docker.com/get-started"
  exit 1
fi

# Ollama check
if ! curl -fsSL http://localhost:11434 &> /dev/null; then
  echo "Ollama not found."
  if [[ "$os" == "darwin" ]]; then
    echo ""
    echo "*** On macOS: ***"
    echo ""
    echo "Ollama can run with GPU acceleration inside Docker containers for Nvidia GPUs."
    echo "To get started using the Docker image, please follow these steps:"
    echo ""
    echo "1. *** CPU only: ***"
    echo "   docker run -d -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama"
    echo ""
    echo "2. *** GPU Acceleration: ***"
    echo "   This option requires running Ollama outside of Docker"
    echo "   To get started, simply download and install Ollama."
    echo "   https://ollama.com/download"
    echo ""
    echo ""
    echo "Installation aborted. Please install Ollama using the methods above and try again."
    exit 1

  elif [[ "$os" == "linux" ]]; then
    echo ""
    echo "*** On Linux: ***"
    echo ""
    echo "Ollama can run with GPU acceleration inside Docker containers for Nvidia GPUs."
    echo "To get started using the Docker image, please follow these steps:"
    echo ""
    echo "1. *** CPU only: ***"
    echo "   docker run -d -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama"
    echo ""
    echo "2. *** Nvidia GPU: ***"
    echo "   docker run -d --gpus=all -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama"
    echo ""
    echo ""
    echo "Installation aborted. Please install Ollama using the methods above and try again."
    exit 1

  fi
fi

# Download the binary
echo "Downloading tlm version ${version} for ${os}/${arch}..."
if ! curl -fsSL -o tlm ${download_url}; then
  echo "Download failed. Please check your internet connection and try again."
  exit 1
fi

# Make executable
chmod +x tlm

# Move to installation directory
echo "Installing tlm..."
if ! mv tlm /usr/local/bin/; then
  echo "Installation requires administrator permissions. Please use sudo or run the script as root."
  exit 1
else
  echo ""
fi

if ! tlm deploy; then
  echo ""
  exit 1
else
  echo ""
fi

echo "Type 'tlm' to get started."
exit 0