package api

import (
	"errors"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/eh-web-viewer/eh-web-viewer/mycurl"
	"golang.org/x/net/html"
)

const InnerText string = "INNER_TEXT"

func findOneAndSelectAttr(top *html.Node, expr string, name string) (v string, err error) {
	elem := htmlquery.FindOne(top, expr)
	if elem == nil {
		err = errors.New(expr + ":" + name + "is null")
		return
	}
	if name == InnerText {
		v = htmlquery.InnerText(elem)
	} else {
		v = htmlquery.SelectAttr(elem, name)
	}
	return
}

func Post() (err error) {
	resp, err := mycurl.Fetch("POST", "https://exhentai.org/uconfig.php",
		map[string]string{"Cookie": COOKIE},
		strings.NewReader(POSTDATA))
	if err != nil {
		return
	}
	return resp.Body.Close()
}
