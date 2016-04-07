# Haribote

Haribote is a Super Simple HTTP Mock Server for your usual development.


### Installation

```
go get github.com/kunihiko-t/haribote
```

### Usage

```
Usage of Haribote:
haribote [OPTIONS] ARGS...
Options
 -f string
     Config file name
 -p int
     Port number (default "9090")
```

Run

```
% haribote -f path/to/config.json -p 9090
2016/04/07 02:39:36 Registered Handler {/text GET text/plain 200 hello }
2016/04/07 02:39:36 Registered Handler {/json GET application/json 200  ./sample/test.json}
2016/04/07 02:39:36 Registered Handler {/image GET image/jpeg 200  ./sample/malta.jpg}
2016/04/07 02:39:36 Registered Handler {/post POST text/plain 200  ./sample/test.txt}
2016/04/07 02:39:36 Registered Handler {/error500 GET text/plain 500 Internal Server Error }
```
You can access localhost like http://localhost:9090/text

### Config file example
```json
{
  "Server": [
    {"Path": "/text", "Method": "GET", "ContentType":"text/plain", "StatusCode":200, "Text": "hello"},
    {"Path": "/json", "Method": "GET", "ContentType":"application/json", "StatusCode":200, "File":"./sample/test.json"},
    {"Path": "/image", "Method": "GET", "ContentType":"image/jpeg", "StatusCode":200, "File":"./sample/malta.jpg"},
    {"Path": "/post", "Method": "POST", "ContentType":"text/plain", "StatusCode":200, "File":"./sample/test.txt"},
    {"Path": "/error505", "Method": "GET", "ContentType":"text/plain", "StatusCode":500, "Text":"Internal Server Error"}
  ]
}
```


## License
MIT
