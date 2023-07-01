package api

import (
	"fmt"

	"github.com/antchfx/htmlquery"
	"github.com/eh-web-viewer/eh-web-viewer/mycurl"
	"golang.org/x/net/html"
)

type Torrent struct {
	Date     string `json:"date"`
	Size     string `json:"size"`
	Seed     string `json:"seed"`
	Peer     string `json:"peer"`
	Download string `json:"download"`
	Uploader string `json:"uploader"`
	Url      string `json:"url"`
	Error    string `json:"error"`
}

type Torrents struct {
	Query    string     `json:"query"`
	Torrents []*Torrent `json:"torrents"`
	Error    string     `json:"error"`
}

func parseNodeToTorrent(n *html.Node) (t *Torrent) {
	defer func() {
		if err := recover(); err != nil {
			t.Error = fmt.Sprintf("RECOVERED: %v\n", err)
			return
		}
	}()
	t = &Torrent{}

	date := htmlquery.FindOne(n, "tr[1]/td[1]")
	t.Date = htmlquery.InnerText(date)

	size := htmlquery.FindOne(n, "tr[1]/td[2]")
	t.Size = htmlquery.InnerText(size)

	seed := htmlquery.FindOne(n, "tr[1]/td[4]")
	t.Seed = htmlquery.InnerText(seed)

	peer := htmlquery.FindOne(n, "tr[1]/td[5]")
	t.Peer = htmlquery.InnerText(peer)

	download := htmlquery.FindOne(n, "tr[1]/td[5]")
	t.Download = htmlquery.InnerText(download)

	uploader := htmlquery.FindOne(n, "tr[2]/td[1]")
	t.Uploader = htmlquery.InnerText(uploader)

	url := htmlquery.FindOne(n, "tr[3]//a")
	t.Url = htmlquery.SelectAttr(url, "href")

	return
}

func queryTorrents(query string) (torrents *Torrents, err error) {
	torrents = &Torrents{}

	torrents.Query = query

	resp, err := mycurl.Fetch("GET", BASE_URL+query,
		map[string]string{"Cookie": COOKIE},
		nil)
	if err != nil {
		torrents.Error = err.Error()
		return
	}

	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		torrents.Error = err.Error()
		return
	}

	list, err := htmlquery.QueryAll(doc, "//tbody")
	if err != nil {
		torrents.Error = err.Error()
		return
	}

	torrents.Torrents = make([]*Torrent, len(list))
	for k, v := range list {
		torrents.Torrents[k] = parseNodeToTorrent(v)
	}

	return
}

func QueryTorrents(query string) (torrents *Torrents, err error) {
	return queryTorrents(query)
}
