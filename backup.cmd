@echo off
:: setlocal enableextensions enabledelayedexpansion

set "nul=nul"
if "%OS%"=="Windows_NT" set "nul="


:: Set variables

set "nativelang_dir=%USERPROFILE%\AppData\Roaming\Notepad++\"
set "nativelang_file=nativeLang.xml"
set "nativelang_full_path=%nativelang_dir%%nativelang_file%"
set "backup_dir=%nativelang_dir%\backup\"


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


:: Temp Output

echo.
echo Doing: echo %%nativelang_full_path%%
echo %nativelang_full_path%
echo.
echo Doing: dir %%backup_dir%%
dir %backup_dir% /w
