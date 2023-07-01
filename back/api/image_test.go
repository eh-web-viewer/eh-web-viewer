package api

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/eh-web-viewer/eh-web-viewer/mycurl"
)

func Test_Image1(t *testing.T) {
	mycurl.SetClient("proxy")
	r, _ := QueryImage("s/da92fa31bc/2558259-1")
	fmt.Println(r)
	j, _ := json.Marshal(r)
	fmt.Println(string(j))
}
func Test_Image2(t *testing.T) {
	mycurl.SetClient("proxy")
	r, _ := QueryImage("s/8681587d21/2005521-5")
	fmt.Println(r)
	j, _ := json.Marshal(r)
	fmt.Println(string(j))
}
