# tlm - Local CLI Copilot, powered by Ollama. 💻🦙

![Latest Build](https://img.shields.io/github/actions/workflow/status/yusufcanb/tlm/build.yaml?style=for-the-badge&logo=github)
[![Sonar Quality Gate](https://img.shields.io/sonar/quality_gate/yusufcanb_tlm?server=https%3A%2F%2Fsonarcloud.io&style=for-the-badge&logo=sonar)](https://sonarcloud.io/project/overview?id=yusufcanb_tlm)
[![Latest Release](https://img.shields.io/github/v/release/yusufcanb/tlm?display_name=release&style=for-the-badge&logo=github&link=https%3A%2F%2Fgithub.com%2Fyusufcanb%2Ftlm%2Freleases)](https://github.com/yusufcanb/tlm/releases)


tlm is your CLI companion which requires nothing except your workstation. It uses most efficient and powerful powerful open-source models like [Llama 3.3](https://ollama.com/library/llama3.3), [Phi4](https://ollama.com/library/phi4), [DeepSeek-R1](https://ollama.com/library/deepseek-r1), [Qwen](https://ollama.com/library/qwen2.5-coder) in your local environment to provide you the best possible command line assistance.

![Suggest](./assets/suggest.gif)

![Explain](./assets/explain2.gif)

![Model Selection](./assets/config.gif)


## Features

- 💸 No API Key (Subscription) is required. (ChatGPT, Claude, Github Copilot, Azure OpenAI, etc.)

- 📡 No internet connection is required.

- 💻 Works on macOS, Linux and Windows.

- 👩🏻‍💻 Automatic shell detection. (Powershell, Bash, Zsh)

- 🚀 One liner generation and command explanation.

- 🧠 Experiment any model. ([Llama 3.3](https://ollama.com/library/llama3.3), [Phi4](https://ollama.com/library/phi4), [DeepSeek-R1](https://ollama.com/library/deepseek-r1), [Qwen](https://ollama.com/library/qwen2.5-coder)) with parameters of your choice.

## Installation

Installation can be done in two ways;

- [Installation script](#installation-script) (recommended)
- [Go Install](#go-install)

###  Installation Script

Installation script is the recommended way to install tlm.
It will recognize the which platform and architecture to download and will execute install command for you.

#### Linux and macOS;

Download and execute the installation script by using the following command;

```bash
curl -fsSL https://raw.githubusercontent.com/yusufcanb/tlm/1.2-pre/install.sh | sudo -E bash
```

#### Windows (Powershell 5.5 or higher)

Download and execute the installation script by using the following command;

```powershell
Invoke-RestMethod -Uri https://raw.githubusercontent.com/yusufcanb/tlm/1.2-pre/install.ps1 | Invoke-Expression
```

### Go Install

If you have Go 1.22 or higher installed on your system, you can easily use the following command to install tlm;

```bash
go install github.com/yusufcanb/tlm@1.2-pre
```

You're ready! Check installation by using the following command;

```bash
tlm
```

## Uninstall

On Linux and macOS;

```bash
rm /usr/local/bin/tlm
rm ~/.tlm.yml
```

On Windows;

```powershell
Remove-Item -Recurse -Force "C:\Users\$env:USERNAME\AppData\Local\Programs\tlm"
Remove-Item -Force "$HOME\.tlm.yml"
```

## Contributors

 <a href = "https://github.com/yusufcanb/tlm/graphs/contributors">
   <img src = "https://contrib.rocks/image?repo=yusufcanb/tlm"/>
 </a>
