package wikitionary

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	userAgent = "WotdBot/0.1 (+https://github.com/cimm/wotd)"
)

type Rss struct {
	ItemList []Item `xml:"channel>item"`
}

type Item struct {
	Description string `xml:"description"`
	Guid        string `xml:"guid"`
}

// fetchFeed loads an RSS feed form a given url. It only parses the RSS item
// descriptions and guids, we are not interested in the other elements.
func fetchFeed(url string) Rss {
	c := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	req.Header.Set("User-Agent", userAgent)

	res, err := c.Do(req)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	rss := Rss{}
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	return rss
}

// lastFeedItem returns the last item from an RSS feed.
func lastFeedItem(r Rss) Item {
	return r.ItemList[len(r.ItemList)-1]
}
