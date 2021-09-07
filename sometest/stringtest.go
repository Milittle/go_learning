package sometest

import "fmt"

func StringTest() {
	a := string("hello world")
	fmt.Println(a[:3], len(a))
	var y = []byte{5: 'y'}
	fmt.Println(y, len(y))
}
