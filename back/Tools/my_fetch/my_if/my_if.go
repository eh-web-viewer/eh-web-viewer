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
var prefix = "2001:470:c:6c:"

const BATCH_SIZE = 4

func AddAddr(IPv6Addr string) error {
	cmd := exec.Command("/sbin/ifconfig", "sit1", "inet6", "add", IPv6Addr+"/64") // command and arguments
	output, err := cmd.Output()
	// log.Println("add addr " + IPv6Addr + " :" + string(output))
	_ = output
	return err
}

func DelAddr(IPv6Addr string) error {
	cmd := exec.Command("/sbin/ifconfig", "sit1", "inet6", "del", IPv6Addr+"/64") // command and arguments
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

func NewAddr(prefix string) net.IP {
	// addr2 := randomInt()
	// addr1 := randomInt()
	// nextIP := prefix + covInt2Hex(addr4) + ":" + covInt2Hex(addr3) + ":" + covInt2Hex(addr2) + ":" + covInt2Hex(addr1)
	nextIP := prefix + randomHex() + ":" + randomHex() + ":" + covInt2Hex(addr4) + ":" + covInt2Hex(addr3)
	addr3 += 1
	addr4 += addr3 / 65536
	addr3 %= 65536
	ip := net.ParseIP(nextIP)
	return ip
}

func generateAddrArr(prefix string, batch int) []net.IP {
	arr := make([]net.IP, batch)
	for i := 0; i < batch; i++ {
		arr[i] = NewAddr(prefix)
	}
	return arr
}

func addAddrArr(arr []net.IP) {
	for _, v := range arr {
		AddAddr(v.String())
		// log.Println("add", v.String())
	}
}

func delAddrArr(arr []net.IP) {
	for _, v := range arr {
		DelAddr(v.String())
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
