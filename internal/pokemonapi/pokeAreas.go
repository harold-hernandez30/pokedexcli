package pokemonapi

type PokeAreas struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	LocationAreas []PokeArea `json:"results"`
}