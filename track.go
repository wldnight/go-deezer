package deezer

type Track struct {
	Id                    int      `json:"id"`
	Readable              string   `json:"readable"`
	Title                 string   `json:"title"`
	TitleShort            string   `json:"title_short"`
	TitleVersion          string   `json:"title_version"`
	Unseen                bool     `json:"unseen"`
	ISRC                  string   `json:"isrc"`
	Link                  string   `json:"link"`
	Share                 string   `json:"share"`
	Duration              int      `json:"duration"`
	TrackPosition         int      `json:"track_position"`
	DiskNumber            int      `json:"disk_number"`
	Rank                  int      `json:"rank"`
	ReleaseDate           string   `json:"release_date"`
	ExplicitLyrics        bool     `json:"explicit_lyrics"`
	ExplicitContentLyrics int      `json:"explicit_content_lyrics"`
	ExplicitContentCover  int      `json:"explicit_content_cover"`
	Preview               string   `json:"preview"`
	BPM                   float64  `json:"bpm"`
	Gain                  float64  `json:"gain"`
	AvailableCountries    []string `json:"available_countries"`
	Alternative           string   `json:"alternative"`
	Contributors          []string `json:"contributors"`
	MD5Image              string   `json:"md5_image"`
	TrackToken            string   `json:"track_token"`
	Artist                Artist   `json:"artist"`
	Album                 Album    `json:"album"`
}
