package echologger

import (
	"io"

	"github.com/fatih/color"
)

type EchoLoggerConfig struct {
	Format string `yaml:"format"`

	TimeFormat string `yaml:"time_format"`

	DisableColors bool `yaml:"colors"`

	Output io.Writer

	colorer *color.Color
}

var (
	DefaultEchoLoggerConfig = EchoLoggerConfig{
		Format:        "method=${method}, uri=${uri}, status=${status}\n",
		TimeFormat:    "dd:MM:yyyy hh:mm:ss",
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

	return cfg
}
