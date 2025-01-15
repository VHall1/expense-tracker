gen:
	@./bin/protoc

compose:
	@docker-compose -f ./docker-compose.yml up -d
