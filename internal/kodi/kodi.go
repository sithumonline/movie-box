package kodi

import (
	"encoding/xml"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

func CreateMovieInfoFile(movie *Movie, filename string) {
	out, _ := xml.MarshalIndent(movie, " ", "  ")
	w := xml.Header + string(out)
	err := ioutil.WriteFile("~/Downloads/movie-box/"+filename+"/artist.nfo", []byte(w), 0644)
	if err != nil {
		log.Error(err)
	}
}
