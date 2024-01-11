package sync

import (
	"fmt"
	"sync"
)

func closureF(x int) func() {
	return func() {
		fmt.Println(x)
	}
}

func OnceMain() {
	var once sync.Once
	x := 4
	once.Do(closureF(x))
}
