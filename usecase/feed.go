package usecase

import (
	"context"
	"fmt"
	"rssplus/domain"
)

type FeedOnceUsecase struct {
	htmlFeedFetcher        domain.IHtmlListFeedFetcher
	htmlFeedItemRepository domain.IHtmlListFeedItemRepository
}

type IFeedOnceUsecase interface {
	NotifyNewItems(ctx context.Context, feeds []domain.Feed) error
}

var _ IFeedOnceUsecase = &FeedOnceUsecase{}

func NewFeedOnceUsecase(
	htmlFeedFetcher domain.IHtmlListFeedFetcher,
	htmlFeedItemRepository domain.IHtmlListFeedItemRepository,
) *FeedOnceUsecase {
	return &FeedOnceUsecase{
		htmlFeedFetcher:        htmlFeedFetcher,
		htmlFeedItemRepository: htmlFeedItemRepository,
	}
}

func (uc *FeedOnceUsecase) NotifyNewItems(ctx context.Context, feeds []domain.Feed) error {
	for _, feed := range feeds {
		switch v := feed.(type) {
		case *domain.HtmlListFeed:
			if err := uc.processHtmlListFeed(ctx, v); err != nil {
				return err
			}
		}
	}
	return nil
}

func (uc *FeedOnceUsecase) processHtmlListFeed(ctx context.Context, feed *domain.HtmlListFeed) error {
	current, err := uc.htmlFeedFetcher.GetItems(feed)
	if err != nil {
		return fmt.Errorf("failed to fetch html items: %w", err)
	}
	previous, err := uc.htmlFeedItemRepository.GetLatest(ctx, feed)
	if err != nil {
		return fmt.Errorf("failed to get latest html items: %w", err)
	}

	diff := domain.DiffFeedItems(previous, current)
	// TODO: 通知を実装する（diff.Addedのみを通知対象とする）。保存より前に行うこと
	fmt.Printf("feed %s: %d added, %d removed\n", feed.Link, len(diff.Added), len(diff.Removed))
	for _, item := range diff.Added {
		fmt.Printf("  + %s\n", item.Title)
	}
	for _, item := range diff.Removed {
		fmt.Printf("  - %s\n", item.Title)
	}

	if err := uc.htmlFeedItemRepository.SaveLatest(ctx, feed, current); err != nil {
		return fmt.Errorf("failed to save latest html items: %w", err)
	}
	return nil
}
