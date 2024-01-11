package example

import (
	"log"
	"os"
	"syscall"
)

func SyscallMain() {
	f, err := os.Create("example.txt")
	if err != nil {
		log.Println("create file example.txt failed", err)
	}
	defer f.Close()
	// 非阻塞模式下，加共享锁
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_SH|syscall.LOCK_NB); err != nil {
		log.Println("add share lock in no block failed", err)
	}
	// 这里进行业务逻辑
	// TODO

	// 解锁
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_UN); err != nil {
		log.Println("unlock share lock failed", err)
	}
	/**
		示例中 LOCK_SH 表示当前获取的是共享锁，
		如果是 LOCK_EX，则表示获取的是排他锁。
		而 LOCK_NB 表示当前获取锁的模式是非阻塞模式，
		如果需要阻塞模式，不加这个参数即可。
		LOCK_UN 则表示解锁，即释放锁。
	**/
	return
}
