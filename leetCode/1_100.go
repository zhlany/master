package main

import "fmt"

func main() {
	println("running...")
	n1 := []int{9, 9, 9, 9, 9, 9, 9}
	n2 := []int{9, 9, 9, 9}
	ll1, ll2 := &ListNode{-1, nil}, &ListNode{-1, nil}
	l1, l2 := ll1, ll2
	j1, j2 := 7, 4
	for i := 0; i < j1; i++ {
		ni := &ListNode{n1[i], nil}
		l1.Next = ni
		l1 = l1.Next
	}
	for i := 0; i < j2; i++ {
		nj := &ListNode{n2[i], nil}
		l2.Next = nj
		l2 = l2.Next
	}
	re := addTwoNumbers(ll1.Next, ll2.Next)
	fmt.Printf("s:: %v \n", re.Val)
	for re.Next != nil {
		fmt.Printf("s:: %v \n", re.Val)
		re = re.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var len1, len2 int = 0, 0
	//WG:=sync.WaitGroup{}
	//WG.Add(1)
	for l1.Next != nil {
		len1++
		l1 = l1.Next
	}
	for l2.Next != nil {
		len2++
		l2 = l2.Next
	}
	//WG.Done()
	//WG.Wait()
	if len1 > len2 {
		for i := 0; i < len1-len2; i++ {
			l2.Next = &ListNode{0, nil}
			l2 = l2.Next
		}
	} else {
		for j := 0; j < len2-len1; j++ {
			l1.Next = &ListNode{0, nil}
			l1 = l1.Next
		}
	}
	result := &ListNode{Val: -1} //存放结果的链表
	conut, n := false, 0
	ladd := result
	for l1 != nil && l2 != nil {
		i := l1.Val + l2.Val
		ladd.Next = &ListNode{i%10 + n, nil}
		//记录下次是否进1
		if i >= 10 {
			n = 1
			conut = true
		} else {
			n = 0
			conut = false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if conut {
		ladd.Next = &ListNode{1, nil}
		ladd = ladd.Next
	}
	return result.Next
}
