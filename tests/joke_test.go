package dto

import (
	"encoding/json"
	"testing"

	"github.com/M-Komorek/joke_norris/dto"
)

func TestChuckNorrisJoke_JSONSerialization(t *testing.T) {
	// Create a sample joke
	joke := dto.ChuckNorrisJoke{
		Id:   "1",
		Text: "Chuck Norris doesn't do push-ups; he pushes the Earth down.",
	}

	// Serialize to JSON
	jokeAsJson, err := json.Marshal(joke)
	if err != nil {
		t.Fatalf("failed to marshal joke to JSON: %v", err)
	}

	// Expected JSON output
	expectedJSON := `{"id":"1","value":"Chuck Norris doesn't do push-ups; he pushes the Earth down."}`
	if string(jokeAsJson) != expectedJSON {
		t.Errorf("expected %s, got %s", expectedJSON, string(jokeAsJson))
	}

	// Deserialize back from JSON
	var decodedJoke dto.ChuckNorrisJoke
	if err := json.Unmarshal(jokeAsJson, &decodedJoke); err != nil {
		t.Fatalf("failed to unmarshal JSON to joke: %v", err)
	}

	// Verify the decoded joke matches the original
	if decodedJoke.Id != joke.Id || decodedJoke.Text != joke.Text {
		t.Errorf("expected %+v, got %+v", joke, decodedJoke)
	}
}
