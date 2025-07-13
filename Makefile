build-linux:
	GOOS=linux GOARCH=amd64 go build -o f2f-linux

build-windows:
	GOOS=windows GOARCH=amd64 go build -o f2f.exe

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o f2f-mac