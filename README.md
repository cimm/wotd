# Slack WOTD

A little Ruby script that fetches the Word Of The Day from [Wikitionary](https://en.wiktionary.org/wiki/Wiktionary:Word_of_the_day), reformats it a little, and posts it to Slack.

## Install

It needs the [Sanitize](https://rubygems.org/gems/sanitize) Ruby gem to remove the Wikitionary HTML so install that first:

    gem install sanitize

Create a [Slack webhook](https://api.slack.com/incoming-webhooks#posting_with_webhooks) and set it as the `SLACK_WEBHOOK` in to the `wotd.rb` file.

Run `./wotd` from this directory and the most recent WOTD should show up in your selected Slack channel. You can optionally run this as a crontab to be greeted by a new WOTD in Slack every day.

## Why

Idea from my copywriter colleague who thought it would be fun to try to use a random WOTD in our discussions in a Slack thread where we were throwing randomly complex words at each other.
