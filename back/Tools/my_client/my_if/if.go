// this is for the ipv6 pool.
// not for general usage.

// https://tunnelbroker.net/tunnel_detail.php?tid=827763

// net6.sh
//
// /sbin/ifconfig sit0 down
// /sbin/ifconfig sit1 down
//
// /sbin/ifconfig sit0 up
// /sbin/ifconfig sit0 inet6 tunnel ::66.220.18.42
// /sbin/ifconfig sit1 up
// /sbin/ifconfig sit1 inet6 add 2001:470:c:6c::2/64
// /sbin/ifconfig sit1 inet6 add 2001:470:c:6c::3/64
// /sbin/ifconfig sit1 inet6 add 2001:470:c:6c::4/64
// /sbin/ifconfig sit1 inet6 add 2001:470:c:6c::5/64
// /sbin/route -A inet6 add ::/0 dev sit1

// why it can work. idk.
// sit0's traffic is always 0 but the previous application can actually visit exhentai with ipv6.
// though, here must be some configurations not set up properly.

package myif

import (
	"os/exec"
	"sync"
)

var ips = make(map[string]struct{})
var mu sync.Mutex

// cmd := exec.Command("ifconfig", "sit1", "inet6", "add", IPv6Addr+"/64") // command and arguments
func Add(IPv6Addr string) error {
	mu.Lock()
	defer mu.Unlock()

	cmd := exec.Command("ifconfig", "sit1", "inet6", "add", IPv6Addr+"/64") // command and arguments
	output, err := cmd.Output()
	// log.Println("add addr " + IPv6Addr + " :" + string(output)) // for debugging
	_ = output

	if err != nil {
		ips[IPv6Addr] = struct{}{}
	}

	return err
}

// cmd := exec.Command("ifconfig", "sit1", "inet6", "del", IPv6Addr+"/64") // command and arguments
func Delete(IPv6Addr string) error {
	mu.Lock()
	defer mu.Unlock()

	cmd := exec.Command("ifconfig", "sit1", "inet6", "del", IPv6Addr+"/64") // command and arguments
	output, err := cmd.Output()
	// log.Println("del addr " + IPv6Addr + " :" + string(output)) // for debugging
	_ = output

	if err != nil {
		delete(ips, IPv6Addr)
	}

	return err
}

func List() []string {
	mu.Lock()
	defer mu.Unlock()
	ipList := make([]string, 0, len(ips))
	for k := range ips {
		ipList = append(ipList, k) // even the official packages implemented in this way as well.
	}
	return ipList
}
