# tlm - Local terminal companion, powered by CodeLLaMa.

tlm is your CLI companion which requires nothing except your workstation. It uses most efficient and powerful [CodeLLaMa](https://ai.meta.com/blog/code-llama-large-language-model-coding/) in your local environment to provide you the best possible command line suggestions.

![Suggest](./assets/suggest.gif)

![Explain](./assets/explain.gif)

![Config](./assets/config.gif)

## Features

- 💸 No API Key (Subscription) is required. (ChatGPT, Github Copilot, Azure OpenAI, etc.)

- 📡 No internet connection is required.

- 💻 Works on macOS, Linux and Windows.

- 👩🏻‍💻 Automatic shell detection.

- 🚀 One liner generation and command explanation.


## Installation

Installation can be done in two ways;

- [Installation script](#installation-script) (recommended)
- [Go Install](#go-install)

### Prerequisites

[Ollama](https://ollama.com/) is needed to download to necessary models.
It can be downloaded with the following methods on different platforms.

- On Linux and macOS;

```bash
curl -fsSL https://ollama.com/install.sh | sh
```

- On Windows;

Download instructions can be followed at the following link: [https://ollama.com/download](https://ollama.com/download)

- Or using official Docker images 🐳;

```bash
# CPU Only
docker run -d -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama

# With GPU (Nvidia only)
docker run -d --gpus=all -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama"
```

###  Installation Script

Installation script is the recommended way to install tlm.
It will recognize the which platform and architecture to download and will execute install command for you.

#### Linux and macOS;


Download and execute the installation script by using the following command;

```bash
curl -fsSL https://raw.githubusercontent.com/yusufcanb/tlm/main/install.sh | sudo bash -E
```

#### Windows (Powershell 5.1 or higher)

Download and execute the installation script by using the following command;

```powershell
Invoke-RestMethod -Uri https://raw.githubusercontent.com/yusufcanb/tlm/main/install.ps1 | Invoke-Expression
```

### Go Install

If you have Go 1.21 or higher installed on your system, you can easily use the following command to install tlm;

```bash
go install github.com/yusufcanb/tlm@latest
```

Check installation by using the following command;

```bash
tlm help
```

## Uninstall

On Linux and macOS;

```bash
rm /usr/local/bin/tlm
```
