package api

import (
	"regexp"

	"github.com/antchfx/htmlquery"
	"github.com/eh-web-viewer/eh-web-viewer/mycurl"
	"golang.org/x/net/html"
)

var (
	rePreview = regexp.MustCompile(`https://s.exhentai.org[-\.\/a-zA-Z0-9_]+`)
)

type Comment struct {
	Username  string `json:"username"`
	Date      string `json:"date"`
	Comment   string `json:"comment"`
	Score     string `json:"score"`
	ScoreLogs string `json:"score_logs"`
}

type GalleryPreview struct {
	Query string `json:"query"` // the original query string

	Preview     string              `json:"preview"`      // the cover url
	Title       string              `json:"title"`        // title in english
	OriginTitle string              `json:"origin_title"` // title in japanese
	Tags        map[string][]string `json:"tags"`         // language:chinese ...
	Category    string              `json:"category"`     // doujin, manga ...
	// Rate        string   `json:"rate"`         // stars from 1 to 10
	Date  string `json:"date"`  // publish date
	Pages string `json:"pages"` // how many pages are there in gallery
	// Seeds       []string `json:"seeds"`        // url to the seeds link
	// Url         string   `json:"url"`          // url to the gallery page

	Images []string `json:"images"` // url of previewing pics
	Urls   []string `json:"urls"`   // url of pics

	Comments []*Comment `json:"comments"` //comments of the gallery

	Error string `json:"error"` // error
}

// func queryGallerySummary(query string) (gallerySummary *GallerySummary, err error) {
// 	return
// }

func parseTags(doc *html.Node) (m map[string][]string, err error) {
	list, err := htmlquery.QueryAll(doc, "//div[@id='taglist']/table/tbody/tr")
	if err != nil {
		return
	}
	m = make(map[string][]string)
	for _, v := range list {
		key, value, err := parseTagTr(v)
		if err != nil {
			return nil, err
		}
		m[key] = value
	}
	return
}

func parseTagTr(doc *html.Node) (k string, v []string, err error) {
	list, err := htmlquery.QueryAll(doc, "//td")
	if err != nil {
		return
	}
	k = htmlquery.InnerText(list[0])

	tags, err := htmlquery.QueryAll(list[1], "//a")
	if err != nil {
		return
	}
	v = make([]string, len(tags))
	for i, tag := range tags {
		v[i] = htmlquery.InnerText(tag)
	}
	return
}

func queryAllStringList(doc *html.Node, xpath, attr string) (l []string, err error) {
	list, err := htmlquery.QueryAll(doc, xpath)
	if err != nil {
		return
	}
	l = make([]string, len(list))
	for k, v := range list {
		l[k] = htmlquery.SelectAttr(v, attr)
	}
	return
}

func queryGalleryPreview(query string) (galleryPreview *GalleryPreview, err error) {
	galleryPreview = &GalleryPreview{}

	galleryPreview.Query = query

	resp, err := mycurl.Fetch("GET", BASE_URL+query,
		map[string]string{"Cookie": COOKIE},
		nil)
	if err != nil {
		galleryPreview.Error = err.Error()
		return
	}
	if resp.StatusCode != 200 {
		galleryPreview.Error = resp.Status
		return
	}

	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		galleryPreview.Error = err.Error()
		return
	}

	galleryPreview.Preview, err = findOneAndSelectAttr(doc, "//div[@id='gd1']/div", "style")
	galleryPreview.Preview = rePreview.FindString(galleryPreview.Preview)
	if err != nil {
		galleryPreview.Error += err.Error() + "\n"
	}
	galleryPreview.Title, err = findOneAndSelectAttr(doc, "//h1[@id='gn']", InnerText)
	if err != nil {
		galleryPreview.Error += err.Error() + "\n"
	}
	galleryPreview.OriginTitle, err = findOneAndSelectAttr(doc, "//h1[@id='gj']", InnerText)
	if err != nil {
		galleryPreview.Error += err.Error() + "\n"
	}
	// tags
	galleryPreview.Tags, err = parseTags(doc)
	if err != nil {
		galleryPreview.Error += err.Error() + "\n"
		return
	}

	galleryPreview.Category, err = findOneAndSelectAttr(doc, "//div[@id='gdc']", InnerText)
	if err != nil {
		galleryPreview.Error += err.Error() + "\n"
	}
	galleryPreview.Date, err = findOneAndSelectAttr(doc, "//div[@id='gdd']/table/tbody/tr[1]/td[2]", InnerText)
	if err != nil {
		galleryPreview.Error += err.Error() + "\n"
	}
	galleryPreview.Pages, err = findOneAndSelectAttr(doc, "//div[@id='gdd']/table/tbody/tr[6]/td[2]", InnerText)
	if err != nil {
		galleryPreview.Error += err.Error() + "\n"
	}
	galleryPreview.Images, err = queryAllStringList(doc, "//div[@class='gdtl']//img", "src")
	if err != nil {
		galleryPreview.Error += err.Error() + "\n"
		return
	}

	galleryPreview.Urls, err = queryAllStringList(doc, "//div[@class='gdtl']//a", "href")
	if err != nil {
		galleryPreview.Error += err.Error() + "\n"
		return
	}

	return
}

func QueryGalleryPreview(query string) (galleryPreview *GalleryPreview, err error) {
	return queryGalleryPreview(query)
}

// func QueryGallerySummary(query string) (gallerySummary *GallerySummary, err error) {
// 	return queryGallerySummary(query)
// }
