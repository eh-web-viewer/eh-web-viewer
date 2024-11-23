package myfetch

import (
	"fmt"
	"io"
	"testing"
)

func TestX(t *testing.T) {

	body, _ := BuildURLEncodedFormReader(map[string]string{
		"client_id":     "agt.client_i",
		"refresh_token": "agt.refresh_toke",
		"grant_type":    "fresh_token",
		"client_secret": "agt.client_secre",
	})
	s, _ := io.ReadAll(body)
	fmt.Printf("%s", s)
}
