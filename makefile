.PHONY: test test-coverage test-benchmark

# Run all tests with coverage profile
test:
	CGO_ENABLED=1 GOARCH=amd64 go test -cover ./... -race

# Run all tests and Show HTML coverage report in the browser
test-coverage:
	CGO_ENABLED=1 GOARCH=amd64 go test -cover -coverprofile=coverage.out ./... -race && go tool cover -html=coverage.out

# Run benchmark tests
test-benchmark:
	CGO_ENABLED=1 GOARCH=amd64 go test ./test/benchmark_test.go -bench=. -benchmem -benchtime=4s -cpu=2 | grep /op