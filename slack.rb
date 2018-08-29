require 'json'
require 'net/http'
require 'net/https'

class Slack
  def initialize(webhook)
    @webhook = URI.parse(webhook)
  end

  def post(text)
    https = Net::HTTP.new(@webhook.host, @webhook.port)
    https.use_ssl = true
    request = Net::HTTP::Post.new(@webhook.request_uri)
    request.body = { text: text }.to_json
    https.request(request)
  end
end
