// 利用channel的缓存区处理协程太多的情况

package main

import (
	"fmt"
	"math"
	"time"
)

func do(i int, ch chan struct{}) {
	fmt.Println(i)
	time.Sleep(time.Second)
	<-ch
}

func main() {
	c := make(chan struct{}, 3000)
	for i := 0; i < math.MaxInt32; i++ {
		c <- struct{}{}
		go do(i, c)
	}
	time.Sleep(time.Hour)
}
