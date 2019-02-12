# Wiktionary WOTD to Slack

A little Go script that posts the Wiktionary [Word of the Day](https://en.wiktionary.org/w/index.php?title=Wiktionary:Word_of_the_day) to a Slack room.

![WOTD Slack message](screenshot.png)

It uses Wiktionary’s [RSS feed](https://en.wiktionary.org/w/api.php?action=featuredfeed&feed=wotd) to extract the most recent entry, wraps the WOTD and its definition in a nicely formatted Slack message and posts it to the Slack room.

## Install

Create a [Slack webhook](https://api.slack.com/incoming-webhooks#posting_with_webhooks) and provide it to the script with the `hook` flag.

Run `wotd` from the download directory (or throw it anywhere in your `PATH`) and the most recent WOTD should show up in the selected Slack channel.

    wotd -hook https://hooks.slack.com/services/some/code/here

You can optionally run this as a crontab to be greeted by a new WOTD in Slack every day.

## Compile

Requires the [bluemonday](https://github.com/microcosm-cc/bluemonday) Go package to sanitize the complex Wiktionary HTML, get that first:

    go get -u github.com/microcosm-cc/bluemonday

Now compile the script from the project root directory:

    go build

## Why

Idea from my copywriter colleague who thought it would be fun to try to use a random WOTD in our discussions in a Slack thread where we were throwing randomly complex words at each other.

There is feature complete Ruby port available as will (see the [`ruby` branch](https://github.com/cimm/wotd/tree/ruby)). I wrote both as an exercise.
