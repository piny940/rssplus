package main

import (
	"flag"
	"log"
	"os"
	"rssplus/domain"
	"rssplus/infrastructure"
	"rssplus/usecase"

	"go.yaml.in/yaml/v4"
)

type Config struct {
	Feeds []*Feed `yaml:"feeds"`
}

type Feed struct {
	Type            domain.FeedType `yaml:"type"`
	Link            string          `yaml:"link"`
	LiSelector      string          `yaml:"liSelector"`
	TitleSelector   string          `yaml:"titleSelector"`
	ContentSelector string          `yaml:"contentSelector"`
}

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
	var conf Config
	err = d.Decode(&conf)
	if err != nil {
		log.Fatalf("failed to decode config file: %v", err)
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
	uc := usecase.NewFeedOnceUsecase(
		&infrastructure.HTMLListFeedFetcher{},
	)
	if err := uc.NotifyNewItems(feeds); err != nil {
		log.Fatalf("failed to notify new items: %v", err)
	}
}
