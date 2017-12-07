
# demo range

Create a range of integers one through five

```shell
	range 1 5
```

Yields 1 2 3 4 5

Create a range of integer negative two to six

```shell
	range -- -2 6
```

Yields -2 -1 0 1 2 3 4 5 6

Create a range of even integers two to ten

```shell
	range -increment=2 2 10
```

Yields 2 4 6 8 10

Create a descending range of integers ten down to one

```shell
	range 10 1
```

Yields 10 9 8 7 6 5 4 3 2 1


Pick a random integer between zero and ten

```shell
	range -r 0 10
```

Yields a random integer from 0 to 10

