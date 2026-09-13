package domain

import "time"

type XmlFeedItem struct {
	GUID        string    `json:"guid"`
	Title       string    `json:"title"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published_at"`
}

var _ FeedItem = &XmlFeedItem{}

func (i *XmlFeedItem) ID() FeedItemID {
	if i.GUID != "" {
		return FeedItemID(sha256Hex(i.GUID))
	}
	return FeedItemID(sha256Hex(i.Link))
}
