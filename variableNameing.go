package main

import (
	"fmt"
)

func main() {

	//以大写字母打头的标识符将被导出
	var Foo string = "bar" //Exported
	//以小写字母打头的不会被导出
	var foo string = "bar" //Not Exported

	var fileName string //Camel Case
	var FileNAME string //Pascal Case

	/**
	 * 		在Go程序中，经常使用指出了数据类型的简单变量名，这让程序员能够专注于逻辑而不是变量。在这种情况下，i表示整数（Integer），s表示字符串
	 * 	 	func a(f float64) float64
	 * 		func (t *Triangle) Area() float64
	 */

	/**
	 * 		接口，它们是具名的方法签名集合。在Go源代码中，接口名通常是这样得到的：在动词后面加上后缀er，形成一个名词。后缀er通常表示操作，因此这种命名方式表示表示操作，如Reader、
	 *	Writer和ByteWriter。在有些情况下，这样生成的接口名可能不是现成的英语单词，如果您在Go源代码中搜索，将找到诸如Twoer这样的接口名
	 */

	/**
	 *		对于要导出的函数，命名约定是尊重这样的事实的，即导入包后，就可通过包名和函数名来访问它。例如，在标准库中，math包就遵循了这样的约定：将计算平方根的函数命名为Sqrt而
	 * 	不是MathSqrt。这合乎情理，因为使用这个函数时，代码为math.Sqrt而不是math.MathSqrt，另外，只要通过这个函数名就能知道它是做什么的。即便不查看函数的实现，程序员也能轻易
	 * 	知道它是做什么的。
	 */

	/**
	 *	命名总是带有一定的主观性的，但花点时间考率如何命名总是值得的。给变量、函数和接口命名时，需要考虑的一些因素包括以下几点
	 *
	 * 	谁将使用这些代码？只是我自己还是整个团队？
	 * 	是否为当前项目制定了命名约定？
	 * 	不熟悉代码的人是否只需看一眼就能大致知道它是做什么的？
	 */

	/**
	 *	golint
	 */

	/**
	 *	使用godoc
	 * 		随着要开发的程序越来越复杂，要确保其品质优良，编写文档很重要。
	 * 		godoc是一款官方工具，可通过分析Go语言源代码及其中的注释来生成文档。由于文档是根据源代码生成的，这很大程度上避免了文档不同步的问题。
	 * 		godoc需要单独安装： go get golang.org/x/tools/cmd/godoc
	 * 		godoc能够生成很多不同的输出，其中包括html。标准库文档就是使用godoc生成的，而很多第三方网站都提供了有关开源项目的HTML格式文档。遵循godoc指定的约定意味着可以通过
	 * 	标准方式生成文档
	 * 		要学习如何编写文档，一种绝妙的方式是研究标准库编写文档的做法，安装godoc后，就可将任何标准库的文档输出到终端中。例如，要查看strings包的文档，可执行
	 * 	godoc strings
	 * 		这将把strings包的文档打印到标准输出中。linux中可结合grep命令
	 * 	godoc strings | grep "func Replace"
	 * 		通过启动一个Web服务器来查看标准库文档。在您没有连接到网络或连接速度有限时，这种做法是一个不错的选择
	 * 	godoc -http=":6060"
	 * 		执行这个命令后，您就可在浏览器输入地址http://localhost:6060/pkg/来查看标准库的文档
	 *
	 *
	 * 	Makefile
	 * 	第14章草过、工具篇、遇到再学
	 */

}
