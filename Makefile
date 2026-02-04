.PHONY: test test-coverage test-benchmark encoders-benchmark loggers-comparison-benchmark

ARCH=amd64

# Run all tests with coverage profile
test:
	CGO_ENABLED=1 GOARCH=${ARCH} go test -cover ./... -race

# Run all tests and Show HTML coverage report in the browser
test-coverage:
	CGO_ENABLED=1 GOARCH=${ARCH} go test -cover -coverprofile=coverage.out ./... -race && go tool cover -html=coverage.out

# Run encoders benchmark tests
encoders-benchmark:
	CGO_ENABLED=1 GOARCH=${ARCH} go test ./test/encoders_benchmark_test.go ./test/functions.go -bench=. -benchmem -benchtime=5s -cpu=8

# Run commercial loggers benchmark comparison tests
loggers-comparison-benchmark:
	CGO_ENABLED=1 GOARCH=${ARCH} go test ./test/loggers_comparison_test.go ./test/functions.go -bench=. -benchmem -benchtime=5s -cpu=8