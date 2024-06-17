package album

import (
	"encoding/json"
	apirequest "spotify-api/src/api-request"
	env "spotify-api/src/config"
)

type ExternalURLs struct {
	Spotify string `json:"spotify"`
}

type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Artist struct {
	ExternalURLs ExternalURLs `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

type Restrictions struct {
	Reason string `json:"reason"`
}

type LinkedFrom struct {
	ExternalURLs ExternalURLs `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

type Track struct {
	Artists          []Artist     `json:"artists"`
	AvailableMarkets []string     `json:"available_markets"`
	DiscNumber       int          `json:"disc_number"`
	DurationMs       int          `json:"duration_ms"`
	Explicit         bool         `json:"explicit"`
	ExternalURLs     ExternalURLs `json:"external_urls"`
	Href             string       `json:"href"`
	ID               string       `json:"id"`
	IsPlayable       bool         `json:"is_playable"`
	LinkedFrom       LinkedFrom   `json:"linked_from"`
	Restrictions     Restrictions `json:"restrictions"`
	Name             string       `json:"name"`
	PreviewURL       string       `json:"preview_url"`
	TrackNumber      int          `json:"track_number"`
	Type             string       `json:"type"`
	URI              string       `json:"uri"`
	IsLocal          bool         `json:"is_local"`
}

type Album struct {
	AlbumType            string       `json:"album_type"`
	TotalTracks          int          `json:"total_tracks"`
	AvailableMarkets     []string     `json:"available_markets"`
	ExternalURLs         ExternalURLs `json:"external_urls"`
	Href                 string       `json:"href"`
	ID                   string       `json:"id"`
	Images               []Image      `json:"images"`
	Name                 string       `json:"name"`
	ReleaseDate          string       `json:"release_date"`
	ReleaseDatePrecision string       `json:"release_date_precision"`
	Restrictions         Restrictions `json:"restrictions"`
	Type                 string       `json:"type"`
	URI                  string       `json:"uri"`
	Artists              []Artist     `json:"artists"`
	Tracks               struct {
		Href     string  `json:"href"`
		Limit    int     `json:"limit"`
		Next     string  `json:"next"`
		Offset   int     `json:"offset"`
		Previous string  `json:"previous"`
		Total    int     `json:"total"`
		Items    []Track `json:"items"`
	} `json:"tracks"`
	Copyrights []struct {
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"copyrights"`
	ExternalIDs struct {
		Isrc string `json:"isrc"`
		Ean  string `json:"ean"`
		Upc  string `json:"upc"`
	} `json:"external_ids"`
	Genres     []string `json:"genres"`
	Label      string   `json:"label"`
	Popularity int      `json:"popularity"`
}

func GetData(albumId string) (Album, error) {
	envConfig := env.LoadEnvConfig()
	url := envConfig.SpotifyBaseURL + "/albums/" + albumId
	response, err := apirequest.MakeRequest("GET", url, "")
	if err != nil {
		return Album{}, err
	}

	var album Album
	err = json.Unmarshal([]byte(response), &album)
	if err != nil {
		return Album{}, err
	}
	return album, nil
}
