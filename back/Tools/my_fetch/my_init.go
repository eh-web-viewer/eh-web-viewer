// gin-pack @ 2024-04-06

package myfetch

func init() {
	DefaultClientPool = NewClientPool(nil)
	DefaultFetcher = NewFetcher(nil, DefaultClientPool)
}
