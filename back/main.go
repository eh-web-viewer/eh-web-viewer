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
	app.Get("/api/*", func(c *fiber.Ctx) (err error) {
		//
		originalURL := c.OriginalURL()
		query := originalURL[4:]

		var r any
		if strings.HasPrefix(query, "/g/") { // gallery
			r, err = api.QueryGalleryPreview(query)
		} else if strings.HasPrefix(query, "/s/") { // image
			r, err = api.QueryImage(query)
		} else if strings.HasPrefix(query, "/gallerytorrents.php") {
			r, err = api.QueryTorrents(query)
		} else { // index
			r, err = api.QueryIndex(query)
		}
		if err != nil {
			return
		}

		return c.JSON(r)

	})
	app.Get("/image/*", func(c *fiber.Ctx) error {
		originalURL := c.OriginalURL()
		query := originalURL[6:]

		if strings.HasPrefix(query, "/s/") { // image
			r, err := api.QueryImage(query)
			if err != nil {
				return err
			}
			return c.Redirect(r.Image, fiber.StatusFound) // 使用 fiber.StatusFound (302) 作为状态码
		} else {
			return c.JSON("nop")
		}
	})
	app.Get("/post/", func(c *fiber.Ctx) error {
		return api.Post()
	})
	app.Static("/", ".", fiber.Static{
		Compress: false, // it will generate [filename].tar.gz files.
		Browse:   false,
		Download: false, // it will make all things being downloaded. (false is default)
	})

	return app
}

func main() {
	my_if.SetPrefix(my_if.PREFIX)
	defer my_if.Cleanup() // though it has never run

	// // set my_if before mycurl.V6POOL
	mycurl.SetClient(mycurl.V6POOL) // curl with different ip

	// commit above and run this to run a vanilla client
	// mycurl.SetClient(mycurl.VANILLA)

	// set mycurl before this line
	// raw mirror
	exproxy.SetClient(func() *http.Client { return mycurl.Client() })
	go exproxy.Proxy(MIRROR_ADDR)

	// set mycurl before this line
	// api server
	app := App()
	err := app.Listen(API_ADDR)
	log.Println(err)
}
