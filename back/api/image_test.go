package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/eh-web-viewer/eh-web-viewer/mycurl"
)

// has origin
func Test_Image1(t *testing.T) {
	mycurl.SetClient("proxy")
	r, _ := QueryImage("s/da92fa31bc/2558259-1")
	fmt.Println(r)
	j, _ := json.Marshal(r)
	fmt.Println(string(j))
}

// no origin
func Test_Image2(t *testing.T) {
	mycurl.SetClient("proxy")
	r, _ := QueryImage("s/8681587d21/2005521-5")
	fmt.Println(r)
	j, _ := json.Marshal(r)
	fmt.Println(string(j))
}
func Test_Image(t *testing.T) {
	q := "s/b565325cb5/2599812-2"
	mycurl.SetClient("proxy")
	r, _ := QueryImage(q)
	fmt.Println(r)
	j, _ := json.Marshal(r)
	fmt.Println(string(j))
}

func printResp(resp *http.Response) {
	respText, _ := io.ReadAll(resp.Body)
	log.Println(string(respText))
}

func TestFetch2(t *testing.T) {
	mycurl.SetClient("proxy") // dont forget it
	q := "s/b565325cb5/2599812-2"
	resp, err := mycurl.Fetch("GET", BASE_URL+q,
		map[string]string{"Cookie": COOKIE},
		nil)
	if err != nil {
		t.Error(err)
	}
	printResp(resp)
}
