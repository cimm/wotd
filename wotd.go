package main

import (
	"flag"
	"github.com/cimm/wotd/slack"
	"github.com/cimm/wotd/wikitionary"
	"net/url"
	"log"
)

func main() {
	wotd := wikitionary.Fetch()
	message := slack.FormatWotd(wotd)
	slack.Post(webhook(), message)
}

func webhook() string {
	hookPtr := flag.String("hook", "", "The webhook URL of the room the WOTD message should be posted in. See https://api.slack.com/incoming-webhooks to learn how to create a new Slack webhook.")
	flag.Parse()
	_, err := url.ParseRequestURI(*hookPtr)
	if err != nil {
    log.Fatalf("err: not a valid Slack webhook")
	}
	return *hookPtr
}
