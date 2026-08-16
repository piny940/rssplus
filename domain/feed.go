package domain

import "time"

type FeedItem struct {
	Title       string    `json:"title"`
	Link        string    `json:"link"`
	PublishedAt time.Time `json:"published_at"`
	Content     string    `json:"content"`
}

type FeedItems struct {
	Items []*FeedItem
}

type FeedType string

const (
	FeedTypeXML  FeedType = "xml"
	FeedTypeHTML FeedType = "html"
)

type HtmlListFeed struct {
	Link            string `json:"link"`
	LiSelector      string `json:"liSelector"`
	TitleSelector   string `json:"titleSelector"`
	ContentSelector string `json:"contentSelector"`
}

var _ Feed = &HtmlListFeed{}

type IHtmlListFeedFetcher interface {
	GetItems(feed *HtmlListFeed) (*FeedItems, error)
}

type XmlFeed struct {
	Link string `json:"link"`
}

var _ Feed = &XmlFeed{}

type Feed interface {
}
