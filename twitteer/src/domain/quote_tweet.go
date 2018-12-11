package domain

type QuoteTweet struct{
	TextTweet
	QuotedTweet Tweet
}

func NewQuoteTweet(user string, text string, tweet Tweet) (*QuoteTweet){
	return &QuoteTweet{ *NewTextTweet(user, text), tweet };
}

func (qt *QuoteTweet)PrintableTweet() (string){
	return "@"+qt.User+": "+qt.Text+" \""+qt.QuotedTweet.PrintableTweet()+"\""
}

func (t *QuoteTweet)String()(string){
	return t.PrintableTweet();
}