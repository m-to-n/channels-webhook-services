set GOARCH=amd64
set GOOS=linux
go build -ldflags="-s -w" -o bin/_version src/_version/main.go
go build -ldflags="-s -w" -o bin/whatsapp-twilio src/whatsapp-twilio/main.go
