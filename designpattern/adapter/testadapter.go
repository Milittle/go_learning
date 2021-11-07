package adapter

func TestAdapter() {
	client := &client{}
	mac := &Mac{}
	client.ConnectInterfaceToComputer(mac)

	win := &Windows{}
	winAdapter := &WindowsAdapter{win: win}
	client.ConnectInterfaceToComputer(winAdapter)
}
