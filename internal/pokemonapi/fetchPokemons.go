package pokemonapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"pokedexcli/internal/pokecache"
)

func FetchPokemons(locationAreaName string, pokemoncache *pokecache.Cache) (PokemonEncounters, error) {
	url := pokeapi_url + "location-area/" + locationAreaName
	cachedRes, ok := pokecache.Get(pokemoncache, url)
	pokemonEncounters := PokemonEncounters{}

	if ok {
		json.Unmarshal(cachedRes, &pokemonEncounters)
		return pokemonEncounters, nil
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	pokecache.Add(pokemoncache, url, body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return pokemonEncounters, fmt.Errorf("sesponse failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		return pokemonEncounters, fmt.Errorf("error found: %v", err)
	}
	json.Unmarshal(body, &pokemonEncounters)

	return pokemonEncounters, nil

}
