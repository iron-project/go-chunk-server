### setup

```bash

$ go get -u -v github.com/gin-gonic/gin
$ go run incs.go
```

### interface

- upload

```bash
$ curl -X PUT 'http://127.0.0.1:10011/v1/chunks?name=xxx' \
-F "file=@/User/rain/incs/incs.go" \
-H "Content-Type: multipart/form-data"
```

- download

```bash
$ curl -X GET 'http://127.0.0.1:10011/v1/chunks?name=incs.go'
```
