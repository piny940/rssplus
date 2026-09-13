package domain

import (
	"crypto/sha256"
	"encoding/hex"
)

type FeedItemID string

type FeedItem interface {
	ID() FeedItemID
}

func sha256Hex(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}
