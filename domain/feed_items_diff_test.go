package domain

import (
	"slices"
	"testing"
)

func item(title string) *HtmlListFeedItem {
	return &HtmlListFeedItem{Title: title, Content: title + " content"}
}

func titles(items []*HtmlListFeedItem) []string {
	res := make([]string, 0, len(items))
	for _, i := range items {
		res = append(res, i.Title)
	}
	return res
}

func TestDiffFeedItems(t *testing.T) {
	tests := []struct {
		name        string
		previous    []*HtmlListFeedItem
		current     []*HtmlListFeedItem
		wantAdded   []string
		wantRemoved []string
	}{
		{
			name:        "previousが空なら全件Added",
			previous:    nil,
			current:     []*HtmlListFeedItem{item("a"), item("b")},
			wantAdded:   []string{"a", "b"},
			wantRemoved: []string{},
		},
		{
			name:        "変化なし",
			previous:    []*HtmlListFeedItem{item("a"), item("b")},
			current:     []*HtmlListFeedItem{item("b"), item("a")},
			wantAdded:   []string{},
			wantRemoved: []string{},
		},
		{
			name:        "追加のみ",
			previous:    []*HtmlListFeedItem{item("a")},
			current:     []*HtmlListFeedItem{item("c"), item("a"), item("b")},
			wantAdded:   []string{"c", "b"},
			wantRemoved: []string{},
		},
		{
			name:        "削除のみ",
			previous:    []*HtmlListFeedItem{item("a"), item("b"), item("c")},
			current:     []*HtmlListFeedItem{item("b")},
			wantAdded:   []string{},
			wantRemoved: []string{"a", "c"},
		},
		{
			name:        "追加と削除の両方",
			previous:    []*HtmlListFeedItem{item("a"), item("b")},
			current:     []*HtmlListFeedItem{item("b"), item("c")},
			wantAdded:   []string{"c"},
			wantRemoved: []string{"a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diff := DiffFeedItems(tt.previous, tt.current)
			if got := titles(diff.Added); !slices.Equal(got, tt.wantAdded) {
				t.Errorf("Added = %v, want %v", got, tt.wantAdded)
			}
			if got := titles(diff.Removed); !slices.Equal(got, tt.wantRemoved) {
				t.Errorf("Removed = %v, want %v", got, tt.wantRemoved)
			}
		})
	}
}
