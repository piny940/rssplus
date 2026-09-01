package lambdaonce

import (
	"context"
	"fmt"
	"log"
	"rssplus/cmd"
	"rssplus/domain"
	"rssplus/infrastructure"
	"rssplus/samples"
	"rssplus/usecase"

	"github.com/aws/aws-lambda-go/lambda"
	"go.yaml.in/yaml/v4"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(_ context.Context) error {
	var conf cmd.Config
	if err := yaml.Unmarshal(samples.Simple, &conf); err != nil {
		return fmt.Errorf("failed to unmarshal config file: %w", err)
	}
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
	uc := usecase.NewFeedOnceUsecase(infrastructure.NewHtmlListFeedFetcher())
	if err := uc.NotifyNewItems(feeds); err != nil {
		log.Fatalf("failed to notify new items: %v", err)
	}
	return nil
}
