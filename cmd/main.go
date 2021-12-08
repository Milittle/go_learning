package main

import (
	"LC/designpattern/adapter"
	"LC/designpattern/bridge"
	"LC/designpattern/composite"
	"LC/designpattern/singleton"
	"LC/sometest"
	"fmt"
	"log"
	"math"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sort"
)

type IntSlice []int

func (a IntSlice) Len() int {
	return len(a)
}

func (a IntSlice) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a IntSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	x := 1
	x, y := 2, x
	fmt.Println(x, y)

	sometest.MapTest()
	sometest.ChannelTest()
	sometest.RedisTest()
	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println(math.Inf(-1), math.Inf(1))
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	runtime.Gosched()
	fmt.Println(sort.IsSorted(IntSlice{2, 2, 2}))
	for i := 0; i < 5; i++ {
		singleton.GetSingleton()
	}
	adapter.TestAdapter()
	bridge.TestBridge()
	composite.TestComposite()
	fmt.Scanln()
	select {}
}
