package tools

import "net/http"

type Header struct {
	http.Header
}

func (h Header) Add(key, value string) {
	if value == "" {
		return
	}
	h.Header.Add(key, value)
}

func (h Header) GetAllKeys() []string {
	s := make([]string, len(h.Header))
	for k := range h.Header {
		s = append(s, k)
	}
	return s
}

func NewHeader(header http.Header) Header {
	if header == nil {
		header = http.Header{}
	}
	return Header{Header: header}
}
