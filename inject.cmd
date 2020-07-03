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
set "user_input_file_ext=%user_input:~-4%"
if not "%user_input_file_ext%"==".xml" (
	echo [error] Not a .xml file
	exit /b 1
)


:: Copy nativeLang.xml

echo.
echo It will overwrite existing nativeLang.xml in Notepad++.
set /p "start_copy=Are you sure? (type [y] to continue) "
if "%start_copy%"=="y" goto confirm_pass
:: else
echo.
echo [error] This script will continue only if you answer "y"
exit /b 1

:confirm_pass
echo.
copy /y /v "%src%" "%dest%"
echo [done]
