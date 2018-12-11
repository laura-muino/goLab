package service

import "github.com/golab/twitteer/src/domain"

type MemoryTweetWriter struct{
	tweets []domain.Tweet
}

func NewMemoryTweetWriter()( *MemoryTweetWriter ){
	return &MemoryTweetWriter{}
}

func (mtw *MemoryTweetWriter) Write(tweet domain.Tweet){
	// Agregamos al array
	mtw.tweets = append(mtw.tweets, tweet)
}

func (mtw *MemoryTweetWriter) GetLastSavedTweet()(domain.Tweet){
	return mtw.tweets[len(mtw.tweets)-1]
}