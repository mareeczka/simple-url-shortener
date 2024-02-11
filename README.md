#simple-url-shortener
Simple URL shortener API written in Golang while I was learning the language. Outputs JSON. 

Uses SQLite to store data and thus requires https://github.com/mattn/go-sqlite3 package.

## Build locally

```
go build
```
## Docker

```
docker run -d -p 127.0.0.1:8080:8080 <image-name>
```
