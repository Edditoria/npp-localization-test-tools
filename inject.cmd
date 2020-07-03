@echo off
:: setlocal enableextensions enabledelayedexpansion

set "nul=nul"
if "%OS%"=="Windows_NT" set "nul="


:: Set variables

set "pwd=%cd%"
set "user_input=%~1"
set "file_full_path=%pwd%\%user_input%"

if "[%user_input%]"=="[]" (
	echo [error] Missing user argument
	exit /b 1
)
:: else
echo.
echo Full path of user input:
echo %file_full_path%
echo.

if exist "%file_full_path%" (
	echo [success] File is found!
) else (
	echo [error] File not found!
)
