package domain

import "context"

type IHtmlListFeedItemRepository interface {
	GetLatest(ctx context.Context, feed *HtmlListFeed) ([]*HtmlListFeedItem, error)
	SaveLatest(ctx context.Context, feed *HtmlListFeed, items []*HtmlListFeedItem) error
}
