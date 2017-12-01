
# Re-order a comma delimiter string

The examples below are each borken down in two parts. The first
part shows a short version of the commands you might use in practice.
This is followed by a step by step implementation to take some
of the mystery out of the short version.

## Problem

How to easily convert a name in "FAMILY, GIVEN" form
to "GIVEN FAMILY" form? 

## The traditional shell approach

Traditionally this can be done using a number of Unix commands such as
_echo_ and _cut_. If you break down the task into getting each name separately it is pretty easy using a couple shell variables.

```shell
    NAME="Doiel, Robert"
    GIVEN_FAMILY="$(echo -n "$NAME" | cut -d , -f 1) $(echo -n "$NAME" | cut -d , -f 2)"
    echo "$FAMILY $GIVEN"
```

This uses the subshell syntax and two separate pipe lines. Let's break it down by parts.

```shell
    NAME="Doiel, Robert"
    echo "Step 1: [$NAME]"
    FAMILY_NAME=$(echo -n "$NAME" | cut -d , -f 1)
    echo "Step 2: [$FAMILY_NAME]
    GIVEN_NAME=$(echo -n "$NAME" | cut -d , -f 2)
    echo "Step 3: [$GIVEN_NAME $FAMILY_NAME]"
```

Each pipeline builds up a name (family and given) and the final _echo_
displays them.

NOTE: The trouble is this doesn't give you what you want. 
Notice the leading space. You can fix that but that is just the start of 
the rabbit whole.

## The datatools aproach

If we think about the family name and given name as elements of an
array we can easily reorder them. In this approach we'll use
two commands.  The _string_ and _jsoncols_ commands from _datatools_.

```shell
    NAME="Doiel, Robert"
    string split ", " "$NAME" | jsoncols -i - -d ' ' '.[0]' '.[1]'
```

First difference you'll notice is we're using an Unix pipe to send
the output of one command to another. While you can eventually do that
in the traditional approach it becomes very complicated very quickly.
Using _datatools_ it is easy to move from strings to JSON and back.

Let's take the _datatools_ approach and output the results of each step
rather than using a pipeline.

```shell
    NAME="Doiel, Robert"
    echo "Step 1: [$NAME]"
    JSON_ARRAY=$(string split ", " "$NAME")
    echo "Step 2: $JSON_ARRAY"
    GIVEN_FAMILY=$(jsoncols -d ' ' '.[0]' '.[1]' "$JSON_ARRAY")
    echo "Step 3: [$GIVEN_FAMILY]"
```

NOTE: In this version there is no leading space issue. _string_ command
can split on multiple characters and in our case it is splitting on 
", " not just on the comma like we get with _cut_.

