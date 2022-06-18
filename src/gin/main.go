package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("程序入口...")
	str := "barfoothefoobarman"
	arr := []string{"foo", "ba"}
	//arr2 := []string{"foo", "bar"}
	fmt.Println("i::", strings.Index(str, "barfoo"))
	fmt.Println("s::", str)
	fmt.Println(findSubstring(str, arr))
}
func findSubstring(s string, words []string) []int {
	l := len(words)
	ll := make([]int, 0)
	list := make([]string, 0)

	for i := 0; i < l; i++ {
		list2 := words
		for j := i; j < l; j++ {
			if j != i {
				list2[i], list2[j] = list2[j], list2[i]
			}
			list = append(list, addlist(list2)...)
			fmt.Println("i::", i, "----all::", list)
		}
	}
	/*	list := make([]string, 0)
		var str string
		k := 0
		for i := 0; i < l; i++ {
			for j := 1; j < l; j++ {
				str = words[i] + words[j]
				list = append(list, str)
				if j == l-1 {
					k++
				}
			}
			words[0], words[k] = words[k], words[0]
		}*/
	fmt.Println("list::", list)
	for i := 0; i < len(list); i++ {
		j := strings.Index(s, list[i])
		fmt.Println("vvv2::", list[i])
		ll = append(ll, j)
	}
	return ll
}

func addlist(words []string) []string {
	list := make([]string, 0)
	var st string
	for _, v := range words {
		st += v
	}
	list = append(list, st)
	for i := len(words) - 1; i > 1; i-- {
		st = st[0:len(words[i])]
		list = append(list, st)
	}
	return list
}
