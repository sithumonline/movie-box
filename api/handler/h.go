package handler

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/sithumonline/movie-box/internal/torrent"
	"github.com/sithumonline/movie-box/internal/yts"

	"github.com/go-chi/chi"
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

	t, _ := torrent.GetTorrent().AddTorrent(tx)
	t.DownloadAll()

	RespondWithText(w, http.StatusOK, resLog.String())
}
