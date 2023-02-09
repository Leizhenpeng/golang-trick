package main

type someStruct struct{}

func NewPointer() interface{} {
	var c *someStruct = nil
	return c
}

func main() {
	p := NewPointer()
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

type iface struct {
    tab *itab
    data unsafe.Pointer
}


interface 变量的 nil 判断，汇编逻辑是判断首 8 字节是否是 0 值

interface 变量定义是一个 16 个字节的结构体，首 8 字节是类型字段，后 8 字节是数据指针。

interface 变量新创建的时候是 nil ，则这 16 个字节是全 0 值；

函数 NewPointer 中p是一个具体类型指针。所以，它一定会把接口变量 iface （返回的接口）前 8 字节设置非零字段的，因为有具体类型呀（无论具体类型是否是 nil 指针）。

具体类型到接口的赋值一定会导致接口非零（不考虑编译不过等问题）
*/

/*
In Go programming language, nil can refer to two things: a nil pointer or a nil interface value.

A nil pointer is a special value of a pointer data type that indicates the pointer doesn't point to a valid memory address. When a pointer is declared but not initialized, it's zero value is nil.

A nil interface value is a special value of an interface type that has no underlying concrete value. An interface value is nil if both its dynamic type and dynamic value are nil.

In summary, nil pointer represents the absence of an underlying memory address, while a nil interface value represents the absence of a concrete value of any type.


Yes, there is a difference in memory representation for a nil pointer and a nil interface value.

A nil pointer occupies a memory location that holds a special value indicating that the pointer doesn't point to a valid memory address. The size of the memory occupied by a nil pointer is dependent on the pointer data type and typically equals to the size of a memory address on the platform the code is running on.

A nil interface value, on the other hand, consists of two parts: a dynamic type and a dynamic value. The dynamic type part is a nil pointer to a type descriptor that describes the type of the concrete value that would be stored in the interface if it had a non-nil value. The dynamic value part is also a nil pointer, indicating the absence of a concrete value. So, a nil interface value takes up twice the size of a memory address.

In summary, the memory representation of a nil pointer is different from that of a nil interface value. The former takes up the size of a memory address, while the latter takes up twice the size of a memory address.Yes, there is a difference in memory representation for a nil pointer and a nil interface value.

nil 指针占用的内存大小取决于指针数据类型,通常等于代码运行的平台上的内存地址的大小。

另一方面，nil 接口值由两部分组成：动态类型和动态值。动态类型部分是指向类型描述符的 nil 指针，该类型描述符描述了接口如果具有非 nil 值时存储的具体值的类型。动态值部分也是一个 nil 指针，表示没有具体值。因此，nil 接口值占用了两倍的内存地址大小。

总的来说，nil 指针和 nil 接口值的内存表示是不同的。前者占用的内存大小为一个内存地址的大小，而后者占用的内存大小为两个内存地址的大小。
*/
