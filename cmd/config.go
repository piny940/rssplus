package cmd

import "rssplus/domain"

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
