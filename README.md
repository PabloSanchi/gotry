# GoTry - Simple Retry Module for Go

GoTry is a lightweight, configurable retry module for Go, designed to make retrying operations like HTTP requests easier. It allows developers to specify custom retry logic, backoff strategies, and error handling in a simple and flexible way.

## Features
- Configurable Retry Logic: Set the number of retries, backoff duration, and maximum jitter.
- Custom Retry Conditions: Define conditions under which a retry should be attempted.
- Retry Callbacks: Execute custom code on each retry attempt.

## Installation
To install GoTry, use go get:

```sh
go get github.com/PabloSanchi/gotry
```

## Usage
```go
package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/PabloSanchi/gotry"
)

func main() {
    retryableFunc := func() (*http.Response, error) {
        // replace with your actual HTTP request logic
        return http.Get("https://example.com")
    }

    responseBody, err := gotry.Retry(retryableFunc,
        gotry.WithRetries(5),
        gotry.WithBackoff(2*time.Second),
        gotry.WithMaxJitter(500*time.Millisecond),
        gotry.WithOnRetry(func(attempt uint, err error) {
            fmt.Printf("Retry attempt %d: %v\n", attempt, err)
        }),
    )

    if err != nil {
        fmt.Printf("Failed to get response: %v\n", err)
        return
    }

    fmt.Printf("Response: %s\n", string(responseBody))
}
```

## Configuration Options
- `WithRetries(uint)`: Set the number of retry attempts.
- `WithBackoff(time.Duration)`: Set the backoff duration between retries.
- `WithMaxJitter(time.Duration)`: Add jitter to the backoff duration.
- `WithOnRetry(OnRetryFunc)`: Define a callback function to be executed on each retry.
- `WithRetryIf(RetryIfFunc)`: Specify the condition under which retries should be attempted.

## License
This project is licensed under the Apache License. See the [LICENSE](LICENSE) file for details.

## 

Simplification of the Avast [retry-go](https://github.com/avast/retry-go) module
