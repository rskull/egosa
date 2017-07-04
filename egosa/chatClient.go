package main

import (
	"github.com/rskull/go-twitter/twitter"
)

type ChatClient interface {
	send(message string)
	makeMessage(tweet twitter.Tweet) string
	isEnable() bool
}
