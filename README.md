# tlm - Local CLI Copilot, powered by Ollama. 💻🦙

![Latest Build](https://img.shields.io/github/actions/workflow/status/yusufcanb/tlm/build.yaml?style=for-the-badge&logo=github)
[![Sonar Quality Gate](https://img.shields.io/sonar/quality_gate/yusufcanb_tlm?server=https%3A%2F%2Fsonarcloud.io&style=for-the-badge&logo=sonar)](https://sonarcloud.io/project/overview?id=yusufcanb_tlm)
[![Latest Release](https://img.shields.io/github/v/release/yusufcanb/tlm?display_name=release&style=for-the-badge&logo=github&link=https%3A%2F%2Fgithub.com%2Fyusufcanb%2Ftlm%2Freleases)](https://github.com/yusufcanb/tlm/releases)

tlm is your CLI companion which requires nothing except your workstation. It uses most efficient and powerful open-source models like [Llama 3.3](https://ollama.com/library/llama3.3), [Phi4](https://ollama.com/library/phi4), [DeepSeek-R1](https://ollama.com/library/deepseek-r1), [Qwen](https://ollama.com/library/qwen2.5-coder) of your choice in your local environment to provide you the best possible command line assistance.

| Get a suggestion                 | Explain a command                |
| -------------------------------- | -------------------------------- |
| ![Suggest](./assets/suggest.gif) | ![Explain](./assets/explain.gif) |

| Ask with context (One-liner RAG) | Configure your favorite model  |
| -------------------------------- | ------------------------------ |
| ![Ask](./assets/ask.gif)         | ![Config](./assets/config.gif) |

## Features

- 💸 No API Key (Subscription) is required. (ChatGPT, Claude, Github Copilot, Azure OpenAI, etc.)

- 📡 No internet connection is required.

- 💻 Works on macOS, Linux and Windows.

- 👩🏻‍💻 Automatic shell detection. (Powershell, Bash, Zsh)

- 🚀 One liner generation and command explanation.

- 🖺 No-brainer RAG (Retrieval Augmented Generation)

- 🧠 Experiment any model. ([Llama3](https://ollama.com/library/llama3.3), [Phi4](https://ollama.com/library/phi4), [DeepSeek-R1](https://ollama.com/library/deepseek-r1), [Qwen](https://ollama.com/library/qwen2.5-coder)) with parameters of your choice.

## Installation

Installation can be done in two ways;

- [Installation script](#installation-script) (recommended)
- [Go Install](#go-install)

### Installation Script

Installation script is the recommended way to install tlm.
It will recognize the which platform and architecture to download and will execute install command for you.

#### Linux and macOS;

Download and execute the installation script by using the following command;

```bash
curl -fsSL https://raw.githubusercontent.com/yusufcanb/tlm/1.2/install.sh | sudo -E bash
```

#### Windows (Powershell 5.5 or higher)

Download and execute the installation script by using the following command;

```powershell
Invoke-RestMethod -Uri https://raw.githubusercontent.com/yusufcanb/tlm/1.2/install.ps1 | Invoke-Expression
```

### Go Install

If you have Go 1.22 or higher installed on your system, you can easily use the following command to install tlm;

```bash
go install github.com/yusufcanb/tlm@1.2
```

You're ready! Check installation by using the following command;

```bash
tlm
```

## Usage

```
$ tlm
NAME:
   tlm - terminal copilot, powered by open-source models.

USAGE:
   tlm suggest "<prompt>"
   tlm s --model=qwen2.5-coder:1.5b --style=stable "<prompt>"

   tlm explain "<command>" # explain a command
   tlm e --model=llama3.2:1b --style=balanced "<command>" # explain a command with a overrided model

   tlm ask "<prompt>" # ask a question
   tlm ask --context . --include *.md "<prompt>" # ask a question with a context

VERSION:
   1.2

COMMANDS:
   ask, a      Asks a question (beta)
   suggest, s  Suggests a command.
   explain, e  Explains a command.
   config, c   Configures language model, style and shell
   version, v  Prints tlm version.
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### Ask - Ask something with or without context

Ask a question with context. Here is an example question with a context of this repositories Go files under ask package.

```
$ tlm ask --help
NAME:
   tlm ask - Asks a question (beta)

USAGE:
   tlm ask "<prompt>" # ask a question
   tlm ask --context . --include *.md "<prompt>" # ask a question with a context

OPTIONS:
   --context value, -c value                                context directory path
   --include value, -i value [ --include value, -i value ]  include patterns. e.g. --include=*.txt or --include=*.txt,*.md        
   --exclude value, -e value [ --exclude value, -e value ]  exclude patterns. e.g. --exclude=**/*_test.go or --exclude=*.pyc,*.pyd
   --interactive, --it                                      enable interactive chat mode (default: false)
   --model value, -m value                                  override the model for command suggestion. (default: qwen2 5-coder:3b)
   --help, -h                                               show help
```

### Suggest - Get Command by Prompt

```
$ tlm suggest --help
NAME:
   tlm suggest - Suggests a command.

USAGE:
   tlm suggest <prompt>
   tlm suggest --model=llama3.2:1b <prompt>
   tlm suggest --model=llama3.2:1b --style=<stable|balanced|creative> <prompt>

DESCRIPTION:
   suggests a command for given prompt.

COMMANDS:
   help, h  Shows a list of commands or help for one command

OPTIONS:
   --model value, -m value  override the model for command suggestion. (default: qwen2.5-coder:3b)
   --style value, -s value  override the style for command suggestion. (default: balanced)        
   --help, -h               show help
```

### Explain - Explain a Command

```
$ tlm explain --help
NAME:
   tlm explain - Explains a command.

USAGE:
   tlm explain <command>
   tlm explain --model=llama3.2:1b <command>
   tlm explain --model=llama3.2:1b --style=<stable|balanced|creative> <command>

DESCRIPTION:
   explains given shell command.

COMMANDS:
   help, h  Shows a list of commands or help for one command

OPTIONS:
   --model value, -m value  override the model for command suggestion. (default: qwen2.5-coder:3b)
   --style value, -s value  override the style for command suggestion. (default: balanced)        
   --help, -h               show help
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
