# GophKeeper

GophKeeper yandex practicum

https://app.diagrams.net/#G1904ze4l_bOQ40mxAZn48c9aqsFlkO3gO

https://github.com/Vasily-van-Zaam/GophKeeper/blob/server/GophKeeper.drawio.svg

![DRAW](https://github.com/Vasily-van-Zaam/GophKeeper/blob/server/GophKeeper.drawio.svg?raw=true)

## Для подтверждения почты пароль статичный - 1234

GOOS=windows GOARCH=386 go build -o client-win-386.exe ./cmd/client  
GOOS=windows GOARCH=amd64 go build -o client-win-amd64.exe ./cmd/client
GOOS=darwin GOARCH=amd64 go build -o client-macos-amd64 ./cmd/client
GOOS=linux GOARCH=amd64 go build -o client-linux-amd64 ./cmd/client

GOOS=windows GOARCH=amd64 go build -o server-win-amd64.exe ./cmd/server
GOOS=darwin GOARCH=amd64 go build -o server-macos-amd64 ./cmd/server
GOOS=linux GOARCH=amd64 go build -o server-linux-amd64 ./cmd/server
