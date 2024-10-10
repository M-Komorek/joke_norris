package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/M-Komorek/joke_norris/dto"
)

type ChuckNorrisJokeClient struct {
	baseUrl string
}

func NewChuckNorrisJokeClient(baseUrl string) *ChuckNorrisJokeClient {
	return &ChuckNorrisJokeClient{
		baseUrl: baseUrl,
	}
}

func (chuckNorrisJokeClient *ChuckNorrisJokeClient) GetChuckNorrisJoke(ctx context.Context) (*dto.ChuckNorrisJoke, error) {
	baseUrl := fmt.Sprintf("%s/chuckNorrisJoke", chuckNorrisJokeClient.baseUrl)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dhuckNorrisJoke := &dto.ChuckNorrisJoke{}
	if err := json.NewDecoder(resp.Body).Decode(dhuckNorrisJoke); err != nil {
		return nil, err
	}

	return dhuckNorrisJoke, nil
}
