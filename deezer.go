package deezer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	http    *http.Client
	baseURL string
}

func New(httpClient *http.Client) *Client {
	c := &Client{
		http:    httpClient,
		baseURL: "https://api.deezer.com",
	}

	return c
}

func (c *Client) get(ctx context.Context, url string, result interface{}) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		return c.decodeError(resp)
	}

	//byt, err := io.ReadAll(resp.Body)
	//log.Println(string(byt))
	//log.Println(err)

	return json.NewDecoder(resp.Body).Decode(result)
}

func (c *Client) decodeError(resp *http.Response) error {
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(responseBody) == 0 {
		return fmt.Errorf("deezer: HTTP %d (body empty)", resp.StatusCode)
	}

	buf := bytes.NewBuffer(responseBody)

	return fmt.Errorf("deezer: HTTP %d: %s", resp.StatusCode, buf.String())
}
