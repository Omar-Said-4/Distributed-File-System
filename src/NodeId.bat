@echo off
setlocal enabledelayedexpansion

set "config_file=node\config\config.json"
set "temp_file=%config_file%.tmp"

:: Check if the file exists
if not exist "%config_file%" (
    echo Config file not found: %config_file%
    exit /b 1
)

:: Read and modify the file
(for /f "usebackq delims=" %%A in ("%config_file%") do (
    set "line=%%A"
    echo %%A | findstr /C:"\"nodeID\"" >nul
    if not errorlevel 1 (
        echo "nodeID": -1,
    ) else (
        echo %%A
    )
)) > "%temp_file%"

:: Replace original file
move /y "%temp_file%" "%config_file%" > nul

echo Updated nodeID to -1 in %config_file%
