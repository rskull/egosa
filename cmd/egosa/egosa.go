package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/dghubble/oauth1"
	"github.com/rskull/egosa/egosa"
	chatwork "github.com/rskull/go-chatwork"
	"github.com/rskull/go-twitter/twitter"
)

const version = "1.0.0"

type Sender struct {
	conf    *egosa.Config
	sinceId int64
}

func (s *Sender) run() {
	// Ouath
	oauthConfig := oauth1.NewConfig(s.conf.Twitter.ConsumerKey, s.conf.Twitter.ConsumerKeySecret)
	token := oauth1.NewToken(s.conf.Twitter.AuthKey, s.conf.Twitter.AuthKeySecret)
	httpClient := oauthConfig.Client(oauth1.NoContext, token)

	// Twitter client
	twitterClient := twitter.NewClient(httpClient)

	// Chatwork client
	chatworkClient := chatwork.NewClient(s.conf.Chatwork.ApiKey)

	tickChan := time.NewTicker(time.Second * time.Duration(s.conf.Core.IntervalSec)).C

	params := &twitter.SearchTweetsParams{
		Query:      s.conf.Twitter.SearchQuery,
		Count:      s.conf.Core.Count,
		ResultType: s.conf.Twitter.ResultType,
		Lang:       s.conf.Twitter.Lang,
		SinceID:    s.sinceId,
	}

	for {
		select {
		case <-tickChan:
			tweets, _, err := twitterClient.Searches.Tweets(params)

			if err != nil {
				log.Fatal(err)
			}

			if len(tweets.Statuses) > 0 {
				s.sinceId = tweets.Statuses[0].ID
			}

			for _, tweet := range tweets.Statuses {
				url := fmt.Sprintf("https://twitter.com/%s/status/%s\n", tweet.User.ScreenName, tweet.IDStr)
				message := fmt.Sprintf(s.conf.Chatwork.SendBody, tweet.User.Name, tweet.User.ScreenName, tweet.Text, url)
				chatworkClient.PostRoomMessage(s.conf.Chatwork.RoomID, message)
			}
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

	var conf egosa.Config
	if _, err := toml.DecodeFile(*confPath, &conf); err != nil {
		log.Fatal(err)
	}

	var sender = Sender{
		conf: &conf,
	}

	sender.run()
}
