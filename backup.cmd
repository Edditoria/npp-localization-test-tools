@echo off
:: setlocal enableextensions enabledelayedexpansion

set "nul=nul"
if "%OS%"=="Windows_NT" set "nul="

:: Set variables

set "nativelang_dir=%USERPROFILE%\AppData\Roaming\Notepad++\"
set "nativelang_file=nativeLang.xml"
set "nativelang_full_path=%nativelang_dir%%nativelang_file%"
set "backup_dir=%nativelang_dir%\backup\"

:: Temp Output

echo.
echo Doing: echo %%nativelang_full_path%%
echo %nativelang_full_path%
echo.
echo Doing: dir %%backup_dir%%
dir %backup_dir% /w
