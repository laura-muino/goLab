package main

import (
	"github.com/golab/twitteer/src/domain"
	"github.com/golab/twitteer/src/service"
	"testing"
)

func BenchMarkPublishedTweet(b *testing.B){
	// Initialization
	tweetManager := service.NewTweetManager(service.NewFileTweetWriter())

	tweet := domain.NewTextTweet("grupoesfera", "this is my tweet")

	// operation
	for n := 0; n < b.N; n++ {
		tweetManager.PublishTweet(tweet)
	}
}
