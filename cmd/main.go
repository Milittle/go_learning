package main

import (
	"LC/sometest"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func main() {
	sometest.MapTest()
	sometest.ChannelTest()
	fmt.Println(runtime.GOMAXPROCS(0))

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	select {}
}
