package echologger

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
)

func New(config ...EchoLoggerConfig) echo.MiddlewareFunc {
	cfg := configDefault(config...)
	if strings.Contains(cfg.Format, "${latency}") {
		cfg.enableLatency = true
	}

	cfg.timeZoneLocation = time.Local
	cfg.tags = findTags(cfg.Format)

	var timestamp atomic.Value
	timestamp.Store(time.Now().Format(cfg.TimeFormat))

	// Update time every 250 milliseconds in a seperate goroutine
	if strings.Contains(cfg.Format, TagTime) {
		go func() {
			for {
				time.Sleep(250 * time.Millisecond)
				timestamp.Store(time.Now().Format(cfg.TimeFormat))
			}
		}()
	}

	// Get PID of current process
	pid := strconv.Itoa(os.Getpid())

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			var start, end time.Time

			if cfg.enableLatency {
				start = time.Now()
			}

			errChain := next(c)

			var (
				statusCode  int
				errorString string
			)

			if errChain != nil {
				fmt.Printf("errChain: %v\n", errChain)
				var err error
				parts := strings.Split(errChain.Error(), ", ")

				codeParts := strings.Split(strings.TrimSpace(parts[0]), "=")
				messageParts := strings.Split(strings.TrimSpace(parts[1]), "=")

				statusCode, err = strconv.Atoi(codeParts[1])
				if err != nil {
					statusCode = 500
				}
				errorString = messageParts[1]
			} else {
				statusCode = c.Response().Status
			}

			if cfg.enableLatency {
				end = time.Now()
			}

			tagsToReplace := map[string]string{}

			for _, tag := range cfg.tags {
				switch tag {
				case TagPid:
					tagsToReplace[tag] = pid
				case TagLatency:
					tagsToReplace[tag] = end.Sub(start).String()
				case TagTime:
					tagsToReplace[tag] = timestamp.Load().(string)
				case TagStatus:
					tagsToReplace[tag] = strconv.Itoa(statusCode)
				case TagMethod:
					tagsToReplace[tag] = c.Request().Method
				case TagPath:
					tagsToReplace[tag] = c.Request().URL.Path
				case TagHost:
					tagsToReplace[tag] = c.RealIP()
				case TagError:
					tagsToReplace[tag] = color.RedString(errorString)
				}
			}

			log := formatLog(cfg.Format, tagsToReplace, errorString, !cfg.DisablePadding, !cfg.DisableColors)

			cfg.output.Printf(log)

			return nil
		}

	}
}
