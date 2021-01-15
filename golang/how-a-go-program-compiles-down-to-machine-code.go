# 一个 Go 程序是如何编译成机器码的
> Koen V. 约2年前
> translated by Turing Zhu

在 [Stream](https://getstream.io/) 里，我们广泛使用 Go，因为它大大提高了我们的生产效率。而且我们还发现通过使用Go，速度非常出色，并且从开始使用Go以来，我们已经实现了堆栈的核心任务部分，例如由gRPC，Raft和RocksDB提供支持的内部存储引擎。今天我们研究Go 1.11编译器及其将Go源代码编译为可执行文件的方式，以了解我们如何使用日常工作工具。
我们还将看到为什么Go代码如此之快以及编译器如何提供帮助。 我们将看一下编译器的三个阶段：

- 扫描器，将源代码转为符号表，以供解析器使用。
- 解析器，将符号转为抽象语法树，以供代码生成器使用。
- 代码生成器，将抽象语法树翻译为机器码。

> 注意：我们即将使用的包，如 `go/sanner`、`go/parser`、`go/token`、`go/ast`等，go 编译器并不用这些，主要只是用于工具操作 go 的源代码。但实际的 go 编译器有与之相似的语义。go编译器不使用这些包的原因是编译器曾经是由C写的，然后才转为go代码的，因此实际的go编译器仍能让人想起这些结构。

## 扫描器

各个编译器的第一步便是将原始源代码文本分解为符号，这由扫描器（也称为词法分析器）完成。符号可以是关键字、字符串、变量名、函数名等。每个合法的程序“单词”都由符号表示。具体的对于go来说，这可能意味着我们有 "package"，"main"，"func"等符号。每个符号都由其位置，类型和Go中的原始文本表示。Go 甚至允许我们使用`go/scanner`和`go/token`程序包在Go程序中自行执行扫描器。这意味着我们的代码在扫描之后，能检查我们的程序对于go编译器来说看起来是什么样。为此，我们会创建一个简单的程序，该程序打印Hello World程序的所有符号。该程序看起来像这样：

```golang
package main

import (
  "fmt"
  "go/scanner"
  "go/token"
)

func main() {
  src := []byte(`package main
import "fmt"
func main() {
  fmt.Println("Hello, world!")
}
`)

  var s scanner.Scanner
  fset := token.NewFileSet()
  file := fset.AddFile("", fset.Base(), len(src))
  s.Init(file, src, nil, 0)

  for {
     pos, tok, lit := s.Scan()
     fmt.Printf("%-6s%-8s%q\n", fset.Position(pos), tok, lit)

     if tok == token.EOF {
        break
     }
  }
}
```

我们会创建源代码字符串并初始化扫描源代码的`scanner.Scanner`结构体。 我们会多次调用`Scan()`，并打印符号的位置，类型和字面字符串，直到到达文件结尾（`EOF`）标记。 当我们运行该程序时，它将打印以下内容：

```golang
2:1   package "package"
2:9   IDENT   "main"
2:13  ;       "\n"
3:1   import  "import"
3:8   STRING  "\"fmt\""
3:13  ;       "\n"
4:1   func    "func"
4:6   IDENT   "main"
4:10  (       ""
4:11  )       ""
4:13  {       ""
5:3   IDENT   "fmt"
5:6   .       ""
5:7   IDENT   "Println"
5:14  (       ""
5:15  STRING  "\"Hello, world!\""
5:30  )       ""
5:31  ;       "\n"
6:1   }       ""
6:2   ;       "\n"
6:2   EOF     ""
[zxd:logs]$ go run turing.go 
1:1   package "package"
1:9   IDENT   "main"
1:13  ;       "\n"
2:1   import  "import"
2:8   STRING  "\"fmt\""
2:13  ;       "\n"
3:1   func    "func"
3:6   IDENT   "main"
3:10  (       ""
3:11  )       ""
3:13  {       ""
4:3   IDENT   "fmt"
4:6   .       ""
4:7   IDENT   "Println"
4:14  (       ""
4:15  STRING  "\"Hello, world!\""
4:30  )       ""
4:31  ;       "\n"
5:1   }       ""
5:2   ;       "\n"
5:2   EOF     ""
```

这里我们可以看到go解释器在编译程序时使用的内容。我们还可以看到，扫描器添加了分号，而分号通常在其他编程语言（如C）中放置。这解释了Go不需要分号的原因：扫描器可以智能地放置分号。

## 解释器

扫描完源代码之后，会将扫描后的符号传给解析器。解析器是编译器的一个阶段，它将符号转成抽象语法树（AST，Abstract Syntax Tree）。AST是源代码的抽象化表示。在AST中，我们能够看到程序结构，例如函数和常量声明。Go也向我们提供了用于解析程序和查看AST的软件包：`go/parser`和`go/ast`。 我们可以像这样使用它们来打印完整的AST：

```golang
package main

import (
  "go/ast"
  "go/parser"
  "go/token"
  "log"
)

func main() {
  src := []byte(`package main
import "fmt"
func main() {
  fmt.Println("Hello, world!")
}
`)

  fset := token.NewFileSet()

  file, err := parser.ParseFile(fset, "", src, 0)
  if err != nil {
     log.Fatal(err)
  }

  ast.Print(fset, file)
}
```

输出：

```golang
     0  *ast.File {
     1  .  Package: 1:1
     2  .  Name: *ast.Ident {
     3  .  .  NamePos: 1:9
     4  .  .  Name: "main"
     5  .  }
     6  .  Decls: []ast.Decl (len = 2) {
     7  .  .  0: *ast.GenDecl {
     8  .  .  .  TokPos: 2:1
     9  .  .  .  Tok: import
    10  .  .  .  Lparen: -
    11  .  .  .  Specs: []ast.Spec (len = 1) {
    12  .  .  .  .  0: *ast.ImportSpec {
    13  .  .  .  .  .  Path: *ast.BasicLit {
    14  .  .  .  .  .  .  ValuePos: 2:8
    15  .  .  .  .  .  .  Kind: STRING
    16  .  .  .  .  .  .  Value: "\"fmt\""
    17  .  .  .  .  .  }
    18  .  .  .  .  .  EndPos: -
    19  .  .  .  .  }
    20  .  .  .  }
    21  .  .  .  Rparen: -
    22  .  .  }
    23  .  .  1: *ast.FuncDecl {
    24  .  .  .  Name: *ast.Ident {
    25  .  .  .  .  NamePos: 3:6
    26  .  .  .  .  Name: "main"
    27  .  .  .  .  Obj: *ast.Object {
    28  .  .  .  .  .  Kind: func
    29  .  .  .  .  .  Name: "main"
    30  .  .  .  .  .  Decl: *(obj @ 23)
    31  .  .  .  .  }
    32  .  .  .  }
    33  .  .  .  Type: *ast.FuncType {
    34  .  .  .  .  Func: 3:1
    35  .  .  .  .  Params: *ast.FieldList {
    36  .  .  .  .  .  Opening: 3:10
    37  .  .  .  .  .  Closing: 3:11
    38  .  .  .  .  }
    39  .  .  .  }
    40  .  .  .  Body: *ast.BlockStmt {
    41  .  .  .  .  Lbrace: 3:13
    42  .  .  .  .  List: []ast.Stmt (len = 1) {
    43  .  .  .  .  .  0: *ast.ExprStmt {
    44  .  .  .  .  .  .  X: *ast.CallExpr {
    45  .  .  .  .  .  .  .  Fun: *ast.SelectorExpr {
    46  .  .  .  .  .  .  .  .  X: *ast.Ident {
    47  .  .  .  .  .  .  .  .  .  NamePos: 4:3
    48  .  .  .  .  .  .  .  .  .  Name: "fmt"
    49  .  .  .  .  .  .  .  .  }
    50  .  .  .  .  .  .  .  .  Sel: *ast.Ident {
    51  .  .  .  .  .  .  .  .  .  NamePos: 4:7
    52  .  .  .  .  .  .  .  .  .  Name: "Println"
    53  .  .  .  .  .  .  .  .  }
    54  .  .  .  .  .  .  .  }
    55  .  .  .  .  .  .  .  Lparen: 4:14
    56  .  .  .  .  .  .  .  Args: []ast.Expr (len = 1) {
    57  .  .  .  .  .  .  .  .  0: *ast.BasicLit {
    58  .  .  .  .  .  .  .  .  .  ValuePos: 4:15
    59  .  .  .  .  .  .  .  .  .  Kind: STRING
    60  .  .  .  .  .  .  .  .  .  Value: "\"Hello, world!\""
    61  .  .  .  .  .  .  .  .  }
    62  .  .  .  .  .  .  .  }
    63  .  .  .  .  .  .  .  Ellipsis: -
    64  .  .  .  .  .  .  .  Rparen: 4:30
    65  .  .  .  .  .  .  }
    66  .  .  .  .  .  }
    67  .  .  .  .  }
    68  .  .  .  .  Rbrace: 5:1
    69  .  .  .  }
    70  .  .  }
    71  .  }
    72  .  Scope: *ast.Scope {
    73  .  .  Objects: map[string]*ast.Object (len = 1) {
    74  .  .  .  "main": *(obj @ 27)
    75  .  .  }
    76  .  }
    77  .  Imports: []*ast.ImportSpec (len = 1) {
    78  .  .  0: *(obj @ 12)
    79  .  }
    80  .  Unresolved: []*ast.Ident (len = 1) {
    81  .  .  0: *(obj @ 46)
    82  .  }
    83  }
```

在上面的输出中，可以看到很多有关该程序的信息。 在`Decls`字段中，有文件中所有声明的列表，例如导入，常量，变量和函数。 在该示例中，我们只有两个：导入的`fmt`包和main函数。 为了进一步了解它，我们可以看一下该图，它是上述数据的表示，但仅包括类型，红色表示与节点相对应的代码：

main函数由三部分组成：函数名，函数声明和函数体。函数名以main标识符表示。函数声明由类型字段指定，如果我们指明的话，还包括参数列表和返回值类型。函数体由一系列的语句组成，包括了程序的所有代码行，在此示例中只有一行。在AST中，我们唯一的一行语句`fmt.Println`由很少的几个部分组成。该语句是一个`ExprStmt`，它是一个表达式，比如这里可以是函数调用，或者字面量，二元操作（如加减法），一元操作（如取负值）等等。 其他在函数调用参数中的也可以是一个表达式。我们的`ExprStmt`包含一个`CallExpr`，实际的函数调用。这又包含几部分，最主要的便是`Fun`和`Args`。Fun包含函数调用的引用，该示例中是`SlectorExpr`，因为我们从`fmt`包中选择了`Println`标识符。但是在AST中，编译器还不知道`fmt`是个包，它也是个变量。Args包含一个表达式的列表，这是函数的参数。该示例中我们将一个字面量字符串传递给函数，因此它由`STRING`类型的`BasicLit`表示。显而易见，从AST中我们可以推断出来很多有用的东西。这意味着之后我们也可以检查并找到AST，例如，文件中的函数调用。为此，我们打算使用`ast`包中的检查函数。 该函数将递归遍历抽象语法树并让我们检查所有节点的信息。要提取所有的函数调用，我们会用下面的代码：

```
package main

import (
  "fmt"
  "go/ast"
  "go/parser"
  "go/printer"
  "go/token"
  "log"
  "os"
)

func main() {
  src := []byte(`package main
import "fmt"
func main() {
  fmt.Println("Hello, world!")
}
`)

  fset := token.NewFileSet()

  file, err := parser.ParseFile(fset, "", src, 0)
  if err != nil {
     log.Fatal(err)
  }

  ast.Inspect(file, func(n ast.Node) bool {
     call, ok := n.(*ast.CallExpr)
     if !ok {
        return true
     }

     printer.Fprint(os.Stdout, fset, call.Fun)
     fmt.Println()

     return false
  })
}
```

我们在这里所做的工作是查找所有的节点并判断是否是**ast.CallExpr**，就是我们刚刚看到的表示函数调用的。我们将会使用printer包打印**Fun**成员中的函数名，如果他们是的话。该代码的输出将会是：**fmt.Println**。这实际上示例程序的唯一函数调用，因此事实上我们找到了所有的函数调用。在构造出来AST之后，编译器会用GOPATH或者对于Go 1.11以及更高版本来说可能是[模块](https://github.com/golang/go/wiki/Modules)来处理所有的导入。之后会做类型检查，以及应用一些初步优化以让程序运行更快。

## 代码生成

在处理完导入以及类型检查之后，我们确认该程序是合规的Go代码然后可以开始将AST转换为机器码（或伪代码）的过程。该处理过程中的第一步是将AST转换为该程序的低级表示，尤其是转换为单一静态赋值(SSA，Static Single Assignment)表。该中间表示并非最终的机器码，但已经表示了大多数的最终机器码了。SSA有个属性集，它能更轻松地应用优化，其中最重要的是始终在使用变量之前定义变量，并且每个变量只分配一次。当SSA的初始版本构造之后，会应用一系列的优化。应用这些优化以确保处理器执行这些代码片能够更快更简单。比如说，死代码，如**if (false) { fmt.Println("test")}** 可以淘汰掉因为该代码不会执行。另一个优化的例子是确保移除nil检查因为编译器能够证明这些不为false。我们看下这个示例程序的SSA和一些优化过程：

```
package main

import "fmt"

func main() {
  fmt.Println(2)
}
```

可以看到，该程序只有一个函数与导入。当运行时会打印2。但对于查看SSA来说这个示例程序足够了。*注意：只有main函数的SSA会显示，因为这是最有意思的部分。*要展示生成的SSA，需要给我们想看SSA的函数设置**GOSSAFUNC**环境变量，本示例中是main函数。我们也需给编译器传递-S标示，这样编译器才能打印代码并创建HTML文件。我们也会为Linux 64-bit编译该文件，以保证机器码和你在这里看到的一样。所以，要编译这个文件，会这样运行：**$ GOSSAFUNC=main GOOS=linux GOARCH=amd64 go build -gcflags "-S" simple.go** 这会打印所有的SSA，但只生成一个ssa.html文件来交互，所以我们会用这个文件。

当打开ssa.html文件时，会展现一系列步骤，大多数都是折叠的。第一步是从AST生成SSA；较低的步骤将非机器专用的SSA转换为机器专用的SSA，同时genssa是最终生成的机器码。开始阶段的代码如下：

```
b1:
	v1  = InitMem <mem>
	v2  = SP <uintptr>
	v3  = SB <uintptr>
	v4  = ConstInterface <interface {}>
	v5  = ArrayMake1 <[1]interface {}> v4
	v6  = VarDef <mem> {.autotmp_0} v1
	v7  = LocalAddr <*[1]interface {}> {.autotmp_0} v2 v6
	v8  = Store <mem> {[1]interface {}} v7 v5 v6
	v9  = LocalAddr <*[1]interface {}> {.autotmp_0} v2 v8
	v10 = Addr <*uint8> {type.int} v3
	v11 = Addr <*int> {"".statictmp_0} v3
	v12 = IMake <interface {}> v10 v11
	v13 = NilCheck <void> v9 v8
	v14 = Const64 <int> [0]
	v15 = Const64 <int> [1]
	v16 = PtrIndex <*interface {}> v9 v14
	v17 = Store <mem> {interface {}} v16 v12 v8
	v18 = NilCheck <void> v9 v17
	v19 = IsSliceInBounds <bool> v14 v15
	v24 = OffPtr <*[]interface {}> [0] v2
	v28 = OffPtr <*int> [24] v2
If v19 → b2 b3 (likely) (line 6)

b2: ← b1
	v22 = Sub64 <int> v15 v14
	v23 = SliceMake <[]interface {}> v9 v22 v22
	v25 = Copy <mem> v17
	v26 = Store <mem> {[]interface {}} v24 v23 v25
	v27 = StaticCall <mem> {fmt.Println} [48] v26
	v29 = VarKill <mem> {.autotmp_0} v27
Ret v29 (line 7)

b3: ← b1
	v20 = Copy <mem> v17
	v21 = StaticCall <mem> {runtime.panicslice} v20
Exit v21 (line 6)
```

该示例程序已经生成了相当多的SSA（总共35行）。但是大多数是模版并且可以删除的（SSA的最终版本只有28行，同时机器代码最终版本只有18行）。每个v都是一个新的变量且可以点开看下它是怎么用的。**b\*\*表示代码块，该示例中有三个代码块：b1\*\*，b2\*\*和b3\*\*。b1是一直执行的。b2和b3是条件代码块，在b1的结尾处可以通过`If v19 -> b2 b3(或者)`查看。可以点开`v19`查看下它的定义。我们可以看到，它被定义为`IsSliceInBounds v14 v15`，且通过查看[go编译器源代码](https://github.com/golang/go/blob/3fd364988ce5dcf3aa1d4eb945d233455db30af6/src/cmd/compile/internal/ssa/gen/genericOps.go#L411)我们知道`IsSliceInBounds`检查**`0 <= arg0 <= arg1`。也可以点开`v14`和`v15`查看其定义，在这里我们看到`v14 = Const64 [0]`；**Const64** 是个64位常量整数。**v15**以同样方式定义但值是**1**。实际上我们有**0 <= 0 <= 1**，很明显是**true**。编译器也能证明这个，且当查看**opt**语句（机器的单独优化）时，**v19**被重写为**ConstBool [true]**。这会在**opt 死代码**中使用，**b3**被删除因为因为之前显示的条件中的**v19**始终为true。现在，我们看看在把SSA转换为特定于机器的SSA之后，Go编译器进行的另一种更简单的优化，因此这是amd64体系结构的机器码。为此，我们会比较较低的死代码和较低的死代码。 这是较低阶段的内容：

```
b1:
	BlockInvalid (6)
b2:
	v2 (?) = SP <uintptr>
	v3 (?) = SB <uintptr>
	v10 (?) = LEAQ <*uint8> {type.int} v3
	v11 (?) = LEAQ <*int> {"".statictmp_0} v3
	v15 (?) = MOVQconst <int> [1]
	v20 (?) = MOVQconst <uintptr> [0]
	v25 (?) = MOVQconst <*uint8> [0]
	v1 (?) = InitMem <mem>
	v6 (6) = VarDef <mem> {.autotmp_0} v1
	v7 (6) = LEAQ <*[1]interface {}> {.autotmp_0} v2
	v9 (6) = LEAQ <*[1]interface {}> {.autotmp_0} v2
	v16 (+6) = LEAQ <*interface {}> {.autotmp_0} v2
	v18 (6) = LEAQ <**uint8> {.autotmp_0} [8] v2
	v21 (6) = LEAQ <**uint8> {.autotmp_0} [8] v2
	v30 (6) = LEAQ <*int> [16] v2
	v19 (6) = LEAQ <*int> [8] v2
	v23 (6) = MOVOconst <int128> [0]
	v8 (6) = MOVOstore <mem> {.autotmp_0} v2 v23 v6
	v22 (6) = MOVQstore <mem> {.autotmp_0} v2 v10 v8
	v17 (6) = MOVQstore <mem> {.autotmp_0} [8] v2 v11 v22
	v14 (6) = MOVQstore <mem> v2 v9 v17
	v28 (6) = MOVQstoreconst <mem> [val=1,off=8] v2 v14
	v26 (6) = MOVQstoreconst <mem> [val=1,off=16] v2 v28
	v27 (6) = CALLstatic <mem> {fmt.Println} [48] v26
	v29 (5) = VarKill <mem> {.autotmp_0} v27
Ret v29 (+7)
```

在HTML文件中，一些行输出是灰色的，这意味着这些行在下面的某一语句中会移除掉。比如，**v15 (MOVQconst [1])**便是灰色输出。随后点击**v15**我们看到v15没有用到，且**MOVQconst**和我们之前看到的**Const64**本质上是一样的指令，只是特定于机器的**amd64**。因此将**v15**设置为**1**。但**v15**没有被用到，所以它是没用的（死的）代码且可以移除掉。go编译器应用大量的这种优化。因此，从AST生成的第一个SSA可能并不是最快的实现，编译器会将SSA优化为一个更快的版本。HTML文件中的每一个语句都是可能发生提速的语句。如果你对Go编译器的SSA感兴趣，可以阅读[Go编译器的SSA源代码](https://github.com/golang/go/tree/master/src/cmd/compile/internal/ssa)。这里定义了所有的操作，以及优化。

## 总结
有了Go编译器及其优化的支持，Go是一种非常高效的语言。要学习更多Go编译的知识，[源代码](https://github.com/golang/go/tree/master/src/cmd/compile)是最好的README。如果想了解为何Stream使用Go，尤其是从Python切换到Go，请阅读[我们发布的博客-切换到Go](https://getstream.io/blog/switched-python-go/)。
