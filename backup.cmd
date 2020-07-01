@echo off
:: setlocal enableextensions enabledelayedexpansion

set "nul=nul"
if "%OS%"=="Windows_NT" set "nul="

:: Set variables
set "npp_lang_dir=%USERPROFILE%\AppData\Roaming\Notepad++"
set "npp_lang_file=nativeLang.xml"
set "npp_lang_path=%npp_lang_dir%\%npp_lang_file%"

echo.
echo Doing: echo %%npp_lang_path%%
echo %npp_lang_path%
echo.
echo Doing: dir %%npp_lang_dir%%
dir %npp_lang_dir% /w
