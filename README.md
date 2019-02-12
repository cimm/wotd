# Wiktionary WOTD to Slack

A little script that posts the Wiktionary [Word of the Day](https://en.wiktionary.org/w/index.php?title=Wiktionary:Word_of_the_day) to a Slack room.

![WOTD Slack message](screenshot.png)

It uses Wiktionary’s [RSS feed](https://en.wiktionary.org/w/api.php?action=featuredfeed&feed=wotd) to extract the most recent entry, wraps the WOTD and its definition in a nicely formatted Slack message and posts it to the Slack room.

## Implementations

There are two different feature identical implementations: one in [Go](https://github.com/cimm/wotd/tree/go) and another in [Ruby](https://github.com/cimm/wotd/tree/ruby). No reason really, only the fun of trying to implement the same program in two different languages. Turns out both implementations have the same single non-standard library extension: an HTML sanitizer. All other functionality was easily covered by the standard features of the language.
