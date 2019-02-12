require 'rss'
require 'open-uri'
require 'sanitize'

FEED = 'https://en.wiktionary.org/w/api.php?action=featuredfeed&feed=wotd'

class Wikitionary
  def fetch
    response = open(FEED)
    RSS::Parser.parse(response)
  end

  def feed_items
    fetch.items
  end

  def self.word_of_the_day
    new.word_of_the_day
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
    word = word_and_type.split[0..-2].join(' ')
    Struct::Wotd.new(word, descriptions, item.link)
  end
end
