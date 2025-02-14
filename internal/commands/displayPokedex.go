package commands

func displayPokedexCommand(cmConfig *Config) error {
	cmConfig.Pokedex.DisplayPokedex()
	return nil
}
