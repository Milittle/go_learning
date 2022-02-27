package channel

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"
)

// 使用通道发送通知，往往是请求和响应的一种特殊情况
// 因为不注重响应的具体内容，而是关心响应是否发生，所以一般使用struct{}作为通道元素类型

//向一个通道发送一个值来实现通知

func Notification01() {
	values := make([]byte, 32*1024*1024)

	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	done := make(chan struct{})

	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		done <- struct{}{}
	}()

	///Do some other concurrent thing

	<-done

	fmt.Println(values[0], values[len(values)-1])
}

//从一个通道接受一个值来实现通知 这种方式不如上面第一种描述的常用

func Notification02() {
	done := make(chan struct{})

	go func() {
		fmt.Println("hello ")
		time.Sleep(time.Second * 4)
		<-done
	}()

	done <- struct{}{}

	fmt.Println("world")
}

//以上两者通知没有本质区别，都是较快者等待较慢者

//多对单和 单对多通知

func worker(id int, ready <-chan T, done chan<- T) {
	<-ready
	log.Println("Worker#", id, " start to work")
	time.Sleep(time.Second * time.Duration(id+1))
	log.Println("Worker#", id, " end of work")
	done <- T{}
}

func Notification03() {
	log.SetFlags(0)
	ready, done := make(chan T), make(chan T)
	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)

	//模拟一个初始化过程
	time.Sleep(time.Second * 2)
	ready <- T{}
	ready <- T{}
	ready <- T{}
	<-done
	<-done
	<-done
}

//这种特点的多通知单的方式，一般不怎么使用，正常情况下使用sync.WaitGroup
//以上的单对多的通知示例，使用的非常少，用的多的反而是使用close通道的方式，然后利用已经close的通道可以接收到无穷值的特性来实现通知：
// 所以我们可以通过修改ready的三行通道发送语句为close(ready)
/*
···
close(ready)
···
*/

//定时通知
func afterDuration(d time.Duration) <-chan struct{} {
	c := make(chan struct{})
	go func() {
		time.Sleep(d)
		c <- struct{}{}
	}()
	return c
}

func Notification04() {
	fmt.Println("Hi")
	<-afterDuration(time.Second)
	fmt.Println("Hello")
	<-afterDuration(time.Second)
	fmt.Println("World")
}

//在time包中有一个After函数和我们的afterDuration一致，尽量以后使用time包中的函数
//time.After函数会阻塞当前协程，而time.Sleep函数不会
