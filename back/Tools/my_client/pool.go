package myclient

import (
	"net/http"
	"sync"
)

type ClientPool struct {
	index   int
	clients []*http.Client
	sync.Mutex
}

func NewClientPool() *ClientPool {
	return &ClientPool{
		index:   0,
		clients: make([]*http.Client, 0),
	}
}

func (p *ClientPool) Add(client *http.Client) {
	p.Lock()
	defer p.Unlock()
	p.clients = append(p.clients, client)
}

func (p *ClientPool) Get() *http.Client {
	p.Lock()
	defer p.Unlock()
	clen := len(p.clients)
	if clen == 0 {
		return http.DefaultClient
	}
	p.index = (p.index + 1) % clen
	return p.clients[p.index]
}
