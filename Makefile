.PHONY: run-dev audit lint

run-dev:
	go run cmd/app/main.go

audit:
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...

lint:
	golangci-lint run ./...
