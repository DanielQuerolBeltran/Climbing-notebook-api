dev:
	docker-compose up

build:
	docker build -t go-docker-test-image .

rebuild:
	docker-compose up --build

restart:
	docker-compose down && docker-compose up -d	