package main

import (
	"fmt"
	"sort"
)

func main() {
	/*	candidates := []int{2, 3, 6, 7}
		target := 7
		fmt.Println(combinationSum(candidates, target))*/
	fmt.Println(jump([]int{2, 0, 2, 1, 4}))
}

func jump(nums []int) (ans int) {
	maxIndex, judgeIndex := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxIndex = max(maxIndex, i+nums[i])      // 更新能跳到的极限位置
		if i == judgeIndex || i == len(nums)-1 { // 到达judge临界点，跳跃次数加一，临界点设置成当前极限距离
			ans++ // 延迟统计跳跃次数，即在每次临界点时前面必须跳一次了
			judgeIndex = maxIndex
		}
	}
	return
}

func combinationSum2(candidates []int, target int) [][]int {
	// 检查输入是否合法
	if candidates == nil || len(candidates) == 0 || target < 0 {
		return [][]int{}
	}
	var res [][]int
	sort.Ints(candidates)

	//回溯算法
	var backtrack func(int, int, []int)
	backtrack = func(sum, index int, path []int) {
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			if NotInArray(temp, res) {
				res = append(res, temp)
			}
			return
		} else if sum > target {
			return
		}

		for i := index; i < len(candidates); i++ {
			path = append(path, candidates[i])
			// 递归，注意可以继续使用同一元素
			backtrack(sum+candidates[i], i, path)
			//重点： 恢复现场，撤销操作
			path = path[:len(path)-1]
		}

	}
	backtrack(0, 0, []int{})

	return res
}

func NotInArray(arr1 []int, arr2 [][]int) bool {
	if arr2 == nil {
		return true
	}
	for _, v := range arr2 {
		if len(v) != len(arr1) {
			continue
		}
		Map := make(map[int]int)
		for _, v1 := range arr1 {
			Map[v1]++
		}
		for _, v2 := range v {
			if Map[v2] == 0 {
				return false
			}
		}
	}
	return true
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) // 排序便于剪枝
	var result [][]int
	var path []int

	// 定义内置的 backtrack 函数（闭包）
	var backtrack func(start int, target int)

	backtrack = func(start int, target int) {
		if target == 0 {
			// 找到一个有效组合，复制到结果中
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		} else if target < 0 {
			// 剪枝：目标和已小于 0，结束递归
			return
		}

		for i := start; i < len(candidates); i++ {
			// 跳过重复的候选值
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			if candidates[i] > target {
				// 剪枝：因为已排序，后续数字更大，直接跳过
				break
			}
			// 选择当前数字
			path = append(path, candidates[i])
			// 递归：可以重复选择当前数字，所以从 i 开始
			backtrack(i, target-candidates[i])
			// 回溯：撤销选择
			path = path[:len(path)-1]
		}
	}

	backtrack(0, target) // 从第 0 个数字开始，目标和为 target
	return result
}
