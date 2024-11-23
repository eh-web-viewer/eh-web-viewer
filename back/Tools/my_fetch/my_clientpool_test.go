// gin-pack @ 2024-04-06
package myfetch

import (
	"net/http"
	"testing"
)

// 没用到啊好像
func TestProxy(t *testing.T) {
	client := NewProxyClient("http://localhost:10809")
	clientPool := NewClientPool(nil)
	clientPool.SetClients([]*http.Client{client})
}
