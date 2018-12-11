package domain

import "time"

type Tweet interface{
	String() string
	PrintableTweet() string
	IsValid() (int, error)
	GetUser() string
	GetText() string
	GetDate() *time.Time
	GetDateString() string
	GetId() int
}
