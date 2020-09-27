package main

import (
	"context"
	"fmt"
)

//返回生成自然数序列的管道
func GenerateNatural(ctx context.Context) chan int {
	ch := make (chan int)
	go func() {
		for i := 2; ;i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()

	return ch;
}

//管道过滤器：删除能被素数整除的数
func PrimeFilter(ctx context.Context, in <- chan int, prime int) chan int {
    out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				select {
				case <-ctx.Done():
					return
				case out <- i:
				}
			}
		}
	}()

    return out;
}

func main() {
	// 通过 Context 控制后台Goroutine状态
	ctx, cancle := context.WithCancel(context.Background())

	ch := GenerateNatural(ctx)

	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("%v: %v\n", i + 1, prime)
		ch = PrimeFilter(ctx, ch, prime)
	}

	//ch2 := make(chan int)
	//for{
	//	select {
	//	case v := <- ch2:
	//		fmt.Println(v)
	//	case <- time.After(time.Second): //若ch2阻塞，超时退出
	//		return
	//		//default:
	//		//	fmt.Println("default")
	//	}
	//}

	cancle()

}
