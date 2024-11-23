// timeline-deamon @ 2023-12-26
// azure-go @ 2023-12-21

package myfetch

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/eh-web-viewer/eh-web-viewer/Tools/orderedmap"
	"github.com/andybalholm/brotli"
	"github.com/klauspost/compress/zstd"
)

// this function receive json request.
func ResponseToObject(r *http.Response) (o *orderedmap.OrderedMap, err error) {
	o = orderedmap.New()
	err = json.NewDecoder(r.Body).Decode(&o)
	return o, err
}

func ResponseToObjectArray(r *http.Response) (arr []*orderedmap.OrderedMap, err error) {
	err = json.NewDecoder(r.Body).Decode(&arr)
	return arr, err
}

func ResponseToReader(r *http.Response) (reader io.Reader, err error) {
	switch r.Header.Get("Content-Encoding") {
	case "gzip":
		return gzip.NewReader(r.Body)
	case "br":
		r := brotli.NewReader(r.Body)
		if r == nil {
			err = fmt.Errorf("header is br and failed to make new reader")
		}
		return r, err
	case "zstd":
		return zstd.NewReader(r.Body)
	default:
		return r.Body, nil
	}
}
