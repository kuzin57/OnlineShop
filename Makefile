PATH_TO_FILE := "./.database"

build:
	docker-compose build bshop

run:
	docker-compose up bshop

clean:
	docker rm $(shell sudo docker ps -aqf "name=onlineshop-bshop-1")
	if [ -d $(PATH_TO_FILE) ]; then \
		sudo rm -r .database; \
	fi

test:
	go test -v ./...