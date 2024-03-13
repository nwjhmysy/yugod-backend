### 1.go中的三个环境变量

- `GO111MODULE`

  **作用**：

  Go语言的`GO111MODULE`环境变量用于控制Go模块的行为。

  - `auto`: 在Go1.11之前，当工作目录在GOPATH/src之外或者有vendor目录时启用模块支持。在Go1.11及以后版本，它会根据当前目录是否包含go.mod文件来自动启用或禁用模块支持。
  - `on`: 强制启用模块支持，无论是否在GOPATH内。
  - `off`: 禁用模块支持，即使用旧的GOPATH模式。

  如果没有设置`GO111MODULE`，Go1.11及以后的版本会默认根据当前目录是否包含go.mod文件来启用或禁用模块支持。

  **查看方式**：

  ```
  go env GO111MODULE
  ```

  **修改**：

  ```
  go env -w GO111MODULE=on
  ```

  将 "on" 替换为 "off" 或 "auto"。

- `GOPATH`

  **作用**：

  `GOPATH` 是一个环境变量，用于指定你的 Go 工作空间的根目录。在 Go 1.11 之前，`GOPATH` 是主要用于存放项目代码、依赖包和构建输出的地方。从 Go 1.11 开始，Go 引入了模块（modules）的概念，允许在任何目录中工作，不再强制依赖于 `GOPATH`。

  **默认路径**：

  默认情况下，Go 模块会将下载的依赖包存放在 `GOPATH/pkg/mod` 目录下。具体路径可能是类似于：

  - Windows: `C:\Users\YourUsername\go\pkg\mod`
  - macOS/Linux: `/Users/YourUsername/go/pkg/mod`

  **查看方式**：

  ```
  go env GOPATH
  ```

  **修改**：

  ```
  go env -w GOPATH=/your/go/workspace
  ```

- `GOPROXY`

  **作用**：

  下载依赖的地址。国内的代理是：https://goproxy.cn,direct

  **查看方式**：

  ```
  go env GOPROXY
  ```

  **修改**：

  ```
  go env -w GOPATH=/your/go/workspace
  ```

### 2.src/vendor文件夹的作用

在 Go 1.11 版本之前，Go 使用 `GOPATH` 环境变量来指定工作空间，并要求所有的 Go 项目都位于 `GOPATH` 目录下。在这个工作空间中，通常会有一个 `src` 目录用于存放源代码，而 `vendor` 目录则用于存放项目的依赖包。

`vendor` 目录的作用是存放项目所依赖的外部包的源码。在这个目录中，你可以手动将项目所依赖的第三方库的源代码复制到 `vendor` 目录下。这样做的目的是确保项目的依赖项的版本是可控的，不会受到外部变化的影响。

在 Go 1.11 版本之后，Go 引入了模块（modules）的概念，不再强制要求项目必须放在 `GOPATH` 下，也不再依赖 `vendor` 目录来存放依赖包。模块支持更灵活地管理依赖项，不再要求将所有项目都放在相同的工作空间中。

如果你使用 Go 模块，通常不需要手动维护 `vendor` 目录，因为 Go 模块会自动下载和管理依赖项。如果你的项目是一个 Go 模块，可以在项目的根目录下使用 `go mod vendor` 命令，将依赖项拷贝到 `vendor` 目录中。这可以用于创建一个本地的 `vendor` 目录，以便在没有网络连接的环境中构建项目。

### 3.结构体指针可以直接“.”出某字段

在Go语言中，使用结构体指针直接通过`.`访问字段的原因是语法糖和方便性。Go语言在这方面设计上进行了一些智能化处理，以使代码更加简洁和易读。

当你使用结构体指针访问字段时，Go语言会自动进行解引用（dereference）操作，以便访问结构体的字段。这使得你可以通过指针直接访问字段，而无需显式使用`*`来解引用指针。

考虑以下示例：

```go
package main

import "fmt"

type Person struct {
    FirstName string
    LastName  string
}

func main() {
    // 创建结构体的指针
    personPtr := &Person{FirstName: "John", LastName: "Doe"}

    // 直接通过指针访问字段，无需显式解引用
    fmt.Println(personPtr.FirstName)
    fmt.Println(personPtr.LastName)
}
```

在上面的示例中，`personPtr.FirstName`直接访问了结构体指针`personPtr`中的`FirstName`字段，而不需要显式地解引用指针。

这种设计使得代码更加简洁，同时也减少了犯错的可能性。Go语言的设计哲学之一是保持简单性和可读性，这种语法糖有助于实现这一目标。当然，这也意味着在使用指针时，程序员需要注意确保指针不为`nil`，以避免潜在的空指针异常。

### 4.`goroutine`和线程的本质区别

`goroutine` 和传统的操作系统线程（比如在C++或Java中创建的线程）之间有一些关键的区别。这些区别主要涉及到调度、并发性、内存模型等方面。以下是一些主要区别：

1. **调度：**
   - **Goroutines：** Go语言的`goroutine`是由Go运行时（goroutine的调度器）管理的，而不是由操作系统的线程调度器管理。Go调度器使用一种称为"M:N"调度的技术，即将多个`goroutine`映射到少量的操作系统线程上，从而更有效地使用系统资源。
   - **线程：** 传统线程通常由操作系统的线程调度器管理。
2. **并发性：**
   - **Goroutines：** 轻量级，创建和销毁的成本较低。Go语言鼓励通过`goroutines`实现并发，因为它们相对轻量且易于管理。
   - **线程：** 创建和销毁线程的成本较高，因为线程是由操作系统调度器管理的重量级结构。
3. **内存模型：**
   - **Goroutines：** 共享内存模型，但通过通道（channels）来实现数据共享和通信，以避免竞态条件。
   - **线程：** 通常通过共享内存进行通信，需要使用锁等同步机制来保护共享数据。
4. **通信：**
   - **Goroutines：** 使用通道（channels）进行通信，通过在`goroutine`之间传递消息来实现协作。
   - **线程：** 通常使用共享内存进行通信，需要显式的同步机制。
5. **异常处理：**
   - **Goroutines：** Go语言使用`defer`和`panic/recover`机制进行异常处理，它允许在`goroutine`中恢复从异常中恢复。
   - **线程：** 传统线程通常使用try-catch块或其他异常处理机制。

总体来说，`goroutines`是Go语言特有的一种轻量级并发机制，它在语言层面提供了高效的并发支持。这种方式更易于使用和管理，而不需要开发者手动处理线程的复杂性。

注意：如果主线程执行完毕，那么子线程也会结束。

### 5.`Channels`

如果说goroutine是Go语音程序的并发体的话，那么channels它们之间的通信机制。一个channels是一个通信机制，它可以让一个goroutine通过它给另一个goroutine发送值信息。每个channel都有一个特殊的类型，也就是channels可发送数据的类型。一个可以发送int类型数据的channel一般写为chan int。

使用内置的make函数，我们可以创建一个channel：

```go
ch := make(chan int) // ch has type 'chan int'
```

和map类似，channel也一个对应make创建的底层数据结构的引用。当我们复制一个channel或用于函数参数传递时，我们只是拷贝了一个channel引用，因此调用者何被调用者将引用同一个channel对象。和其它的引用类型一样，channel的零值也是nil。

两个相同类型的channel可以使用==运算符比较。如果两个channel引用的是相通的对象，那么比较的结果为真。一个channel也可以和nil进行比较。

一个channel有发送和接受两个主要操作，都是通信行为。一个发送语句将一个值从一个goroutine通过channel发送到另一个执行接收操作的goroutine。发送和接收两个操作都是用`<-`运算符。在发送语句中，`<-`运算符分割channel和要发送的值。在接收语句中，`<-`运算符写在channel对象之前。一个不使用接收结果的接收操作也是合法的。

```go
ch <- x  // a send statement

x = <-ch // a receive expression in an assignment statement

<-ch     // a receive statement; result is discarded
```

**Channel还支持close操作，用于关闭channel，随后对基于该channel的任何发送操作都将导致panic异常。**对一个已经被close过的channel之行接收操作依然可以接受到之前已经成功发送的数据；如果channel中已经没有数据的话讲产生一个零值的数据。

使用内置的close函数就可以关闭一个channel：

```go
close(ch)
```

以最简单方式调用make函数创建的时一个无缓冲的channel，但是我们也可以指定第二个整形参数，对应channel的容量。如果channel的容量大于零，那么该channel就是带缓冲的channel。

```go
ch = make(chan int)    // unbuffered channel

ch = make(chan int, 0) // unbuffered channel

ch = make(chan int, 3) // buffered channel with capacity 3
```

### 6.`goroutine`和`Channels`

`goroutine`和`Channels`结合使用用于处理并发和多个线程之间的通信。

```go
package main

import (
	"fmt"
	"time"
)

func asyncOperation(c chan string) {
	// 模拟异步操作
	time.Sleep(2 * time.Second)
	c <- "Async operation result"

	close(c)
}

func main() {
	resultChannel := make(chan string)

	go asyncOperation(resultChannel)

	// 主程序可以继续执行其他操作
	fmt.Println("other Controls")

	// 等待异步操作完成并获取结果
	result := <-resultChannel
	fmt.Println(result)

	// 继续执行非异步操作
	fmt.Println("Non-async operation")
}

// 输出
// other Controls
// 2s后
// Async operation result
// Non-async operation
```

`go asyncOperation(resultChannel)` 启动了一个新的 goroutine，它会在后台执行 `asyncOperation` 函数。主程序继续执行其他操作，而不会等待异步操作完成。

然后，通过 `<-resultChannel` 语句等待异步操作的结果。这个语句会阻塞，直到有数据从 `resultChannel` 中可用。当异步操作完成时，它会将结果发送到 `resultChannel`，这时 `<-resultChannel` 解除阻塞，主程序继续执行。

因此，非异步操作会在等待异步操作完成后执行。这就是为什么 "Non-async operation" 的输出会在异步操作结果之后的原因。如果不使用 `go` 关键字启动 goroutine，那么异步操作会在主程序的其他操作之前执行，但由于 `go` 启动了新的 goroutine，所以异步操作可以在后台执行。

### 7.defer和panic/recover机制进行异常处理

在Go语言中，`defer` 、`panic` 和 `recover` 是一套用于处理异常的机制。这些机制用于在程序执行过程中处理不可预测的错误或异常情况。

1. **defer：**

   - `defer` 用于确保函数调用结束时发生清理操作，无论函数是否发生了错误。
   - `defer`语句会在包含它的函数执行完成之前执行，而不论函数是正常返回还是发生了 panic 异常。

   ```go
   func exampleFunction() {
       defer fmt.Println("This will be executed last.")
       fmt.Println("This will be executed first.")
   }
   ```

2. **panic：**

   - `panic` 是一个内建函数，用于引发运行时错误。通常用于表示不可恢复的错误。
   - 当 `panic` 函数被调用时，当前函数的执行被停止，然后开始沿着调用堆栈向上执行函数，执行每个函数的 `defer` 语句，然后程序终止。

   ```go
   func examplePanic() {
       defer fmt.Println("This will be executed before panic.")
       panic("Something went wrong!")
   }
   ```

3. **recover：**

   - `recover` 是一个内建函数，用于从 panic 中恢复，并且只能在 `defer` 语句中调用。
   - 当 `recover` 函数被调用时，它会停止 panic 的传播，并返回传递给 panic 的值。如果没有 panic，则 `recover` 返回 `nil`。

   ```go
   func exampleRecover() {
       defer func() {
           if r := recover(); r != nil {
               fmt.Println("Recovered from panic:", r)
           }
       }()
       panic("Something went wrong!")
   }
   ```

综合使用 `defer`、`panic` 和 `recover` 可以用于处理一些意外情况，保证程序在发生异常时能够做一些必要的清理工作。需要注意的是，`panic` 和 `recover` 应该谨慎使用，因为过度使用它们可能导致代码难以理解和维护。通常，应该优先使用错误值（error）来处理错误，而不是 `panic`。

### 8.`goroutine`重要特性之一

在Go语言中，如果在一个方法中启动一个goroutine（Go线程），该方法执行完毕后，**这个goroutine并不会自动结束**。**主程序会等待所有的goroutine完成后再退出**。这就是Go语言中的并发模型的一部分。

当一个Go程序启动时，它会创建一个主goroutine，然后在这个goroutine中执行`main`函数。如果在`main`函数或其他goroutine中启动了新的goroutine，这些goroutine会在它们的任务完成之前，或者主goroutine退出之前，保持运行。

如果在一个方法中启动了一个goroutine，而且这个方法的调用者没有等待该goroutine完成，那么这个goroutine可能在主goroutine退出时被中断，从而导致未完成的任务。这种情况可能会导致程序中的未定义行为，因此在设计并发程序时要特别注意这一点。

为了确保在程序退出之前等待goroutine完成，可以使用一些同步机制，例如`sync.WaitGroup`，`channel`等。通过这些机制，您可以在主goroutine中等待其他goroutine完成它们的任务，以确保程序退出时所有任务都已完成。