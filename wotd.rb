#!/usr/bin/env ruby

require_relative "./parser"
require_relative "./slack"
require_relative "./wikitionary"

Struct.new('Wotd', :word, :descriptions, :link)

parser = Parser.run
wotd = Wikitionary.word_of_the_day
text = [
  "*<#{wotd.link}|#{wotd.word}>* ",
  wotd.descriptions.map{ |line| "\n> #{line}" }.join
].join
Slack.new(parser.webhook).post(text)
