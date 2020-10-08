@echo off
:: setlocal enableextensions enabledelayedexpansion

set "nul=nul"
if "%OS%"=="Windows_NT" set "nul="


:: Set variables

set "pwd=%cd%"
set "nativelang_dir=%USERPROFILE%\AppData\Roaming\Notepad++\"
set "nativelang_file=nativeLang.xml"
set "nativelang_full_path=%nativelang_dir%%nativelang_file%"
set "backup_dir=%cd%\backup\"


:: Health check

if not exist "%nativelang_dir%%nul%" (
	echo [error] Directory not found: %nativelang_dir%
	exit /b 1
)
if not exist "%nativelang_full_path%" (
	echo [error] File not found: %nativelang_file%
	exit /b 1
)
if not exist "%backup_dir%%nul%" (
	echo [error] Directory not found: %backup_dir%
	exit /b 1
)


:: Backup nativeLang.xml

echo.
copy /-y /v "%nativelang_full_path%" "%backup_dir%"
echo.
dir %backup_dir% /w
