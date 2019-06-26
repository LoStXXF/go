package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex
var many = 10

func test(handle bool, n int) {
	lock.Lock()
	defer lock.Unlock()
	if many <= 0 {
		fmt.Println("你的余额不足，不能消费", n)
	} else {
		time.Sleep(10 * time.Microsecond)
		many -= 10
		if !handle {
			many += 10
		}

	}
}
func main() {
	fmt.Println(many)
	go test(true, 1)
	go test(true, 2)
	go test(true, 3)
	go test(true, 4)
	go test(true, 5)
	go test(true, 6)
	go test(false, 7)
	go test(true, 8)
	go test(true, 9)
	go test(true, 10)
	go test(true, 11)
	go test(true, 12)
	go test(true, 13)
	go test(true, 14)
	go test(true, 15)
	go test(true, 16)
	go test(true, 17)
	go test(true, 18)
	go test(true, 19)
	go test(false, 20)
	go test(true, 21)
	go test(true, 22)
	go test(true, 23)
	go test(true, 24)
	go test(true, 25)
	go test(true, 26)
	go test(true, 27)
	go test(true, 28)
	go test(true, 29)
	go test(true, 30)
	go test(true, 31)
	go test(true, 32)
	go test(false, 33)
	go test(true, 34)
	go test(true, 35)
	go test(true, 36)
	go test(true, 37)
	go test(true, 38)
	go test(true, 39)
	go test(true, 40)
	go test(true, 41)
	go test(true, 42)
	go test(true, 43)
	go test(true, 44)
	go test(true, 45)
	go test(false, 46)
	go test(true, 47)
	go test(true, 48)
	go test(true, 49)
	go test(true, 50)
	go test(true, 51)
	go test(true, 52)
	time.Sleep(1 * time.Second)
	fmt.Println(many)
}
