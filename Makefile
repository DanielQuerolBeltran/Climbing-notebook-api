dev:
	docker-compose up

build-api:
	docker-compose build api

build:
	docker build -t go-docker-test-image .

rebuild:
	docker-compose build --no-cache

restart:
	docker-compose down && docker-compose up -d	