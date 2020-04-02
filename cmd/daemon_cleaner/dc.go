package main

import (
	"flying-castle/app"
	"flying-castle/business"
	"flying-castle/castle"
	"flying-castle/cmd"
	"time"
)

type CleanerFlags struct {
	ScanInterval  time.Duration `flag:"scan_interval" required:"true" default:"1s"`
	EnableWebHook bool          `flag:"enable_webhook"`
}

func (c *CleanerFlags) Validate() {

}

func clean(chunkBusiness *business.ChunkBusiness) error {
	deletedChunks, err := chunkBusiness.DeleteOrphanChunks()
	if err != nil {
		return err
	}
	if len(deletedChunks) == 0 {
		return nil
	}
	var deletedKeys = []string{}
	for _, deletedChunk := range deletedChunks {
		deletedKeys = append(deletedKeys, deletedChunk.Key)
	}
	err = castle.DeleteChunks(deletedKeys)
	if err != nil {
		return err
	}
	return nil
}

func daemonCleaner(config *app.Config, flags CleanerFlags) error {
	var ticker = time.NewTicker(flags.ScanInterval)
	var chunkBusiness = business.NewChunkBusiness()
	for range ticker.C {
		err := clean(chunkBusiness)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	var config = app.GetConfig()
	var flags = CleanerFlags{}
	app.ReadFlags(&flags)
	var err = cmd.SetupApp(config)
	if err != nil {
		panic(err)
	}
	err = daemonCleaner(config, flags)
	if err != nil {
		panic(err)
	}
}
