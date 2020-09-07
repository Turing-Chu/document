# Go中的调度：第二部分 - Go 调度器

> [https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html][part1]
> 
> 如果您计算机专业英文阅读能力不错，建议您阅读原文。

## 序

这是一个由三部分组成的系列文章中的第二篇，它将提供对Go调度程序背后的机制和语义的理解。 本篇着重于go调度。


三部分系列的索引：

- [Go 中的调度 : 第一部分 - 操作系统调度][part1]
- [Go 中的调度 : 第二部分 - Go 调度][part2]
- [Go 中的调度 : 第三 部分 - 并发][part3]

[part1]: https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html
[part2]: https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html
[part3]: https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part3.html


## 简介

在该调度系列的[第一部分][part1]中，我在操作系统调度层面解释了调度器，并认为理解并领会go调度器的语义学是非常重要的。在本篇中，我会在语义学层面解释go调度器如何运行并着重于高级行为。go调度器是一个复杂的系统，因此机制上的小细节并不重要。重要的是对go调度器的工作原理与行为有一个好的模型。这会让你做出更好的工程决策。

## 程序启动

当go程序启动时，会为主机上标识的每个虚拟核心分配逻辑处理器（P）。如果你的处理器的每个物理核（[超线程，Hyper-Threading][Hyper-Threading]）上有多个硬件线程，那么对于你的go程序来说每个硬件线程就是一个虚拟核心。为了更好的理解，看下我的MacBook Pro的系统报告。

**图1**

![](https://www.ardanlabs.com/images/goinggo/94_figure1.png)

可以看到，我的单处理器有4个物理核。该报告并未暴露每个物理核有几个硬件线程。Intel Core i7 处理器有超线程，这意味着每个物理核有2个硬件线程。这也对go程序反映出当并行执行OS线程时，有8个虚拟核可用。

要对此进行测试，请考虑以下程序：

**清单1**

```golang
package main

import (
	"fmt"
	"runtime"
)

func main() {

    // NumCPU returns the number of logical
    // CPUs usable by the current process.
    fmt.Println(runtime.NumCPU())
}
```

当我在我的机器上运行这个程序时，NumCPU()函数的调用结果是8。在我机器上运行的任何go程序都会有8个P。

每个P被分配一个OS线程（“M”）。‘M’代表机器。对每个可执行程序来说，该线程仍然由OS管理且OS仍然负责将该线程放在核上，这会在[最后一部分][part3]说明。这意味着当我在我本机上运行go程序时，我有8个线程可用于执行我的任务，每个线程都单独绑定到一个P上。

每个go程序也都会初始化为一个go协程（“G”），这是go程序的执行路径。go协程本质上是一个[协程][coroutine]，但这是go的，我们字母“G”代替“C”，于是得到单词go协程（goroutine）。你可以理解为go协程是应用级别的线程，在许多地方和OS的线程很像。只不过OS线程是在核上进行上下文切换，但go协程是在M上上下文切换。

最后一个难题是运行队列。在go调度器中有两种不同的运行队列：全局运行队列（Global Run Queue，GRQ）和本地运行队列（Local Run Queue，LRQ）。每个P都有一个LRQ来管理被分配到P上下文中的go协程。这些go协程在分配给该P的M上来回进行上下文切换。GRQ用于尚未分配P的go协程。将go协程从GRQ挪到LRQ有个过程，我们稍后探讨。

Figure 2 provides an image of all these components together.
图2提供了所有这些部分的图像。

**图2**

![图2](https://www.ardanlabs.com/images/goinggo/94_figure2.png)

[Hyper-Threading]: https://en.wikipedia.org/wiki/Hyper-threading
[coroutine]: https://en.wikipedia.org/wiki/Coroutine

## 协作式调度器


正如在第一部分中所讨论的，OS调度器是抢占式调度器。这实际上意味着在任何给定的时间无法预测调度器正在做什么。内核正在做决策并且任何事情都是不确定的。运行在OS之上的应用程序无法通过调度来控制内核内部发生的事情，除非它们利用像[原语][atomic]指令和[互斥][mutex]调用之类的同步原语。

go调度器是go运行时的一部分，且go运行时被编译进应用程序中。这意味着go调度器运行在内核之上的[用户空间][user space]中。当前go调度器的实现是[协作式][cooperating]调度而非抢占式调度（译者注：此文系2018-09-27发表，此时go最新版本为go1.11）。协作式调度器意味着需要在代码安全点发生的用户空间事件定义明确，以制定调度策略。

go协作式调度器的闪亮点在于它看起来以及感觉上是抢先式的。你无法预测go调度器正在做什么。这是因为为该协作式调度器制定的策略并不掌握在开发人员手中，而是在go运行时中。将go调度器视为抢占式调度非常重要，因为调度器是不确定的，这并不是一件容易的事。

[atomic]: https://en.wikipedia.org/wiki/Linearizability
[mutex]: https://en.wikipedia.org/wiki/Lock_(computer_science)
[user space]: https://en.wikipedia.org/wiki/User_space
[cooperating]: https://en.wikipedia.org/wiki/Cooperative_multitasking

## Go 协程状态

和线程一样，go协程也有相同的三种高级状态。这说明了go调度器在任何go协程中所扮演的角色。go协程可以是这三种状态之一：等待态，可执行态，执行态。

**等待态**：这意味着go协程已经停止然后等待某些事情以便继续。这可能是因为操作系统（系统调用）或同步调用（原语和互斥操作）。这些[延迟][latencies]类型是造成低性能的根因。

**执行态**：这意味着go协程需要M上的时间片以执行其指令。如果有大量协程等待，那么go协程获取时间片会等待更久。而且，随着更多的go协程争夺计算时间，每个协程所获得得时间都会变短。这种调度延迟类型也能引起低性能。

**可执行态**：这意味着go协程已经在M上运行指令。应用的相关任务即将结束。这是都希望的。

[latencies]: https://en.wikipedia.org/wiki/Latency_(engineering)

## 上下文切换

go调度器需要在代码安全点发生的用户空间事件定义明确，以进行上下文切换。这些时间和安全点在函数调用中体现出来。函数调用对go调度器健康来说至关重要。当前（使用go1.11或更低版本），如果运行任何不进行函数调用的紧密循环，会引起调度器和垃圾回收中的延迟。在适当的时间范围内进行函数调用直至关重要。

> 注意：:在1.12中有一个已经接受的[提案][proposal]，该提案在go调度器中为紧密循环的抢占使用非合作-抢占式技术。
 
 在go程序中有四类事件发生时需要go调度器制定调度策略。这并不意味着这些事件之一总是发生。这只意味着调度器有机会制定调度策略了。

- `go`关键字的使用
- 垃圾收集
- 系统调用
- 同步与编排

**`go`关键字的使用**

`go`关键字是创建go协程的方法。一旦协程创建，调度器便有机会做出调度策略。

**垃圾收集**

因为GC用其自己的协程运行，这些协程需要M上的时间运行。这造成GC产生大量混乱的调度。但是调度器对于go协程所做的事是非常聪明的，它会利用该情报做出明智决策。一个明智的决策是在GC中将接触堆的协程与那些不接触堆的协程进行上下文切换。当GC运行时，会制定很多调度决策。

**系统调用**

当go协程进行系统调用时，这会让go协程阻塞在M上，调度器有时能通过上下文切换将这个协程从M上换下来并换成一个新的协程。但是，有时需要一个新的M来继续执行在P中排队的go协程。在下一部分中将更详细地说明其工作原理。

**同步与编排**

如果原语、互斥调用或通道操作调用引起go协程阻塞，调度器可以通过上下文切换成一个新的协程运行。一旦该协程再次可运行，它会被重新放回队列并最终通过上下文切换回到M上。

[proposal]: https://github.com/golang/go/issues/24543
 
## 异步系统调用

当所使用的OS能处理异步系统调用时，一些诸如网络轮询器等可以更高效的处理系统调用。这通过 kqueue（MacOS）、epoll（Linux）或者 iocp（Windows）在各自的OS上完成。

我们今天所使用的好多系统都可以异步处理基于网络的系统调用。此即为网络轮询器名称的来源，因其主要用处既是处理网络操作。通过使用网络轮询器完成网络系统调用，当系统调用产生时，调度器可以阻止go协程阻塞M。其作用便是保证M能够执行P的LRQ中的协程，而不用创建新的M。这也降低了OS上的调度负载。

最好通过一个运行示例来了解其运行方式。

**图3**
![图三](https://www.ardanlabs.com/images/goinggo/94_figure3.png)

图3展示了基本调度图。goroutine-1 正在M上运行，另有超过3个go协程在LRQ中等待以获取M上的时间片。网络轮询器空闲着，无事可做。

**图4**

![图4](https://www.ardanlabs.com/images/goinggo/94_figure4.png)

在图4中，goroutine-1 需要网络系统调用，因此goroutine-1被挪到网络轮续器中，然后开始异步处理网络系统调用。当goroutine-1挪到网络轮询器之后，M便可以执行LRQ中另一个go协程。在该示例中，goroutine-2通过上下文切换到M上。

**图5**

![图5](https://www.ardanlabs.com/images/goinggo/94_figure5.png)

在图5中，网络轮询器完成了异步网路调用，goroutine-1被移回到P的LRQ中。当goroutine-1可以上下文切换回M上时，其所负责的相关go代码可以再次执行。最大的优点是，通过执行网络系统调用，不再需要额外的M。网络轮询器有个OS线程来高效地处理事件循环。


## 同步系统调用

要是go协程需要一个无法异步完成的系统调用该怎么办？在该示例中，无法使用网络轮询器，go协程的系统调用将会阻塞M。这很不幸但并无其他办法阻止这样的事情发生。一个无法异步完成的系统调用的示例是基于文件的系统调用。要是你正在使用CGO的话，类似调用C函数的其他情况也会阻塞M。

> 注意： Windows系统无法异步进行基于文件系统调用。在技术上来说，当在Windows上运行时，可以使用网络轮询器。

我们看一下同步系统调用（如文件I/O）是如何阻塞M的。

**图6**

![图6](https://www.ardanlabs.com/images/goinggo/94_figure6.png)

图6再次展示了基本调度图，但这次goroutine-1进行了同步系统调用，将会阻塞M1。i

**图7**

![图7](https://www.ardanlabs.com/images/goinggo/94_figure7.png)

在图7中，调度器会识别出goroutine-1已经引起了M的阻塞。此时，调度器将仍然绑定到正在阻塞的goroutine-1的M1从P中分离出来。然后调度器创建一个新的M2为P提供服务。此时，从LRQ中选择的goroutine-2被上下文切换到M2上。如果因为之前的交换M已经存在了，那么这个转换要比创建新的M要快。

**图8**

![图8](https://www.ardanlabs.com/images/goinggo/94_figure8.png)

图8中，goroutine-1产生的阻塞状态的系统调用完成。此时，goroutine-1可以移回LRQ然后再次由P提供服务。随后M1被放到一边以供后续类似这样的场景需求再次发生的时候使用。

## 任务窃取

调度器的另一方面是任务窃取调度。其在某些方面提供帮助，以保持调度高效。首先，所需要的最后一件事是M切回等待态，一旦这样的事情发生，操作系统回通过上下文切换将M从内核上取下。这意味着P无法再作任何工作，即使有go协程处于可运行态，一直到M被上下文切换到内核上。任务窃取还有助于平衡所有P上的go协程，从而更好的分配任务并更高效地完成。


我们来看一个示例。

**图9**

![图9](https://www.ardanlabs.com/images/goinggo/94_figure9.png)

图9中，我们有个多线程go程序，有2个P提供服务且每个P有4个go协程，其中一个go协程在GRQ中。如果P的一个服务迅速地执行其所有的go协程会发生什么？

**图10**

![图10](https://www.ardanlabs.com/images/goinggo/94_figure10.png)

图10中，P1没有go协程可执行。但P2的LRQ以及GRQ中仍有go协程处于可执行态，这时P1需要[窃取任务][steal work]。窃取任务规则如下。

**清单2**

```golang
runtime.schedule() {
    // only 1/61 of the time, check the global runnable queue for a G.
    // if not found, check the local queue.
    // if not found,
    //     try to steal from other Ps.
    //     if not, check the global runnable queue.
    //     if not found, poll network.
}
```

因此，基于清单2中的这些规则，P1需要检查在P2的LRQ中的go协程并取走一半。

**图11**

![图11](https://www.ardanlabs.com/images/goinggo/94_figure11.png)

在图11中，P1上有一半的go协程被取走，P2现在可以执行这些go协程。

如果P2完成了其所服务的所有go协程，但此时P1的LRQ空了会发生什么？

**图12**

![图12](https://www.ardanlabs.com/images/goinggo/94_figure12.png)

在图12中，P2完成了所有的任务，现在需要再窃取一些。首先，P2会先查找P1的LRQ，但并未发现go协程。接下来，P2会查找GRQ。然后找到goroutineI-9。

**图13**

![图13](https://www.ardanlabs.com/images/goinggo/94_figure13.png)

早图13中，P2从GRQ窃取到goroutine-9然后开始执行其任务。任务窃取的好处在于其让M保持繁忙而非空闲。这种任务窃取在内部被认为是周转M。JBD在其[任务窃取][working-steal]的博客文章中很好地解释了这种周转的其他好处。

[steal work]: https://golang.org/src/runtime/proc.go
[working-steal]: https://rakyll.org/scheduler/

## 实践示例
有了合适的机制和语义之后，我会向你展示这些机制和语义是如何结合在一起的，以及随着时间的推移，go调度器如何执行更多任务。想像一下用C语言编写的多线程程序，该程序管理2个相互传递消息的OS线程。

**图14**

![图14](https://www.ardanlabs.com/images/goinggo/94_figure14.png)

在图14中，两个线程相互传递消息。线程1获得了内核上的上下文切换，现在正在运行，并将信息发送给线程2。

> 注意：消息如何传递并不是重点。重点是随着编排的进行的线程状态。

**图15**

![图15](https://www.ardanlabs.com/images/goinggo/94_figure15.png)

在图15中，当线程1发送完消息时，它需等待回应。这会让线程1通过上下文切换从内核上取下并放到等待态。当线程2收到消息通知时，它被放到可运行态。现在OS可以运行上下文切换并将线程2放到内核上执行，这都是在内核2上发生的。接下来，线程2处理消息并向线程1发送一条新消息。

**图16**

![图16](https://www.ardanlabs.com/images/goinggo/94_figure16.png)

在图16中，随着线程2发送给线程1的消息被接受，又发生一次上下文切换。现在线程2通过上下文切换从执行态切换到等待态，同时线程1通过上下文切换从等待态切换为可执行态并最终进入执行态，这让它处理消息并返回一条新的消息。

所有的这些上下文切换和状态变更都需要执行时间，这限制了任务需要多快完成。因为每次上下文切换潜在性的增加了大约1000纳秒的延迟，每纳秒大约执行12条机器指令，在执行上下文切换时，你大约有1.2万条指令没有执行。因为这些线程也在不同的内核上来回切换，这些额外增加的延迟也导致缓存总线错误的机率较高。

我们来探讨一下用go协程和go调度器实现的同样的示例。

**图17**

![图17](https://www.ardanlabs.com/images/goinggo/94_figure17.png)

在图17中，两个go协程被编排成相互传递消息。G1在M1上进行上下文切换，并在内核1上运行，这可让G1执行其任务。G1的这些任务是将消息发送给G2。

**图18**

![图18](https://www.ardanlabs.com/images/goinggo/94_figure18.png)

在图18中，当发送完成时，G1需要等待响应。这导致G1从M1上通过上下文切换为等待态。当收到消息通知时，G2被放到可运行态。现在Go调度器执行上下文切换并将G2放到M1上运行，同样仍在内核1上运行。接下来G2处理消息并向G1返回一条新的消息。

**图19**

![图19](https://www.ardanlabs.com/images/goinggo/94_figure19.png)

在图19中，随着G2发送的消息被G1接收到，又发生了一次上下文切换。现在G2通过上下文切换从执行态切换为等待态，然后G1通过上下文切换从等待态切换为可执行态并最终进入执行态，这可以让它处理并返回一条新的消息。

表面上的事情并没有什么不同。无论使用线程还是go协程，都会发生同样的上下文切换与状态变更。然而，在使用线程和go协程上最主要的差异猛一看并不明显。

在使用go协程的案例上，使用同一个OS线程和内核被用于所有的处理。从OS的层面看，这意味着OS线程并没有被切换到等待态；因此，在使用线程进行上下文切换所损失的所有指令在go协程进行上下文切换时并没有损失。


本质上来说，go能在操作系统级别将IO/阻塞任务变成了CPU密集型任务。由于所有的上限为切换都发生在应用级别，对于每个上下文切换来说，我们并没有像线程那样损失1.2万左右的指令（平均）。在go中，同样的上下文切换只消耗约200纳秒或约2400条指令。调度器同时也有助于提高缓存总线效率与[NUMA][NUMA]。因此，在虚拟内核上无需更多线程。因为go调度器尝试用更少的线程并让每个线程来做更多的工作，随着时间的推移有可能完成更多任务，这有益于降低OS和硬件上的负载。

[NUMA]: http://frankdenneman.nl/2016/07/07/numa-deep-dive-part-1-uma-numa

## 总结


在考虑操作系统和硬件运行的错综复杂性上，go调度器在如何设计上真的非常了不起。
<!--TODO: 原句：The Go scheduler is really amazing in how the design takes into account the intricacies of how the OS and the hardware work. 认为翻译不太好，请后续优化 -->
随着时间的推移，在操作系统级别将IO/阻塞任务转为CPU密集型任务的能力使我们在利用更多PU容量方面获得了巨大胜利。所以，在具有虚拟内核的操作系统上无需更多线程。在每个虚拟内核只有一个OS线程的情况下期待完成更多工作（CPU和IO/阻塞绑定）是合理的。网络嘤嘤以及其他不需要阻塞OS线程的系统调用都可以这样做。

作为一个开发者，根据所处理的任务类型仍然需要理解你的应用正在做什么。你不能创建一个无限数量go协程的但又期待惊人性能的应用。少即是多，但理解这些go调度器语义，你可以做出更好的工程决策。在接下来的文章中，我会探讨以保守的方式利用并发性来获取更好性能的想法，同时仍然平衡添加到代码中复杂量。
