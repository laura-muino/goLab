package service_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/golab/twitteer/src/domain"
	"github.com/golab/twitteer/src/service"
	"strings"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T){
	var tweet domain.Tweet;
	var tweetWriter service.TweetWriter
	tweetWriter =  service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	var tweet1user, tweet1text string = "Luciano", "HolaMundo!"
	tweet = domain.NewTextTweet(tweet1user, tweet1text)

	// Operation
	tweetManager.PublishTweet(tweet);

	publishedTweet := tweetManager.GetLastTweet()
	if publishedTweet.GetUser() != tweet1user &&
		publishedTweet.GetText() != tweet1text {
		t.Errorf("Expected tweet should be %s: %s \n but is %s: %s",
			tweet1user, tweet1text, publishedTweet.GetUser(), publishedTweet.GetText());
	}
	if publishedTweet.GetDate() == nil {
		t.Errorf("Excepted date can't be nil");
	}
}

func TestWithoutUserIsNotPublished(t *testing.T){
	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter =  service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	var tweet domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet( tweet )

	// validation
	if err != nil && err.Error() != "user is required"{
		t.Error("Expected error is user is requiered")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T){
	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter =  service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	var tweet domain.Tweet
	assert := assert.New(t)

	var text string
	user := "usuario"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet( tweet )

	assert.NotNil(err, "Expected error, but the tweet was published" )
	//assert.Equal(err != nil && err.Error() != "text is required",true, "Expected error is text is requiered")
	assert.Equal(err.Error(), "text is required")
}

func TestTweetWichExceeding140CharactersIsNotPublished(t *testing.T){
	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter =  service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	var tweet domain.Tweet
	assert := assert.New(t)

	text := strings.Repeat("a", 141)
	user := "usuario"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet( tweet )

	assert.NotNil(err, "Expected error, but the tweet was published")
	//assert.Equal(err != nil && err.Error() != "text exceding 140 characters",true, "Expected error is text exceding 140 characters")
	assert.Equal(err.Error(), "text exceding 140 characters")


}

func TestCanPublishAndretrieveMoreThanOneTweet(t *testing.T){
	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter =  service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	assert := assert.New(t)
	var tweet, secondTweet domain.Tweet // Fill the tweets

	var tweet1user, tweet1text string = "Luciano", "HolaMundo!"
	var tweet2user, tweet2text string = "Diana", "Hoy hice arroz!"
	tweet = domain.NewTextTweet(tweet1user, tweet1text)
	secondTweet = domain.NewTextTweet(tweet2user, tweet2text)


	// Operation
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)

	// Validation
	publishedTweets := tweetManager.GetTweets()
	assert.Equal(len(publishedTweets)!=2,false, "Expected size is 2, but was %d", len(publishedTweets) )

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]
	assert.Equal(isValidTweet(firstPublishedTweet, tweet1user, tweet1text), true, "first Published Tweet is invalid")
	assert.Equal(isValidTweet(secondPublishedTweet, tweet2user, tweet2text), true, "second Published Tweet is invalid")

}

func isValidTweet( tweet domain.Tweet, user string, text string) (bool){
	return (  tweet.GetUser() == user && tweet.GetText() == text  )
}


func TestCanRetrieveTweetById(t *testing.T){

	// Init
	assert := assert.New(t)
	var tweetWriter service.TweetWriter
	tweetWriter =  service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet1user, tweet1text string = "Luciano", "Soy un tweet con id!"
	tweet := domain.NewTextTweet(tweet1user, tweet1text)
	var id int

	// Operation
	id, _ = tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweetById(id)

	assert.Equal(isValidTweet(publishedTweet, tweet1user, tweet1text), true, "Published Tweet is invalid")
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T){
	// Init
	assert := assert.New(t)
	var tweetWriter service.TweetWriter
	tweetWriter =  service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	var tweet1user, tweet1text string = "Luciano", "Hello World!"
	var tweet2user, tweet2text string = "Daiana", "Hoy, Hice arroz!"
	var tweet3user, tweet3text string = "Liugi", "Its Me Mario!"
	var tweet4user, tweet4text string = "Luciano", "Hakuna Matatata"

	tweetManager.PublishTweet( domain.NewTextTweet(tweet1user, tweet1text) )
	tweetManager.PublishTweet( domain.NewTextTweet(tweet2user, tweet2text) )
	tweetManager.PublishTweet( domain.NewTextTweet(tweet3user, tweet3text) )
	tweetManager.PublishTweet( domain.NewTextTweet(tweet4user, tweet4text) )

	// Operation
	count := tweetManager.CountTweetsByUser(tweet1user)
	// Validation
	assert.Equal( tweetManager.CountTweetsByUser(tweet1user)==2, true,"Expected count is 2 but was %d", count )


	count = tweetManager.CountTweetsByUser(tweet2user)
	assert.Equal( tweetManager.CountTweetsByUser(tweet2user)==1, true,"Expected count is 2 but was %d", count )

	count = tweetManager.CountTweetsByUser(tweet3user)
	assert.Equal( tweetManager.CountTweetsByUser(tweet3user)==1, true,"Expected count is 2 but was %d", count )

}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T ){
	// Init
	assert := assert.New(t)
	var tweetWriter service.TweetWriter
	tweetWriter =  service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	var tweet1user, tweet1text string = "Luciano", "Hello World!"
	var tweet2user, tweet2text string = "Daiana", "Hoy, Hice arroz!"
	var tweet3user, tweet3text string = "Liugi", "Its Me Mario!"
	var tweet4user, tweet4text string = "Luciano", "Hakuna Matatata"
	var tweet5user, tweet5text string = "Luciano", "Hotel California Rocks!"

	tweetManager.PublishTweet( domain.NewTextTweet(tweet1user, tweet1text) )
	tweetManager.PublishTweet( domain.NewTextTweet(tweet2user, tweet2text) )
	tweetManager.PublishTweet( domain.NewTextTweet(tweet3user, tweet3text) )
	tweetManager.PublishTweet( domain.NewTextTweet(tweet4user, tweet4text) )
	tweetManager.PublishTweet( domain.NewTextTweet(tweet5user, tweet5text) )

	// Operation
	tweets := tweetManager.GetTweetsByUser(tweet1user)

	// Validation
	assert.Equal( len(tweets)==3, true, "Tweets len(%d) is not 2", len(tweets) )

}

func TestPublishedTweetIsSavedToExternalResources(t *testing.T){
	// Initialization
	assert := assert.New(t)
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweet // Fill the tweet with data
	var tweet1user, tweet1text string = "Luciano", "Hello World!"
	tweet = domain.NewTextTweet(tweet1user, tweet1text)

	// Operation

	id,_ := tweetManager.PublishTweet(tweet)

	// Validation
	memoryWriter := (tweetWriter).(*service.MemoryTweetWriter)
	savedTweet := memoryWriter.GetLastSavedTweet()

	assert.Equal(savedTweet!=nil,true, "Error, tweet was not saved")
	assert.Equal(savedTweet.GetId()==id,true, "Error, Tweet id do not match")
}


func TestCanSearchForTweetContainingText(t *testing.T){
	// Initialization
	assert := assert.New(t)
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)

	/// Create and Publish a Tweet
	var tweet1user, tweet1text string = "Luciano", "Hello World!"
	tweetManager.PublishTweet( domain.NewTextTweet(tweet1user, tweet1text) )

	// Operation
	searchResult := make(chan domain.Tweet)
	quit := make(chan int)
	query := "Hello"
	tweetManager.SearchTweetsContaining(query, searchResult)

	// Validation
	var foundTweet domain.Tweet
	select {
		case foundTweet = <-searchResult:
		case notFound := <-quit:
			print("notFound<%d>", notFound)
	}
	assert.Equal(foundTweet != nil, true, "Error, foundTweet is nil")
	assert.Equal(!strings.Contains(foundTweet.GetText(), query), false, "Error, foundTweet dont march with query")


		// if foundTweet == nil ..
		// if !strings.Contains(foundTweet.GetText(), query) ...
}


