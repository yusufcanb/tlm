# tlm - Local terminal companion, powered by CodeLLaMa.

tlm is your CLI companion which requires nothing then your workstation. It uses most efficient and powerful [CodeLLaMa](https://ai.meta.com/blog/code-llama-large-language-model-coding/) in your local environment to provide you the best possible command line suggestions.

![](./assets/suggest.gif)

![](./assets/explain.gif)

![](./assets/config.gif)

## Features

- 💸 No API Key (Subscription) is required. (ChatGPT, Github Copilot, Azure OpenAI, etc.) 

- 📡 No internet connection is required.

- 💻 Works on MacOS, Linux and Windows.

- 👩🏻‍💻 Automatic shell detection.
 
- 🚀 One liner generation and command explanation.


## Installation

Installation can be done in two ways;

- Installation script (recommended)
- Go Install

### Go Install

If you Go 1.21 or higher installed on your system, you can easily use the following command to install tlm;

```bash
go install github.com/yusufcanb/tlm@latest
```

Check installation by using the following command;

```bash
tlm help
```

###  Installation Script

Installation script is the recommended way to install tlm.
It will recognize the which platform and architecture to download and will execute install command for you. 

#### Linux and MacOS;


Download and execute the installation script by using the following command;

```bash 
curl -fsSL https://raw.githubusercontent.com/yusufcanb/tlm/main/install.sh | bash
```

#### Windows (Powershell 5.1 or higher)

Download and execute the installation script by using the following command;

```powershell
Invoke-RestMethod -Uri https://raw.githubusercontent.com/yusufcanb/tlm/main/install.ps1 | Invoke-Expression
```

## Uninstall

On Linux and MacOS;

```bash
rm /usr/local/bin/tlm
```
