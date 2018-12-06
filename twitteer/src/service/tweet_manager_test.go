package service_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/golab/twitteer/src/domain"
	"github.com/golab/twitteer/src/service"
	"strings"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T){
	var tweet *domain.Tweet;
	service.InitializeService()
	user := "grupoesfera";
	text := "This is my first tweet";
	tweet = domain.NewTweet(user, text);

	// Operation
	service.PublishTweet(tweet);

	publishedTweet := service.GetTweet()
	if publishedTweet.User != user &&
		publishedTweet.Text != text {
		t.Errorf("Expected tweet should be %s: %s \n but is %s: %s",
			user, text, publishedTweet.User, publishedTweet.Text);
	}
	if publishedTweet.Date == nil {
		t.Errorf("Excepted date can't be nil");
	}
}

func TestWithoutUserIsNotPublished(t *testing.T){
	// Initialization
	service.InitializeService()
	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = service.PublishTweet( tweet )

	// validation
	if err != nil && err.Error() != "user is required"{
		t.Error("Expected error is user is requiered")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T){
	// Initialization
	service.InitializeService()
	var tweet *domain.Tweet
	assert := assert.New(t)

	var text string
	user := "usuario"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = service.PublishTweet( tweet )

	assert.NotNil(err, "Expected error, but the tweet was published" )
	//assert.Equal(err != nil && err.Error() != "text is required",true, "Expected error is text is requiered")
	assert.Equal(err.Error(), "text is required")
}

func TestTweetWichExceeding140CharactersIsNotPublished(t *testing.T){
	// Initialization
	service.InitializeService()
	var tweet *domain.Tweet
	assert := assert.New(t)

	text := strings.Repeat("a", 141)
	user := "usuario"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = service.PublishTweet( tweet )

	assert.NotNil(err, "Expected error, but the tweet was published")
	//assert.Equal(err != nil && err.Error() != "text exceding 140 characters",true, "Expected error is text exceding 140 characters")
	assert.Equal(err.Error(), "text exceding 140 characters")


}

func TestCanPublishAndretrieveMoreThanOneTweet(t *testing.T){
	// Initialization
	service.InitializeService()
	assert := assert.New(t)
	var tweet, secondTweet *domain.Tweet // Fill the tweets

	var tweet1user, tweet1text string = "Luciano", "HolaMundo!"
	var tweet2user, tweet2text string = "Diana", "Hoy hice arroz!"
	tweet = domain.NewTweet(tweet1user, tweet1text)
	secondTweet = domain.NewTweet(tweet2user, tweet2text)


	// Operation
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)

	// Validation
	publishedTweets := service.GetTweets()
	assert.Equal(len(publishedTweets)!=2,false, "Expected size is 2, but was %d", len(publishedTweets) )

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]
	assert.Equal(isValidTweet(firstPublishedTweet, tweet1user, tweet1text), true, "first Published Tweet is invalid")
	assert.Equal(isValidTweet(secondPublishedTweet, tweet2user, tweet2text), true, "second Published Tweet is invalid")

}

func isValidTweet( tweet *domain.Tweet, user string, text string) (bool){
	return (  tweet.User == user && tweet.Text == text  )
}


func TestCanRetrieveTweetById(t *testing.T){

	// Init
	assert := assert.New(t)
	service.InitializeService()

	var tweet1user, tweet1text string = "Luciano", "Soy un tweet con id!"
	tweet := domain.NewTweet(tweet1user, tweet1text)
	var id int

	// Operation
	id, _ = service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweetById(id)

	assert.Equal(isValidTweet(publishedTweet, tweet1user, tweet1text), true, "Published Tweet is invalid")
}



func TestCanCountTheTweetsSentByAnUser(t *testing.T){
	// Init
	assert := assert.New(t)
	service.InitializeService()
	var tweet1user, tweet1text string = "Luciano", "Hello World!"
	var tweet2user, tweet2text string = "Daiana", "Hoy, Hice arroz!"
	var tweet3user, tweet3text string = "Liugi", "Its Me Mario!"
	var tweet4user, tweet4text string = "Luciano", "Hakuna Matatata"

	service.PublishTweet( domain.NewTweet(tweet1user, tweet1text) )
	service.PublishTweet( domain.NewTweet(tweet2user, tweet2text) )
	service.PublishTweet( domain.NewTweet(tweet3user, tweet3text) )
	service.PublishTweet( domain.NewTweet(tweet4user, tweet4text) )

	// Operation
	count := service.CountTweetsByUser(tweet1user)
	// Validation
	assert.Equal( service.CountTweetsByUser(tweet1user)==2, true,"Expected count is 2 but was %d", count )


	count = service.CountTweetsByUser(tweet2user)
	assert.Equal( service.CountTweetsByUser(tweet2user)==1, true,"Expected count is 2 but was %d", count )

	count = service.CountTweetsByUser(tweet3user)
	assert.Equal( service.CountTweetsByUser(tweet3user)==1, true,"Expected count is 2 but was %d", count )

}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T ){
	// Init
	assert := assert.New(t)
	service.InitializeService()
	var tweet1user, tweet1text string = "Luciano", "Hello World!"
	var tweet2user, tweet2text string = "Daiana", "Hoy, Hice arroz!"
	var tweet3user, tweet3text string = "Liugi", "Its Me Mario!"
	var tweet4user, tweet4text string = "Luciano", "Hakuna Matatata"
	var tweet5user, tweet5text string = "Luciano", "Hotel California Rocks!"

	service.PublishTweet( domain.NewTweet(tweet1user, tweet1text) )
	service.PublishTweet( domain.NewTweet(tweet2user, tweet2text) )
	service.PublishTweet( domain.NewTweet(tweet3user, tweet3text) )
	service.PublishTweet( domain.NewTweet(tweet4user, tweet4text) )
	service.PublishTweet( domain.NewTweet(tweet5user, tweet5text) )

	// Operation
	tweets := service.GetTweetsByUser(tweet1user)

	// Validation
	assert.Equal(  len(tweets)==3, true, "Tweets len(%d) is not 2", len(tweets) )





}
