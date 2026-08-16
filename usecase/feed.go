package usecase

import (
	"fmt"
	"rssplus/domain"
)

type FeedOnceUsecase struct {
	htmlFeedFetcher domain.IHtmlListFeedFetcher
}

type IFeedOnceUsecase interface {
	NotifyNewItems(feeds []domain.Feed) error
}

func NewFeedOnceUsecase(htmlFeedFetcher domain.IHtmlListFeedFetcher) *FeedOnceUsecase {
	return &FeedOnceUsecase{htmlFeedFetcher: htmlFeedFetcher}
}

func (uc *FeedOnceUsecase) NotifyNewItems(feeds []domain.Feed) error {
	for _, feed := range feeds {
		var items *domain.FeedItems
		var err error
		switch v := feed.(type) {
		case *domain.HtmlListFeed:
			items, err = uc.htmlFeedFetcher.GetItems(v)
			if err != nil {
				return fmt.Errorf("failed to fetch html items: %w", err)
			}
		}
		fmt.Println(items)
	}
	return nil
}
