package adapter

import "fmt"

type WindowsAdapter struct {
	win *Windows
}

func (w *WindowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.win.insertIntoUSBPort()
}