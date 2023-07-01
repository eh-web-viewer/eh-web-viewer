package api

import (
	"github.com/antchfx/htmlquery"
	"github.com/eh-web-viewer/eh-web-viewer/mycurl"
)

type Image struct {
	Query       string `json:"query"`        // which query it executes
	GalleryPage string `json:"gallery_page"` // url to gallery
	NextPage    string `json:"next_page"`    // url to next page
	PrevPage    string `json:"prev_page"`    // url to previous page
	Image       string `json:"image"`        // url to image
	AltImage    string `json:"alt_image"`    // url to alternative image
	OriginImage string `json:"origin_image"` // url to origin image
	Error       string `json:"error"`        // error
}

func queryImage(query string) (image *Image, err error) {
	image = &Image{}

	image.Query = query

	resp, err := mycurl.Fetch("GET", BASE_URL+query,
		map[string]string{"Cookie": COOKIE},
		nil)
	if err != nil {
		image.Error = err.Error()
		return
	}

	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		image.Error = err.Error()
		return
	}

	prevPage := htmlquery.FindOne(doc, "//a[@id='prev']")
	image.PrevPage = htmlquery.SelectAttr(prevPage, "href")

	nextPage := htmlquery.FindOne(doc, "//a[@id='next']")
	image.NextPage = htmlquery.SelectAttr(nextPage, "href")

	gallery := htmlquery.FindOne(doc, "//div[@class='sb']/a")
	image.GalleryPage = htmlquery.SelectAttr(gallery, "href")

	img := htmlquery.FindOne(doc, "//img[@id='img']")
	image.Image = htmlquery.SelectAttr(img, "src")

	loadfail := htmlquery.FindOne(doc, "//a[@id='loadfail']")
	image.AltImage = htmlquery.SelectAttr(loadfail, "onclick")

	origin := htmlquery.FindOne(doc, "//div[@id='i7']/a")
	image.OriginImage = htmlquery.SelectAttr(origin, "href")

	return
}

func QueryImage(query string) (image *Image, err error) {
	return queryImage(query)
}
