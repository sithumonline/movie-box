package kodi

import (
	"encoding/xml"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

func CreateMovieInfoFile(movie *Movie, filename string) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Error(err)
	}

	path := home + "/Downloads/movie-box/" + filename + "/artist.nfo"
	f, err := os.Create(path)
	if err != nil {
		log.Error(err)
	}
	defer f.Close()

	out, _ := xml.MarshalIndent(movie, " ", "  ")
	w := xml.Header + string(out)
	err = ioutil.WriteFile(path, []byte(w), 0644)
	if err != nil {
		log.Error(err)
	}
}
