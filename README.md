# About This Repo

This repo contains some simple batch scripts for testing your localization XML file in [Notepad++][npp_site]. It can:

- [x] Backup current nativeLang.xml in Notepad++.
- [ ] Restore a nativeLang.xml in Notepad++.
- [ ] Customize path in your localization repo.


## Usage

Assume that you are working on a localization file `<my_lang.xml>` in a directory `<C:\dev\npp_my_lang\>`:

```cmd
:: For example, in your repo: C:\dev\npp_my_lang\>

:: Add this repo as git submodule named "tools"
git submodule add https://github.com/Edditoria/npp-localization-test-tools.git tools

:: Backup current nativeLang.xml in Notepad++
.\tools\backup.cmd

:: Inject (Force copy) the XML file to Notepad++
.\tools\inject.cmd .\my_lang.xml

:: Update submodule(s)
git submodule update
```

You may want to add instructions in your localization repo:

```cmd
:: Get submodule
git submodule init
git submodule update
```

Quick and dirty. Welcome issue and pull requests ;-)

## Copyright and License

Copyright (c) Edditoria. All rights reserved. Code released under the [MIT License](LICENSE.txt). Docs released under [Creative Commons](https://creativecommons.org/licenses/by/4.0/).

As human-readable summary (but not a substitute for the license):

You can use it, share it, modify the codes and distribute your work for private and commercial uses. If you like, please share your work with me. :pizza:

> Notepad++ is a free (free as in both "free speech" and "free beer") source code editor and Notepad replacement that supports several programming languages and natural languages. Running in the MS Windows environment, its use is governed by [GPL License](https://github.com/notepad-plus-plus/notepad-plus-plus/blob/master/LICENSE).


[npp_site]: https://notepad-plus-plus.org
