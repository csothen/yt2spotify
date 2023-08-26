package integrations

import "github.com/csothen/yt2spotify/pkg/configuration"

func Configure(config *configuration.Configuration) error {
	if err := configureYoutube(config); err != nil {
		return err
	}

	if err := configureSpotify(config); err != nil {
		return err
	}

	return nil
}
