PROJECT=github-fetch
OUTPUT=bin/$(PROJECT)

.PHONY: build
build:
	go mod vendor
	CGO_ENABLED=0 go build -mod=vendor -trimpath -o $(OUTPUT)

clean:
	go clean
	rm -f $(OUTPUT)
	rm -rf vendor/

run: build
	$(OUTPUT)

build-linux:
	go mod vendor
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod=vendor -trimpath -o $(OUTPUT)

build-windows:
	go mod vendor
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -mod=vendor -trimpath -o $(OUTPUT).exe
	upx -6 $(OUTPUT).exe
