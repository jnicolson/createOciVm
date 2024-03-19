package main

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/pelletier/go-toml"
)

func ReadConfig(configFile string) *Config {
	log.Info().Str("file", configFile).Msg("Reading Config File")

	file, err := os.Open(configFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Loading config file")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			
		}
	}(file)

	config := Config{}

	configParser := toml.NewDecoder(file)
	err = configParser.Decode(&config)
	if err != nil {
		return nil
	}

	return &config

}
