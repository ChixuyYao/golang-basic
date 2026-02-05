package channel

import (
	"testing"
	"time"
)

// channel的基础使用
func TestChannel(t *testing.T) {
	// channel声明
	//var ch chan struct{}

	// channel 声明并创建
	//ch1 := make(chan int)

	// channel 声明并给定 buffer(缓存)
	ch2 := make(chan int, 1)

	// 向 channel 中存值
	ch2 <- 1
	// 从 channel 中取值
	data := <-ch2

	t.Log(data)

	// 关闭 channel
	close(ch2)
}

func TestChannelClose(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 0
	// val:从channel中读取出的值
	// ok:从channel读取是否成功
	val, ok := <-ch
	t.Log(val, ok)
	close(ch)

	// 当channel被close掉后,向channel中存值会panic
	// ch <- 1

	val, ok = <-ch
	t.Log(val, ok)

	// 当channel被close掉后,再次close会panic
	// close(ch)
}

func TestChannelLoop(t *testing.T) {
	ch := make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Second * 1)
		}
		close(ch)
	}()

	for val := range ch {
		t.Log(val)
	}
}

type BigStruct struct{}

// channel 阻塞(休眠)
func TestChannelBlocking(t *testing.T) {
	ch := make(chan int) // 尚未分配 Buffer
	b1 := BigStruct{}
	go func() {
		var b BigStruct
		// go routine 泄露
		ch <- 111
		t.Log(b, b1)
	}()
}

// channel的select使用
func TestChannelSelect(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 2)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 2
	}()

	select {
	case val := <-ch1:
		t.Log("命中channel01:", val)
	case val := <-ch2:
		t.Log("命中channel02:", val)
	}
}
