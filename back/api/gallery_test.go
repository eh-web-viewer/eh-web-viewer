package api

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/antchfx/htmlquery"
	"github.com/eh-web-viewer/eh-web-viewer/mycurl"
)

func Test_re(t *testing.T) {
	s := rePreview.FindString("width:250px; height:354px; background:transparent url(https://s.exhentai.org/t/e4/e2/e4e2e4a025387ba8c2b5a087ab67d89c72d2b7e1-1445020-1488-2105-jpg_250.jpg) no-repeat")
	fmt.Println(s)
}

func Test_galleryPreviews(t *testing.T) {
	mycurl.SetClient("proxy")
	r, _ := queryGalleryPreview("g/2557923/cd8d42df44/?p=1")
	fmt.Println(r)
	j, _ := json.Marshal(r)
	fmt.Println(string(j))
}

func TestFetch(t *testing.T) {
	mycurl.SetClient("proxy")
	query := "g/2559337/4d1b221fc9/?p=1"
	galleryPreview := &GalleryPreview{}

	resp, _ := mycurl.Fetch("GET", BASE_URL+query,
		map[string]string{"Cookie": COOKIE},
		nil)
	defer resp.Body.Close()

	respText, _ := io.ReadAll(resp.Body)
	respStr := string(respText)

	fmt.Println(respStr)

	doc, err := htmlquery.Parse(strings.NewReader(respStr))
	if err != nil {
		galleryPreview.Error = err.Error()
		return
	}

	gd1 := htmlquery.FindOne(doc, "//div[@id='gd1']/div")
	galleryPreview.Preview = rePreview.FindString(htmlquery.SelectAttr(gd1, "style"))

	gn := htmlquery.FindOne(doc, "//h1[@id='gn']")
	galleryPreview.Title = htmlquery.InnerText(gn)

	gj := htmlquery.FindOne(doc, "//h1[@id='gj']")
	galleryPreview.OriginTitle = htmlquery.InnerText(gj)

	// tags
	galleryPreview.Tags, err = parseTags(doc)
	if err != nil {
		return
	}

	gdc := htmlquery.FindOne(doc, "//div[@id='gdc']")
	galleryPreview.Category = htmlquery.InnerText(gdc)

	pages := htmlquery.FindOne(doc, "//div[@id='gdd']/table/tbody/tr[6]/td[2]")
	galleryPreview.Pages = htmlquery.InnerText(pages)

	galleryPreview.Images, err = queryAllStringList(doc, "//div[@class='gdtl']//img", "src")
	if err != nil {
		return
	}

	galleryPreview.Url, err = queryAllStringList(doc, "//div[@class='gdtl']//a", "href")
	if err != nil {
		return
	}

	fmt.Println(galleryPreview)

}

// func Test_gallerySummery(t *testing.T) {
// 	mycurl.SetClient("proxy")
// 	// r, _ := queryGallerySummary("g/2557923/cd8d42df44/?p=1")
// 	// fmt.Println(r)
// }
