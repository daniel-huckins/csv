
PWD=$$(pwd)

client:
	@echo building client
	@cd $(PWD)/app/ && \
	npm install && \
	npm run build

run:
	go run main.go client.go template.go
