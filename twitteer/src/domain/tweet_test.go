package domain_test

import (
	"github.com/golab/twitteer/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanGetAPrintableTweet(t *testing.T) {
	// Initialization
	assert := assert.New(t)
	tweet := domain.NewTweet("Luciano", "Hello World!")

	// Operation
	text := tweet.String()
	expectedText := "@Luciano: Hello World!"
	assert.Truef(text==expectedText, "Error, '"+text+"' != '"+expectedText+"'")
}

