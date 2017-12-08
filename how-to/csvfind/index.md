
# Using csvfind

Find the rows where the third column matches "The Red Book of Westmarch" exactly

```shell
    csvfind -i books.csv -o result1.csv -col=2 "The Red Book of Westmarch"
```

Find the rows where the third column (colums numbered 0,1,2) matches approximately 
"The Red Book of Westmarch"

```shell
    csvfind -i books.csv -o result2.csv -col=2 -levenshtein \
       -insert-cost=1 -delete-cost=1 -substitute-cost=3 \
       -max-edit-distance=50 -append-edit-distance \
       "The Red Book of Westmarch"
```

In this example we've appended the edit distance to see how close the matches are.

You can also search for phrases in columns.

```shell
    csvfind -i books.csv -o result3.csv -col=2 -contains "Red Book"
```

