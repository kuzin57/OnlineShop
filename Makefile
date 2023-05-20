build:
	sudo rm -r .database
	docker-compose build bshop

run:
	docker-compose up bshop

test:
	go test -v ./...