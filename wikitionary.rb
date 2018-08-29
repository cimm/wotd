require 'rss'
require 'open-uri'
require 'sanitize'

class Wikitionary
  def initialize(feed)
    @feed = feed
  end

  def fetch
    response = open(@feed)
    RSS::Parser.parse(response)
  end

  def feed_items
    fetch.items
  end

  def word_of_the_day
    item = feed_items.last
    body = Sanitize.clean(item.description)
    parts = body.lines.reject do |line|
      line.strip.empty? ||
      line.start_with?(' edit') ||
      line.start_with?('← yesterday')
    end
    parts.map!{ |line| line.chomp.strip }
    word_and_type, *descriptions = parts
    type = word_and_type.split.last
    word = word_and_type.split[0..-2].join
    Struct::Wotd.new(word, type, descriptions, item.link)
  end
end
