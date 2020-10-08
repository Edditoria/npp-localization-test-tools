# About This Repo

This repo contains some simple batch scripts for testing your localization XML file in Notepad++. It can:

- [x] Backup current nativeLang.xml in Notepad++.
- [ ] Restore a nativeLang.xml in Notepad++.
- [ ] Customize path in your localization repo.


## Usage

Assume that you are working on a localization file `<my_lang.xml>` in a directory `<C:\dev\npp_my_lang\>`

```cmd
:: In your repo: C:\dev\npp_my_lang\>

:: Add this repo as git submodule named "tools"
git submodule add https://github.com/Edditoria/npp-localization-test-tools.git tools

:: Backup current nativeLang.xml in Notepad++
.\tools\backup.cmd

:: Inject (Force copy) the XML file to Notepad++
.\tools\inject.cmd .\my_lang.xml
```

You may want to add instruction for your repo:

```cmd
:: Get submodule
git submodule init
git submodule update
```

Quick and dirty. Welcome issue and pull requests ;-)
