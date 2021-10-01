package main

import (
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
	sometest.MapTest()
	sometest.ChannelTest()
	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println(math.Inf(-1), math.Inf(1))
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	runtime.Gosched()
	fmt.Println(sort.IsSorted(IntSlice{2, 2, 2}))
	select {}
}
