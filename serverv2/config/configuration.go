package config

type Configuration struct {
	serverConfiguration
	databaseConfiguration
	spotifyConfiguration
	youtubeConfiguration
}

func Load() Configuration {
	return Configuration{
		serverConfiguration:   newServerConfiguration(),
		databaseConfiguration: newDatabaseConfiguration(),
		spotifyConfiguration:  newSpotifyConfiguration(),
		youtubeConfiguration:  newYoutubeConfiguration(),
	}
}
