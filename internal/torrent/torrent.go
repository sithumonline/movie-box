package torrent

import (
	"os"

	"github.com/anacrolix/torrent"
	log "github.com/sirupsen/logrus"
)

func GetTorrent() *torrent.Client {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Error(err)
	}

	cfg := torrent.NewDefaultClientConfig()
	cfg.DataDir = home + "/Downloads/movie-box"
	c, _ := torrent.NewClient(cfg)

	if c == nil {
		return c
	}

	return c
}
