package api

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/eh-web-viewer/eh-web-viewer/mycurl"
	"golang.org/x/net/html"
)

type GallerySummary struct {
	// Query       string `json:"query"`        // the original query string
	Preview string `json:"preview"` // the cover url
	Title   string `json:"title"`   // title in english
	// OriginTitle string `json:"origin_title"` // title in japanese
	// Tags        []string `json:"tags"`         // language:chinese ...
	Category string `json:"category"` // doujin, manga ...
	Rate     string `json:"rate"`     // stars from 1 to 10
	Date     string `json:"date"`     // publish date
	Pages    string `json:"pages"`    // how many pages are there in gallery
	// Seeds    []string `json:"seeds"`    // url to the seeds link
	Seeds string `json:"seeds"` // url to the seeds link
	Url   string `json:"url"`   // url to the gallery page
	Error string `json:"error"` // error

	// Comments []*Comment `json:"comments"` //comments of the gallery
}
type Index struct {
	Query     string            `json:"query"`     // which query it executes
	Results   string            `json:"results"`   // how many results
	NextPage  string            `json:"next_page"` // url to next page
	PrevPage  string            `json:"prev_page"` // url to previous page
	Galleries []*GallerySummary `json:"galleries"` // search results
	Error     string            `json:"error"`     // error
}

var (
	reResults  = regexp.MustCompile(`([,+\d]+) results.`)
	rePrevPage = regexp.MustCompile(`var prevurl="https://exhentai.org/?(.*)";`)
	reNextPage = regexp.MustCompile(`var nexturl="https://exhentai.org/?(.*)";`)
	// reGalleries = regexp.MustCompile(``)
)

func parseIndexNodeToSummary(n *html.Node) (gs *GallerySummary) {
	defer func() {
		if err := recover(); err != nil {
			gs.Error = fmt.Sprintf("RECOVERED: %v\n", err)
			return
		}
	}()
	gs = &GallerySummary{}

	url := htmlquery.FindOne(n, "//a")
	gs.Url = htmlquery.SelectAttr(url, "href")

	name := htmlquery.FindOne(n, "//div[@class='gl4t glname glink']")
	gs.Title = htmlquery.InnerText(name)

	img := htmlquery.FindOne(n, "//img")
	gs.Preview = htmlquery.SelectAttr(img, "src")

	gl5t := htmlquery.FindOne(n, "//div[@class='gl5t']")
	gl5t_1 := htmlquery.FindOne(gl5t, "div[1]")
	category := htmlquery.FindOne(gl5t_1, "//div[contains(@class, 'cs')]")
	gs.Category = htmlquery.InnerText(category)

	date := htmlquery.FindOne(gl5t_1, "div[2]")
	gs.Date = htmlquery.InnerText(date)

	gl5t_2 := htmlquery.FindOne(gl5t, "div[2]")
	rate := htmlquery.FindOne(gl5t_2, "div[1]")
	gs.Rate = htmlquery.SelectAttr(rate, "style")

	pages := htmlquery.FindOne(gl5t_2, "div[2]")
	gs.Pages = htmlquery.InnerText(pages)

	seed := htmlquery.FindOne(gl5t_2, "//img")
	// gs.Seeds = []string{htmlquery.SelectAttr(seed, "src")}
	gs.Seeds = htmlquery.SelectAttr(seed, "src")
	// Error       string   `json:"error"`        // error

	return
}

// query: should contain '?'
func queryIndex(query string) (index *Index, err error) {
	index = &Index{}

	index.Query = query

	resp, err := mycurl.Fetch("GET", BASE_URL+query,
		map[string]string{"Cookie": COOKIE},
		nil)
	if err != nil {
		index.Error = err.Error()
		return
	}

	respText, err := io.ReadAll(resp.Body)
	if err != nil {
		index.Error = err.Error()
		return
	}

	respStr := string(respText)
	index.Results = (reResults.FindString(respStr))
	index.NextPage = (reNextPage.FindString(respStr))
	index.PrevPage = (rePrevPage.FindString(respStr)) // 麻了，golang的re包太麻了，传回去js里面处理吧。

	doc, err := htmlquery.Parse(strings.NewReader(respStr))
	if err != nil {
		index.Error = err.Error()
		return
	}
	list, err := htmlquery.QueryAll(doc, "//div[@class='gl1t']")
	if err != nil {
		index.Error = err.Error()
		return
	}
	index.Galleries = make([]*GallerySummary, len(list))
	for k, v := range list {
		index.Galleries[k] = parseIndexNodeToSummary(v)
	}

	return
}

func QueryIndex(query string) (index *Index, err error) {
	return queryIndex(query)
}
