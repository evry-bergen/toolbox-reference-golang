# Golang Reference Project

REST API written in Go using [Gin](https://github.com/gin-gonic/gin) HTTP
framework.

## Development

### Go Modules

Dependencies are handled by [Go Modules](https://blog.golang.org/modules2019),
and a good example of how to use them are avaiable
[here](https://www.kablamo.com.au/blog-1/2018/12/10/just-tell-me-how-to-use-go-modules).

Install new dependnecy:

```
go get -u <repo url>
```

Vendor dependencies:

```
go mod vendor
```


# Topics

## Graceful shutdown for http, grpc etc
https://gist.github.com/akhenakh/38dbfea70dc36964e23acc19777f3869

## Graceful shutdown for http
https://medium.com/honestbee-tw-engineer/gracefully-shutdown-in-go-http-server-5f5e6b83da5a
