package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dghubble/oauth1"
	twitter "github.com/g8rswimmer/go-twitter/v2"
)

type authorize struct {
}

func (a authorize) Add(req *http.Request) {
}

func main() {
	var clientToken, clientSecret, accessToken, accessSecret string
	flag.StringVar(&clientToken, "client-token", os.Getenv("XPOST_CLIENT_TOKEN"), "Twitter ClientToken")
	flag.StringVar(&clientSecret, "client-secret", os.Getenv("XPOST_CLIENT_SECRET"), "Twitter ClientSecret")
	flag.StringVar(&accessToken, "access-token", os.Getenv("XPOST_ACCESS_TOKEN"), "Twitter AccessToken")
	flag.StringVar(&accessSecret, "access-secret", os.Getenv("XPOST_ACCESS_SECRET"), "Twitter AccessSecret")

	flag.Parse()

	client := &twitter.Client{
		Authorizer: authorize{},
		Client: oauth1.NewConfig(clientToken, clientSecret).Client(oauth1.NoContext, &oauth1.Token{
			Token:       accessToken,
			TokenSecret: accessSecret,
		}),
		Host: "https://api.twitter.com",
	}

	req := twitter.CreateTweetRequest{
		Text: strings.Join(flag.Args(), " "),
		/*
			Reply: &twitter.CreateTweetReply{
				InReplyToTweetID: entry.ID,
			},
		*/
	}
	_, err := client.CreateTweet(context.Background(), req)
	if err != nil {
		log.Println(err)
	}
}
