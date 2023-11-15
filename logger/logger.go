package echologger

import (
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
			// next(c)

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
					tagsToReplace[tag] = strconv.Itoa(c.Response().Status)
				case TagMethod:
					tagsToReplace[tag] = c.Request().Method
				case TagPath:
					tagsToReplace[tag] = c.Request().URL.Path
				case TagHost:
					tagsToReplace[tag] = c.RealIP()
				case TagError:
					if errChain != nil {
						tagsToReplace[tag] = color.RedString(errChain.Error())
					} else {
						tagsToReplace[tag] = ""
					}
				}
			}

			log := formatLog(cfg.Format, tagsToReplace, !cfg.DisablePadding, !cfg.DisableColors)

			cfg.output.Printf(log)

			return nil
		}

	}
}
