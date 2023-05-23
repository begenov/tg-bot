.PHONY:

build: 
	go build -o ./.bin/telegram cmd/main.go
run: build
	./.bin/telegram

Image: 
	docker pull keinos/sqlite3
	
Container:
	docker run --rm --name=tg-db -it  keinos/sqlite3
Migrate:
	migrate create -ext sql -dir ./schema -seq init

Excec:
	docker exec -it tg-db bash
# psql -U postgres