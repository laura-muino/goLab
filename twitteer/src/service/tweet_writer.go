package service

import "github.com/golab/twitteer/src/domain"

type TweetWriter interface{
	Write( tweet domain.Tweet )
}
