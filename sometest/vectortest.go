package sometest

import "fmt"

// Just have one direct part.
// The deep indirectly part will be shared.

const Size = 32

type Person struct {
	name string
	age  int
}

// Vector type
var _ [5]string   // be a string vector.
var _ [Size]int   // be a int vector.
var _ [16][]byte  // be a []byte vector.
var _ [100]Person // be a People struct vector.

// Slice type
var _ []bool         // be a bool slice.
var _ []int64        // be a int64 slice.
var _ []map[int]bool //be a map[int]bool slice.
var _ []*int         // be a *int slice.

// Map type
var _ map[string]int            // be a key is string, value is int map.
var _ map[int]bool              // be a key is int, value is bool.
var _ map[int16][6]string       // be a key is int16, value is string vector map.
var _ map[struct{ x int }]*int8 // be a key is struct, value is *int map.

// every slice size is same.(because this type has indirectly part)
// every map size is same.(because this type has indirectly part)
// every vector size equal elementSize * elementCount.(because this type just only has direct part)

// Constant variable

func TestVector() {
	a := [10]bool{}
	b := [5]int{}
	fmt.Println(a)
	fmt.Println(b)
	c := a[:3]
	fmt.Println(c)
	c[0] = true
	fmt.Println(c, a)
}
