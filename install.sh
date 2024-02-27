#!/bin/bash

set -eu

status() { echo ">>> $*" >&2; }
error() { echo "ERROR $*"; }
warning() { echo "WARNING: $*"; }


# OS and Architecture Detection
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  os="linux"
elif [[ "$OSTYPE" == "darwin"* ]]; then
  os="darwin"
else
  error "Unsupported operating system. Only Linux and macOS are currently supported."
  exit 1
fi

if [[ "$(uname -m)" == "x86_64" ]]; then
  arch="amd64"
elif [[ "$(uname -m)" == "aarch64" || "$(uname -m)" == "arm64" ]]; then
  arch="arm64"
else
  error "Unsupported architecture. tlm requires a 64-bit system (x86_64 or arm64)."
  exit 1
fi

# Download URL Construction
version="1.0-rc2"
base_url="https://github.com/yusufcanb/tlm/releases/download"
download_url="${base_url}/${version}/tlm_${version}_${os}_${arch}"

# Docker check
if ! command -v docker &>/dev/null; then
  error "Docker not found. Please install Docker from https://www.docker.com/get-started"
  exit 1
fi

# Ollama check
if ! curl -fsSL http://localhost:11434 &> /dev/null; then
  error "Ollama not found."
  if [[ "$os" == "darwin" ]]; then
    status ""
    status "*** On macOS: ***"
    status ""
    status "Ollama can run with GPU acceleration inside Docker containers for Nvidia GPUs."
    status "To get started using the Docker image, please follow these steps:"
    status ""
    status "1. *** CPU only: ***"
    status "   docker run -d -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama"
    status ""
    status "2. *** GPU Acceleration: ***"
    status "   This option requires running Ollama outside of Docker"
    status "   To get started, simply download and install Ollama."
    status "   https://ollama.com/download"
    status ""
    status ""
    status "Installation aborted. Please install Ollama using the methods above and try again."
    exit 1

  elif [[ "$os" == "linux" ]]; then
    status ""
    status "*** On Linux: ***"
    status ""
    status "Ollama can run with GPU acceleration inside Docker containers for Nvidia GPUs."
    status "To get started using the Docker image, please follow these steps:"
    status ""
    status "1. *** CPU only: ***"
    status "   docker run -d -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama"
    status ""
    status "2. *** Nvidia GPU: ***"
    status "   docker run -d --gpus=all -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama"
    status ""
    status ""
    status "Installation aborted. Please install Ollama using the methods above and try again."
    exit 1

  fi
fi

# Download the binary
status "Downloading tlm version ${version} for ${os}/${arch}..."
if ! curl -fsSL -o tlm ${download_url}; then
  error "Download failed. Please check your internet connection and try again."
  exit 1
fi

# Make executable
chmod +x tlm

# Move to installation directory
status "Installing tlm..."

SUDO=
if [ "$(id -u)" -ne 0 ]; then
    # Running as root, no need for sudo
    if ! available sudo; then
        error "This script requires superuser permissions. Please re-run as root."
    fi

    SUDO="sudo"
fi

$SUDO mv tlm /usr/local/bin/;

if ! tlm deploy; then
  error ""
  exit 1
else
  status ""
fi

status "Type 'tlm' to get started."
exit 0