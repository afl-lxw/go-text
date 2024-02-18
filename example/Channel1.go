package example

import (
	"fmt"
	"time"
)

func startTask(id, n int, chans []chan struct{}) {
	// 每个任务从对应的 chan 读取数据，并传递给下一个chan
	for {
		token := <-chans[id]
		fmt.Printf("%d \n", id+1)
		chans[(id+1)%n] <- token
		time.Sleep(time.Second)
	}
}

func Channel1Main() {
	n := 4
	chans := []chan struct{}{}
	for i := 0; i < n; i++ {
		chans = append(chans, make(chan struct{}))
	}

	for i := 0; i < n; i++ {
		go startTask(i, n, chans)
	}
	chans[0] <- struct{}{}
	select {}
}
