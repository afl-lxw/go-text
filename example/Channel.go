package example

import "fmt"

func CountByCommunicating() {
	w := make(chan int, 100) // 用于确保100个 goroutine 都执行完毕
	ch := make(chan int, 1)
	ch <- 0
	for i := 0; i < 100; i++ {
		fmt.Println("for 中  ", i)
		go func() {
			select {
			case count := <-ch:
				ch <- count + 1
			}
			w <- 0
		}()
	}
	for i := 0; i < 100; i++ {
		fmt.Println("for 输出  ", i)
		<-w
	}
	fmt.Println(<-ch)
}
