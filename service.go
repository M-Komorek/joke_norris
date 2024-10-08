package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/M-Komorek/joke_norris/dto"
)

type Service interface {
	GetChuckNorrisJoke(context.Context) (*dto.ChuckNorrisJoke, error)
}

type ChuckNorrisJokeService struct {
	url string
}

func NewChuckNorrisJokeService(url string) Service {
	return &ChuckNorrisJokeService{
		url: url,
	}
}

func (chuckNorrisJokeService *ChuckNorrisJokeService) GetChuckNorrisJoke(ctx context.Context) (*dto.ChuckNorrisJoke, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", chuckNorrisJokeService.url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received status code %d", resp.StatusCode)
	}

	chuckNorrisJoke := &dto.ChuckNorrisJoke{}
	if err := json.NewDecoder(resp.Body).Decode(chuckNorrisJoke); err != nil {
		return nil, fmt.Errorf("error decoding response body %d", resp.StatusCode)
	}

	return chuckNorrisJoke, nil
}
