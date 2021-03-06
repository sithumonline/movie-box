package handler

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/sithumonline/movie-box/internal/torrent"
	"github.com/sithumonline/movie-box/internal/yts"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func AddMovie(w http.ResponseWriter, r *http.Request) {
	var resLog strings.Builder

	name := chi.URLParam(r, "name")
	yt, err := yts.GetMovie("https://yts.mx/api/v2/list_movies.json?query_term=" + url.QueryEscape(name))
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if yt == nil {
		RespondWithError(w, http.StatusNotFound, "movie not found")
		return
	}

	torr, logs := yts.GetMovieTorrentLink(yt, "1080p")
	for i := range logs {
		resLog.WriteString(logs[i])
	}
	if torr == "" {
		RespondWithText(w, http.StatusBadRequest, resLog.String())
		return
	}

	tx, err := yts.GetMateInfoLink(torr)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	t := func() *torrent.Torrent {
		c := torrent.GetTorrentClient()
		t, err := c().AddTorrent(tx)
		if err != nil {
			log.Fatal(err)
		}

		return t
	}()

	go func() {
		<-t.GotInfo()
		t.DownloadAll()
	}()

	RespondWithText(w, http.StatusOK, resLog.String())
}

func GetLogs(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("/tmp/movie-box.log")

	if err != nil {
		log.Error(err)
	}

	RespondWithText(w, http.StatusOK, string(content))
}
