# Go 与 C++ 的对比和比较

> [ETHAN SCULLY][ethan-scully] at 2020-08-06
> 
> translated by Turing Zhu
>
> source article: [Go vs C++ Compared and Contrasted
](go-vs-cpp)

[ethan-scully]: https://careerkarma.com/blog/author/ethan-scully/
[go-vs-cpp]:  https://careerkarma.com/blog/go-vs-c-plus-plus/

---


- [Go 与 C++ 概要](#summary)
- [Go(Golang) 编程](#go-programming)
- [C++编程](#cpp-programming)
- [Go 与 C++ 比较](#go-cpp-compare)
	- [Go vs C++: 速度与可读性](#go-cpp-speed-readability)
	- [C++ vs Go: 性能](#cpp-go-performace)
	- [Go vs C++: 安全性](#go-cpp-security)
	- [C++ vs Go: 应用](#cpp-go-application)
	- [Go vs C++: 社区](#go-cpp-community)
- [FAQ](#faq)

##  <a name="summary"></a>Go 与 C++ 概要
Go 是一门简单、紧凑且通用的语言。而 C++ 是一门快速且复杂的通用编程语言。Go 和 C++ 都是静态类型语言且都有强大的社区。C++ 广泛用于各种应用，而 Go 主要用于 Web 后端。

---

C++ 使用广泛。作为一门系统编程语言，C++ 是大量程序、计算任务以及其他编程语言的基石。C++ 应用于许多平台之上，并且被用于开发从视频游戏到驱动太空探测器程序的各种程序。C++ 已经被使用了很长时间，并且过了很长时间。

<!-- 这里 许多平台那里， 各种事情，以及最后一句后半句翻译并不是太恰当-->

Go 编程（或着叫 Golang）在编程领域是全新的。Go 由 Google 创造，旨在作为取代 C++ 作为通用型系统编程语言，并且也是为专门解决这一问题而构建的。那么，哪个更好呢：是弱者还是年老的冠军？


##  <a name="go-programming"></a>Go(Golang) 编程
![golang-logo](https://744025.smushcdn.com/1245953/wp-content/uploads/2019/05/go-logo.png?lossy=1&strip=1&webp=1)

作为编程语言来说，Golang 相当的新颖。Go 是由 Rob Pike 、Robert Griesemer 和 Ken Thompson 专门为 Google 创建的。它是一门静态类型、编译型、通用型的编程语言，更像 C++。Go 语言的编译器最初由 C 编写的，但现在也是用 Go 写的，这让 Go 语言能够做到自我做主。

Go，以及其许多IDE和库，都是通过有吸引力的开源许可证进行发布的。

Go 是专为现代多核处理器而设计的。Go 语言支持并发编程，这意味着它可以使用不同的线程同时运行多个处理过程，而不是同一时刻只运行一个任务。它还具有延迟垃圾回收功能，可以进行内存管理以防止内存泄漏。

##  <a name="cpp-programming"></a>C++编程
![cpp-logo](https://744025.smushcdn.com/1245953/wp-content/uploads/2019/12/cplusplus-logo.png?lossy=1&strip=1&webp=1)

C++ 是世界上最广泛使用的编程语言之一。是一门编译型、中级、面向对象的编程语言，在构建时会考虑性能与效率。C++ 用于构建各种程序。快速且友好的 C++（而且是 C 的堂兄弟）构成了计算机世界很大一部分的骨干。

C++ 是很久之前创建出来的，是 1979 年一名丹麦的计算机科学家打算做 一个 C 的扩展以使其能使用类的时候。现在 C++ 被用在各种地方，甚至被用来写其他语言的编译器和解释器。

##  <a name="go-cpp-compare"></a>Go 与 C++ 比较

现在我们对每种语言的起源都有了一点了解，让我们把它们放在一起，看看它们在以下几类中是如何并驾齐驱的：

> **更多**:  [C++ 和 C：你应该学哪一个?][cpp-vs-c]

[cpp-vs-c]: https://careerkarma.com/blog/c-plus-plus-vs-c/

### <a name="go-cpp-speed-readability"></a>Go vs C++: 速度与可读性

C++ 被称为 DIY 语言，虽然他不具备许多特性，但是如果你足够了解这门语言的话，你可以构建你想要的任何特征。

同样要注意的是，C++ 被认为是[_中级语言_][middle-level]，因此它不像高级语言那样具有语言性和直观性，但也不像汇编语言那么粗糙。

然而，这意味着和高级语言比较起来，C++ 编写代码更复杂。在像 Python 这样的语言中，需要几行代码中的东西在 C++ 中可能需要很多。

Go 代码更紧凑。他围绕着简单性和可扩展性而构建，尽管去掉了不必要的方括号和圆括号，但仍然减少了出错的余地。

和 C++ 一样， Go 也是静态类型的，这意味着程序员必须为每个变量定义类型。但是，和 C++ 比起来，Go 更容易学习与编码，因为 Go 更简单与紧凑。Go 同时还有一些内置特性无需为每个项目重写（如垃圾收集），并且这些特性运行很好。

另一个比较是编译时间。C++ 的编译时间非常慢。而编译时间依赖于实际编码的内容。和 C++ 比起来 Go 的编译时间明显更快。

由于所写的代码在运行之前需要进行编译，并且每次修改之后需要重新编译，因此编译时间影响编码速度。当需要一遍又一遍地运行代码来找出 C++ 代码中的一个缺少的分号时，编译时间会很快堆积起来。

也值得提及数据结构。C++ 是众所周知且都熟悉的面向对象结构，而 Go 却是过程式的并发型编程语言。和 C++ 不同，Go 没有带构造器和析构器的类。

[middle-level]: https://careerkarma.com/blog/high-level-and-low-level-languages/

### <a name="cpp-go-performace"></a>C++ vs Go: 性能

和其他的高级语言相比较而言，Go 相当快。其编译、静态类型以及高效的垃圾回收使其变得极其快。Go 也擅长内存管理；它有指针但没有引用。毫不夸张的说，Go 比它的那些解释型和动态型的编程语言小伙伴快了将近四倍。

也就是说，在速度方面，和 C++ 非常接近（以及 C 语言）。所有的时间都花在编码和编译上。因为 C++ 难编码，中级语言，更接近于机器码：当编译时，他更适合于机器码的嵌套。

C++ 也缺少那些让编码更容易但会给生成的程序增加阻力的特性。当提及运行时，C++ 轻量，精简且快速。

Go 配备了所有的能让你在编码过程中使生活更容易的零件和部件，因此，在运行时会慢一些。其中最大的一块是他的虽然很好但很慢的垃圾收集器。

> **更多: ** [Python vs Java vs C](https://careerkarma.com/blog/python-vs-java-vs-c/)

While garbage collecting is normally a death knell— signaling a slow-performing language—Go’s is highly optimized. However, it’s still a garbage collector and it will still slow down the code compared to not having one at all. 

公平的说，Go isn’t magnitudes slower than C++。除非你的程序必须以最大限度的提高速度，否则的话，Go 会和C++运行的一样。除非做大量的计算，不然的话速度上的差异不会引起太大的注意。

### <a name="go-cpp-security"></a>Go vs C++: 安全性

C 语言程序中一些很严重的漏洞包括利用缓冲区溢出。这是当缓冲区加载了很多信息且最后信息被写入到相邻的缓冲区中时。这会造成崩溃，或者像很多人发现的那样，一个可以越权访问的漏洞。

Go 具有内置的限制可以帮助防止这样的问题。例如，Go 不允许指针运算。你无法通过指针的值遍历数组（必须通过索引访问这些元素）。以这种方式执行操作会迫使程序员使用包含边界检查的方法，以防止溢出。

然而也应该注意到，缓冲区溢出并不是所有 C++ 程序固有的漏洞。在 Go 中强制执行的方法在 C++ 中同样可能，唯一的区别是 C++ 允许程序员懒散并制造这些漏洞。

### <a name="cpp-go-application"></a>C++ vs Go: 应用

C++ 依然能抵挡住 Go 冲击的一个主要原因是拥有无穷无尽的应用。C++ 毫无隐瞒。程序员以及随后的程序可以访问源代码自身的任何部分，并且机器也可以运行源代码。

目前尚未有任何开放或关闭的内置特性，这还是一块创建程序与系统的干净板子。这就是为何 C++ 甚至可能创造出开源系统的原因；你有权访问任何地方。

另一方面，Go 更像是一个密封的系统。进入 Go 的内部工作机制要困难的多。例如，Go 的臭名昭著的垃圾收集就是如此。如果程序员想要修改垃圾回收的机制，或者确定它是否还在，他们会很难做到这一点。

尽管 Go 是一门极好的语言，但其并不是为了设计成和 C++ 那样那么“低”的功能。因此， Go 并不像 C++ 那样使用广泛，而且当前最常用到 Go 的地方就是 Web 后端。

### <a name="go-cpp-community"></a>Go vs C++: 社区

C++ 已经使用很长一段时间了。其后是一个很大的社区，因此，这里有几乎所有有关 C++ 问题的答案。如果你需要集成，一些人可能已经做了，或者更可能的事，你现在正在集成的已经有特性集成到你正在写的代码里面了。

> **更多:** [» MORE:  Python vs. Java vs. C++](https://careerkarma.com/blog/python-vs-java-vs-c/)

然而，？？？。C++ 太陈旧了； 因此，许多库、模块和教程都是过时的。找到一种既实用又不过时的方案取决于你。

Go 还比较新颖，在这语言之后使用案例和人员都不多。一直到最近，由于文档不多，许多程序员对这门语言都不大感兴趣。

但是，Go 的库要比 C++ 的库小很多，因为 Go 是一门新语言。Go 并没有遗留在互联网上的一些陈旧的开发工具、建议和集成。你找的所有有关 Go 的东西都是新的且是最前沿的。所有存在的 Go 代码都可以运行且写的都符合现代开发规范。

Go 的社区也比较活跃。毕竟是一门新兴语言， 社区周围仍然期待发现 Go 可以用来做什么，所有 C++ 已经有的内容现在正在由 Go 程序员和开发人员实现。成为一种新语言的一部分是令人兴奋的，因为还有很多角落有待探索以及特性有待开发。

如果对你来说使用任意一门语言开发是有趣的，不要犹豫，选择它。尽管 Go 并不能那么快替代 C++，但是 Go 仍然被经常使用且需求量很大。虽然它们可能有各自不同的语言强项，但可以很好地组合在一起，你不会错的。


##  <a name="faq"></a> FAQ

### C# vs C++: 谁的速度更快？

C# 比 C++ 更快吗？通常来说，C++ 要比 C# 快，因为 C++ 是低级状态。但是，为了获得更高的性能，必须采用 C++ 更低的语言特性，并在微观层面优化。

### C# vs C++: 谁的性能更好？

因为 C# 是比 C++ 更高级的语言，其码代码时间更少。但 C++ 程序更快，在趋势上 C++ 码代码时间更长。

### Go vs Golang？

由于其网站名-[goalng.org](https://tip.golang.org/doc/faq#go_or_golang) 的缘故，该编程语言通常是指“Golang”，但其实际上被称为“Go”。

### C vs C++: 速度

C++ 是 旧语言 C 的强化。因为 C++ 支持面向对象以及像多态、抽象数据类型和封装，在趋势上 C++ 要比 C [更快](https://www.educba.com/c-vs-c-plus-plus-performance/)。

### C vs C++: 性能

对许多人来说 C++ 是[编程语言的一个更好的选择](https://careerkarma.com/blog/c-plus-plus-vs-c/)，因为 C++ 有很多特性以及应用。此外，许多人发现 C++ 比 C 更容易学习和使用。












