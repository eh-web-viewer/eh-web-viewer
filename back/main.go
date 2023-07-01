package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/eh-web-viewer/eh-web-viewer/api"
	"github.com/eh-web-viewer/eh-web-viewer/exproxy"
	"github.com/eh-web-viewer/eh-web-viewer/my_if"
	"github.com/eh-web-viewer/eh-web-viewer/mycurl"
	"github.com/gofiber/fiber/v2"
)

func App() *fiber.App {
	app := fiber.New()
	app.All("/echo/*", func(c *fiber.Ctx) error {
		buf := bytes.NewBuffer(nil)

		// c.Method() is /echo/fsdf/sdf
		if _, err := buf.WriteString(fmt.Sprintf("%s %s\n", c.Method(), c.Path())); err != nil {
			return err
		}

		// GET /echo/fsdf/sdf, * = fsdf/sdf
		if _, err := buf.WriteString(fmt.Sprintf("c.Params('*') = '%s'\n", c.Params("*"))); err != nil {
			return err
		}

		// with out '?'
		if _, err := buf.WriteString(fmt.Sprintf("string(c.Request().URI().QueryString())) = '%s'\n", string(c.Request().URI().QueryString()))); err != nil {
			return err
		}

		headers := c.GetReqHeaders()
		keys := make([]string, 0, len(headers))
		for k := range headers {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			if _, err := buf.WriteString(fmt.Sprintf("%s: %s\n", k, headers[k])); err != nil {
				return err
			}
		}
		buf.WriteString("==============================\n")

		if _, err := buf.Write(c.Body()); err != nil {
			return err
		}

		_, err := c.Write(buf.Bytes())

		return err

	})
	app.Get("/api/*", func(c *fiber.Ctx) error {
		path := c.Params("*")
		query := string(c.Request().URI().QueryString())
		if query != "" && path != "" {
			query = path + "/?" + query
		} else if query != "" && path == "" {
			query = "?" + query
		}
		if query == "" {
			query = path
		}
		// query = exhentai/{here}, includes '?' part

		var r any
		var err error
		if strings.HasPrefix(query, "g/") { // gallery
			r, err = api.QueryGalleryPreview(query)
		} else if strings.HasPrefix(query, "s/") { // image
			r, err = api.QueryImage(query)
		} else if strings.HasPrefix(query, "gallerytorrents.php") {
			r, err = api.QueryTorrents(query)
		} else { // index
			r, err = api.QueryIndex(query)
		}
		if err != nil {
			return err
		}

		return c.JSON(r)

	})
	return app
}

func main() {
	my_if.SetPrefix(my_if.PREFIX)
	defer my_if.Cleanup()

	// set my_if before mycurl.V6POOL
	mycurl.SetClient(mycurl.V6POOL) // curl with different ip

	// set mycurl before this line
	exproxy.SetClient(func() *http.Client { return mycurl.Client() })
	go exproxy.Proxy(MIRROR_ADDR)

	// set mycurl before this line
	app := App()
	err := app.Listen(API_ADDR)
	log.Println(err)
}
