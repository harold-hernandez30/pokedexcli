package pokemonapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"pokedexcli/internal/pokecache"
)

func FetchLocationArea(url string, locationCache *pokecache.Cache) (PokeAreas, error) {

	cachedRes, ok := pokecache.Get(locationCache, url)
	pokeAreas := PokeAreas{}

	if ok {
		json.Unmarshal(cachedRes, &pokeAreas)
		return pokeAreas, nil
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	pokecache.Add(locationCache, url, body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return pokeAreas, fmt.Errorf("sesponse failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		return pokeAreas, fmt.Errorf("error found: %v", err)
	}
	json.Unmarshal(body, &pokeAreas)

	return pokeAreas, nil
}
