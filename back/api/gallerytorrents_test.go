package api

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/antchfx/htmlquery"
	"github.com/eh-web-viewer/eh-web-viewer/mycurl"
)

func Test_Torrents(t *testing.T) {
	mycurl.SetClient("proxy")
	query := "gallerytorrents.php?gid=2558455&t=e86edf3a07"
	torrents := &Torrents{}

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

	j, _ := json.Marshal(torrents)
	fmt.Println(string(j))
}
