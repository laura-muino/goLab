package main

import (
	"fmt"
	"strconv"

	//"github.com/abiosoft/ishell"
	"github.com/gin-gonic/gin"
	"github.com/golab/twitteer/src/domain"
	"github.com/golab/twitteer/src/service"
	"net/http"
)


//func funcionQueHaceGet(c * gin.Context){
//	parametro := c.Param("parametro")
//	c.String(http.StatusOK, "ok")
//}

var Manager = service.NewTweetManager(service.NewFileTweetWriter())

func main() {

	//tweetManager := service.NewTweetManager(service.NewFileTweetWriter())

	var newTweet, newTweet2 domain.Tweet
	newTweet = domain.NewTextTweet("usuario1", "Hola Mundo!")
	newTweet2 = domain.NewTextTweet("usuario2", "Hola Mundo!")
	_, error := Manager.PublishTweet(newTweet)
	Manager.PublishTweet(newTweet2)
	if error == nil {
		fmt.Print("Tweet sent\n")
	} else {
		fmt.Errorf("Error: %s", error)
	}

	tweets := Manager.GetTweets()
	for _,tweet := range tweets{
		fmt.Printf("%v\n", tweet)
		//tweet.PrintableTweet()
	}

	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/getTweets", getTweets)
		v1.GET("/getById/:id", getTweetById)
		v1.GET("/countTweetsByUser/:user", countTweetsByUser)
		v1.POST("/postTextTweet/", postTextTweet)
		v1.POST("/postImageTweet/", postImageTweet)

	}

	//router.POST("/unPost", funcionQueHacePost)

	router.Run(":8080")
}


/*type TextTweet struct {
	User string `json: "user"`
	Text string `json: "text"`
}

type ImageTweet struct {
	TextTweet
	Url string `json: "url"`
}*/

func getTweets(c *gin.Context){
	tweets := Manager.GetTweets()

	if tweets == nil{
		c.String(http.StatusBadRequest, "There are no tweets")
		return
	}
	c.JSON(http.StatusOK, tweets)
}

/*func postQuotedTextTweet(c *gin.Context) {
	var tweet *domain.QuoteTweet

	if err := c.ShouldBindJSON(&tweet); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var tweetToPublish domain.Tweet
	quotedTextTweet = domain.NewImageTweet(tweet.User, tweet.Text, tweet.Url)
	_, error := Manager.PublishTweet(tweetToPublish)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error})
		return
	}
	c.JSON(http.StatusOK, gin.H{ "User": tweet.User, "Text": tweet.Text, "Url": tweet.Url})

}*/

func postImageTweet(c *gin.Context){
	var tweet *domain.ImageTweet
	if err := c.ShouldBindJSON(&tweet); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var tweetToPublish domain.Tweet
	tweetToPublish = domain.NewImageTweet(tweet.User, tweet.Text, tweet.Url)
	_, error := Manager.PublishTweet(tweetToPublish)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error})
		return
	}
	c.JSON(http.StatusOK, gin.H{ "User": tweet.User, "Text": tweet.Text, "Url": tweet.Url})
}

func postTextTweet(c *gin.Context){
	var tweet *domain.TextTweet
	if err := c.ShouldBindJSON(&tweet); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var tweetToPublish domain.Tweet
	tweetToPublish = domain.NewTextTweet(tweet.User, tweet.Text)
	_, error := Manager.PublishTweet(tweetToPublish)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error})
		return
	}
	c.JSON(http.StatusOK, gin.H{ "User": tweet.User, "Text": tweet.Text})
}

func countTweetsByUser(c *gin.Context){
	count := Manager.CountTweetsByUser(c.Param("user"))
	c.String(http.StatusOK, strconv.Itoa(count))

}

func getTweetById(c * gin.Context){
	id, error := strconv.Atoi(c.Param("id"))
	if(error!=nil){
		c.String(http.StatusBadRequest, "BadRequest")
		return
	}

	fmt.Printf("getTweetById: idparam<%s>, id<%d>", c.Param("id"), id)
	resultTweet := Manager.GetTweetById(id)

	if resultTweet==nil{
		c.String(http.StatusBadRequest, "BadRequest, id do not exist")
		return
	}

	c.JSON(http.StatusOK, resultTweet )

}

	/*tweetManager := service.NewTweetManager(service.NewFileTweetWriter())
	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")



	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("User: ")
			user := c.ReadLine()

			c.Print("---- publishTweet ---- \n")
			c.Print("\t 1) Publish a Tweet\n")
			c.Print("\t 2) Publish a Tweet with Image\n")
			c.Print("\t 3) Response to anotherTweet\n")

			c.Print("Option: ")
			option := c.ReadLine()

			var newTweet domain.Tweet
			switch option{
			case "1":
				c.Print("Write your tweet: ")
				tweet := c.ReadLine()
				newTweet = domain.NewTextTweet(user, tweet)

			case "2":
				c.Print("Write your tweet: ")
				tweet := c.ReadLine()
				c.Print("Url: ")
				url := c.ReadLine()
				newTweet = domain.NewImageTweet(user, tweet, url)
			case "3":
				c.Print("Who to reply ?")
				userToReply := c.ReadLine()
				tweetsByUser := tweetManager.GetTweetsByUser(userToReply)
				if tweetsByUser != nil  {
					for index,tweet := range tweetsByUser{
						c.Printf("%d: %v\n", index, tweet)
					}
					c.Print("choose your tweet to reply: ")
					var indexToReply int
					_, err := fmt.Scanf("%d", &indexToReply)
					if indexToReply <= len(tweetsByUser) && err == nil {
						c.Print("reply: ")
						reply := c.ReadLine()
						newTweet = domain.NewQuoteTweet(user, reply, tweetsByUser[indexToReply])
					}else{
						c.Print("Incorrect option of tweet to reply")
						return
					}

				}else{
					c.Print("There is tweets from that user")
					return
				}
			default:
				c.Print("Invalid Option")
				return
			}

			_, error := tweetManager.PublishTweet(newTweet)

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

			tweet := tweetManager.GetLastTweet()
			c.Println(tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context){
			defer c.ShowPrompt(true)
			tweets := tweetManager.GetTweets()
			for _,tweet := range tweets{
				//c.Printf("%v\n", tweet)
				tweet.PrintableTweet()
			}
		},
	})

	shell.Run()

}*/