# WOTD Slack webhook

A little Ruby script that fetches the Word Of The Day from [Wikitionary](https://en.wiktionary.org/wiki/Wiktionary:Word_of_the_day), reformats it a little, and posts it to Slack.

## Install

It needs the [Sanitize](https://rubygems.org/gems/sanitize) Ruby gem to remove the Wikitionary HTML so install that first:

> gem install sanitize

Create a [Slack webhook](https://api.slack.com/incoming-webhooks#posting_with_webhooks) and set it as the the `SLACK_WEBHOOK` in to the `wotd.rb` file.

Run `./wotd` from this directory and the most recent WOTD should show up in your selected Slack channel. You can optionally run this as a crontab to have be greeted by a new WOTD in Slack every day.
