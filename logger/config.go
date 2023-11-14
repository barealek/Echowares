package echologger

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

type loggerOutput interface {
	Printf(string, ...interface{}) (int, error)
}

type EchoLoggerConfig struct {
	Format     string `yaml:"format"`
	TimeFormat string `yaml:"time_format"`

	DisableColors bool `yaml:"colors"`

	output loggerOutput

	TimeZone         string `yaml:"time_zone"`
	timeZoneLocation *time.Location

	enableLatency bool
	tags          []string
}

var (
	DefaultEchoLoggerConfig = EchoLoggerConfig{

		Format:        fmt.Sprintf("%v | %v | %v | %v | %v | %v\n", TagTime, TagStatus, TagLatency, TagHost, TagMethod, TagPath),
		TimeFormat:    "15:04:05",
		TimeZone:      "Local",
		DisableColors: true,
		output:        color.New(),
	}
)

func configDefault(config ...EchoLoggerConfig) EchoLoggerConfig {
	if len(config) != 1 {
		return DefaultEchoLoggerConfig
	}

	cfg := config[0]

	if cfg.Format == "" {
		cfg.Format = DefaultEchoLoggerConfig.Format
	}

	if cfg.TimeFormat == "" {
		cfg.TimeFormat = DefaultEchoLoggerConfig.TimeFormat
	}

	if cfg.output == nil {
		cfg.output = DefaultEchoLoggerConfig.output
	}

	if strings.Contains(cfg.Format, "${latency}") {
		cfg.enableLatency = true
	}

	tz, err := time.LoadLocation(cfg.TimeZone)
	if err != nil || tz == nil {
		cfg.timeZoneLocation = time.Local
	} else {
		cfg.timeZoneLocation = tz
	}

	cfg.tags = findTags(cfg.Format)

	return cfg
}
