%reltime(1) user manual | version 1.3.0 f486d87
% R. S. Doiel
% 2025-01-31

# NAME

reltime

# SYNOPSIS

reltime RELATIVE_TIME_STRING

# DESCRIPTION

reltime provides a relative time string in the "HH:MM:SS" in 24 hour format.

The notation for the relative time string is based on Go's time duration string. From
https://golang.google.cn/pkg/time/#ParseDuration, 

> A duration string is a possibly signed sequence of decimal numbers, each with 
> optional fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m".
> Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".


# EXAMPLES

Get the dime ninety minutes in the past.

~~~shell
    reltime -- -90m
~~~

Get the time 24 hours ago

~~~shell
    reltime -- -24h
~~~

Get the time 16 hours, 23 minutes and 4 seconds in the future.

~~~shell
	reltime 16h23m4s
~~~


