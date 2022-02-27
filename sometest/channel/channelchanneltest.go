package channel

import "fmt"

var counter = func(n int) chan<- chan<- int {
	requests := make(chan chan<- int)
	go func() {
		for request := range requests {
			if request == nil {
				n++ // 递增计数
			} else {
				request <- n // 返回当前计数
			}
		}
	}()
	return requests // 隐式转换到类型chan<- (chan<- int)
}(0)

func ChannelChannelTest() {
	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			counter <- nil
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)
	<-done
	<-done

	request := make(chan int, 1)
	counter <- request
	fmt.Println(<-request) // 2000
}
