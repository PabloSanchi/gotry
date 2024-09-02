package examples

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	retry "github.com/PabloSanchi/gotry"
)

const (
	url = "https://official-joke-api.appspot.com/random_joke"
)

type Joke struct {
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
	Id        uint   `json:"id"`
}

func TestFetchJoke() {
	resp, err := retry.Retry(
		func() (*http.Response, error) {
			return http.Get(url)
		},
		retry.WithRetries(2),
		retry.WithBackoff(2*time.Second),
		retry.WithOnRetry(func(n uint, err error) {
			log.Printf("Retrying request after error: %v", err)
		}),
	)

	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	var joke Joke

	if err := json.NewDecoder(resp.Body).Decode(&joke); err != nil {
		log.Fatalf("Failed to decode response: %v", err)
	}

	log.Printf("Joke: %s", joke.Setup)
	log.Printf("Punchline: %s", joke.Punchline)
}
