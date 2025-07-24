BINARY_NAME=f2f

all: build-linux build-windows

build-linux:
	GOOS=linux GOARCH=amd64 go build -o dist/$(BINARY_NAME)-linux

build-windows:
	GOOS=windows GOARCH=amd64 go build -o dist/$(BINARY_NAME)-windows.exe

clean:
	rm -rf dist