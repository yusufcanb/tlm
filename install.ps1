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
$version = "1.0-rc3"
$base_url = "https://github.com/yusufcanb/tlm/releases/download"
$download_url = "${base_url}/${version}/tlm_${version}_${os}_${arch}.exe"

# Docker check
if (-not (Get-Command docker -ErrorAction SilentlyContinue)) {
    Write-Error "Docker not found. Please install Docker from https://www.docker.com/get-started"
    return -1
}

# Ollama check - For Windows, we'll assume Ollama is installed directly on the system
try {
    Invoke-WebRequest -Uri "http://localhost:11434" -UseBasicParsing -ErrorAction Stop | Out-Null
} catch {
    Write-Host "ERR: Ollama not found." -ForegroundColor red
    Write-Host ""
    Write-Host "*** On Windows: ***" -ForegroundColor green
    Write-Host ""
    Write-Host "Ollama can run with GPU acceleration inside Docker containers for Nvidia GPUs."
    Write-Host "To get started using the Docker image, please follow these steps:"
    Write-Host ""
    Write-Host "1. *** CPU only: ***" -ForegroundColor green
    Write-Host "   docker run -d -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama"
    Write-Host ""
    Write-Host "2. *** Nvidia GPU: ***" -ForegroundColor green
    Write-Host "   docker run -d --gpus=all -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama"
    Write-Host ""
    Write-Host ""
    Write-Host "Installation aborted." -ForegroundColor red
    Write-Host "Please install Ollama using the methods above and try again." -ForegroundColor red
    return -1
}

# Download the binary
Write-Host "Downloading tlm version ${version} for ${os}/${arch}..."
try {
    Invoke-WebRequest -Uri $download_url -OutFile 'tlm.exe' -UseBasicParsing -ErrorAction Stop | Out-Null
} catch {
    Write-Error "Download failed. Please check your internet connection and try again."
    return -1
}

# Move to installation directory
Write-Host "Installing tlm..."
#try {
#    Move-Item -Path 'tlm.exe' -Destination 'C:\Windows\Program Files\tlm\' -Force
#} catch {
#    Write-Error "Installation requires administrator permissions. Please elevate with rights and run the script again."
#    exit 1
#}

# Ollama deployment - specific to the original script, might need modification
try {
    tlm deploy
} catch {
    Write-Error "tlm deploy failed."
    return 1
}

Write-Host "Type 'tlm.exe help' to get started."
return 0