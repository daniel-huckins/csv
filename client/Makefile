
PWD=$$(pwd)

client:
	@echo building client
	@cd $(PWD)/app/ && \
	npm run build

run: client
	go run ../cmd/run/main.go
