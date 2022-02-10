package channel

import (
	"fmt"
	"math/rand"
	"time"
)

//将单向接受通道作为函数返回值
func longTimeRequest() <-chan int32 {
	r := make(chan int32)

	go func() {
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()

	return r
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func FutureTest01() {
	rand.Seed(time.Now().UnixNano())

	a, b := longTimeRequest(), longTimeRequest()
	fmt.Println(sumSquares(<-a, <-b))
}

//将单向发送通道类型作为函数实参
func longTimeRequest01(r chan<- int32) {
	time.Sleep(time.Second * 3)
	r <- rand.Int31n(100)
}

func FutureTest02() {
	rand.Seed(time.Now().UnixNano())

	ra, rb := make(chan int32), make(chan int32)

	go longTimeRequest01(ra)
	go longTimeRequest01(rb)

	fmt.Println(sumSquares(<-ra, <-rb))
}

//采用最快响应
func source(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Intn(3)+1

	time.Sleep(time.Duration(rb) * time.Second)

	c <- ra
}

func FutureTest03() {
	rand.Seed(time.Now().UnixNano())
	startTime := time.Now()
	c := make(chan int32, 5)
	for i := 0; i < cap(c); i++ {
		go source(c)
	}
	rnd := <-c
	fmt.Println(time.Since(startTime))
	fmt.Println(rnd)
}
