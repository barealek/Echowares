package echolimiter

import (
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type visitor struct {
	*sync.RWMutex
	lastAccess time.Time
	requests   int
}

type RateLimiter func(next echo.HandlerFunc) echo.HandlerFunc

func NewRateLimiter(opts ...OptFunction) func(next echo.HandlerFunc) echo.HandlerFunc {
	options := defaultOptions()
	for _, opt := range opts {
		opt(options)
	}

	var visitors = make(map[string]*visitor)
	var mtx sync.Mutex

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ip := c.RealIP()
			mtx.Lock()
			v, exists := visitors[ip]
			if !exists {
				v = &visitor{
					RWMutex:    &sync.RWMutex{},
					lastAccess: time.Now(),
					requests:   1,
				}
				visitors[ip] = v
			}
			mtx.Unlock()

			v.Lock()
			if time.Since(v.lastAccess) > options.Window {
				v.requests = 0
				v.lastAccess = time.Now()
			}
			v.requests++
			v.Unlock()

			if v.requests > options.MaxRequests {
				return echo.ErrTooManyRequests
			}

			return next(c)
		}
	}

}
