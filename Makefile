.PHONY: build clean deploy

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/_version src/_version/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/whatsapp-twilio src/whatsapp-twilio/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
