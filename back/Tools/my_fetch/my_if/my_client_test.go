package my_if

import (
	"fmt"
	"testing"
)

func TestXxfx(t *testing.T) {
	c, e := NewClient("fe80::", nil)
	fmt.Println(e)
	// b := c.(*http.Client)
	// fmt.Println(b)
	fmt.Println(c)
}
