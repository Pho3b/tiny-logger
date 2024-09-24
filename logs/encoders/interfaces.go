package encoders

import "gitlab.com/docebo/libraries/go/tiny-logger/logs/configs"

type Encoder interface {
	LogDebug(conf *configs.TLConfigs, args ...interface{})
}
