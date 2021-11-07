package bridge

import "fmt"

type Mac struct {
	printer Printer
}

func (mac *Mac) Print() {
	fmt.Println("Print by Mac")
	mac.printer.PrintFile()
}

func (mac *Mac) SetPrinter(p Printer) {
	mac.printer = p
}
