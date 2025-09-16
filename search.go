package deezer

import (
	"context"
	"net/url"
)

type SearchResult struct {
	Id             int    `json:"id"`
	Readable       bool   `json:"readable"`
	Title          string `json:"title"`
	TitleShort     string `json:"title_short"`
	TitleVersion   string `json:"title_version"`
	Link           string `json:"link"`
	Duration       int    `json:"duration"`
	Rank           int    `json:"rank"`
	ExplicitLyrics bool   `json:"explicit_lyrics"`
	Preview        string `json:"preview"`
	Artist         Artist `json:"artist"`
	Album          Album  `json:"album"`
}

func (c *Client) Search(ctx context.Context, query string) ([]SearchResult, error) {
	deezerUrl := c.baseURL + "/search?q=" + url.QueryEscape(query)

	var result struct {
		Data []SearchResult `json:"data"`
	}

	err := c.get(ctx, deezerUrl, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}
