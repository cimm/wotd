require "optparse"
require "optparse/uri"

class Parser
  attr_reader :webhook

  def self.run
    parser = new.parse
    parser.validate
    parser
  end

  def parse
    @opt_parser = OptionParser.new do |opts|
      opts.on("", "--hook WEBHOOK", URI, "The webhook URL of the room the WOTD message should be pasted in. See https://api.slack.com/incoming-webhooks to learn how to create a new Slack webhook.") do |v|
        @webhook = v
      end
      opts.on("-h", "--help", "Show this help message") do
        puts opts
        exit
      end
    end
    @opt_parser.parse!
    self
  end

  def validate
    if @webhook.nil?
      STDERR.puts("err: missing webhook")
      puts @opt_parser.parse('-h')
      exit(1)
    end

    unless @webhook.kind_of?(URI::HTTP)
      STDERR.puts("err: invalid webhook")
      exit(1)
    end
  end
end
