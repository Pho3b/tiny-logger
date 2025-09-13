.PHONY: test test-coverage test-benchmark

ARCH=amd64

# Run all tests with coverage profile
test:
	CGO_ENABLED=1 GOARCH=${ARCH} go test -cover ./... -race

# Run all tests and Show HTML coverage report in the browser
test-coverage:
	CGO_ENABLED=1 GOARCH=${ARCH} go test -cover -coverprofile=coverage.out ./... -race && go tool cover -html=coverage.out

# Run benchmark tests
test-benchmark:
	CGO_ENABLED=1 GOARCH=${ARCH} go test ./test/benchmark_test.go -bench=. -benchmem -benchtime=5s -cpu=8 | grep /op
