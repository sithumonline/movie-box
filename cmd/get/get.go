package get

import (
	"net/url"
	"os"

	"github.com/sithumonline/movie-box/internal/yts"

	"github.com/anacrolix/torrent"
	_ "github.com/anacrolix/torrent/metainfo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var name string
var quality string
var destination string

// GetVideoCmd downloads the movie
var GetMovieCmd = &cobra.Command{
	Use:   "get",
	Short: "download movie",
	Run: func(cmd *cobra.Command, args []string) {
		if !(quality == "1080p" || quality == "720p") {
			log.Fatal("movie quality must de 1080p or 720p")
		}
		if name == "" {
			log.Fatal("please enter the name of the move")
		}

		yt, err := yts.GetMovie("https://yts.mx/api/v2/list_movies.json?query_term=" + url.QueryEscape(name))
		if err != nil {
			log.Fatal(err)
		}
		if yt == nil {
			log.Fatal("movie not found")
		}

		torr, logs, _ := yts.GetMovieTorrentLink(yt, quality)

		for i := range logs {
			log.Print(logs[i])
		}

		if torr == "" {
			os.Exit(0)
		}

		cfg := torrent.NewDefaultClientConfig()
		cfg.DataDir = destination
		c, _ := torrent.NewClient(cfg)
		defer c.Close()
		tx, err := yts.GetMateInfoLink(torr)
		if err != nil {
			log.Fatal(err)
		}

		t, err := c.AddTorrent(tx)
		if err != nil {
			log.Fatal(err)
		}
		<-t.GotInfo()
		t.DownloadAll()
		c.WaitAll()

		log.Info("ermahgerd, movie downloaded")
	},
}

func init() {
	GetMovieCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the movie what you what to download")
	GetMovieCmd.Flags().StringVarP(&quality, "quality", "q", "1080p", "Quality of the movie what you what to download")
	GetMovieCmd.Flags().StringVarP(&destination, "out", "o", "./out", "Output directory the movie")
}
