package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {
	// Initialization
	assert := assert.New(t)
	tweet := NewImageTweet("grupoesfera", "This is my image",
		"http://www.grupoesfera.com.ar/common/img/grupoesfera.png")
	// Operation
	text := tweet.PrintableTweet()
	// Validation
	expectedText := "@grupoesfera: This is my image http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
	assert.True(text==expectedText, "Result error: "+text+" != "+expectedText)

}

