package echologger

import (
	"io"

	"github.com/fatih/color"
)

type EchoLoggerConfig struct {
	Format string `yaml:"format"`

	TimeFormat string `yaml:"time_format"`

	Colors bool `yaml:"colors"`

	Output io.Writer

	colorer *color.Color
}

var (
	DefaultEchoLoggerConfig = EchoLoggerConfig{
		Format:     "method=${method}, uri=${uri}, status=${status}\n",
		TimeFormat: "dd:MM:yyyy hh:mm:ss",
		Colors:     true,
		Output:     color.Output,
		colorer:    color.New(),
	}
)
