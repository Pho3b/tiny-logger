module github.com/Pho3b/tiny-logger/test/comparison

go 1.25.0

require (
	github.com/Pho3b/tiny-logger v0.0.0
	github.com/rs/zerolog v1.35.1
	github.com/sirupsen/logrus v1.9.4
	go.uber.org/zap v1.28.0
)

require (
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.22 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/sys v0.43.0 // indirect
)

replace github.com/Pho3b/tiny-logger => ../../
