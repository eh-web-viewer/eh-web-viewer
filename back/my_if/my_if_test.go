package my_if

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestXx2(t *testing.T) {
	SetPrefix(PREFIX)
	// Init()
	// Shift()

	log.Println("===")
	log.Println(GetIPBatchAndShift())
	log.Println("===")
	log.Println(GetIPBatchAndShift())
	log.Println("===")
	log.Println(GetIPBatchAndShift())
	log.Println("===")
	log.Println(GetIPBatchAndShift())
	log.Println("===")

	Cleanup()
}

func TestX2x2(t *testing.T) {
	SetPrefix(PREFIX)
	Init()
	// this is good.
	log.Println("===")
	Shift()
	log.Println(GetIPBatch())
	log.Println("===")
	Shift()
	log.Println(GetIPBatch())
	log.Println("===")
	Shift()
	log.Println(GetIPBatch())
	log.Println("===")
	Shift()
	log.Println(GetIPBatch())
	log.Println("===")

	Cleanup()
}

func TestXxx(t *testing.T) {
	fmt.Println(time.Now().UnixNano())
	for j := 0; j < 3; j++ {
		for i := 0; i < 65536/2; i++ {
			(newAddr(PREFIX))
		}
		fmt.Println(newAddr(PREFIX))
	}
	// fmt.Println(newAddr(PREFIX))
	// fmt.Println(newAddr(PREFIX))
	// fmt.Println(newAddr(PREFIX))
	// fmt.Println(newAddr(PREFIX))
	// fmt.Println(newAddr(PREFIX))
	// fmt.Println(newAddr(PREFIX))
	// fmt.Println(newAddr(PREFIX))
	// fmt.Println(newAddr(PREFIX))
	// fmt.Println(newAddr(PREFIX))
	// fmt.Println(newAddr(PREFIX))
	// fmt.Println(newAddr(PREFIX))
	// fmt.Println(newAddr(PREFIX))
	// fmt.Println(newAddr(PREFIX))
	// fmt.Println(newAddr(PREFIX))
	// fmt.Println(newAddr(PREFIX))
}
