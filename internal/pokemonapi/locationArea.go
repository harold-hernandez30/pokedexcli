package pokemonapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)


func FetchLocationArea(url string) (PokeAreas, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return PokeAreas{}, fmt.Errorf("sesponse failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		return PokeAreas{}, fmt.Errorf("error found: %v", err)
	}
	pokeAreas := PokeAreas{}
	json.Unmarshal(body, &pokeAreas)
	
	return pokeAreas, nil
}

