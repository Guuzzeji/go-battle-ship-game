build-local:
	mkdir -p dist
	go build -o ./dist/main .

build-run:
	./dist/main

run-go-server:
	go run server.go
