
# Convert lines of text into an array

Given the following text in a file named _t.txt_.

```
    one
    two
    three
```

How to you get `["one","two","three"]` using the string command?

On approach that might need like it'd work would be to use the *join*
action word with the _string_ command. But that wouldn't given you what
you want. You actually want to "split" the text file on the new line character.

```shell
   string -i t.txt split "\n" 
```
