package main

import (
	"context"
	"fmt"
	"rssplus/cmd"
	"rssplus/samples"

	"github.com/aws/aws-lambda-go/lambda"
	"go.yaml.in/yaml/v4"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context) error {
	var conf cmd.Config
	if err := yaml.Unmarshal(samples.Simple, &conf); err != nil {
		return fmt.Errorf("failed to unmarshal config file: %w", err)
	}
	uc, err := cmd.NewFeedOnceUsecase(ctx)
	if err != nil {
		return fmt.Errorf("failed to initialize usecase: %w", err)
	}
	if err := uc.NotifyNewItems(ctx, cmd.BuildFeeds(&conf)); err != nil {
		return fmt.Errorf("failed to notify new items: %w", err)
	}
	return nil
}
