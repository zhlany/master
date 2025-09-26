package main

import (
	"fmt"
)

// 构建next数组
// 核心：寻找最长的相同前后缀
func buildNext(pattern string) []int {
	m := len(pattern)
	next := make([]int, m)
	next[0] = 0
	j := 0 // j指向前缀的末尾，同时表示当前公共前后缀长度
	i := 1 // i指向后缀的末尾
	for i < m {
		if pattern[i] == pattern[j] {
			j++         // 公共前后缀长度+1
			next[i] = j // 更新next[i]
			// 末尾向前移动一位，继续寻找
			i++
		} else {
			if j > 0 {
				j = next[j-1] // 回溯到前一个位置
			} else {
				next[i] = 0
				i++
			}
		}
	}
	return next
}

// KMP KMP算法: 匹配模式串pattern在文本串text中的位置
func KMP(text string, pattern string) int {
	n := len(text)
	m := len(pattern)
	if m == 0 {
		return 0
	}
	// 构建next数组
	next := buildNext(pattern)

	i, j := 0, 0
	for i < n {
		if text[i] == pattern[j] {
			i++
			j++
			if j == m {
				return i - j // 匹配成功，返回起始位置
			}
		} else if j > 0 {
			j = next[j-1] // 利用next数组跳过部分匹配
		} else {
			i++ // 第一个字符就失配，移动i
		}
	}
	return -1 // 未找到
}

func main() {
	fmt.Println("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("KMP算法")

	/*	pattern := "acacc"
		next := buildNext(pattern)*/
	ss := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ss = append(ss[:2], ss[3:]...)
	fmt.Println("next::", ss)
	{
		fmt.Println("KMP:", KMP("abcdabcd", "abcd"))
	}
}
