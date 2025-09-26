package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
 *@project: master
 *@author: Administrator
 *@created: 2025/4/3 08:54
 *@updated: 2025/4/3 08:54
 *@description:
	1.桶排序：分布式排序算法，分桶排序好后合并
	2.插入排序：选择当前元素，从前往后插入
	3.快速排序：分治法，数组分两部分，[:pivot]<=pivot,pivot, (pivot:n]>pivot,递归重复至空
	4.归并排序：将待排序数组分为左右两个部分，左边的元素都比右边的元素小，然后递归调用，重复上面的步骤，直到数组为空，有序
*/

// Node represents a node in the linked list
type Node struct {
	Data int
	Next *Node
}

// LinkedList represents the linked list
type LinkedList struct {
	Head *Node
}

func randInt() int {
	return rand.Intn(100)
}
func main() {
	n := 3
	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		arr = append(arr, randInt())
	}
	//fmt.Println("arr::", arr)
	/*m := maopao(arr)
	fmt.Println("m::", m)*/
	//k := quickSort2(arr)
	//fmt.Println("k::", k)

	buildNext("afadd")
}

// buildNext 函数计算KMP算法中的next数组
// arr: 需要构建next数组的字符串
// 返回值: arr对应的next数组
func buildNext(arr string) []int {
	fmt.Println("arr::", arr)
	next := make([]int, len(arr))
	next[0] = -1

	j, k := 0, -1 //  i是text索引，k是pattern索引
	for j < len(arr)-1 {
		if k == -1 || arr[j] == arr[k] { // 匹配成功
			j++
			k++
			next[j] = k
			fmt.Println("next::", next)
		} else {
			//回退K值
			k = next[k]
		}
		fmt.Println("j:: ", j, "   k::", k)
	}
	return next
}

// KMPSearch 使用KMP算法在text中查找pattern的第一次出现位置
// text: 主字符串
// pattern: 需要查找的模式字符串
// 返回值: pattern在text中第一次出现的位置，如果未找到则返回-1
func KMPSearch(text, pattern string) int {
	if len(pattern) == 0 {
		return 0
	}
	if len(text) < len(pattern) {
		return -1
	}

	next := buildNext(pattern)
	i, j := 0, 0 // i是text索引，j是pattern索引

	for i < len(text) && j < len(pattern) {
		if j == -1 || text[i] == pattern[j] {
			i++
			j++
		} else {
			//模式串回退到next[j]
			j = next[j]
		}
	}

	if j == len(pattern) {
		return i - j // 返回匹配的起始位置
	}
	return -1 // 未找到
}

// 冒泡排序
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

// 快速排序
// 使用的分治法，选择一个基准，将数组分为两部分，比基准小的放左边，大的放右边，然后递归调用，重复上面的步骤
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)/2]
	var left, right []int
	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		}
	}
	return append(quickSort(left), append([]int{pivot}, quickSort(right)...)...)
}

// 插入排序
func insertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

// 归并排序
// 分治法，将待排序数组分为左右两个部分，左边的元素都比右边的元素小，然后递归调用，重复上面的步骤，直到数组为空，有序
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	var arr []int
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr = append(arr, left[i])
			i++
		} else {
			arr = append(arr, right[j])
			j++
		}
	}
	arr = append(arr, left[i:]...)
	arr = append(arr, right[j:]...)
	return arr
}

// 随机快速排序
// 1.使用随机数随机选择一个基准值，奖后将数组分割为两部分，左边的子数组比后边的子数组大，
// 2.再递归调用快速排序函数，直到数组为空或者只有一个函数，则返回数组
func quickSort2(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	rag := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 随机选择基准值并将其放到第一个位置
	pivot := rag.Intn(len(arr))
	//把基准放在第一个位置
	arr[0], arr[pivot] = arr[pivot], arr[0]

	//从第二个元素开始
	left := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[0] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}
	// 将基准值放回正确位置
	arr[0], arr[left-1] = arr[left-1], arr[0]
	quickSort2(arr[:left-1])
	quickSort2(arr[left:])
	return arr
}
