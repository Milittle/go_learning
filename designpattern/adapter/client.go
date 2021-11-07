package adapter

import "fmt"

type client struct {
}

func (c *client) ConnectInterfaceToComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.insertIntoLightningPort()
}
