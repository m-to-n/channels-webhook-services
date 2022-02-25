.PHONY: build clean deploy

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/_version lambdas/_version/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/whatsapp-twilio lambdas/whatsapp-twilio/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
