/**
 *	govendor : 常用的包依赖管理工具
 *
 * 	安装
 * 	go get -u -v github.com/kardianos/govendor
 *
 * 	govendor命令文件：{$GOPATH}/go/bin/govendor init
 *
 * 	#初始化vendor目录、此时将创建vendor文件夹
 *	govendor init
 *
 * 	govendor list
 * 	e ： 扩展包、未加进来的
 * 	v ： 已存在的
 *
 * 	govendor add pkg_name
 */
package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
)

func main() {
	/**
	 *	使用第三方包stringutil中Reverse函数反转字符串
	 * 	此包位置 ： vendor/github.com/golang/example/stringutil
	 */
	s := "ti esrever dna tipilf nwod gniht ym tup I"
	fmt.Println(stringutil.Reverse(s))
}
