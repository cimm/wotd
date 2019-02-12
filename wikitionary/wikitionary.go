package wikitionary

import (
	"github.com/microcosm-cc/bluemonday"
	"strings"
)

const (
	endpoint = "https://en.wiktionary.org/w/api.php?action=featuredfeed&feed=wotd"
)

type Wotd struct {
	Word        string
	Definitions []string
	Url         string
}

// Fetch returns the WOTD and its definition from Wikitionary.
func Fetch() (wotd Wotd) {
	f := fetchFeed(endpoint)
	i := lastFeedItem(f)
	wotd.Word = textBetween(i.Description, "<span id=\"WOTD-rss-title\">", "</span>")
	description := textBetween(i.Description, "<div id=\"WOTD-rss-description\">", "</div>")
	wotd.Definitions = cleanDescription(description)
	wotd.Url = i.Guid
	return
}

// cleanDescription removes all HTML tags from the RSS description and splits
// it per line.
func cleanDescription(description string) []string {
	d := strings.TrimSpace(bluemonday.StrictPolicy().Sanitize(description))
	return strings.Split(d, "\n")
}

// textBetween returns the characters between the prefix and suffix in
// the provided text.
func textBetween(text, prefix, suffix string) string {
	rightSlice := strings.Split(text, prefix)[1]
	return strings.Split(rightSlice, suffix)[0]
}
