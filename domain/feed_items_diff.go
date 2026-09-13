package domain

type FeedItemsDiff[T FeedItem] struct {
	Added   []T
	Removed []T
}

// previousが空の場合は、currentの全アイテムがAddedになる
func DiffFeedItems[T FeedItem](previous, current []T) FeedItemsDiff[T] {
	previousIDs := make(map[FeedItemID]struct{}, len(previous))
	for _, item := range previous {
		previousIDs[item.ID()] = struct{}{}
	}
	currentIDs := make(map[FeedItemID]struct{}, len(current))
	for _, item := range current {
		currentIDs[item.ID()] = struct{}{}
	}

	var diff FeedItemsDiff[T]
	for _, item := range current {
		if _, ok := previousIDs[item.ID()]; !ok {
			diff.Added = append(diff.Added, item)
		}
	}
	for _, item := range previous {
		if _, ok := currentIDs[item.ID()]; !ok {
			diff.Removed = append(diff.Removed, item)
		}
	}
	return diff
}
