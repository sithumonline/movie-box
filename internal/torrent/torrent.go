package torrent

import "github.com/anacrolix/torrent"

func GetTorrent() *torrent.Client {
	cfg := torrent.NewDefaultClientConfig()
	cfg.DataDir = "~/Downloads/movie-box"
	c, _ := torrent.NewClient(cfg)

	if c == nil {
		return c
	}

	return c
}
