
# Using timefmt

Format the date July, 7, 2016 in YYYY-MM-DD format

```shell
    timefmt -input "2006-01-02" -output "01/02/2006" "2016-07-02"
```

Yields "07/02/2016"

Format the MySQL date/time of 8:08am, July 2, 2016

```shell
    timefmt -input mysql -output RFC822  "2016-07-02 08:08:08"
```

Yields "02 Jul 16 08:08 UTC"

## example files

- [timefmt-demo.bash](timefmt-demo.bash)

