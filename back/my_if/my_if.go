package my_if

import (
	"math/rand"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var addr4 = randomInt()
var addr3 = randomInt()
var prestBatch = make([]net.IP, 0)
var nextBatch = make([]net.IP, 0)
var prefix = PREFIX

func addAddr(IPv6Addr string) error {
	cmd := exec.Command("ifconfig", "sit1", "inet6", "add", IPv6Addr+"/64") // command and arguments
	output, err := cmd.Output()
	// log.Println("add addr " + IPv6Addr + " :" + string(output))
	_ = output
	return err
}

func delAddr(IPv6Addr string) error {
	cmd := exec.Command("ifconfig", "sit1", "inet6", "del", IPv6Addr+"/64") // command and arguments
	output, err := cmd.Output()
	// log.Println("del addr " + IPv6Addr + " :" + string(output))
	_ = output
	return err
}

func randomInt() int {
	randomInt := r.Intn(65536)
	return randomInt
}

func covInt2Hex(num int) string {
	hex := strconv.FormatInt(int64(num), 16)
	hexLower := strings.ToLower(hex)
	return hexLower
}

func randomHex() string {
	hex := covInt2Hex(randomInt())
	return hex
}

func newAddr(prefix string) net.IP {
	// addr2 := randomInt()
	// addr1 := randomInt()
	// nextIP := prefix + covInt2Hex(addr4) + ":" + covInt2Hex(addr3) + ":" + covInt2Hex(addr2) + ":" + covInt2Hex(addr1)
	nextIP := prefix + covInt2Hex(addr4) + ":" + covInt2Hex(addr3) + ":" + randomHex() + ":" + randomHex()
	addr3 += 1
	addr4 += addr3 / 65536
	addr3 %= 65536
	ip := net.ParseIP(nextIP)
	return ip
}

func generateAddrArr(prefix string, batch int) []net.IP {
	arr := make([]net.IP, batch)
	for i := 0; i < batch; i++ {
		arr[i] = newAddr(prefix)
	}
	return arr
}

func addAddrArr(arr []net.IP) {
	for _, v := range arr {
		addAddr(v.String())
		// log.Println("add", v.String())
	}
}

func delAddrArr(arr []net.IP) {
	for _, v := range arr {
		delAddr(v.String())
		// log.Println("del", v.String())
	}
}

func Cleanup() {
	delAddrArr(prestBatch)
	delAddrArr(nextBatch)
}

func Init() {
	nextBatch = generateAddrArr(prefix, BATCH_SIZE)
	addAddrArr(nextBatch)
}

func Shift() {
	delAddrArr(prestBatch)
	prestBatch = nextBatch
	nextBatch = generateAddrArr(prefix, BATCH_SIZE)
	addAddrArr(nextBatch)
}

func SetPrefix(s string) {
	prefix = s
}

func GetIPBatch() []net.IP {
	return nextBatch
}

func GetIPBatchAndShift() []net.IP {
	Shift()
	return GetIPBatch()
}
