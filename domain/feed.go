package domain

type FeedType string

const (
	FeedTypeXML  FeedType = "xml"
	FeedTypeHTML FeedType = "html"
)

type FeedID string

type Feed interface {
	ID() FeedID
}

type HtmlListFeed struct {
	Link            string `json:"link"`
	LiSelector      string `json:"liSelector"`
	TitleSelector   string `json:"titleSelector"`
	ContentSelector string `json:"contentSelector"`
}

var _ Feed = &HtmlListFeed{}

func (f *HtmlListFeed) ID() FeedID {
	return FeedID(sha256Hex(string(FeedTypeHTML) + "\n" + f.Link))
}

type IHtmlListFeedFetcher interface {
	GetItems(feed *HtmlListFeed) ([]*HtmlListFeedItem, error)
}

type XmlFeed struct {
	Link string `json:"link"`
}

var _ Feed = &XmlFeed{}

func (f *XmlFeed) ID() FeedID {
	return FeedID(sha256Hex(string(FeedTypeXML) + "\n" + f.Link))
}
