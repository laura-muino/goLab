package service

import (
	"errors"
	"github.com/golab/twitteer/src/domain"
)

//var tweet *domain.Tweet

type TweetManager struct{
	tweets []*domain.Tweet
	tweetsByUser map[string][]*domain.Tweet
}

func NewTweetManager()(*TweetManager){
	tweets := make([]*domain.Tweet, 0)
	tweetsByUser := make(map[string][]*domain.Tweet)
	return &TweetManager{tweets, tweetsByUser}
}

func (v *TweetManager) PublishTweet( aTweet *domain.Tweet ) (int, error){
	if aTweet.Text == "" {
		return 0, errors.New("text is required")
	}
	if aTweet.User == "" {
		return 0, errors.New("user is required")
	}

	if len(aTweet.Text) > 140 {
		return 0, errors.New("text exceding 140 characters")
	}

	v.tweets = append(v.tweets, aTweet)

	v.addTweetByUser( aTweet )

	return aTweet.Id, nil
}

func (v *TweetManager)addTweetByUser( aTweet *domain.Tweet ){
	userTweets, exist := v.tweetsByUser[aTweet.User]
	if !exist{
		newTweetsByUser := []*domain.Tweet{aTweet}
		v.tweetsByUser[aTweet.User]=newTweetsByUser
	}else{
		userTweets = append(userTweets, aTweet)
		v.tweetsByUser[aTweet.User]=userTweets
	}
}

func (v *TweetManager)GetTweetById(id int)(*domain.Tweet){
	for i:=0 ; i < len(v.tweets); i++ {
		if v.tweets[i].Id == id{
			return v.tweets[i]
	    }
	}
	return nil
}

func (v *TweetManager)CountTweetsByUser(user string) (int){
/*	var count int
	for _, aTweet := range tweets{
		if aTweet.User == user{
			count++
	}
	}
	return count*/

	userTweets, exist := v.tweetsByUser[user]
	if exist{
		return len(userTweets)
	}
	return 0

}

func (v *TweetManager) GetLastTweet() (*domain.Tweet){
	return v.tweets[len(v.tweets)-1]
}

func (v *TweetManager) GetTweets() ([]*domain.Tweet){
	return v.tweets
}

func (v *TweetManager) GetTweetsByUser(user string) ( []*domain.Tweet ){
	userTweets, exist := v.tweetsByUser[user]
	if exist{
		return userTweets
	}
	return nil
}