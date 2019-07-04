# ushtr (URL shortener)

[![Scc Count Badge](https://sloc.xyz/github.com/vasyahuyasa/ushtr/)](https://github.com/vasyahuyasa/ushtr/)

Ushtr is simple url shortener service with scalability in mind.

## API

### Create URL

```
http://ushtr.org/-?q=google.ru&mode=json
```
Support `GET`, `POST`, `PUT` methods.

If method is `POST` or `PUT` then data content type must be `application/x-www-form-urlencoded`.

__`q`__ is URL for shorten.

__`mode`__ is result type, can be `text` or `json`. If `mode` is `text` then result will be plain text contain short url. If `mode` is `json` then result will be json object in format:

```
{
    short: "https://ushtr.org/aFvb35Q"
}
``` 

By default is `text`.

#### Response codes
| Code   | Comment        |
|--------|----------------|
| 200    | Ok             |
| 400    | Bad param data |
| 503    | Internal server error |
'

## Retrive URL

```

```
