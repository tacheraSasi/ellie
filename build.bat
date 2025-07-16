@echo off
REM build.bat - Builds the Ellie executable and creates the NSIS installer

REM --- Configuration ---
SET NSIS_PATH="C:\Program Files (x86)\NSIS\makensis.exe"
SET PROJECT_ROOT=%~dp0
SET DIST_DIR=%PROJECT_ROOT%dist
SET INSTALLER_DIR=%PROJECT_ROOT%installer
SET EXE_NAME=ellie.exe
SET OUTPUT_EXE=%DIST_DIR%%EXE_NAME%
SET NSIS_DIR="C:\Program Files (x86)\NSIS"

REM --- Build Steps ---

REM 1. Build the Go executable
echo Building Go executable...
cd %PROJECT_ROOT%
set GOOS=windows
set GOARCH=amd64
go build -ldflags="-s -w" -o %OUTPUT_EXE% .
if %errorlevel% neq 0 (
    echo Go build failed!
    exit /b 1
)
echo Go executable built successfully at %OUTPUT_EXE%

REM 2. Copy executable to installer directory
echo Copying executable to installer directory...
copy /Y %OUTPUT_EXE% %INSTALLER_DIR%%EXE_NAME%
echo Executable copied.

REM 3. Compile the NSIS installer
echo Compiling NSIS installer...
cd %INSTALLER_DIR%
set NSISDIR=%NSIS_DIR%
%NSIS_PATH% ellie.nsi
if %errorlevel% neq 0 (
    echo NSIS compilation failed!
    exit /b 1
)
echo Installer created successfully!

echo.
echo Build complete. You can find Ellie-Setup-1.0.0.exe in the 'installer' directory.
pause
