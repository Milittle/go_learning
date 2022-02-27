package channel

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

/*
含有一个default分支和一个case分支的select代码块可以被用做一个尝试发送或者尝试接收操作，取决于case关键字后跟随的是一个发送操作还是一个接收操作。
* 如果case关键字后跟随的是一个发送操作，则此select代码块为一个尝试发送操作。 如果case分支的发送操作是阻塞的，则default分支将被执行，发送失败；否则发送成功，case分支得到执行。
* 如果case关键字后跟随的是一个接收操作，则此select代码块为一个尝试接收操作。 如果case分支的接收操作是阻塞的，则default分支将被执行，接收失败；否则接收成功，case分支得到执行。
*/

func TryTest01() {
	type Book struct{ id int }
	bookshelf := make(chan Book, 3)

	for i := 0; i < cap(bookshelf)*2; i++ {
		select {
		case bookshelf <- Book{id: i}:
			fmt.Println("成功将书放在书架上", i)
		default:
			fmt.Println("书架已经被占满了")
		}
	}

	for i := 0; i < cap(bookshelf)*2; i++ {
		select {
		case book := <-bookshelf:
			fmt.Println("成功从书架上取下一本书", book.id)
		default:
			fmt.Println("书架上已经没有书了")
		}
	}
}

// 无阻塞地检查一个通道是否已经关闭无阻塞地检查一个通道是否已经关闭
// 此方法常用来查看某个期待中的通知是否已经来临。此通知将由另一个协程通过关闭一个通道来发送。
// 一定要注意是没有任何协程给此通道发送数据，而是通过关闭来通知另外一个协程。
func IsClosed(c chan T) bool {
	select {
	case <-c:
		return true
	default:
	}
	return false
}

//峰值限制（peak/burst limiting）
//这个峰值限制可以看作是防止过大的并发请求，可以通过座位的例子
func SemTest06() {
	rand.Seed(time.Now().UnixNano())

	bar24x7 := make(Bar01, 10) // 最对同时服务10位顾客
	for customerId := 0; ; customerId++ {
		time.Sleep(time.Millisecond * 800)
		customer := Customer{customerId}
		select {
		case bar24x7 <- customer:
			go bar24x7.ServeCustomer(customer)
		default:
			log.Print("顾客#", customerId, "不愿等待而离去")
		}
	}
	for {
		time.Sleep(time.Second)
	}
}

//另一种“采用最快回应”的实现方式
//每个数据源协程只需使用一个缓冲为1的通道并向其尝试发送回应数据即可

func source01(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Intn(3)+1
	// 休眠1秒/2秒/3秒
	time.Sleep(time.Duration(rb) * time.Second)
	select {
	case c <- ra:
	default:
	}
}

func TryTest02() {
	rand.Seed(time.Now().UnixNano())

	c := make(chan int32, 1) // 此通道容量必须至少为1
	for i := 0; i < 5; i++ {
		go source(c)
	}
	rnd := <-c // 只采用第一个成功发送的回应数据
	fmt.Println(rnd)
}

// 第三种“采用最快回应”的实现方式
// 如果一个“采用最快回应”用例中的数据源的数量很少，比如两个或三个，
// 我们可以让每个数据源使用一个单独的缓冲通道来回应数据，然后使用一个select代码块来同时接收这三个通道。 示例代码如下：

func source02() <-chan int32 {
	c := make(chan int32, 1) // 必须为一个缓冲通道
	go func() {
		ra, rb := rand.Int31(), rand.Intn(3)+1
		time.Sleep(time.Duration(rb) * time.Second)
		c <- ra
	}()
	return c
}

func TryTest03() {
	rand.Seed(time.Now().UnixNano())

	var rnd int32
	// 阻塞在此直到某个数据源率先回应。
	select {
	case rnd = <-source02():
	case rnd = <-source02():
	case rnd = <-source02():
	}
	fmt.Println(rnd)
}

//注意：如果上例中使用的通道是非缓冲的，未被选中的case分支对应的两个source函数调用中开辟的协程将处于永久阻塞状态，从而造成内存泄露。

//超时机制（timeout）

func RequestWithTimeout(timeout time.Duration) (int, error) {
	c := make(chan int)
	//go doRequest(c) // 可能需要超出预期的时长回应

	select {
	case data := <-c:
		return data, nil
	case <-time.After(timeout):
		return 0, errors.New("超时了！")
	}
}

//脉搏器（ticker）
func tick(d time.Duration) <-chan struct{} {
	c := make(chan struct{}, 1) // 容量最好为1
	go func() {
		for {
			time.Sleep(d)
			select {
			case c <- struct{}{}:
			default:
			}
		}
	}()
	return c
}

func TickTest() {
	t := time.Now()
	for range tick(time.Second) {
		fmt.Println(time.Since(t))
	}
}

//事实上，time标准库包中的Tick函数提供了同样的功能，但效率更高。 我们应该尽量使用标准库包中的实现。

//速率限制（rate limiting）

type Request interface{}

func handle(r Request) { fmt.Println(r.(int)) }

const RateLimitPeriod = time.Minute
const RateLimit = 200 // 任何一分钟内最多处理200个请求

func handleRequests(requests <-chan Request) {
	quotas := make(chan time.Time, RateLimit)

	go func() {
		tick := time.NewTicker(RateLimitPeriod / RateLimit)
		defer tick.Stop()
		for t := range tick.C {
			select {
			case quotas <- t:
			default:
			}
		}
	}()

	for r := range requests {
		<-quotas
		go handle(r)
	}
}

func RateTest() {
	requests := make(chan Request)
	go handleRequests(requests)
	// time.Sleep(time.Minute)
	for i := 0; ; i++ {
		requests <- i
	}
}
