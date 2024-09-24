package encoders

import (
	"fmt"
	"gitlab.com/docebo/libraries/go/tiny-logger/internal/services"
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/configs"
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/log_level"
	"os"
)

type DefaultEncoder struct {
	BaseEncoder
	servicesWrapper services.Wrapper
}

func (d *DefaultEncoder) LogDebug(conf *configs.TLConfigs, args ...interface{}) {
	if len(args) > 0 {
		dateTime := d.servicesWrapper.DateTimePrinter.PrintDateTime(conf)
		colors := d.servicesWrapper.ColorsPrinter.PrintColors(conf, log_level.DebugLvlName)

		_, _ = fmt.Fprintln(
			os.Stdout,
			fmt.Sprintf("%vDEBUG%s:%v %s", colors[0], dateTime, colors[1], d.buildMsg(args...)),
		)
	}
}

func NewDefaultEncoder(servicesWrapper services.Wrapper) *DefaultEncoder {
	return &DefaultEncoder{
		servicesWrapper: servicesWrapper,
	}
}
