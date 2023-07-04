package exproxy

import (
	"net/http"
	"testing"

	"github.com/eh-web-viewer/eh-web-viewer/my_if"
	"github.com/eh-web-viewer/eh-web-viewer/mycurl"
)

func TestUnableToTest(t *testing.T) {
	// initial of my_if
	my_if.SetPrefix(my_if.PREFIX)
	defer my_if.Cleanup()
	// initial of mycurl
	mycurl.SetClient(mycurl.V6POOL) // curl with different ip
	// initial of exproxy
	SetClient(func() *http.Client { return mycurl.Client() })
	Proxy("127.0.0.1:8080") // listen on such addr
}

func TestXxx(t *testing.T) {
	// initial of mycurl
	mycurl.SetClient(mycurl.PROXY) // proxy
	// initial of exproxy
	SetClient(func() *http.Client { return mycurl.Client() })
	Proxy("127.0.0.1:8080") // listen on such addr
}
