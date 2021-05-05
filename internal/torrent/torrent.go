package torrent

import (
	"os"

	"github.com/anacrolix/torrent"
	log "github.com/sirupsen/logrus"
)

type Torrent = torrent.Torrent

func GetTorrentClient() func() *torrent.Client {
	path := "/tmp/movie-box.log"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			log.Error(err)
		}
		defer f.Close()
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Error(err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		log.Error(err)
	}

	cfg := torrent.NewDefaultClientConfig()
	cfg.DataDir = home + "/Downloads/movie-box"

	c := func() *torrent.Client {
		c, _ := torrent.NewClient(cfg)
		c.WriteStatus(file)
		defer c.Close()
		return c
	}

	return c
}
