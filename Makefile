PROJECT=github-fetch
OUTPUT=bin/$(PROJECT)
VERSION=0.0.2

.PHONY: build
build:
	go mod vendor
	CGO_ENABLED=0 go build -ldflags="-s" -mod=vendor -trimpath -o $(OUTPUT)-$(VERSION)-darwin

clean:
	go clean
	rm -f $(OUTPUT)
	rm -rf vendor/

run: build
	$(OUTPUT)

build-linux:
	go mod vendor
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s" -mod=vendor -trimpath -o $(OUTPUT)-$(VERSION)-linux-amd64

build-windows:
	go mod vendor
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s" -mod=vendor -trimpath -o $(OUTPUT)-$(VERSION).exe
	upx -6 $(OUTPUT).exe
