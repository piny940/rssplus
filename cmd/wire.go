package cmd

import (
	"context"
	"fmt"
	"os"
	"rssplus/domain"
	"rssplus/infrastructure"
	"rssplus/usecase"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const feedItemsTableNameEnv = "FEED_ITEMS_TABLE_NAME"

func BuildFeeds(conf *Config) []domain.Feed {
	var feeds []domain.Feed
	for _, f := range conf.Feeds {
		switch f.Type {
		case domain.FeedTypeXML:
			feeds = append(feeds, &domain.XmlFeed{
				Link: f.Link,
			})
		case domain.FeedTypeHTML:
			feeds = append(feeds, &domain.HtmlListFeed{
				Link:            f.Link,
				LiSelector:      f.LiSelector,
				TitleSelector:   f.TitleSelector,
				ContentSelector: f.ContentSelector,
			})
		}
	}
	return feeds
}

func NewFeedOnceUsecase(ctx context.Context) (usecase.IFeedOnceUsecase, error) {
	tableName := os.Getenv(feedItemsTableNameEnv)
	if tableName == "" {
		return nil, fmt.Errorf("environment variable %s is not set", feedItemsTableNameEnv)
	}
	awsConf, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load aws config: %w", err)
	}
	client := dynamodb.NewFromConfig(awsConf)
	return usecase.NewFeedOnceUsecase(
		infrastructure.NewHtmlListFeedFetcher(),
		infrastructure.NewDynamoHtmlListFeedItemRepository(client, tableName),
	), nil
}
