package sometest

import (
	"fmt"
	"time"
)

func ChannelTest() {
	c := make(chan bool)
	go PutValue(c)
	v, Closed := <-c
	fmt.Println(v, Closed)
	v1, Closed1 := <-c
	fmt.Println(v1, Closed1)

	c1 := make(chan int)
	go func(c chan<- int, v int) {
		time.Sleep(time.Second * 2)
		c <- v * v
		fmt.Println("set value to channel done")
	}(c1, 3)
	done := make(chan struct{})
	go func(ch <-chan int) {
		n := <-ch
		fmt.Println("get value from channel done, the value is: ", n)
		time.Sleep(time.Second)
		done <- struct{}{}
	}(c1)
	<-done
	fmt.Println("done")
	forrangetest(fabonaccitestwithfor())
	selecttest()
}

func PutValue(c chan bool) {
	c <- true
	c <- true
}

func fabonaccitestwithfor() chan uint64 {
	fabonacci := func() chan uint64 {
		c := make(chan uint64)
		go func() {
			x, y := uint64(0), uint64(1)
			for ; y < (1 << 63); c <- y {
				x, y = y, x+y
			}
			close(c)
		}()
		return c
	}

	c := fabonacci()

	return c

	for x, ok := <-c; ok; x, ok = <-c {
		//time.Sleep(time.Second)
		fmt.Println("x value is: ", x)
	}
	return c
}

func forrangetest(ch <-chan uint64) {
	c := make(chan uint64)
	go func() {
		x, y := uint64(0), uint64(1)
		for ; y < (1 << 63); c <- y {
			x, y = y, x + y
		}
		close(c)
	}()
	for x := range ch{
		fmt.Println("x value is: ", x)
	}
}

func selecttest() {
	var ch chan struct{}
	select {
	case <- ch:
	case ch <- struct{}{}:
	default:
		fmt.Println("Go here.")
	}

	c := make(chan string, 2)
	trySend := func(v string) {
		select {
		case c <- v:
		default: // 如果c的缓冲已满，则执行默认分支。
		}
	}
	tryReceive := func() string {
		select {
		case v := <-c: return v
		default: return "-" // 如果c的缓冲为空，则执行默认分支。
		}
	}
	trySend("Hello!") // 发送成功
	trySend("Hi!")    // 发送成功
	trySend("Bye!")   // 发送失败，但不会阻塞。
	// 下面这两行将接收成功。
	fmt.Println(tryReceive()) // Hello!
	fmt.Println(tryReceive()) // Hi!
	// 下面这行将接收失败。
	fmt.Println(tryReceive()) // -
}

//keep have noone goruntine send to this channel data
func IsClosed(c chan int) bool {
	select {
	case <-c:
		return true
	default:
	}
	return false
}
