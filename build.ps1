# build.ps1 - Builds the Ellie executable and creates the NSIS installer

# --- Configuration ---
$NsisPath = "C:\Program Files (x86)\NSIS\makensis.exe"
$ProjectRoot = $PSScriptRoot
$DistDir = "$ProjectRoot\dist"
$InstallerDir = "$ProjectRoot\installer"
$InstallerScript = "$InstallerDir\ellie.nsi"
$NsisIncludeDir = "C:\Program Files (x86)\NSIS\Include"
$ExeName = "ellie.exe"
$OutputExe = "$DistDir\$ExeName"

# --- Build Steps ---

# 1. Build the Go executable
Write-Host "Building Go executable..."
Push-Location $ProjectRoot
$env:GOOS = "windows"
$env:GOARCH = "amd64"
go build -ldflags="-s -w" -o $OutputExe .
if ($LASTEXITCODE -ne 0) {
    Write-Error "Go build failed!"
    exit 1
}
Pop-Location
Write-Host "Go executable built successfully at $OutputExe"

# 2. Copy executable to installer directory
Write-Host "Copying executable to installer directory..."
Copy-Item -Path $OutputExe -Destination "$InstallerDir\$ExeName" -Force
Write-Host "Executable copied."

# 3. Compile the NSIS installer
Write-Host "Compiling NSIS installer..."
Push-Location $InstallerDir
& $NsisPath /I"$NsisIncludeDir" $InstallerScript
if ($LASTEXITCODE -ne 0) {
    Write-Error "NSIS compilation failed!"
    exit 1
}
Pop-Location
Write-Host "Installer created successfully!"

Write-Host "Build complete. You can find the installer in the 'installer' directory."
