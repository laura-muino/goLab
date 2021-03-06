package service

import (
	"github.com/golab/twitteer/src/domain"
	"strings"
)

//var tweet *domain.Tweet

type TweetManager struct{
	tweets []domain.Tweet
	tweetsByUser map[string][]domain.Tweet
	writer TweetWriter
}

func NewTweetManager(aTweeterWriter TweetWriter)(*TweetManager){
	tweets := make([]domain.Tweet, 0)
	tweetsByUser := make(map[string][]domain.Tweet)
	return &TweetManager{tweets, tweetsByUser, aTweeterWriter}
}

func (v *TweetManager) PublishTweet( aTweet domain.Tweet ) (int, error){
	id, valid := aTweet.IsValid()
	
	if valid==nil {
		v.tweets = append(v.tweets, aTweet)
		v.addTweetByUser( aTweet )
		v.writer.Write( aTweet )
	}
	return id, valid
}

func (v *TweetManager)addTweetByUser( aTweet domain.Tweet ){
	userTweets, exist := v.tweetsByUser[aTweet.GetUser()]
	if !exist{
		newTweetsByUser := []domain.Tweet{aTweet}
		v.tweetsByUser[aTweet.GetUser()]=newTweetsByUser
	}else{
		userTweets = append(userTweets, aTweet)
		v.tweetsByUser[aTweet.GetUser()]=userTweets
	}
}

func (v *TweetManager)GetTweetById(id int)(domain.Tweet){
	for i:=0 ; i < len(v.tweets); i++ {
		if v.tweets[i].GetId() == id{
			return v.tweets[i]
	    }
	}
	return nil
}

func (v *TweetManager)CountTweetsByUser(user string) (int){

	userTweets, exist := v.tweetsByUser[user]
	if exist{
		return len(userTweets)
	}
	return 0

}

func (v *TweetManager) GetLastTweet() (domain.Tweet){
	return v.tweets[len(v.tweets)-1]
}

func (v *TweetManager) GetTweets() ([]domain.Tweet){
	return v.tweets
}

func (v *TweetManager) GetTweetsByUser(user string) ( []domain.Tweet ){
	userTweets, exist := v.tweetsByUser[user]
	if exist{
		return userTweets
	}
	return nil
}

func (v *TweetManager) SearchTweetsContaining(query string, searchResult chan domain.Tweet) {

	go func(){
		for _, tweet := range v.tweets {
			if strings.Contains(tweet.GetText(), query) {
				searchResult <- tweet
				return
			}
		}
	}()
}