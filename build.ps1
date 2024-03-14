# Parameters
$targets = "darwin", "linux", "windows"
$archs = "amd64", "arm64"
$appName = "tlm"

# Command-Line Version Argument
$version = $args[0]
if (-not $version) {
    Write-Output "Error: Please provide a version number as a command-line argument."
    exit 1
}

# Housekeeping
Remove-Item -Recurse -Force "dist" -ErrorAction SilentlyContinue
New-Item -ItemType Directory -Path "dist"

# Build Function (Helper)
Function Build-Target($os, $version, $arch) {
    $outputName = "${appName}_${version}_${os}_${arch}"
    $sha1 = (git rev-parse --short HEAD).Trim()
    if ($os -eq "windows") {
        $outputName += ".exe"
    }

    Write-Output "Building for $os/$arch (version: $version) -> $outputName"
    # Invokes the Go toolchain (assumes it's in the PATH)
    go build -o "dist/$version/$outputName" -ldflags "-X main.sha1ver=$sha1" "main.go"
}

# Build for each target OS
foreach ($os in $targets) {

    foreach ($arch in $archs) {
        $env:GOOS = $os
        $env:GOARCH = $arch
        $env:CGO_ENABLED = 0
        Build-Target $os $version $arch
    }

}

Write-Output "Done!"