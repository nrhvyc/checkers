build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm cmd/app/*;
	go build cmd/server/server.go;

run:
	make build;
	./server;
