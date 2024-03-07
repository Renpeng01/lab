package go_link_name_interface

import (
	_ "lab/src/go_link_name_src"
)

// go:linkname可以跨包使用，需要导入 unsafe包。
// A所在的包必须要打入B所在的包，这个也好理解，相当于A调用了B。
// go build无法编译go:linkname,必须用单独的compile命令进行编译，因为go build会加上-complete参数，这个参数会检查到没有方法体的方法，并且不通过，报错missing function body。或者，在这个没有方法体的包下添加一个空的aa.s 的汇编文件标示，就可以编译了
// 原文链接：https://blog.csdn.net/raoxiaoya/article/details/122079366
func GolinknameinterfaceTest()
