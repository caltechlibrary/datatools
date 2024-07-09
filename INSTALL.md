
Installation
============

*datatools* is a collection of command line programs run from a shell like Bash. 


Quick install using curl
------------------------

The following experimental installer should work for macOS and Linux
(e.g. Debian, Ubuntu, Raspberry Pi OS)

Copy and run the following command in your shell (e.g. Terminal)

~~~
curl https://caltechlibrary.github.io/datatools/installer.sh | sh
~~~

On Windows we provide a Powershell script that can be run with

~~~
irm https://caltechlibrary.github.io/datatools/installer.ps1 | iex
~~~

Compiled version
----------------

This is generalized instructions for a release. 

Compiled versions are available for Mac OS X (Intel and M1 processor, macos-x86_64 and macOS-arm64), Linux (Intel process, Linux-x86_64), 
Windows (Intel and arm64 processor, windows-x86_64 and Windows-arm64) and Rapsberry Pi (arm7 processor, RaspberryPiOS-arm7)

VERSION_NUMBER is a [symantic version number](http://semver.org/) (e.g. v0.1.2)


For all the released version go to the project page on Github and click latest release

>    https://github.com/caltechlibrary/datatools/releases/latest


| Platform       | Zip Filename                                      |
|----------------|---------------------------------------------------|
| Windows        | datatools-VERSION_NUMBER-Windows-x86_64.zip       |
| Windows        | datatools-VERSION_NUMBER-Windows-arm64.zip        |
| Mac OS X       | datatools-VERSION_NUMBER-macOS-x86_64.zip         |
| Mac OS X       | datatools-VERSION_NUMBER-macOS-arm64.zip          |
| Linux/Intel    | datatools-VERSION_NUMBER-Linux-x86_64.zip         |
| Linux/ARM 64   | datatools-VERSION_NUMBER-Linux-aarch64.zip        |
| Raspbery Pi OS | datatools-VERSION_NUMBER-RaspberryPiOS-arm7.zip   |


The basic recipe
----------------

+ Find the Zip file listed matching the architecture you're running and download it
    + (e.g. if you're on a Windows 10 laptop/Surface with a Intel style CPU you'd choose the Zip file with "windows-x86_64" in the name).
+ Download the zip file and unzip the file.  
+ Copy the contents of the folder named "bin" to a folder that is in your path 
    + (e.g. "$HOME/bin" is common).
+ Adjust your PATH if needed
    + (e.g. export PATH="$HOME/bin:$PATH")
+ Test


### Mac OS

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Make sure the new location in in our path
5. Test

Here's an example of the commands run in the Terminal App after downloading the 
zip file.

#### Intel (x86_64) Hardware

```shell
    cd Downloads/
    unzip datatools-*-macos-x86_64.zip
    mkdir -p $HOME/bin
    mv -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    csvfind -version
```

#### ARM64 (arm64) Hardware

```shell
    cd Downloads/
    unzip datatools-*-macos-arm64.zip
    mkdir -p $HOME/bin
    mv -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    csvfind -version
```


### Windows

(Assumes you're working from Bash as provided by Linux Subsystem for Windows)

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell on Windows 10 after
downloading the zip file.

#### Intel (x86_64) Hardware

```shell
    cd Downloads/
    unzip datatools-*-windows-x86_64.zip
    mkdir -p $HOME/bin
    mv -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    csvfind -version
```

#### ARM64 (arm64) Hardware

```shell
    cd Downloads/
    unzip datatools-*-windows-arm64.zip
    mkdir -p $HOME/bin
    mv -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    csvfind -version
```



### Linux 

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell after
downloading the zip file.

```shell
    cd Downloads/
    unzip datatools-*-linux-x86_64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    csvfind -version
```


### Raspberry Pi

Released version is for a Raspberry Pi 2 or later use (i.e. requires ARM 7 support).

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell after
downloading the zip file.

```shell
    cd Downloads/
    unzip datatools-*-raspberry_pi_os-arm7.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    csvfind -version
```


Compiling from source
---------------------

_datatools_ is "go gettable" if you have previously gotten xlsx v1.0.5 package from [github.com/tealeg/xlsx](https://github.com/tealeg/xlsx).  The datatools package does not support versions v2.x and greater of xlsx. Below are the steps I use today with "go get" command to download the dependant packages as well as _datatools_'s source code.

Setting up the right version of xlsx for datatools

```shell
    cd
    go get github.com/tealeg/xlsx
    cd src/github.com/tealeg
    git checkout v1.0.5
    cd
```

Using `go get` to install datatools using v1.0.5 of xlsx.

```
    go get github.com/caltechlibrary/datatools/...
```

Or clone the repository and then compile

```shell
    cd
    git clone https://github.com/caltechlibrary/datatools src/github.com/caltechlibrary/datatools
    cd src/github.com/caltechlibrary/datatools
    make
    make test
    make install
```

