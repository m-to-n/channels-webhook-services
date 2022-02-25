set GOARCH=amd64
set GOOS=linux
go build -ldflags="-s -w" -o bin/_version lambdas/_version/main.go
go build -ldflags="-s -w" -o bin/whatsapp-twilio lambdas/whatsapp-twilio/main.go
