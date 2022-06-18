package main

import (
	"fmt"
	"math/rand"
)

func randInt() int {
	return rand.Intn(100)
}
func main() {
	arr := make([]int, 0, 20)
	for i := 0; i < 20; i++ {
		arr = append(arr, randInt())
	}
	fmt.Println("arr::", arr)
	/*m := maopao(arr)
	fmt.Println("m::", m)*/
	k := kuaisu(arr, 0, len(arr)-1)
	fmt.Println("k::", k)

}

//冒泡排序
func maopao(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

//快速排序
func kuaisu(arr []int, start, end int) []int {
	if start >= end {
		return arr
	}
	//选择p:= arr[0]为基准,l为左标， r为右标
	p, l, r := arr[start], start+1, end
	//fmt.Println("p:", p, l, r)
	for l < r {
		for arr[l] <= p && l < r {
			l++
		}
		for arr[r] >= p && l < r {
			r--
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
			fmt.Println("::", l, arr[l], "::", r, arr[r], "\n", arr)
		}
		if l == r-1 {
			break
		}
	}
	arr[l], arr[start] = arr[start], arr[l]
	kuaisu(arr, start, l-1)
	kuaisu(arr, l+1, end)
	return arr
}
