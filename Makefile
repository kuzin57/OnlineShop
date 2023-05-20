build:
	docker-compose build bshop

run:
	docker-compose up bshop

test:
	go test -v ./...