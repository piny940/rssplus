package domain

import "testing"

func TestHtmlListFeedItemID(t *testing.T) {
	a := &HtmlListFeedItem{Title: "title", Content: "content"}
	sameAsA := &HtmlListFeedItem{Title: "title", Content: "content"}
	differentContent := &HtmlListFeedItem{Title: "title", Content: "other"}

	if a.ID() != sameAsA.ID() {
		t.Errorf("same title and content should have same ID")
	}
	if a.ID() == differentContent.ID() {
		t.Errorf("different content should have different ID")
	}
}
