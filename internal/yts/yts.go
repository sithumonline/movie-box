package yts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/sirupsen/logrus"
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
			logrus.Error(err)
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
			logrus.Error(err)
		}
	}()

	ytRes := Response{}

	if err := json.NewDecoder(resp.Body).Decode(&ytRes); err != nil {
		return nil, err
	}

	return &ytRes, nil
}

//GetMovieTorrentLink find torrent url that matches to given quality
func GetMovieTorrentLink(res *Response, quality string) string {
	var links string
	movies := res.Data.Movies

	if len(movies) > 0 {
		logrus.Info("there are few movies you can download, please enter the exact movie that you need to download")
		for i := range movies {
			logrus.Println(movies[i].TitleEnglish)
		}
		os.Exit(0)
	}

	for i := range movies {
		for i2 := range movies[i].Torrents {
			matched, _ := regexp.MatchString(quality, movies[i].Torrents[i2].Quality)
			if !matched {
				continue
			}
			links = movies[i].Torrents[i2].Url
			logrus.Info("movie : " + movies[i].Url)
			break
		}
	}

	return links
}
