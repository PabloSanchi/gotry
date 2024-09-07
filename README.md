<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/PabloSanchi/gotry" alt="Go Version">
<a href="https://pkg.go.dev/github.com/PabloSanchi/gotry"><img src="https://pkg.go.dev/badge/github.com/PabloSanchi/gotry" alt="PkgGoDev"></a>
<a href="https://opensource.org/license/apache-2-0"><img src="https://img.shields.io/badge/License-Apache_2.0-blue.svg" alt="License"></a>
</p>

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

    resp, err := gotry.Retry(retryableFunc,
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

    // we assume there is a body from this request;
    defer resp.Body.Close()
    body := resp.Body
    
    // do something with the body
    // ...
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
