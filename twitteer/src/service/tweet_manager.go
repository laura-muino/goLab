package service

import (
	"errors"
	"github.com/golab/twitteer/src/domain"
)

//var tweet *domain.Tweet

var tweets []*domain.Tweet
var tweetsByUser map[string][]*domain.Tweet

func InitializeService(){
	tweets = make([]*domain.Tweet, 0)
	tweetsByUser = make(map[string][]*domain.Tweet)
}

func PublishTweet( aTweet *domain.Tweet ) (int, error){

	if aTweet.Text == "" {
		return 0, errors.New("text is required")
	}
	if aTweet.User == "" {
		return 0, errors.New("user is required")
	}

	if len(aTweet.Text) > 140 {
		return 0, errors.New("text exceding 140 characters")
	}

	tweets = append(tweets, aTweet)

	addTweetByUser( aTweet )

	return aTweet.Id, nil
}

func addTweetByUser( aTweet *domain.Tweet ){
	userTweets, exist := tweetsByUser[aTweet.User]
	if !exist{
		newTweetsByUser := []*domain.Tweet{aTweet}
		tweetsByUser[aTweet.User]=newTweetsByUser
	}else{
		userTweets = append(userTweets, aTweet)
		tweetsByUser[aTweet.User]=userTweets
	}
}

func GetTweetById(id int)(*domain.Tweet){
	for i:=0 ; i < len(tweets); i++ {
		if tweets[i].Id == id{
			return tweets[i]
	    }
	}
	return nil
}

func CountTweetsByUser(user string) (int){
/*	var count int
	for _, aTweet := range tweets{
		if aTweet.User == user{
			count++
	}
	}
	return count*/

	userTweets, exist := tweetsByUser[user]
	if exist{
		return len(userTweets)
	}
	return 0

}

func GetTweet() (*domain.Tweet){
	return tweets[len(tweets)-1]
}

func GetTweets() ([]*domain.Tweet){
	return tweets
}

func GetTweetsByUser(user string) ( []*domain.Tweet ){
	userTweets, exist := tweetsByUser[user]
	if exist{
		return userTweets
	}
	return nil
}