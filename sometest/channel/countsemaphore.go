package channel

import (
	"log"
	"math/rand"
	"time"
)

//计数信号量通常限制最大并发数

//和将通道用做互斥锁一样，也有两种方式用来获取一个用做计数信号量的通道的一份所有权。
//1. 通过发送操作来获取所有权，通过接收操作来释放所有权

type Seat int
type Bar chan Seat

func (bar Bar) ServeCustomer(c int) {
	log.Print("顾客#", c, "进入酒吧")
	seat := <-bar // 需要一个位子来喝酒
	log.Print("++ customer#", c, " drinks at seat#", seat)
	log.Print("++ 顾客#", c, "在第", seat, "个座位开始饮酒")
	time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
	log.Print("-- 顾客#", c, "离开了第", seat, "个座位")
	bar <- seat // 释放座位，离开酒吧
}

//此代码中，同时饮酒的人不会超过10人，喝完酒的人就会离开酒吧，释放椅子（可以将椅子理解为一些公共资源，比如线程池中的线程）

func SemTest01() {
	rand.Seed(time.Now().UnixNano())

	bar24x7 := make(Bar, 10) // 此酒吧有10个座位
	// 摆放10个座位。
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat(seatId) // 均不会阻塞
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		go bar24x7.ServeCustomer(customerId)
	}

	select {}
}

//下面是一个优化版本，每个协程都在创建好的协程中获取seat，就会阻塞在那里，我们采用提前获取seat，如果没有空作为，阻塞在创建协程之前

func (bar Bar) ServeCustomerOptimizer(c int, seat Seat) {
	log.Print("++ customer#", c, " drinks at seat#", seat)
	log.Print("++ 顾客#", c, "在第", seat, "个座位开始饮酒")
	time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
	log.Print("-- 顾客#", c, "离开了第", seat, "个座位")
	bar <- seat // 释放座位，离开酒吧
}

func SemTest02() {
	rand.Seed(time.Now().UnixNano())

	bar24x7 := make(Bar, 10) // 此酒吧有10个座位
	// 摆放10个座位。
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat(seatId) // 均不会阻塞
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		seat := <-bar24x7
		go bar24x7.ServeCustomerOptimizer(customerId, seat)
	}

	select {}
}

//在上面这个修改后的例子中，在任一时刻最多只有10个顾客协程在运行（但是在程序的生命期内，仍旧会有大量的顾客协程不断被创建和销毁）

func (bar Bar) ServeCustomerAtSeat(consumers chan int) {
	for c := range consumers {
		seatId := <-bar
		log.Print("++ 顾客#", c, "在第", seatId, "个座位开始饮酒")
		time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
		log.Print("-- 顾客#", c, "离开了第", seatId, "个座位")
		bar <- seatId // 释放座位，离开酒吧
	}
}

// 创建一个10个协程的池子，然后通过通道来传递信息
// 在协程内部函数中，使用for-range的方式来实现循环

func SemTest03() {
	rand.Seed(time.Now().UnixNano())

	bar24x7 := make(Bar, 10)
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat(seatId)
	}

	consumers := make(chan int)
	for i := 0; i < cap(bar24x7); i++ {
		go bar24x7.ServeCustomerAtSeat(consumers)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		consumers <- customerId
	}
}

func ServeCustomer(consumers chan int) {
	for c := range consumers {
		log.Print("++ 顾客#", c, "开始饮酒")
		time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
		log.Print("-- 顾客#", c, "离开座位")
	}
}

//当我们在使用资源的时候，其实我们是没必要关注资源的序号的

func SemTest04() {
	rand.Seed(time.Now().UnixNano())

	const BarSeatCount = 10
	consumers := make(chan int)
	for i := 0; i < BarSeatCount; i++ {
		go ServeCustomer(consumers)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		consumers <- customerId
	}
}

type Customer struct{ id int }
type Bar01 chan Customer

func (bar Bar01) ServeCustomer(c Customer) {
	log.Print("++ 顾客#", c.id, "开始饮酒")
	time.Sleep(time.Second * time.Duration(2+rand.Intn(10)))
	log.Print("-- 顾客#", c.id, "离开酒吧")
	<-bar // 离开酒吧，腾出位子
}

func SemTest05() {
	rand.Seed(time.Now().UnixNano())

	bar24x7 := make(Bar01, 10) // 最对同时服务10位顾客
	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second * 2)
		customer := Customer{customerId}
		bar24x7 <- customer // 等待进入酒吧
		go bar24x7.ServeCustomer(customer)
	}
	for {
		time.Sleep(time.Second)
	}
}
