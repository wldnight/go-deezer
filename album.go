package deezer

type Album struct {
	Id                    int      `json:"id"`
	Title                 string   `json:"title"`
	UPC                   string   `json:"upc"`
	Link                  string   `json:"link"`
	Share                 string   `json:"share"`
	Cover                 string   `json:"cover"`
	CoverSmall            string   `json:"cover_small"`
	CoverMedium           string   `json:"cover_medium"`
	CoverBig              string   `json:"cover_big"`
	CoverXl               string   `json:"cover_xl"`
	MD5Image              string   `json:"md5_image"`
	GenreId               int      `json:"genre_id"`
	Genres                Genre    `json:"genres"`
	Label                 string   `json:"label"`
	Provider              string   `json:"provider"`
	TrackCount            int      `json:"nb_tracks"`
	Duration              int      `json:"duration"`
	FanCount              int      `json:"fans"`
	ReleaseDate           string   `json:"release_date"`
	RecordType            string   `json:"record_type"`
	Available             bool     `json:"available"`
	Tracklist             string   `json:"tracklist"`
	ExplicitLyrics        bool     `json:"explicit_lyrics"`
	ExplicitContentLyrics int      `json:"explicit_content_lyrics"`
	ExplicitContentCover  int      `json:"explicit_content_cover"`
	Contributors          []string `json:"contributors"`
	Fallback              *Album   `json:"fallback"`
	Artist                Artist   `json:"artist"`
	Tracks                []Track  `json:"tracks,omitempty"`
}
