package internal

import (
	"encoding/xml"
	"fmt"
)

const RSS_URL = "https://trends.google.es/trends/trendingsearches/daily/rss?geo=MX"

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

func (r RSS) Print() string {
	var items string
	for _, item := range r.Channel.ItemList {
		items += fmt.Sprintf("\n\t%s", item.Print())
	}
	return fmt.Sprintf("%s - %s\nURL: %v \n%s", r.Channel.Title, r.Channel.Description, r.Channel.Link.Url, items)
}

type ChannelLink struct {
	XMLName xml.Name `xml:"link"`
	Url     string   `xml:"href,attr"`
}
type Channel struct {
	Title       string      `xml:"title"`
	Description string      `xml:"description"`
	Link        ChannelLink `xml:"link"`
	ItemList    []Item      `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Traffic     string `xml:"approx_traffic"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	NewsList    []News `xml:"news_item"`
}

func (i Item) Print() string {
	var news string
	for _, n := range i.NewsList {
		news += fmt.Sprintf("\n\t\t- %s [%s]", n.Headline, n.Source)
	}
	return fmt.Sprintf("* %s [%s] on %s (%s)\n%s\n", i.Title, i.Traffic, i.PubDate, i.Link, news)
}

type News struct {
	Headline    string `xml:"news_item_title"`
	Description string `xml:"news_item_snippet"`
	Source      string `xml:"news_item_source"`
	Link        string `xml:"news_item_url"`
}

// Path: internal/types.go
