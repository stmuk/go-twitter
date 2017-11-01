package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

var CONSUMER_KEY = os.Getenv("CONSUMER_KEY")
var CONSUMER_SECRET = os.Getenv("CONSUMER_SECRET")
var ACCESS_TOKEN = os.Getenv("ACCESS_TOKEN")
var ACCESS_TOKEN_SECRET = os.Getenv("ACCESS_TOKEN_SECRET")

func main() {
	anaconda.SetConsumerKey(CONSUMER_KEY)
	anaconda.SetConsumerSecret(CONSUMER_SECRET)
	api := anaconda.NewTwitterApi(ACCESS_TOKEN, ACCESS_TOKEN_SECRET)

	// people I follow ("friends")
	result := api.GetFriendsListAll(nil)
	i := 0

	for page := range result {
		if i == 2 {
			return
		}

		friends := page.Friends

		for _, friend := range friends {
			fmt.Printf("name: %s screenname: %s\n", friend.Name, friend.ScreenName)

		}

		if page.Error != nil {
			log.Fatal(page.Error) // XXX
		}

		if page.Friends == nil || len(page.Friends) == 0 {
			log.Fatal("err")
		}
		i++
	}
}
