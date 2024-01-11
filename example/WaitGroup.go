package example

import (
	"fmt"
	"sync"
	"time"
)

func GroupMain() {
	var waitGroup sync.WaitGroup

	start := time.Now()
	waitGroup.Add(5)
	fmt.Println("开始")

	for i := 0; i < 5; i++ {
		go func() {
			defer waitGroup.Done()
			time.Sleep(time.Second)
			fmt.Println("done")
		}()
	}
	fmt.Println("开始中")

	waitGroup.Wait()
	fmt.Println(time.Now().Sub(start).Seconds())
}
