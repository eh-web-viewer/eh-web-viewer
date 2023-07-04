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
	if resp.StatusCode != 200 {
		image.Error = resp.Status
		return
	}

	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		image.Error = err.Error()
		return
	}

	image.PrevPage, err = findOneAndSelectAttr(doc, "//a[@id='prev']", "href")
	if err != nil {
		image.Error += err.Error() + "\n"
	}
	image.NextPage, err = findOneAndSelectAttr(doc, "//a[@id='next']", "href")
	if err != nil {
		image.Error += err.Error() + "\n"
	}
	image.GalleryPage, err = findOneAndSelectAttr(doc, "//div[@class='sb']/a", "href")
	if err != nil {
		image.Error += err.Error() + "\n"
	}
	image.Image, err = findOneAndSelectAttr(doc, "//img[@id='img']", "src")
	if err != nil {
		image.Error += err.Error() + "\n"
	}
	image.AltImage, err = findOneAndSelectAttr(doc, "//a[@id='loadfail']", "onclick")
	if err != nil {
		image.Error += err.Error() + "\n"
	}
	image.OriginImage, err = findOneAndSelectAttr(doc, "//div[@id='i7']/a", "href")
	if err != nil {
		image.Error += err.Error() + "\n"
		err = nil
	}

	return
}

func QueryImage(query string) (image *Image, err error) {
	return queryImage(query)
}
