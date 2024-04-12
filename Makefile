build:
	go build -o ./bin/goquote

run: build
	bin/goquote