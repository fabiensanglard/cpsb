# The Book of CP-System

A study of the revolutionary hardware used by Capcom from 1989 to 1995 to release breath-taking arcade games such as Street Fighter II, Final Fight, and Ghouls 'n Ghosts.

The source code is open-source but I retain ownership of all drawings. They are only provided here so the project can be compiled.

## How to build (Linux)

- Setup env
```
mkdir ~/.fonts
cp src/fonts/* ~/.fonts
fc-cache -f -v
sudo apt-get update
sudo apt install -y golang-go
go get golang.org/x/image/draw
sudo add-apt-repository universe
sudo add-apt-repository ppa:inkscape.dev/stable
sudo apt-get update
sudo apt install inkscape
sudo apt install -y texlive-full
```

- Build
```
# for english
./make.sh release

# or for french translation
./make.sh release fr
```

Note: Make sure Inkscape version is at least v1.0.0 or SVG -> PNG conversion will fail.

## How to build (Windows)

- Install WSL1/WSL2
- Follow Linux steps.

## How to build (MacOS X)
- Install fonts (in `src/fonts` folder)
- Get MacTeX.
- Install golang (https://golang.org/doc/install)
- Install Inkscape > v1.0
- Symlink inkscape binary to `/usr/local/bin` or add it to your PATH
- Get build system dependencies 
```
go mod init "image/draw"
go get golang.org/x/image/draw
```
- Build
```
./make.sh release
```



### Debug mode

Building is slow when using 300dpi assets. To develop, the build system generates low dpi assets. 
Debug is the default mode.

```
./make.sh debug

# or for french translation
./make.sh debug fr
```

### Force asset regeneration
```
./make.sh debug|release -f
```

### Poor typesetting

If upon compilation the typesetting is "off" with blank pages, it may be that the wrong version of TexLive is used to compile. Run the following command and make sure you are using version 2023.

```
$ tex --version
TeX 3.141592653 (TeX Live 2023)
```

This book is verified to compile properly with TexLive 2023. In case you have another version, use the following steps to install the lastest.

[Uninstall TexLive](https://puttym.github.io/update-texlive)

[Download TexLive NetInstall](https://www.tug.org/texlive/acquire-netinstall.html)

[Install TexLive](https://www.tug.org/texlive/quickinstall.html)
