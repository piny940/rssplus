package infrastructure

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"rssplus/domain"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36"

var browserHeaders = map[string]string{
	"User-Agent":                userAgent,
	"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	"Accept-Language":           "ja,en-US;q=0.9,en;q=0.8",
	"Upgrade-Insecure-Requests": "1",
	"Sec-Fetch-Dest":            "document",
	"Sec-Fetch-Mode":            "navigate",
	"Sec-Fetch-Site":            "none",
	"Sec-Fetch-User":            "?1",
}

var ErrNoItemsFound = errors.New("no items found")

const (
	maxAttempts        = 3
	initialRetryWait   = 30 * time.Second
	errorBodyMaxLength = 512
)

type HTMLListFeedFetcher struct {
}

var _ domain.IHtmlListFeedFetcher = &HTMLListFeedFetcher{}

func (h *HTMLListFeedFetcher) GetItems(feed *domain.HtmlListFeed) (*domain.FeedItems, error) {
	wait := initialRetryWait
	for attempt := 1; ; attempt++ {
		items, err := h.getItemsOnce(feed)
		if !errors.Is(err, ErrNoItemsFound) || attempt == maxAttempts {
			return items, err
		}
		time.Sleep(wait)
		wait *= 2
	}
}

func (h *HTMLListFeedFetcher) getItemsOnce(feed *domain.HtmlListFeed) (*domain.FeedItems, error) {
	req, err := http.NewRequest(http.MethodGet, feed.Link, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	for k, v := range browserHeaders {
		req.Header.Set(k, v)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to http Get: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		body, _ := io.ReadAll(io.LimitReader(res.Body, errorBodyMaxLength))
		return nil, fmt.Errorf("status code was not 200, was %d: %s", res.StatusCode, body)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to get document: %w", err)
	}
	items := make([]*domain.FeedItem, 0)
	doc.Find(feed.LiSelector).Each(func(i int, s *goquery.Selection) {
		title := s.Find(feed.TitleSelector).Text()
		content := s.Find(feed.ContentSelector).Text()
		items = append(items, &domain.FeedItem{
			Title:   title,
			Content: content,
		})
	})
	if len(items) == 0 {
		return nil, fmt.Errorf("%w: %q matched no element", ErrNoItemsFound, feed.LiSelector)
	}
	return &domain.FeedItems{Items: items}, nil
}
