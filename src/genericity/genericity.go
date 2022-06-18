package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

//go1.18泛型使用

// C ** T 及为类型//
//创建任意类型管道
/*type C[T any] chan T

var C2 chan string

// M map类型
type M[k int, v any] map[k]v

var M2 map[int]string

var i = 2

// S 切片
type S[T any] []T

var S2 = make([]int, i, i+2)
*/
//**//
////////////////

//My64BitsLongNum 自定义约束
type My64BitsLongNum interface {
	int32
}

func MyCompare[T My64BitsLongNum](a, b T) bool {
	return a < b
}

// MyMap 自定义类型 comparable:可比较的
type MyMap[K string, V constraints.Integer | constraints.Float] map[K]V

func Sli[T any](s []T) {
	for _, t := range s {
		fmt.Printf("%v \n", t)
	}
}

func main() {
	Sli[int]([]int{33, 11, 55, 44, 11})
	b := MyCompare(2, 4)
	fmt.Println("MyCompare bool: ", b)
	m := make(MyMap[string, int])
	m["11"] = 11
	m["dd"] = 22
	for k, v := range m {
		fmt.Printf("k:%v, v:%v\n", k, v)
	}

}
