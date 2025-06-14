.PHONY: test coverage clean

# Run all tests with coverage profile
test:
	CGO_ENABLED=1 GOARCH=amd64 go test -cover ./... -race

# Run all tests and Show HTML coverage report in the browser
test-coverage:
	CGO_ENABLED=1 GOARCH=amd64 go test -cover -coverprofile=coverage.out ./... -race && go tool cover -html=coverage.out
