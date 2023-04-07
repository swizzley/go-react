PROFILE ?= Local

dev:
	@cd app && npm run dev

run:
	@cd app && npm run build
	@PROFILE=$(PROFILE) go run *.go

install:
	@cd app && npm install
	@go mod init
	@go mod tidy
	@go mod vendor
