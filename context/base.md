# context 用法

## 概览

> 在 Go 语言中，`context` 是一个非常重要的工具，尤其是在并发编程中。它的主要作用是跨 API 边界传递请求范围、取消信号和超时信息。`context` 主要用于在 Go 中管理长时间运行的任务（如 HTTP 请求处理），尤其是当多个 goroutine 需要共同协作时。

## 1. context 的基本概念

`context` 类型在 context 包中定义，通常用来传递关于请求生命周期的信息，比如：

* **取消信号**：让 goroutine 知道是否需要停止当前操作。
* **超时控制**：设置一个时间限制，超时后任务会自动取消。
* **请求范围数据**：可以携带一些与请求相关的数据，比如用户 ID 或请求 ID。

## 2. context 的常见使用场景

### 2.1. HTTP 请求处理

在 web 服务中，context 通常用来传递 HTTP 请求的相关信息，尤其是取消信号和超时。例如，在处理一个用户请求时，如果用户在请求完成前中途取消，context 可以告诉我们需要停止当前的操作。

### 2.2. 并发任务控制

在并发编程中，context 可以在多个 goroutine 之间传递取消信号或超时信号，从而有效控制 goroutine 的生命周期。

## 3. 如何使用 context

在 Go 中，context 的使用主要围绕以下几个函数：

* context.Background()
* context.TODO()
* context.WithCancel()
* context.WithTimeout()
* context.WithDeadline()
* context.WithValue()

### 3.1. context.Background()

这是一个根 context，通常作为第一个 context 被传递给其他函数。它一般用于启动顶级请求的根 context。

```go
ctx := context.Background()
```

### 3.2. context.TODO()

这是一个占位符，表示你还没有确定要使用的 context，通常在程序的初始化阶段使用，待开发完成时再进行替换。

```go
ctx := context.TODO()
```

### 3.3. context.WithCancel()

WithCancel 可以派生出一个可取消的子 context。当你不再需要某个操作时，可以通过 cancel() 来通知所有相关的 goroutine 停止工作。

```go

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

```
