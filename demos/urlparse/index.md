
# demo urlparse

With no options returns "http\texample.com\t/my/page.html"

```shell
    urlparse http://example.com/my/page.html
```

Get protocol. Returns "http".

```shell
    urlparse -protocol http://example.com/my/page.html
```

Get host or domain name.  Returns "example.com".

```shell
    urlparse -host http://example.com/my/page.html
```

Get path. Returns "/my/page.html".

```shell
    urlparse -path http://example.com/my/page.html
```

Get dirname. Returns "my"

```shell
    urlparse -dirname http://example.com/my/page.html
```

Get basename. Returns "page.html".

```shell
    urlparse -basename http://example.com/my/page.html
```

Get extension. Returns ".html".

```shell
    urlparse -extname http://example.com/my/page.html
```

Without options urlparse returns protocol, host and path
fields separated by a tab.

