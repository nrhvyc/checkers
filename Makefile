build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm cmd/server/server.go;
	go build cmd/server/server.go;

run:
	make build;
	./server;
