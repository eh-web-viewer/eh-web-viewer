package exproxy

import (
	"log"
	"net/http"
	"testing"

	"github.com/eh-web-viewer/eh-web-viewer/my_if"
	"github.com/eh-web-viewer/eh-web-viewer/mycurl"
)

func TestUnableToTest(t *testing.T) {
	my_if.SetPrefix(my_if.PREFIX)
	defer my_if.Cleanup()

	mycurl.SetClient(mycurl.V6POOL) // curl with different ip

	SetClient(func() *http.Client { return mycurl.Client() })
	Proxy("127.0.0.1:8080")
}

func TestXxx(t *testing.T) {
	log.Println("????")
	mycurl.SetClient(mycurl.PROXY) // proxy
	SetClient(func() *http.Client { return mycurl.Client() })
	Proxy("127.0.0.1:8080")
}
