package echologger

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/fatih/color"
)

type EchoLoggerConfig struct {
	Format     string `yaml:"format"`
	TimeFormat string `yaml:"time_format"`

	DisableColors bool `yaml:"colors"`

	Output  io.Writer
	colorer *color.Color

	TimeZone         string `yaml:"time_zone"`
	timeZoneLocation *time.Location

	enableLatency bool
}

var (
	DefaultEchoLoggerConfig = EchoLoggerConfig{
		Format:        fmt.Sprintf("[%v] %v - %v %v %v\n", TagTime, TagStatus, TagLatency, TagMethod, TagPath),
		TimeFormat:    "15:04:05",
		TimeZone:      "Local",
		DisableColors: true,
		Output:        color.Output,
		colorer:       color.New(),
	}
)

func configDefault(config ...EchoLoggerConfig) EchoLoggerConfig {
	if len(config) < 1 {
		return DefaultEchoLoggerConfig
	}

	cfg := config[0]

	if cfg.Format == "" {
		cfg.Format = DefaultEchoLoggerConfig.Format
	}

	if cfg.TimeFormat == "" {
		cfg.TimeFormat = DefaultEchoLoggerConfig.TimeFormat
	}

	if cfg.Output == nil {
		cfg.Output = DefaultEchoLoggerConfig.Output
	}

	if cfg.colorer == nil {
		cfg.colorer = DefaultEchoLoggerConfig.colorer
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

	return cfg
}
