package sometest

import "fmt"

// Map key must be comparable
// also have direct part and indirectly part.

func MapTest() {
	m := map[int]string{1: "hello", 2: "world"}
	v := m[1]
	v = "test"
	fmt.Println(m, v)
}
