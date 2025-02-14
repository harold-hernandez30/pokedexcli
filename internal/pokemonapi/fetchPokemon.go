package pokemonapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"pokedexcli/internal/pokecache"
)

func FetchPokemon(pokemonName string, pokemonCache *pokecache.Cache) (Pokemon, error) {
	url := pokeapi_url + "pokemon/" + pokemonName
	cachedRes, ok := pokecache.Get(pokemonCache, url)
	pokemon := Pokemon{}

	if ok {
		json.Unmarshal(cachedRes, &pokemon)
		return pokemon, nil
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	pokecache.Add(pokemonCache, url, body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return pokemon, fmt.Errorf("sesponse failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		return pokemon, fmt.Errorf("error found: %v", err)
	}
	json.Unmarshal(body, &pokemon)

	return pokemon, nil
}
