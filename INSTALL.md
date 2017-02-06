
# Installation

## Compiled version

*datatools* is a collection of command line programs run from a shell like Bash (or Powershell in Windows). If you download the repository a compiled version is in the dist directory. The compiled VERSION_NO matching your computer type and operating system can be copied to a bin directory in your PATH.

Compiled versions are available for Mac OS X (amd64 processor), Linux (amd64), Windows (amd64) and Rapsberry Pi (both ARM6 and ARM7)

### Mac OS X

1. Download **datatools-VERSION_NO-release.zip** from [https://github.com/caltechlibrary/datatools/releases/latest](https://github.com/caltechlibrary/datatools/releases/latest)
2. Open a finder window, find and unzip **datatools-VERSION_NO-release.zip**
3. Look in the unziped folder and find *dist/macosx-amd64/*
4. Drag (or copy) *jsoncols*, *jsonrange*, etc. to a "bin" directory in your path
5. Open and "Terminal" and run `jsoncols -h` to confirm you were successful

### Windows

1. Download **datatools-VERSION_NO-release.zip** from [https://github.com/caltechlibrary/datatools/releases/latest](https://github.com/caltechlibrary/datatools/releases/latest)
2. Open the file manager find and unzip **datatools-VERSION_NO-release.zip**
3. Look in the unziped folder and find *dist/windows-amd64/*
4. Drag (or copy) *jsoncols.exe*, *jsonrange.exe*, etc. to a "bin" directory in your path
5. Open Bash and and run `jsoncols -h` to confirm you were successful

### Linux

1. Download **datatools-VERSION_NO-release.zip** from [https://github.com/caltechlibrary/datatools/releases/latest](https://github.com/caltechlibrary/datatools/releases/latest)
2. Find and unzip **datatools-VERSION_NO-release.zip**
3. In the unziped directory and find for *dist/linux-amd64/*
4. Copy *jsoncols*, *jsonrange*, etc. to a "bin" directory (e.g. cp ~/Downloads/datatools-VERSION_NO-release/dist/linux-amd64/\* ~/bin/)
5. From the shell prompt run `jsoncols -h` to confirm you were successful

### Raspberry Pi

If you are using a Raspberry Pi 2 or later use the ARM7 VERSION_NO, ARM6 is only for the first generaiton Raspberry Pi.

1. Download **datatools-VERSION_NO-release.zip** from [https://github.com/caltechlibrary/datatools/releases/latest](https://github.com/caltechlibrary/datatools/releases/latest)
2. Find and unzip **datatools-VERSION_NO-release.zip**
3. In the unziped directory and find for *dist/raspberrypi-arm7/*
4. Copy *jsoncols*, *jsonrange*, etc. to a "bin" directory (e.g. cp ~/Downloads/datatools-VERSION_NO-release/dist/raspberrypi-arm7/\* ~/bin/)
    + if you are using an original Raspberry Pi you should copy the ARM6 version instead
5. From the shell prompt run `jsoncols -h` to confirm you were successful


## Compiling from source

If you have go v1.7.4 or better installed then should be able to "go get" to install all the **datatools** utilities and

```
    go get -u github.com/caltechlibrary/datatools/...
```

Or for Windows 10 Powershell (assumes the Windows versions of Go and Git are previously installed)


```powershell
    $Env:GOPATH = "$HOME"
    go get -u github.com/caltechlibrary/datatools/...
```

or to install from source on Linux/Unix systems

```bash
    git clone https://github.com/caltechlibrary/datatools src/github.com/caltechlibrary/datatools
    cd src/github.com/caltechlibrary/datatools
    make
    make test
    make install
```

Or for Windows 10 Powershell

```powershell
    $Env:GOBIN = "$HOME\bin"
    git clone https://github.com/caltechlibrary/datatools src/github.com/caltechlibrary/datatools
    cd src\github.com\caltechlibrary\datatools
    go install cmds\jsoncols\filefile.go
```

