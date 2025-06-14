.PHONY: test coverage clean

# Run all tests with verbose output and coverage profile
test:
	CGO_ENABLED=1 GOARCH=amd64 go test -cover ./... -race

# Run all tests and Show HTML coverage report in the browser
test-coverage: test
	CGO_ENABLED=1 GOARCH=amd64 go test -v -cover -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
