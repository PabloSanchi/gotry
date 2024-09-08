package gotry

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestRetryWithSuccessOnFirstRequest tests the retry mechanism with a server that always returns a 200 status code.
func TestRetryWithSuccessOnFirstRequest(t *testing.T) {
	type ResponseType struct {
		Content string `json:"content"`
		Status  int    `json:"status"`
	}

	expectedResponse := ResponseType{Content: "success"}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"content":"success"}`))
	}))
	defer server.Close()

	var retryCount uint = 0
	var expectedRetries uint = 0

	sendRequest := func() (*http.Response, error) {
		return http.Get(server.URL + "/success")
	}

	resp, err := Retry(
		sendRequest,
		WithOnRetry(func(n uint, err error) {
			retryCount = n
			log.Printf("Retrying request after error: %v", err)
		}),
	)

	if err != nil {
		t.Errorf("Expected nil, got error: %v", err)
	}

	if retryCount != expectedRetries {
		t.Errorf("Expected 0 retries, but got %d", retryCount)
	}

	defer resp.Body.Close()
	var response ResponseType
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}

	if response != expectedResponse {
		t.Errorf("Expected response %v, but got %v", expectedResponse, response)
	}
}

// TestRetryWithNoSuccessStatusOnAnyRequest tests the retry mechanism with a server that always returns a 429 status code.
func TestRetryWithNoSuccessStatusOnAnyRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer server.Close()

	var backoffTime = 2 * time.Second
	var retryCount uint = 0
	var expectedRetries uint = 3

	sendRequest := func() (*http.Response, error) {
		return http.Get(server.URL + "/notfound")
	}

	startTime := time.Now()

	resp, err := Retry(
		sendRequest,
		WithBackoff(backoffTime),
		WithOnRetry(func(n uint, err error) {
			retryCount = n
			log.Printf("Retrying request after error: %v", err)
		}),
	)

	endTime := time.Now()

	elapsedTime := endTime.Sub(startTime)
	expectedMinimumTime := backoffTime * time.Duration(expectedRetries)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if resp != nil {
		defer resp.Body.Close()
		t.Errorf("Expected nil response, got %v", resp.Body)
	}

	if retryCount != expectedRetries {
		t.Errorf("Expected 3 retries, but got %d", retryCount)
	}

	if elapsedTime < expectedMinimumTime {
		t.Errorf("Expected minimum time of %v, but got %v", expectedMinimumTime, elapsedTime)
	}

	// check if the difference between times is not too high (as is a controlled environment the difference should be minimal)
	if elapsedTime-expectedMinimumTime > time.Second {
		t.Errorf("Expected minimum time of %v, but got %v", expectedMinimumTime, elapsedTime)
	}
}

// TestRetryWithBackoffLimit tests the retry mechanism with a server that always returns a 429 status code and a backoff limit.
func TestRetryWithBackoffLimit(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer server.Close()

	var backoffTime = 2 * time.Second
	var retryCount uint = 0
	var expectedRetries uint = 3

	sendRequest := func() (*http.Response, error) {
		return http.Get(server.URL + "/notfound")
	}

	startTime := time.Now()

	resp, err := Retry(
		sendRequest,
		WithBackoff(backoffTime),
		WithBackoffLimit(1*time.Second),
		WithOnRetry(func(n uint, err error) {
			retryCount = n
			log.Printf("Retrying request after error: %v", err)
		}),
	)

	endTime := time.Now()

	elapsedTime := endTime.Sub(startTime)
	expectedMinimumTime := 3 * time.Second

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if resp != nil {
		defer resp.Body.Close()
		t.Errorf("Expected nil response, got %v", resp.Body)
	}

	if retryCount != expectedRetries {
		t.Errorf("Expected 3 retries, but got %d", retryCount)
	}

	if elapsedTime < expectedMinimumTime {
		t.Errorf("Expected minimum time of %v, but got %v", expectedMinimumTime, elapsedTime)
	}

	// check if the difference between times is not too high (as is a controlled environment the difference should be minimal)
	if elapsedTime-expectedMinimumTime > time.Second {
		t.Errorf("Expected minimum time of %v, but got %v", expectedMinimumTime, elapsedTime)
	}

}

// TestRetryWithWithLinearBackoff tests the retry mechanism with a server that always returns a 429 status code.
func TestRetryWithWithLinearBackoff(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer server.Close()

	var backoffTime = 2 * time.Second
	var retryCount uint = 0
	var expectedRetries uint = 3

	sendRequest := func() (*http.Response, error) {
		return http.Get(server.URL + "/notfound")
	}

	startTime := time.Now()

	resp, err := Retry(
		sendRequest,
		WithBackoff(backoffTime),
		WithLinearBackoff(),
		WithOnRetry(func(n uint, err error) {
			retryCount = n
			log.Printf("Retrying request after error: %v", err)
		}),
	)

	endTime := time.Now()

	elapsedTime := endTime.Sub(startTime)
	expectedMinimumTime := backoffTime + (backoffTime * 2) + (backoffTime * 3)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if resp != nil {
		defer resp.Body.Close()
		t.Errorf("Expected nil response, got %v", resp.Body)
	}

	if retryCount != expectedRetries {
		t.Errorf("Expected 3 retries, but got %d", retryCount)
	}

	if elapsedTime < expectedMinimumTime {
		t.Errorf("Expected minimum time of %v, but got %v", expectedMinimumTime, elapsedTime)
	}

	// check if the difference between times is not too high (as is a controlled environment the difference should be minimal)
	if elapsedTime-expectedMinimumTime > time.Second {
		t.Errorf("Expected minimum time of %v, but got %v", expectedMinimumTime, elapsedTime)
	}

}

// TestRetryWithWithExponentialBackoff tests the retry mechanism with a server that always returns a 429 status code.
func TestRetryWithWithExponentialBackoff(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer server.Close()

	var backoffTime = 2 * time.Second
	var retryCount uint = 0
	var expectedRetries uint = 3

	sendRequest := func() (*http.Response, error) {
		return http.Get(server.URL + "/notfound")
	}

	startTime := time.Now()

	resp, err := Retry(
		sendRequest,
		WithBackoff(backoffTime),
		WithExponentialBackoff(),
		WithOnRetry(func(n uint, err error) {
			retryCount = n
			log.Printf("Retrying request after error: %v", err)
		}),
	)

	endTime := time.Now()

	elapsedTime := endTime.Sub(startTime)
	expectedMinimumTime := (backoffTime * 2) + (backoffTime * 4) + (backoffTime * 8)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if resp != nil {
		defer resp.Body.Close()
		t.Errorf("Expected nil response, got %v", resp.Body)
	}

	if retryCount != expectedRetries {
		t.Errorf("Expected 3 retries, but got %d", retryCount)
	}

	if elapsedTime < expectedMinimumTime {
		t.Errorf("Expected minimum time of %v, but got %v", expectedMinimumTime, elapsedTime)
	}

	// check if the difference between times is not too high (as is a controlled environment the difference should be minimal)
	if elapsedTime-expectedMinimumTime > time.Second {
		t.Errorf("Expected minimum time of %v, but got %v", expectedMinimumTime, elapsedTime)
	}

}

// TestRetryWithWithCustomBackoff tests the retry mechanism with a server that always returns a 429 status code.
func TestRetryWithWithCustomBackoff(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer server.Close()

	var backoffTime = 2 * time.Second
	var retryCount uint = 0
	var expectedRetries uint = 3

	sendRequest := func() (*http.Response, error) {
		return http.Get(server.URL + "/notfound")
	}

	startTime := time.Now()

	resp, err := Retry(
		sendRequest,
		WithBackoff(backoffTime),
		// Dummy custom backoff function that always returns the same time as the backoffTime
		WithCustomBackoff(func(base time.Duration, n uint) time.Duration {
			return backoffTime
		}),
		WithOnRetry(func(n uint, err error) {
			retryCount = n
			log.Printf("Retrying request after error: %v", err)
		}),
	)

	endTime := time.Now()

	elapsedTime := endTime.Sub(startTime)
	expectedMinimumTime := backoffTime * 3

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if resp != nil {
		defer resp.Body.Close()
		t.Errorf("Expected nil response, got %v", resp.Body)
	}

	if retryCount != expectedRetries {
		t.Errorf("Expected 3 retries, but got %d", retryCount)
	}

	if elapsedTime < expectedMinimumTime {
		t.Errorf("Expected minimum time of %v, but got %v", expectedMinimumTime, elapsedTime)
	}

	// check if the difference between times is not too high (as is a controlled environment the difference should be minimal)
	if elapsedTime-expectedMinimumTime > time.Second {
		t.Errorf("Expected minimum time of %v, but got %v", expectedMinimumTime, elapsedTime)
	}
}
