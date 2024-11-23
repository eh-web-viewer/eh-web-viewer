// gin-pack @ 2024-04-06
// azure-go @ 2023-12-21

package myfetch

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/eh-web-viewer/eh-web-viewer/Tools/orderedmap"
)

// requests
// 就加了个header。
func NewRequest(method, url string, header http.Header, body io.Reader) (*http.Request, error) {

	req, err := http.NewRequest(
		method,
		url,
		body,
	)
	if err != nil {
		return nil, err
	}

	if header != nil {
		req.Header = header
	}

	return req, nil
}

// 下面的不知道是啥东西。没实现吧。
// 应该是合并到URLEncodedForm里面了
// func BuildPlainReader(s any) io.Reader {
// 	switch v := s.(type) {
// 	case string:
// 		return strings.NewReader(v)
// 	case []byte:
// 		return bytes.NewReader(v)
// 	}
// 	return nil
// }

// support string, []byte, map[string]string, map[string][]string, url.Value, OrderedMap
type URLEncodedForm struct {
	data any
}

func (f *URLEncodedForm) Reader() (io.Reader, error) {

	switch bv := f.data.(type) {
	case string:
		return strings.NewReader(bv), nil
	case []byte:
		return bytes.NewReader(bv), nil
	case map[string]string:
		data := make(url.Values)
		for k, v := range bv {
			data.Set(k, v)
		}
		return strings.NewReader(data.Encode()), nil
	case map[string][]string:
		return strings.NewReader(url.Values(bv).Encode()), nil
	case url.Values:
		return strings.NewReader(bv.Encode()), nil
	case *orderedmap.OrderedMap:
		data := make(url.Values)
		for _, k := range bv.Keys() {
			switch v, _ := bv.Get(k); sv := v.(type) { // switched v
			case string:
				data.Set(k, sv)
			case []string:
				data[k] = sv
			}
		}
		return strings.NewReader(data.Encode()), nil
	default:
		return nil, fmt.Errorf("unknown urlencoded type: %T", f.data)
	}

}

// Apply application/x-www-form-urlencoded
// support string, []byte, map[string]string, map[string][]string, url.Value, OrderedMap
func BuildURLEncodedFormReader(data any) (io.Reader, error) {
	return (&URLEncodedForm{data: data}).Reader()
}

// func BuildJsonReader(data any) (io.Reader, error) {
// 	b := new(bytes.Buffer)
// 	err := json.NewEncoder(b).Encode(data)
// 	return b, err
// }
