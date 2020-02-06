CC=go
SRC_FILES=hello.go

all:
	GOOS=js GOARCH=wasm go build -o hello.wasm

clean:
	-rm hello.wasm