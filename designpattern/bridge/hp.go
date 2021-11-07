package bridge

import "fmt"

type HP struct {
}

func (h *HP) PrintFile() {
	fmt.Println("print file by HP")
}
