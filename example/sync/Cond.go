package sync

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func CondMain() {
	c := sync.NewCond(&sync.Mutex{})
	var ready int

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			// 加锁更改等待条件
			c.L.Lock()
			ready++
			c.L.Unlock()

			log.Printf("运动员#%d 已准备就绪\n", i)
			// 广播唤醒所有的等待者
			c.Broadcast()
		}(i)
	}

	c.L.Lock()
	for ready != 10 {
		c.Wait()
		log.Println("裁判员被唤醒一次")
	}
	c.L.Unlock()

	//所有的运动员是否就绪
	log.Println("所有运动员都准备就绪。比赛开始，3，2，1, ......")

	/**
	三个方法
	Cond 有三个方法：Signal、Broadcast 和 Wait。
	这三个方法名是计算机科学中条件变量的通用方法名。
	我们分别看一下这三个方法。
	Signal: 调用者通过该方法唤醒一个正在等待此 Cond 的 goroutine。如果此时没有等待的 goroutine，则无须通知 waiter，如果等待队列中有一个或者多个等待者，则移出第一个 goroutine 并唤醒它。
	Broadcast: 和上面的 Signal 方法类似，只是 Broadcast 会清空队列，唤醒全部的等待中的 goroutine。
	调用 Signal 和 Broadcast 的时候，不要求调用者持有锁 c.L 。
	Wait: 调用该方法的 goroutine 会被放到 Cond 的等待队列中并阻塞，直到被 Signal 或者 Broadcast 方法唤醒。
	调用 Wait 方法的时候一定要持有锁 c.L 。
	**/

}
