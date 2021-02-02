# Golang 和 Dart 的区别

> [https://www.geeksforgeeks.org/difference-between-golang-and-dart/][source-article]
> 
> 作者：[sam_2200][sam_2200]，翻译: Turing Zhu
> 
> 上次更新: 2020-11-24


[source-article]: https://www.geeksforgeeks.org/difference-between-golang-and-dart/
[sam_2200]: https://auth.geeksforgeeks.org/user/sam_2200/articles


[Go][Go] 是一种程序化编程语言。由 Google 的罗伯特·格里塞默(Robert Griesemer)，罗伯·派克(Rob Pike)和肯·汤普森(Ken Thompson)在2007开发并于2009年作为一门开源编程语言发布。 由于高效的依赖管理,程序由包组成。该语言也支持和动态语言相似的采用模式环境。 Go在语法上和C相似，但有内存安全，垃圾收集，结构化类型以及CSP风格的并发。Go也叫**Golang**。

[Go]: https://www.geeksforgeeks.org/go-programming-language-introduction/

[Dart][Dart] 也是一门最初由Google开发的开源编程语言。它适用于服务端和用户端。Dart SDK来自其编译器--Dart 虚拟机和 dart2js 实用工具（旨在生成与 Dart 脚本等效的 JS，以便可以在不支持 Dart 的网站上运行）。Dart支持面向对象编程语言特性，比如类、对象、接口等。

[Dart]: https://www.geeksforgeeks.org/introduction-to-dart-programming-language/

## Golang VS Dart

| Go | Dart |
| :--- | :--- |
| Go是一种并发和过程化编程语言。 | Dart 是一种面向对象编程语言。 | 
| Go用于跨越大规模网络服务器和大型分布式系统编程。 |  现如今 Dart 与 Flutter 一起广泛用于开发移动应用程序。 |
| Go 没有带有构造器和析构器的类。| Dart 有带有构造器和析构器的类。|
| Go 语言为内存分配提供自动垃圾收集。 | 垃圾收集由 Dart 虚拟机执行。|
| Go 语言有指针，但不包括任意指针。| Dart 也是只有指针但不包括任意指针。|
| Go 语言中 map 以引用方式传递。 | Dart 中 map 以值的方式传递。 |  
| 既不支持函数重载也不支持自定义操作符。| Dart 也是既不支持函数重载也不支持自定义操作符。|
| 不支持常量与易失性限定词。| Dart支持常量和默认值未包含在此列表中，因为Dart的未来版本可能支持非常量默认值。|
| 不使用头文件。go 使用包替代头文件，使用 import 来导入外部包。| Dart 也使用包。|
| 没有 while 和 do-while 语句。但 for loop 可以作为 while 循环使用。| Dart 有 while 和 do-while 语句。|
| Go 有 go协程和通道。| Dart/Flutter 是单线程的，所以无法共享全局变量。|
| Go 不支持继承。 但提供了嵌入式的替代方案。| Dart 支持继承。|  
