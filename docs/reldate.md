
# USAGE

## reldate [OPTIONS] [TIME_DESCRPTION]

## SYNOPSIS

reldate is a small command line utility which returns the relative date in 
YYYY-MM-DD format. This is helpful when scripting various time 
relationships. The difference in time returned are determined by 
the time increments provided.

Time increments are a positive or negative integer. Time unit can be
either day(s), week(s), month(s), or year(s). Weekday names are
case insentive (e.g. Monday and monday). They can be abbreviated
to the first three letters of the name, e.g. Sunday can be Sun, Monday
can be Mon, Tuesday can be Tue, Wednesday can be Wed, Thursday can
be Thu, Friday can be Fri or Saturday can be Sat.

## OPTIONS

	-e	Display the end of month day. E.g. 2012-02-29
	-end-of-month	Display the end of month day. E.g. 2012-02-29
	-f	Date the relative time is calculated from.
	-from	Date the relative time is calculated from.
	-h	display help
	-help	display help
	-l	display license
	-license	display license
	-v	display version
	-version	display version

## EXAMPLES

If today was 2014-08-03 and you wanted the date three days in the past try–

```
    reldate 3 days
```

The output would be

```
    2014-08-06
```

TIME UNITS

Supported time units are

+ day(s)
+ week(s)
+ year(s)

Specifying a date to calucate from

reldate handles dates in the YYYY-MM-DD format (e.g. March 1, 2014 would be 
2014-03-01). By default reldate uses today as the date to calculate relative 
time from. If you use the –from option you can it will calculate the 
relative date from that specific date.

```
   reldate --from=2014-08-03 3 days
```

Will yield

2014-08-06

NEGATIVE INCREMENTS

Command line arguments traditionally start with a dash which we also use to 
denote a nagative number. To tell the command line process that to not treat 
negative numbers as an “option” preceed your time increment and time unit 
with a double dash.

```
    reldate --from=2014-08-03 -- -3 days
```

Will yield

```
    2014-07-31
```

RELATIVE WEEK DAYS

You can calculate a date from a weekday name (e.g. Saturday, Monday, Tuesday) 
knowning a day (e.g. 2015-02-10 or the current date of the week) occuring in 
a week. A common case would be wanting to figure out the Monday date of a week 
containing 2015-02-10. The week is presumed to start on Sunday (i.e. 0) and 
finish with Saturday (e.g. 6).

```
    reldate --from=2015-02-10 Monday
```

will yeild

```
    2015-02-09
```

As that is the Monday of the week containing 2015-02-10. Weekday names case 
insensitive and can be the first three letters of the English names or full 
English names (e.g. Monday, monday, Mon, mon).


reldate v0.0.17
