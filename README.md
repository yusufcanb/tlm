# tlm - Local CLI Copilot, powered by CodeLLaMa. 💻🦙

![Latest Build](https://img.shields.io/github/actions/workflow/status/yusufcanb/tlm/build.yaml?style=for-the-badge&logo=github)
[![Sonar Quality Gate](https://img.shields.io/sonar/quality_gate/yusufcanb_tlm?server=https%3A%2F%2Fsonarcloud.io&style=for-the-badge&logo=sonar)](https://sonarcloud.io/project/overview?id=yusufcanb_tlm)
[![Latest Release](https://img.shields.io/github/v/release/yusufcanb/tlm?display_name=release&style=for-the-badge&logo=github&link=https%3A%2F%2Fgithub.com%2Fyusufcanb%2Ftlm%2Freleases)](https://github.com/yusufcanb/tlm/releases)
![Downloads](https://img.shields.io/github/downloads/yusufcanb/tlm/total.svg?style=for-the-badge&logo=github&color=orange)



tlm is your CLI companion which requires nothing except your workstation. It uses most efficient and powerful [CodeLLaMa](https://ai.meta.com/blog/code-llama-large-language-model-coding/) in your local environment to provide you the best possible command line suggestions.

![Suggest](./assets/suggest.gif)

![Explain](./assets/explain.gif)


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

- On macOs and Windows;

Download instructions can be followed at the following link: [https://ollama.com/download](https://ollama.com/download)

- On Linux;

```bash
curl -fsSL https://ollama.com/install.sh | sh
```

- Or using official Docker images 🐳;

```bash
# CPU Only
docker run -d -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama

# With GPU (Nvidia only)
docker run -d --gpus=all -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama
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

Then, deploy tlm modelfiles.

> :memo: **Note:** If you have Ollama deployed on somewhere else. Please first run `tlm config` and configure Ollama host.

```bash
tlm deploy
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

On Windows;

Remove the directory under;

```
C:\Users\<username>\AppData\Local\Programs\tlm
```
