package commands

import (
	"io"
	"fmt"
	"context"
	"net/http"
	"encoding/xml"
	"html"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

//Fetch RSSFeed from URL, takes time_between time.ParseDuration
func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
    if err != nil {
        return nil, fmt.Errorf("new request: %w", err)
    }

	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("http do: %w", err)
    }
    defer resp.Body.Close()

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("read body: %w", err)
    }

	var out RSSFeed
	err = xml.Unmarshal(data, &out)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON, %s", err)
	}

	return &out, nil
}

//Remove html escape characters from RSSFeed
func (f *RSSFeed) Unescape() {
	f.Channel.Title = html.UnescapeString(f.Channel.Title)
	f.Channel.Description = html.UnescapeString(f.Channel.Description)

	for i := range f.Channel.Item {
		f.Channel.Item[i].Title = html.UnescapeString(f.Channel.Item[i].Title)
		f.Channel.Item[i].Description = html.UnescapeString(f.Channel.Item[i].Description)
	}
}