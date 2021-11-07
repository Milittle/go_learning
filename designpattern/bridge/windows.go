package bridge

import "fmt"

type Windows struct {
	printer Printer
}

func (mac *Windows) Print() {
	fmt.Println("Print by Windows")
	mac.printer.PrintFile()
}

func (mac *Windows) SetPrinter(p Printer) {
	mac.printer = p
}
