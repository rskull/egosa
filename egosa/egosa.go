package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/dghubble/oauth1"
	"github.com/rskull/go-twitter/twitter"
)

const version = "2.0.1"

type Sender struct {
	chatClients   *ChatClients
	twitterClient *twitter.Client
	sinceId       int64
}

func arrayContains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func newSender(chatClients *ChatClients, twitterClient *twitter.Client) *Sender {
	return &Sender{
		chatClients:   chatClients,
		twitterClient: twitterClient,
	}
}

func (s *Sender) run() {

	params := &twitter.SearchTweetsParams{
		Query:      ConfEgosa.Twitter.SearchQuery,
		Count:      ConfEgosa.Core.Count,
		ResultType: ConfEgosa.Twitter.ResultType,
		Lang:       ConfEgosa.Twitter.Lang,
		SinceID:    s.sinceId,
	}

	tweets, _, err := s.twitterClient.Searches.Tweets(params)

	if err != nil {
		log.Fatal(err)
	}

	if len(tweets.Statuses) > 0 {
		s.sinceId = tweets.Statuses[0].ID
	}

	log.Println("Egosa: send tweets")

	for _, tweet := range tweets.Statuses {
		if arrayContains(ConfEgosa.Twitter.IgnoreUsers, tweet.User.ScreenName) == false {
			s.chatClients.Tweet <- tweet
		}
	}
}

func main() {
	versionPrinted := flag.Bool("v", false, "egosa version")
	confPath := flag.String("c", "", "configuration file path for egosa")
	flag.Parse()
	if *versionPrinted {
		fmt.Printf("egosa version %s\n", version)
		return
	}

	var conf Config
	if _, err := toml.DecodeFile(*confPath, &conf); err != nil {
		log.Fatal(err)
	}

	ConfEgosa = conf

	log.Println("Egosa start")

	// Twitter Client
	oauthConfig := oauth1.NewConfig(ConfEgosa.Twitter.ConsumerKey, ConfEgosa.Twitter.ConsumerKeySecret)
	token := oauth1.NewToken(ConfEgosa.Twitter.AuthKey, ConfEgosa.Twitter.AuthKeySecret)
	httpClient := oauthConfig.Client(oauth1.NoContext, token)
	twitterClient := twitter.NewClient(httpClient)

	// Chat Clients
	chatClients := newChatClients()

	// Sender
	sender := newSender(chatClients, twitterClient)

	// Timer
	tickChan := time.NewTicker(time.Second * time.Duration(ConfEgosa.Core.IntervalSec)).C

	go chatClients.run()

	for {
		select {
		case <-tickChan:
			sender.run()
		}
	}
}
