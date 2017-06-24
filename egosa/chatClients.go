package main

import (
	"github.com/rskull/go-twitter/twitter"
)

type ChatClients struct {
	conf    *Config
	clients []ChatClient
	Tweet   chan twitter.Tweet
}

func newChatClients(conf *Config) *ChatClients {
	return &ChatClients{
		conf: conf,
		clients: []ChatClient{
			newChatwork(conf),
		},
		Tweet: make(chan twitter.Tweet),
	}
}

func (c *ChatClients) run() {
	for {
		select {
		case tweet := <-c.Tweet:
			go func(tweet twitter.Tweet) {
				for _, client := range c.clients {
					message := client.makeMessage(tweet)
					client.send(message)
				}
			}(tweet)
		}
	}
}