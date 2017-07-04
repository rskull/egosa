package main

import (
	"fmt"

	chatwork "github.com/rskull/go-chatwork"
	"github.com/rskull/go-twitter/twitter"
)

type Chatwork struct {
	client *chatwork.Client
}

func newChatwork() *Chatwork {
	client := chatwork.NewClient(ConfEgosa.Chatwork.ApiKey)
	return &Chatwork{
		client: client,
	}
}

func (c *Chatwork) makeMessage(tweet twitter.Tweet) string {
	url := fmt.Sprintf("https://twitter.com/%s/status/%s\n", tweet.User.ScreenName, tweet.IDStr)
	message := fmt.Sprintf(ConfEgosa.Chatwork.SendBody, tweet.User.Name, tweet.User.ScreenName, tweet.Text, url)
	return message
}

func (c *Chatwork) send(message string) {
	c.client.PostRoomMessage(ConfEgosa.Chatwork.RoomID, message)
}
