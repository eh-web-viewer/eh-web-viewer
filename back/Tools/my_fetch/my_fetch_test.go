// gin-pack @ 2024-04-06

package myfetch

import (
	"log"
	"testing"
)

func TestFetch(t *testing.T) {

	resp, _ := Fetch("GET", "https://static.cloudflareinsights.com/beacon.min.js/v84a3a4012de94ce1a686ba8c167c359c1696973893317", nil, nil)
	log.Println(resp.Header)
	log.Println(resp.Header.Get("Content-Type"))
}
