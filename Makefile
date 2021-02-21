run:
	@docker-compose run search

test:
	@docker-compose run tests

update:
	@docker-compose run update

run_binary:
	@docker-compose run build_binary
