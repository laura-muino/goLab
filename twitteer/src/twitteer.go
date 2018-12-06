package main

import (
	"github.com/abiosoft/ishell"
	"github.com/golab/twitteer/src/domain"
	"github.com/golab/twitteer/src/service"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Who you are?: ")

			user := c.ReadLine()

			c.Print("Write your tweet: ")

			tweet := c.ReadLine()

			//dateNow := time.Now().Local()
			var newTweet = domain.NewTweet(user, tweet)

			error := service.PublishTweet(newTweet)

			if error == nil {
				c.Print("Tweet sent\n")
			}else{
				c.Print("Error: %s", error)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetTweet()

			c.Println(tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context){
			defer c.ShowPrompt(true)
			tweet := service.GetTweets()
			c.Printf("%v", tweet)
		},
	})

	shell.Run()

}