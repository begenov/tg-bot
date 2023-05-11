.PHONY:

build: 
	go build -o ./.bin/telegram cmd/main.go
run: build
	./.bin/telegram