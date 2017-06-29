package main

import (
	"fmt"

	chatwork "github.com/rskull/go-chatwork"
	"github.com/rskull/go-twitter/twitter"
)

type Chatwork struct {
	conf   *Config
	client *chatwork.Client
}

func newChatwork(conf *Config) *Chatwork {
	client := chatwork.NewClient(conf.Chatwork.ApiKey)
	return &Chatwork{
		conf:   conf,
		client: client,
	}
}

func (c *Chatwork) makeMessage(tweet twitter.Tweet) string {
	url := fmt.Sprintf("https://twitter.com/%s/status/%s\n", tweet.User.ScreenName, tweet.IDStr)
	message := fmt.Sprintf(c.conf.Chatwork.SendBody, tweet.User.Name, tweet.User.ScreenName, tweet.Text, url)
	return message
}

func (c *Chatwork) send(message string) {
	c.client.PostRoomMessage(c.conf.Chatwork.RoomID, message)
}
