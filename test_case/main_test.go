package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func test() {
	fmt.Println("===================")
}

func TestContext(t *testing.T) {
	// c := context.Background()
	// c = context.WithValue(c, "worker", "log_processor")
	// fmt.Println("===================", c)
	var a atomic.Bool
	var wg sync.WaitGroup
	var so sync.Once
	var b atomic.Int32
	var sm sync.Map
	wg.Add(10)
	for v := range 10 {
		go func(v int32) {
			defer wg.Done()
			fmt.Println("v:", v)
			a.Store(true)
			so.Do(test)
			sm.Store(v, true)
			// b = v
			// atomic.StoreInt32(&b, 1)
			b.Store(4)
		}(int32(v))
	}
	wg.Wait()
	time.Sleep(60 * time.Second)
	sm.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true // 返回 true 继续遍历
	})
	time.Sleep(60 * time.Second)
}
