BINARY_NAME = gatus-cli
VERSION = $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS = -ldflags="-s -w -X main.version=$(VERSION)"

.PHONY: all
all: clean build

.PHONY: build
build:
	mkdir -p bin
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o bin/$(BINARY_NAME)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o bin/$(BINARY_NAME)-darwin-arm64 .
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/$(BINARY_NAME)-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o bin/$(BINARY_NAME)-linux-arm64 .
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/$(BINARY_NAME)-windows-amd64.exe .

.PHONY: clean
clean:
	rm -rf bin/

.PHONY: compress
build-and-compress: build
	cd bin && for file in *; do zip "$${file}.zip" "$$file"; done

.PHONY: checksums
checksums: build
	cd bin && sha256sum * > checksums.txt