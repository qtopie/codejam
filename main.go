package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// document https://go.dev/blog/strings

// why rune use int32 not uint32 https://stackoverflow.com/questions/24714665/why-is-rune-in-golang-an-alias-for-int32-and-not-uint32
func main() {
	var v int64
	v = 0
	fmt.Println(v)

	abc := byte(1)
	var fed int
	fed = int(abc)
	fmt.Println(byte('1') == 1, fed)

	fmt.Println(strconv.Itoa(-1))
	fmt.Println(3 ^ 1)
	src := []int{1, 2, 3, 4, 5}
	fmt.Println(src[5:])
	src = src[0:]
	fmt.Println(len(src[0:]))
	testSlice(src)

	copy(src[0:], src[2:])
	fmt.Println(src) // [3 4 5 4 5]

	//	fmt.Println(rune('a') == byte('a'))
	fmt.Println(1 << 31)
	fmt.Println('a', rune('a'), byte('a')) // 97, same int value
	fmt.Println('a' == rune("abc"[0]))     // print true

	fmt.Println(string("abc你好"[1])) // print b , byte

	fmt.Println("abc"[0] == 'a') // true

	fmt.Println(reflect.TypeOf('a'), reflect.TypeOf(byte('a'))) // rune(int32), byte(uint8)

	fmt.Println(string([]rune("abc你好")[4])) // print 好

	for _, v := range "abc你好" {
		fmt.Print(string(v), "\t") // rune
	}
	fmt.Println()

	var a rune
	var b byte
	a = 'a'
	b = 'b'
	fmt.Println(a, b)
	// a = b // compile error

}

func testSlice(nums []int) {
	fmt.Println(len(nums))
	fmt.Println(nums)
}
