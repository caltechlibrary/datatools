
# How to trim a trailing newline from a text file.

Given the following text in a file name _t.txt_ where the last line contains a trailing newline.

```
    one
    two
    three
```

Running *split* to create an JSON array yields an extra empty string.

```shell
    string -i t.txt split '\n'
```

Yields

```json
    ["one","two","three",""]
```

To avoid the trailing empty string in the array you can *trimspace* first then do your
*split* on newlines.

```shell
    string -i t.txt trimspace | split -i - '\n'
```

Yields

```json
    ["one","two","three"]
```
