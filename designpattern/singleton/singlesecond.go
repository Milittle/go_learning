package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once
var singleinstance *single

func GetSingleInstance() * single {
	if singleinstance == nil {
		once.Do(func() {
			fmt.Println("Create the single instance")
			singleinstance = &single{}
		})
	} else {
		fmt.Println("The instance has been created")
	}
	return singleinstance
}
