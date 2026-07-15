.PHONY: lint vet test all

# Runs the spell-checker
lint:
	golangci-lint run ./...

# Checks for suspicious code
vet:
	go vet ./...

# Runs your tests with the race detector
test:
	go test -race ./...

# A shortcut to run all three at once
all: vet lint test