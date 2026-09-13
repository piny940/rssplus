package infrastructure

import (
	"context"
	"fmt"
	"rssplus/domain"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type htmlListFeedSnapshotRecord struct {
	FeedID    string                   `dynamodbav:"feed_id"`
	FeedType  string                   `dynamodbav:"feed_type"`
	Link      string                   `dynamodbav:"link"`
	Items     []htmlListFeedItemRecord `dynamodbav:"items"`
	UpdatedAt time.Time                `dynamodbav:"updated_at"`
}

type htmlListFeedItemRecord struct {
	Title   string `dynamodbav:"title"`
	Content string `dynamodbav:"content"`
}

type DynamoHtmlListFeedItemRepository struct {
	client    *dynamodb.Client
	tableName string
}

var _ domain.IHtmlListFeedItemRepository = &DynamoHtmlListFeedItemRepository{}

func NewDynamoHtmlListFeedItemRepository(client *dynamodb.Client, tableName string) *DynamoHtmlListFeedItemRepository {
	return &DynamoHtmlListFeedItemRepository{client: client, tableName: tableName}
}

func (r *DynamoHtmlListFeedItemRepository) GetLatest(ctx context.Context, feed *domain.HtmlListFeed) ([]*domain.HtmlListFeedItem, error) {
	out, err := r.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"feed_id": &types.AttributeValueMemberS{Value: string(feed.ID())},
		},
		ConsistentRead: aws.Bool(true),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get item from dynamodb: %w", err)
	}
	if len(out.Item) == 0 {
		return []*domain.HtmlListFeedItem{}, nil
	}
	var record htmlListFeedSnapshotRecord
	if err := attributevalue.UnmarshalMap(out.Item, &record); err != nil {
		return nil, fmt.Errorf("failed to unmarshal snapshot record: %w", err)
	}
	items := make([]*domain.HtmlListFeedItem, 0, len(record.Items))
	for _, i := range record.Items {
		items = append(items, &domain.HtmlListFeedItem{
			Title:   i.Title,
			Content: i.Content,
		})
	}
	return items, nil
}

func (r *DynamoHtmlListFeedItemRepository) SaveLatest(ctx context.Context, feed *domain.HtmlListFeed, items []*domain.HtmlListFeedItem) error {
	record := htmlListFeedSnapshotRecord{
		FeedID:    string(feed.ID()),
		FeedType:  string(domain.FeedTypeHTML),
		Link:      feed.Link,
		Items:     make([]htmlListFeedItemRecord, 0, len(items)),
		UpdatedAt: time.Now(),
	}
	for _, i := range items {
		record.Items = append(record.Items, htmlListFeedItemRecord{
			Title:   i.Title,
			Content: i.Content,
		})
	}
	av, err := attributevalue.MarshalMap(record)
	if err != nil {
		return fmt.Errorf("failed to marshal snapshot record: %w", err)
	}
	if _, err := r.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      av,
	}); err != nil {
		return fmt.Errorf("failed to put item to dynamodb: %w", err)
	}
	return nil
}
