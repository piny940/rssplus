package main

import (
	"context"
	"flag"
	"log"
	"os"
	"rssplus/cmd"

	"go.yaml.in/yaml/v4"
)

var (
	confFile string
)

func main() {
	flag.StringVar(&confFile, "c", "", "Path to config file")
	flag.Parse()
	if confFile == "" {
		log.Fatal("config file path not provided")
	}
	f, err := os.Open(confFile)
	if err != nil {
		log.Fatalf("failed to open config file: %v", err)
	}
	defer f.Close()
	d := yaml.NewDecoder(f)
	var conf cmd.Config
	err = d.Decode(&conf)
	if err != nil {
		log.Fatalf("failed to decode config file: %v", err)
	}
	ctx := context.Background()
	uc, err := cmd.NewFeedOnceUsecase(ctx)
	if err != nil {
		log.Fatalf("failed to initialize usecase: %v", err)
	}
	if err := uc.NotifyNewItems(ctx, cmd.BuildFeeds(&conf)); err != nil {
		log.Fatalf("failed to notify new items: %v", err)
	}
}
