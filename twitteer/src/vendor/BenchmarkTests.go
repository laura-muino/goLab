package vendor

import "testing"

func BenchMarkPublishedTweet(b *testing.B){
	// Initialization
	fileTweetWriter :=
	tweetManager := NewTweetManager(NewFileTweetWriter())


}
