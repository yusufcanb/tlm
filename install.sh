#!/bin/bash

set -e

status() { echo ">>> $*" >&2; }
error() { echo "ERROR $*"; }
warning() { echo "WARNING: $*"; }

print_message() {
    local message="$1"
    local color="$2"
    echo -e "\e[${color}m${message}\e[0m"
}

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
version="1.2"
base_url="https://github.com/yusufcanb/tlm/releases/download"
download_url="${base_url}/${version}/tlm_${version}_${os}_${arch}"

if [ -n "${OLLAMA_HOST+x}" ]; then
    ollama_host=$OLLAMA_HOST
else
    ollama_host="http://localhost:11434"
fi

# Ollama check
if ! curl -fsSL $ollama_host &>/dev/null; then
    if [[ "$os" == "darwin" ]]; then
        print_message "ERR: Ollama not found." "31" # Red color
        print_message "If you have Ollama installed, please make sure it's running and accessible at ${ollama_host}" "31"
        print_message "or configure OLLAMA_HOST environment variable." "31"
        echo """
>>> If have Ollama on your system or network, you can set the OLLAMA_HOST like below;

    $ export OLLAMA_HOST=http://localhost:11434

>>> If you don't have Ollama installed, you can install it using the following methods;

    $(print_message "*** macOS: ***" "32")

    Download instructions can be followed at the following link: https://ollama.com/download

    $(print_message "*** Official Docker Images: ***" "32")

    Ollama can run with GPU acceleration inside Docker containers for Nvidia GPUs.
    To get started using the Docker image, please follow these steps:

        $(print_message "1. *** CPU only: ***" "32")

        $ docker run -d -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama


        $(print_message "2. *** Nvidia GPU: ***" "32")

        $ docker run -d --gpus=all -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama
        """
        print_message "Installation aborted..." "31"
        print_message "Please install or configure Ollama using the methods above and try again." "31"
        exit 1

    elif [[ "$os" == "linux" ]]; then
        print_message "ERR: Ollama not found." "31" # Red color
        print_message "If you have Ollama installed, please make sure it's running and accessible at ${ollama_host}" "31"
        print_message "or configure OLLAMA_HOST environment variable." "31"
        echo """
>>> If have Ollama on your system or network, you can set the OLLAMA_HOST like below;

    $ export OLLAMA_HOST=http://localhost:11434

>>> If you don't have Ollama installed, you can install it using the following methods;

    $(print_message "*** Linux: ***" "32")

    Download instructions can be followed at the following link: https://ollama.com/download

    $(print_message "*** Official Docker Images: ***" "32")

    Ollama can run with GPU acceleration inside Docker containers for Nvidia GPUs.
    To get started using the Docker image, please follow these steps:

        $(print_message "1. *** CPU only: ***" "32")

        $ docker run -d -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama


        $(print_message "2. *** Nvidia GPU: ***" "32")

        $ docker run -d --gpus=all -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama
        """
        print_message "Installation aborted..." "31"
        print_message "Please install or configure Ollama using the methods above and try again." "31"
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
        exit 1
    fi

    SUDO="sudo"
fi

$SUDO mv tlm /usr/local/bin/

# set shell auto
if ! $SUDO tlm config set shell auto &>/dev/null; then
    error "tlm config set shell <auto> failed."
    exit 1
fi

# change ownership of tlm config file to user
$SUDO chown $SUDO_USER ~/.tlm.yaml

status "Type 'tlm' to get started."
exit 0
