package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {
	// Initialization
	assert := assert.New(t)
	quotedTweet := NewTextTweet("grupoesfera", "This is my tweet")
	tweet := NewQuoteTweet("nick", "Awesome", quotedTweet )

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@nick: Awesome \"@grupoesfera: This is my tweet\""
	assert.True(text==expectedText, "Result error: "+text+" != "+expectedText)

}