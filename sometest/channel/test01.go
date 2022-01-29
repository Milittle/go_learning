package channel

import "fmt"

func Test01() {
	c := make(chan *int, 10)
	a := 1
	c <- &a
	a = 2
	d := <-c

	fmt.Println(*d)
}
