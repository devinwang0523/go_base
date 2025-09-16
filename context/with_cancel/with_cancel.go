package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1) // 等待一个 goroutine 完成

	go func(ctx context.Context) {
		// 在 goruntime 完成时调用 Done
		defer wg.Done()
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("完成任务")
		case <-ctx.Done():
			fmt.Println("任务被取消")
		}
	}(ctx)

	// 模拟延迟
	time.Sleep(1 * time.Second)
	// 取消上下文
	cancel()

	// 等待 goroutine 完成
	wg.Wait()
}
