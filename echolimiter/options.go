package echolimiter

import "time"

type OptFunction func(*RateLimitOptions)

type RateLimitOptions struct {
	MaxRequests int
	Window      time.Duration
}

func defaultOptions() *RateLimitOptions {
	return &RateLimitOptions{
		MaxRequests: 10,
		Window:      60 * time.Second,
	}
}

func WithMaxRequests(max int) OptFunction {
	return func(o *RateLimitOptions) {
		o.MaxRequests = max
	}
}

func WithWindow(window time.Duration) OptFunction {
	return func(o *RateLimitOptions) {
		o.Window = window
	}
}
