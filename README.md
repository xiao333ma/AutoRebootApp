在 windows 上自动重启崩溃的程序

修改 config.json 来配置需要崩溃重启的 APP

``` sh
GOOS=windows GOARCH=amd64 go build .
```