build-local:
	mkdir -p dist
	go build -o ./dist/server .

build-run:
	./dist/server

go-run:
	go run main.go
