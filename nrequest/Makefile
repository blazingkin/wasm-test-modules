CC=go
SRC_FILES=nrequest.go

all:
	GOOS=js GOARCH=wasm go build -o nrequest.wasm

clean:
	-rm nrequest.wasm