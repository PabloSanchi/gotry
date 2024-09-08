package gotry

import (
	"context"
	"errors"
	"math/rand"
	"net/http"
	"time"
)

// RetryableFuncWithResponse represents a function that returns an HTTP response or an error.
type RetryableFuncWithResponse func() (*http.Response, error)

// Retry retries the provided retryableFunc according to the retry configuration options.
func Retry(retryableFunc RetryableFuncWithResponse, options ...RetryOption) (*http.Response, error) {
	opts := newDefaultRetryConfig()

	for _, opt := range options {
		if opt != nil {
			opt(opts)
		}
	}

	var lastErr error
	for n := uint(0); n < opts.retries; n++ {
		if err := opts.context.Err(); err != nil {
			return nil, err
		}

		resp, err := retryableFunc()
		if err == nil && resp != nil && resp.StatusCode == http.StatusOK {
			return resp, nil
		}

		if err == nil && resp != nil {
			err = errors.New(resp.Status)
		}

		if !opts.retryIf(err) {
			return nil, err
		}

		lastErr = err

		backoffDuration := getBackoffDuration(opts, n+1)

		select {
		case <-time.After(backoffDuration):
		case <-opts.context.Done():
			return nil, opts.context.Err()
		}

		opts.onRetry(n+1, err)
	}

	return nil, lastErr
}

// getBackoffDuration calculates the backoff duration based on the retry configuration and attempt number.
func getBackoffDuration(config *RetryConfig, attempt uint) time.Duration {
	backoffDuration := config.backoffStrategy(config.backoff, attempt)

	if config.maxJitter > 0 {
		jitter := time.Duration(rand.Int63n(int64(config.maxJitter)))
		backoffDuration += jitter
	}

	if config.backoffLimit > 0 && backoffDuration > config.backoffLimit {
		backoffDuration = config.backoffLimit
	}

	return backoffDuration
}

// WithRetries sets the number of retries for the retry configuration.
func WithRetries(retries uint) RetryOption {
	return func(cfg *RetryConfig) {
		cfg.retries = retries
	}
}

// WithBackoff sets the backoff duration between retries.
func WithBackoff(backoff time.Duration) RetryOption {
	return func(cfg *RetryConfig) {
		cfg.backoff = backoff
	}
}

// WithBackoffLimit sets the maximum backoff duration allowed for the retry configuration.
func WithBackoffLimit(backoffLimit time.Duration) RetryOption {
	return func(cfg *RetryConfig) {
		cfg.backoffLimit = backoffLimit
	}
}

// WithMaxJitter sets the maximum jitter duration to add to the backoff.
func WithMaxJitter(maxJitter time.Duration) RetryOption {
	return func(cfg *RetryConfig) {
		cfg.maxJitter = maxJitter
	}
}

// WithOnRetry sets the callback function to execute on each retry.
func WithOnRetry(onRetry OnRetryFunc) RetryOption {
	return func(cfg *RetryConfig) {
		cfg.onRetry = onRetry
	}
}

// WithRetryIf sets the condition to determine whether to retry based on the error.
func WithRetryIf(retryIf RetryIfFunc) RetryOption {
	return func(cfg *RetryConfig) {
		cfg.retryIf = retryIf
	}
}

// WithContext sets the context for the retry configuration.
func WithContext(ctx context.Context) RetryOption {
	return func(cfg *RetryConfig) {
		cfg.context = ctx
	}
}

// WithLinearBackoff sets the backoff duration between retries to a linear duration.
func WithLinearBackoff() RetryOption {
	return func(cfg *RetryConfig) {
		cfg.backoffStrategy = func(base time.Duration, n uint) time.Duration {
			return base * time.Duration(n)
		}
	}
}

// WithExponentialBackoff sets the backoff duration between retries to an exponential duration.
func WithExponentialBackoff() RetryOption {
	return func(cfg *RetryConfig) {
		cfg.backoffStrategy = func(base time.Duration, n uint) time.Duration {
			return base * time.Duration(1<<n)
		}
	}
}

// WithCustomBackoff sets a custom backoff strategy for the retry configuration.
func WithCustomBackoff(strategy func(base time.Duration, n uint) time.Duration) RetryOption {
	return func(cfg *RetryConfig) {
		cfg.backoffStrategy = strategy
	}
}
