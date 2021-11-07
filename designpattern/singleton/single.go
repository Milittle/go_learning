package singleton

import (
	"fmt"
	"sync"
)

type single struct {
}

var singleton *single
var lock = &sync.Mutex{}

func GetSingleton() *single {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleton == nil {
			fmt.Println("Creating the singleton instance")
			singleton = &single{}
		} else {
			fmt.Println("The instance has been created")
		}
	} else {
		fmt.Println("The instance has been created")
	}
	return singleton
}
