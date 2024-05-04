package pokeapi

import (
	"net/http"
	"time"

	"github.com/pcoelho00/pokedex/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(checkInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(checkInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

const baseURL = "https://pokeapi.co/api/v2"
