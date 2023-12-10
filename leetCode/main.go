package main

func main() {

}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var tmpList []int
	for {
		n1, n2 := 0, 0
		if l1.Next != nil {
			n1 = l1.Val
		} else {
			n1 = 0
		}
		if l2.Next != nil {
			n2 = l2.Val
		} else {
			n2 = 0
		}
		tmpList = append(tmpList, n1+n2)

		if l1.Next == nil && l2.Next == nil {
			break
		}
	}
	tmp := false
	retNone := ListNode{Val: -1}
	for _, v := range tmpList {
		tmp, v = check(tmp, v)
		newNone := ListNode{v, nil}
		retNone.Next = &newNone
	}
	return &retNone
}

func check(tmp bool, a int) (bool, int) {
	if tmp {
		a += 1
	}
	if a >= 10 {
		return true, a - 10
	}
	return false, a
}
