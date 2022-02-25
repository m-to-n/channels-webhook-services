set GOARCH=amd64
set GOOS=linux
go build -ldflags="-s -w" -o bin/hello src/hello/main.go
go build -ldflags="-s -w" -o bin/world src/world/main.go
go build -ldflags="-s -w" -o bin/whatsapp-twilio src/whatsapp-twilio/main.go
