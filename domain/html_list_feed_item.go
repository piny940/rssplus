package domain

type HtmlListFeedItem struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

var _ FeedItem = &HtmlListFeedItem{}

// HTMLのリストには安定したキーがないため、内容全体から同一性を判定する
func (i *HtmlListFeedItem) ID() FeedItemID {
	return FeedItemID(sha256Hex(i.Title + "\x00" + i.Content))
}
