run:
	@docker-compose run search

test:
	@docker-compose run tests

build:
	@go build -v -o search

run_binary:
	@./search

update:
	@go mod tidy