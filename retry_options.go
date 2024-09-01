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
	retries   uint
	backoff   time.Duration
	maxJitter time.Duration
	onRetry   OnRetryFunc
	retryIf   RetryIfFunc
	timer     Timer
	context   context.Context
}

// RetryOption is a function type for modifying RetryConfig options.
type RetryOption func(*RetryConfig)

// newDefaultRetryConfig creates a default RetryConfig with sensible defaults.
func newDefaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		retries:   3,
		backoff:   1 * time.Second,
		maxJitter: 0 * time.Second,                            // no jitter by default
		onRetry:   func(n uint, err error) {},                 // no-op onRetry by default
		retryIf:   func(err error) bool { return err != nil }, // retry on any error by default
		timer:     &timerImpl{},
		context:   context.Background(),
	}
}
