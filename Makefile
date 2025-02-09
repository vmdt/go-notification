.PHONY:

## choco install make
# ========
# Run service

run:
	@go run ./cmd/main.go

build:
	@go build -o ./bin/app ./cmd/main.go

run_build:
	@./bin/app