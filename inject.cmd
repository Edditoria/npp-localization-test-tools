@echo off
:: setlocal enableextensions enabledelayedexpansion

set "nul=nul"
if "%OS%"=="Windows_NT" set "nul="


:: Set variables

set "pwd=%cd%"
set "user_input=%~1"
set "src=%pwd%\%user_input%"
set "nativelang_dir=%USERPROFILE%\AppData\Roaming\Notepad++\"
set "nativelang_filename=nativeLang.xml"
set "dest=%nativelang_dir%%nativelang_filename%"

:: Health check

if "[%user_input%]"=="[]" (
	echo [error] Missing user argument
	exit /b 1
)
if exist "%user_input%\*" (
	echo [error] It is a directory
	exit /b 1
)
if not exist "%src%" (
	echo [error] File not found!
	exit /b 1
)


:: Copy nativeLang.xml

echo.
copy /y /v "%src%" "%dest%"
echo [done]
