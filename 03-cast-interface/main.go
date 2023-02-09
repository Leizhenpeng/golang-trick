package main

type someStruct struct{}

func NewPointer() *someStruct {
	var c *someStruct = nil
	return c
}

func main() {
	p := NewPointer()
	// 首 8 字节是否是 0 值
	if p == nil {
		println("p is nil")
	} else {
		println("p is not nil")
	}
}

/*
# which output?
# A. p is nil
# B. p is not nil
# C. panic
*/

/*

interface 结构体定义

type iface struct {
    tab *itab
    data unsafe.Pointer
}

interface 一个 16 个字节的结构体，首 8 字节是类型字段，后 8 字节是数据指针。

interface 变量新创建的时候是 nil ，则这 16 个字节是全 0 值；

interface 变量的 nil 判断，汇编逻辑是判断首 8 字节是否是 0 值

函数 NewPointer 中p是一个具体类型指针。所以，它一定会把接口变量 iface （返回的接口）前 8 字节设置非零字段的，因为有具体类型呀（无论具体类型是否是 nil 指针）。

具体类型到接口的赋值一定会导致接口非零
*/

/*
In Go programming language, nil can refer to two things: a nil pointer or a nil interface value.

A nil pointer is a special value of a pointer data type that indicates the pointer doesn't point to a valid memory address. When a pointer is declared but not initialized, it's zero value is nil.

nil 指针占用的内存大小取决于指针数据类型,通常等于代码运行的平台上的内存地址的大小。

nil 接口值由两部分组成：动态类型和动态值。动态类型部分是指向类型描述符的 nil 指针，该类型描述符描述了接口如果具有非 nil 值时存储的具体值的类型。动态值部分也是一个 nil 指针，表示没有具体值。因此，nil 接口值占用了两倍的内存地址大小。

nil 指针和 nil 接口值的内存表示是不同的。前者占用的内存大小为一个内存地址的大小，而后者占用的内存大小为两个内存地址的大小。
*/
