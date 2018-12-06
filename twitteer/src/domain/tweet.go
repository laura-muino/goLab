package domain

import "time"

type Tweet struct {
	Id int
	User string;
	Text string;
	Date *time.Time;
}

var id int

func NewTweet(user, text string) (*Tweet){
	nowDate := time.Now().Local();
	id++
	return &Tweet{ id,  user, text, &nowDate };

}

func (t *Tweet)PrintableTweet()(string){
	return "@"+t.User+": "+t.Text+"\n"
}

func (t *Tweet)String()(string){
	return t.PrintableTweet()
}