%codemeta2cff(1) user manual | version 1.3.4 4312aaa
% R. S. Doiel
% 2025-05-15

# NAME

codemeta2cff

# SYSNOPSIS

codemeta2cff [OPTIONS] [CODEMETA_JSON CITATION_CFF]

# DESCRIPTION

Reads codemeta.json file and writes CITATION.cff. By default
it assume both are in the current directory.  You can also
provide the name and path to both files.

# OPTIONS

-help
: display help

# EXAMPLE

Generating the CITATION.cff from the codemeta.json file the current
working directory.

~~~
codemeta2cff
~~~

Specifying the full paths.

~~~
codemeta2cff /opt/local/myproject/codemeta.json /opt/local/myproject/CITATION.cff
~~~


