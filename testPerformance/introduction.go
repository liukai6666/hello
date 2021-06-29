/**
 * 	测试的重要性
 * 		通过测试代码的功能，开发人员能够在很大程度上确定程序是有效的。另外，每次修改代码后，开发人员都可运行测试，确认没有引入Bug和衰退。
 * 		常用的测试有多种
 * 			单元测试
 * 			功能测试
 * 			集成测试
 * 			
 * 		单元测试
 * 			单元测试针对一小部分代码，并独立地对它们进行测试。通常，这一小部分代码可能是单个函数，而要测试的是其输入和输出。典型的单元测试可能指出，如果给函数x提供这些值，它应返回
 * 		这个值。在确认程序的最小的构件按期望的方式运行。在程序增大和变化过程中，单元测试是发现衰退的绝佳方式。衰退是修改过程中引入的Bug或错误。
 * 			衰退意味着代码在修改前有效，但修改后无效了。单元测试通常能够发现衰退，因为它们测试的是程序的最小组成部分。
 * 
 * 		集成测试
 * 			集成测试通常测试的是应用程序各部分协同工作的情况。集成测试还检查诸如网络调用和数据库连接等方面，以确保整个系统按期望的那样工作。通常，集成测试比单元测试更难编写，因为这
 * 		些测试需要评估应用程序依赖的各个部分。
 * 
 * 		功能测试
 * 			功能测试通常被称为端到端测试或由外向内的测试。这些测试从最终用户的角度核实软件按期望的那样工作，他们评估从外部看到程序的运行情况，而不关心软件内部的工作原理。
 * 	testing包
 * 		Go语言在标准库中提供了testing包，它还支持命令go
 * 		与testing包相关的设计良好的约定
 * 			第一个约定
 * 				Go测试与其测试的代码在一起。测试不是放在独立的测试目录中，而是与它们要测试的代码放在同一个目录中。
 * 			测试文件的命名
 * 				在测试的文件的名称后面加上后缀_test
 * 				例如 strings.go、测试文件名为strings_test.go、并位于文件strings.go所在的目录中。
 * 			第二个约定：
 * 				测试函数以单词Test开头
 * 				例如
 * 				func TestTruth(t *testing.T) {}
 * 		测试失败时，命令go test提供了一些有用的信息：测试名、文件名以及导致测试失败的代码所在行。
 * 	基准测试
 * 		Go提供了功能强大的基准测试框架，能够让您使用基准测试程序来确定完成特定任务时性能最佳的方式是哪一种，它能够让您反复的运行函数，从而建立基准。您无须指定运行函数的次数，因为
 * 	基准测试框架将通过调整它来获得可靠的数据集。基准测试结束后，将生成一个报告，指出每次操作耗用了多少ns
 * 		基准测试名以关键字Benchmark开头，它们接受一个类型为B的参数，并对函数进行基准测试。
 * 		要运行基准测试，必须在命令go test中指定参数-bench
 * 	提供测试覆盖率
 * 		测试覆盖率是度量代码测试详尽程度的指标，它指出了被测试执行了的代码所在的百分比值
 * 		go test -cover
 */ 