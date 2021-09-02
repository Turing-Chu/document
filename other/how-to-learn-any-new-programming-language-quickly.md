# 如何快速学习一门新的编程语言

> [https://medium.com/better-programming/how-to-learn-any-new-programming-language-quickly-94996895669b][src]
>
> 原作者: [Bob Roebling][author]
>
> 如果您计算机英语阅读能力不错，建议您阅读原文


![img-1](https://miro.medium.com/max/1400/0*UE_JjOy_bS3mJzvY)

<center>*图片由[Unsplash][unsplash]上的[Clément H][clement]提供*</center>


本文假设你已经至少会一门编程语言；但这里的概念会帮助你开始编程。

我将我还上学的时候的一位老师告诉我的一些话分享给新开发者：**你所学的最难的编程语言将会是你的第二语言**。

先别沮丧 - 这意味着，当你开始学习编程时，关于编程的这些想法都会先入为主。最终，你会做更多的语法连接和假设。因此，当学习第二语言时，你必须“忘记”这些假设。如果你打算尝试学习第二或第三语言，请试着记住这个。


[src]: https://medium.com/better-programming/how-to-learn-any-new-programming-language-quickly-94996895669b
[author]: https://medium.com/@broebling?source=post_page-----94996895669b----------------------
[clement]: https://unsplash.com/@clemhlrdt?utm_source=medium&utm_medium=referral
[unsplash]: https://unsplash.com/?utm_source=medium&utm_medium=referral

--- 

## 编程解剖

编程语言很多，[轻易就上5000][more_lang]，但[TIOBE 清单][tiobe]只列出来前250种。前20种语言都有相似的标准库，除非特例。

我认为，考虑编程的最佳方法是剔除所有多余的“东西”，这样就只剩下必要的了。

[more_lang]: https://codelani.com/posts/how-many-programming-languages-are-there-in-the-world.html
[tiobe]: https://www.tiobe.com/tiobe-index/

--- 

## 基本知识

每种语言的各个方面都可以降低到`true`或`false`。为什么？因为电这样工作 - 你或者充电或者不充。内存以`0`和`1`的格式存储值，这个位要么被充电了要么没有。

8个位等于1个字节，这足以列出[ASCII表][ascii]中的任何字符。位以此顺序翻转，这可以表示字符的`十`进制值。计算机知道如何将十进制转换为字符。

![](https://miro.medium.com/max/1400/1*V6gi4Fk0uy-c-xkFaB88AQ.png)

<center>*基本二进制表示展示了单词 Hello 如何创建的*</center>

理解了这个概念之后，“为什么”会让剩下的部分对你来说更容易。

[ascii]: http://www.asciitable.com/

--- 

## 工具

![](https://miro.medium.com/max/1400/0*PkXTSWp4sLcu6xvt)
<center>*图片由[Unsplash][unsplash]上的[Fleur][fleur]提供*</center>

工具都一样，而且可以按任意顺序学习，这是我通常采用的顺序。

[fleur]: https://unsplash.com/@yer_a_wizard?utm_source=medium&utm_medium=referral

### 变量

看起来足够简单，但严格来说，怎么创建一个变量？

### 操作符

什么是操作符？怎么使用操作符？可以假设有基本数学操作符，但是逻辑操作符呢？“ AND”操作符是拼写为 “and” 还是 “AND”，还是使用诸如 “&&” 之类的符号？


### 条件

出乎意料的是，我在Swift和Python上阅读最多的文章都与策略制定有关。 接下来需要知道的是如何在程序中做出决定。 您尝试学习的语言是否使用传统的“ if / else if / else”或更Python化的语言（例如“ if / elif / else”）？ 您的语言是否带有“ switch”或“ guard”声明？

###  循环

如何遍历重复的任务？ 语言是否包含for循环，while循环，do-while循环或for-each语句？

### 函数

是否可以创建函数？ 如果可以，该怎么做？ 您如何在这些函数中包含参数？ 知道如何正确使用函数将节省您的时间，并使您的生活变得更加轻松。

### 类和结构体

这种语言是否理解类或结构体的概念？ 这看起来是一个愚蠢的问题，但某些语言都没有，或只有一种。 如果可以，如何创建类或结构体？ 类是否需要构造函数或init方法？

### 错误处理

错误不可避免。 当错误发生时，该语言是否具有强大的错误处理解决方案，将如何使用它？ 是“ try / catch”，“ try / except”还是别的？ 是否有其他条款（例如“其他”或“最终”）允许使用其他错误选项？

### 测试

你如何测试你的代码？是否有有内置库来做测试或是否必须下载单独的工具？

所有这些工具都应该使用最流行的编程语言。 即使是最老的的语言（例如COBOL）也拥有大多数这些工具，但是它们可能被称为不同的名称，例如段落或习字簿

![](https://miro.medium.com/max/1400/0*yhimPCQv2w-pMUlV)

<center>*图片由[Darius Soodmand][Darius Soodmand]提供于[Unsplash][unsplash]*</center>

[Darius Soodmand]: https://unsplash.com/@dsoodmand?utm_source=medium&utm_medium=referral

## 越来越好

一旦理解了这些工具，接下来你需要做的就是使用这些工具来写 APP。你可以通过阅读文档了解语言，但直到使用这门语言写一些小应用之前你都不是真正了解这门语言。

编写应用可以强迫你像 X 编程人员一样思考。可以这么说，因为我在 C 中用了 C++ 的类并且看了些 C++ 的文档而熟悉 C++，但直到我用 C++ 写了个应用，并且使用了该语言的一些特定特性我才真正的了解 C++。

21点是不错的入门项目。21点需要变量、操作符、条件语句、循环（基于玩家数）、函数、类/结构体以及错误处理。可以包含潜在故障的测试用例，如执行卡片越界。

其他合适的入门项目可以是回忆对对碰、快艇骰子或者自动售货机。

其他更多高级特性，试着重写大富翁的游戏。更多的是要关注机制与文本化。

关键是要记住，如果低估任务的难度（比如跳过21点的双下或分割功能），你只是限制了自己对语言的理解

## 还有其他的么？

上面所列的并不是一门语言所提供的所有东西。事实上，你可以使用上面列出的工具来编写任何内容，但是标准库中的其他额外功能会更容易些。大多数标准库都有一样的功能，所以你可以依赖于语言之间的相似名字。

一门语言使用的越多，标准库信息发现的就愈多，但要保证先学工具。当学习一门语言时，试着找出它的优点和缺点。这些将帮助你理解针对特定问题使用哪种语言。


需要快速进行数据科学？看些 Python 包或者 R 语言。需要写一个快速的服务？可以看看 C 或者 Go。web服务怎么样？看看 Java 或者 Python。

我不是只看语言就知道的。我通过使用这些语言才学会了这些。

由于到目前为止这可能是我最短的一篇文章，我将给你一个学习一门新语言的挑战。祝你好运！


