#!/usr/bin/env ruby

require_relative "./wikitionary"
require_relative "./slack"

WOTD_FEED = 'https://en.wiktionary.org/w/api.php?action=featuredfeed&feed=wotd'
SLACK_WEBHOOK = 'COPY_YOUR_WEBHOOK_URL_HERE'

Struct.new('Wotd', :word, :type, :descriptions, :link)

wotd = Wikitionary.new(WOTD_FEED).word_of_the_day
text = [
  "*<#{wotd.link}|#{wotd.word}>* ",
  wotd.type,
  wotd.descriptions.map{ |line| "\n> #{line}" }.join
].join
Slack.new(SLACK_WEBHOOK).post(text)
