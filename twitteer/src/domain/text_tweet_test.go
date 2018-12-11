package domain_test

import (
	"github.com/golab/twitteer/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTextTweetPrintUserAndText(t *testing.T) {
	// Initialization
	assert := assert.New(t)
	tweet := domain.NewTextTweet("Luciano", "Hello World!")

	// Operation
	text := tweet.PrintableTweet()
	expectedText := "@Luciano: Hello World!"
	assert.Truef(text==expectedText, "Error, '"+text+"' != '"+expectedText+"'")
}
