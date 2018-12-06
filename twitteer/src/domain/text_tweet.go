package domain

import (
	"errors"
	"time"
)

type TextTweet struct {
	Id int
	User string
	Text string
	Date *time.Time
}

var id int

func NewTextTweet(user, text string) (*TextTweet){
	nowDate := time.Now().Local();
	id++
	return &TextTweet{ id,  user, text, &nowDate };
}

func (t *TextTweet)String()(string){
	return "@"+t.User+": "+t.Text+"\n"
}

func (t *TextTweet) IsValid() (int, error){
	if t.Text == "" {
		return 0, errors.New("text is required")
	}

	if t.User == "" {
		return 0, errors.New("user is required")
	}

	if len(t.Text) > 140 {
		return 0, errors.New("text exceding 140 characters")
	}
	return t.GetId(), nil
}

func (t *TextTweet) GetDate() (*time.Time){
	return t.Date
}

func (t *TextTweet) GetDateString() (string){
	return t.Date.String()
}
func (t *TextTweet) GetUser() (string){
	return t.User
}
func (t *TextTweet) GetText() (string){
	return t.Text
}
func (t *TextTweet) GetId() (int){
	return t.Id
}