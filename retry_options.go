package gotry

import (
	"context"
	"time"
)

// OnRetryFunc is a function type that is called on each retry attempt.
type OnRetryFunc func(attempt uint, err error)

// RetryIfFunc determines whether a retry should be attempted based on the error.
type RetryIfFunc func(error) bool

// RetryConfig contains configuration options for the retry mechanism.
type RetryConfig struct {
	retries         uint
	backoff         time.Duration
	backoffStrategy func(base time.Duration, n uint) time.Duration
	backoffLimit    time.Duration // maximum backoff time allowed
	maxJitter       time.Duration
	onRetry         OnRetryFunc
	retryIf         RetryIfFunc
	context         context.Context
}

// RetryOption is a function type for modifying RetryConfig options.
type RetryOption func(*RetryConfig)

// newDefaultRetryConfig creates a default RetryConfig with sensible defaults.
func newDefaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		retries:         3,
		backoff:         1 * time.Second,
		backoffStrategy: func(base time.Duration, n uint) time.Duration { return base },
		backoffLimit:    0,                                          // no limit by default
		maxJitter:       0,                                          // no jitter by default
		onRetry:         func(n uint, err error) {},                 // no-op onRetry by default
		retryIf:         func(err error) bool { return err != nil }, // retry on any error by default
		context:         context.Background(),
	}
}
