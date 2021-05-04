package yts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/sithumonline/movie-box/internal/kodi"

	"github.com/anacrolix/torrent/metainfo"
	log "github.com/sirupsen/logrus"
)

//getMateInfo extract *metainfo.MetaInfo from torrent file
func getMateInfo(body io.Reader) (*metainfo.MetaInfo, error) {
	m, err := metainfo.Load(body)
	if err != nil {
		return nil, err
	}

	return m, nil
}

//GetMateInfoLink download torrent file from given url
func GetMateInfoLink(url string) (*metainfo.MetaInfo, error) {
	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-Type", "application/text")

	if reqErr != nil {
		return nil, reqErr
	}

	client := &http.Client{}
	resp, respErr := client.Do(req)

	if respErr != nil {
		return nil, respErr
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("torrent download err : %d", resp.StatusCode)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Error(err)
		}
	}()

	m, err := getMateInfo(resp.Body)

	if err != nil {
		return nil, err
	}

	return m, nil
}

//GetMovie fetch json file from YTS API
//https://yts.mx/api#list_movies
func GetMovie(url string) (*Response, error) {
	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-Type", "application/text")

	if reqErr != nil {
		return nil, reqErr
	}

	client := &http.Client{}
	resp, respErr := client.Do(req)

	if respErr != nil {
		return nil, respErr
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("torrent list err : %d", resp.StatusCode)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Error(err)
		}
	}()

	ytRes := Response{}

	if err := json.NewDecoder(resp.Body).Decode(&ytRes); err != nil {
		return nil, err
	}

	return &ytRes, nil
}

//GetMovieTorrentLink find torrent url that matches to given quality
func GetMovieTorrentLink(res *Response, quality string) (string, []string, *kodi.Movie) {
	var links string
	var logs = []string{""}
	var k *kodi.Movie
	movies := res.Data.Movies

	if len(movies) > 1 {
		logs = append(logs, "there are few movies you can download, please enter the exact movie that you need to download\n")
		links = ""
		for i := range movies {
			logs = append(logs, movies[i].TitleEnglish+"\n")
		}
		return "", logs, k
	}

	for i := range movies {
		k = &kodi.Movie{
			Title:         movies[i].TitleEnglish,
			OriginalTitle: movies[i].Title,
			UserRating:    movies[i].Rating,
			Plot:          movies[i].Summary,
			Runtime:       movies[i].Runtime,
			Thumb:         movies[i].LargeCoverImage,
		}
		for i2 := range movies[i].Torrents {
			matched, _ := regexp.MatchString(quality, movies[i].Torrents[i2].Quality)
			if !matched {
				continue
			}
			links = movies[i].Torrents[i2].Url
			logs = append(logs, "movie : "+movies[i].Url+"\n", "torrent : "+links+"\n")
			break
		}
	}

	return links, logs, k
}
