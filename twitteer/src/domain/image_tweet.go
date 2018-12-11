package domain

type ImageTweet struct{
	TextTweet
	Url string `json: "url"`
}

func NewImageTweet(user, text, url string) (*ImageTweet){
	return &ImageTweet{ *NewTextTweet(user, text), url  };
}

func (it *ImageTweet)PrintableTweet() (string){
	return "@"+it.User+": "+it.Text+" "+it.Url
}

func (t *ImageTweet)String()(string){
	return t.PrintableTweet();
}