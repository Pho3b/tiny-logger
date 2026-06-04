.PHONY: test test-coverage test-benchmark encoders-benchmark loggers-comparison-benchmark

GO_TOOLCHAIN=go1.25.0
ARCH=$(shell go env GOARCH)
GO_TEST_SETTINGS=GOTOOLCHAIN=${GO_TOOLCHAIN}+auto CGO_ENABLED=1 GOARCH=${ARCH}

# Run all tests with coverage profile
test:
	${GO_TEST_SETTINGS} go test -cover ./... -race

# Run all tests and Show HTML coverage report in the browser
test-coverage:
	${GO_TEST_SETTINGS} go test -cover -coverprofile=coverage.out ./... -race && go tool cover -html=coverage.out

# Run encoders benchmark tests
encoders-benchmark:
	${GO_TEST_SETTINGS} go test ./test/encoders_benchmark_test.go ./test/functions.go -bench=. -benchmem -benchtime=5s -cpu=8

# Run commercial loggers benchmark comparison tests.
# These live in their own module (test/comparison) so that zerolog/zap/logrus never become
# dependencies of the main, zero-dependency tiny-logger module.
loggers-comparison-benchmark:
	cd test/comparison && CGO_ENABLED=1 GOARCH=${ARCH} go test ./... -bench=. -benchmem -benchtime=5s -cpu=8