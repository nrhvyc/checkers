build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm cmd/server/routes.go;
	go build -o server cmd/server/routes.go;

run:
	make build;
	./server;
