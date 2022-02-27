package channel

import "fmt"

//通过大小为1的缓冲通道作为mutex

// 这种方式没有sync包中自带的mutex效率高，所以使用场景有限

//1. 通过发送操作来获取锁，通过接受操作来释放锁
func MutexTest01() {
	mutex := make(chan T, 1) //must be set 1

	counter := 0

	increase := func() {
		mutex <- T{}
		counter++
		<-mutex
	}

	increase1000 := func(done chan<- T) {
		for i := 0; i < 1000; i++ {
			increase()
		}
		done <- T{}
	}

	done := make(chan T)
	go increase1000(done)
	go increase1000(done)
	<-done
	<-done
	fmt.Println(counter)
}

//2. 通过接受操作来获取锁，通过发送操作来释放锁
func MutexTest02() {
	mutex := make(chan T, 1) // Must be set to 1
	mutex <- T{}             // It must have this sentence

	counter := 0
	increase := func() {
		<-mutex
		counter++
		mutex <- T{}
	}

	increase1000 := func(done chan<- T) {
		for i := 0; i < 1000; i++ {
			increase()
		}
		done <- T{}
	}

	done := make(chan T)
	go increase1000(done)
	go increase1000(done)

	<-done
	<-done
	fmt.Println(counter)
}
