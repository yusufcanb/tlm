# OS and Architecture Detection
if ($env:OS -like 'Windows*') {
    $os = 'windows'
} else {
    Write-Error "Unsupported operating system. Only Windows is currently supported."
    return -1
}

if ($env:PROCESSOR_ARCHITECTURE -eq 'AMD64') {
    $arch = 'amd64'
} elseif ($env:PROCESSOR_ARCHITECTURE -eq 'ARM64') {
    $arch = 'arm64'
} else {
    Write-Error "Unsupported architecture. tlm requires a 64-bit system (x86_64 or arm64)."
    return -1
}

# Download URL Construction
$version = "1.2-pre"
$base_url = "https://github.com/yusufcanb/tlm/releases/download"
$download_url = "${base_url}/${version}/tlm_${version}_${os}_${arch}.exe"

# Ollama check
$ollamaHost = $env:OLLAMA_HOST
if (-not $ollamaHost) {
    $ollamaHost = "http://localhost:11434"
}

try {
    Invoke-WebRequest -Uri $ollamaHost -UseBasicParsing -ErrorAction Stop | Out-Null
} catch {
    Write-Host "ERR: Ollama not found." -ForegroundColor red
    Write-Host "If you have Ollama installed, please make sure it's running and accessible at $ollamaHost" -ForegroundColor red
    Write-Host "or configure OLLAMA_HOST environment variable." -ForegroundColor red
    Write-Host ""
    Write-Host ">>> If have Ollama on your system or network, you can set the OLLAMA_HOST like below;"
    Write-Host "    `$env:OLLAMA_HOST` = 'http://localhost:11434'"
    Write-Host ""
    Write-Host ""
    Write-Host ">>> If you don't have Ollama installed, you can install it using the following methods;"
    Write-Host ""
    Write-Host "    *** Windows: ***" -ForegroundColor green
    Write-Host "    Download instructions can be followed at the following link: https://ollama.com/download"
    Write-Host ""
    Write-Host "    *** Official Docker Images: ***" -ForegroundColor green
    Write-Host ""
    Write-Host "    Ollama can run with GPU acceleration inside Docker containers for Nvidia GPUs."
    Write-Host "    To get started using the Docker image, please follow these steps:"
    Write-Host ""
    Write-Host "        1. *** CPU only: ***" -ForegroundColor green
    Write-Host "            docker run -d -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama"
    Write-Host ""
    Write-Host "        2. *** Nvidia GPU: ***" -ForegroundColor green
    Write-Host "            docker run -d --gpus=all -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama"
    Write-Host ""
    Write-Host ""
    Write-Host "Installation aborted." -ForegroundColor red
    Write-Host "Please install or configure Ollama using the methods above and try again." -ForegroundColor red
    return
}

# Create Application Directory
$install_directory = "C:\Users\$env:USERNAME\AppData\Local\Programs\tlm"
if (-not (Test-Path $install_directory)) {
    New-Item -ItemType Directory -Path $install_directory | Out-Null
}

# Download the binary
Write-Host "Downloading tlm version ${version} for ${os}/${arch}..."
try {
    Invoke-WebRequest -Uri $download_url -OutFile "$install_directory\tlm.exe" -UseBasicParsing -ErrorAction Stop | Out-Null
} catch {
    Write-Error "Download failed. Please check your internet connection and try again."
    return -1
}

# Add installation directory to PATH

$user_env = [System.Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::User)

# Check if the installation directory is already in the PATH
if ($user_env -notcontains $install_directory) {
    # Add the installation directory to the PATH
    [System.Environment]::SetEnvironmentVariable("Path", "$user_env;$install_directory", [System.EnvironmentVariableTarget]::User)

    # Display a message indicating success
    Write-Host "Installation directory added to user PATH."
} else {
    # Display a message indicating that the directory is already in the PATH
    Write-Host "Installation directory is already in user PATH."
}

# Configure tlm to use Ollama
try {
    ."$install_directory\tlm.exe" config
} catch {
    Write-Error "tlm config set llm.host failed."
    return 1
}

Write-Host ""
Write-Host "Installation completed successfully."
Write-Host "Type 'tlm help' to get started."