Installation for development of **datatools**
===========================================

**datatools** A set of command line tools for working with CSV, Excel Workbooks, JSON and structured text documents.

Quick install with curl or irm
------------------------------

There is an experimental installer.sh script that can be run with the following command to install latest table release. This may work for macOS, Linux and if you’re using Windows with the Unix subsystem. This would be run from your shell (e.g. Terminal on macOS).

~~~shell
curl https://caltechlibrary.github.io/datatools/installer.sh | sh
~~~

This will install the programs included in datatools in your `$HOME/bin` directory.

If you are running Windows 10 or 11 use the Powershell command below.

~~~ps1
irm https://caltechlibrary.github.io/datatools/installer.ps1 | iex
~~~

Installing from source
----------------------

### Required software

- Golang &gt;&#x3D; 1.23.5
- Pandoc &gt;&#x3D; 3.1

### Steps

1. git clone https://github.com/caltechlibrary/datatools
2. Change directory into the `datatools` directory
3. Make to build, test and install

~~~shell
git clone https://github.com/caltechlibrary/datatools
cd datatools
make
make test
make install
~~~

