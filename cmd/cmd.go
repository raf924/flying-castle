package cmd

import (
	"flying-castle/app"
	"flying-castle/castle"
	"flying-castle/db"
)

func SetupApp(config *app.Config) error {
	err := db.LoadDB(config.DbUrl)
	if err != nil {
		return db.ConnectionError
	}
	backend, err := castle.NewBackend(config.DataPath, config)
	if err != nil {
		return err
	}
	castle.SetStorageBackend(backend)
	return nil
}
