package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 流程控制
// 只出现一次的数字
func singleNumPicker(nums []int) int {
	countMap := make(map[int]int)
	for _, k := range nums {
		_, ok := countMap[k]
		if ok {
			countMap[k] += 1
		} else {
			countMap[k] = 1
		}
	}

	for k, v := range countMap {
		if v == 1 {
			return k
		}
	}
	return 0
}

// 回文数
func isPalidrome(x int) bool {
	if x == 0 {
		return true
	}

	if x < 0 || x%10 == 0 {
		return false
	}

	str := strconv.Itoa(x)
	length := len(str)

	if length == 1 {
		return true
	}

	i := 0
	j := length - 1

	for j > i {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}

	return true
}

// 字符串
// 有效的括号		字符串处理、栈的使用
func isValid(s string) bool {
	stack := make([]rune, 0, len(s))
	for _, v := range s {
		switch v {
		case '(', '[', '{':
			stack = append(stack, v)
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pop == '(' && v != ')' ||
				pop == '[' && v != ']' ||
				pop == '{' && v != '}' {
				return false
			}
		default:
			return false
		}
	}
	return len(stack) == 0
}

// 最长公共前缀		字符串处理、循环嵌套
func longestCommonPrefix(strs []string) string {
	var pre []byte = make([]byte, 0, len(strs[0]))

	for i := 0; i >= 0; i++ {
		for j, s := range strs {

			if i >= len(s) {
				if j != 0 {
					pre = pre[:len(pre)-1]
				}
				i = -2
				break
			}

			if j == 0 {
				pre = append(pre, s[i])
				continue
			}

			if pre[len(pre)-1] != s[i] {
				pre = pre[:len(pre)-1]
				i = -2
				break
			}
		}
	}

	return string(pre)
}

// 基本值类型
// 加一 				数组操作、进位处理
func plusOne(digits []int) []int {

	len := len(digits)

	forward := true

	if len > 0 {
		for i := len - 1; forward; i-- {
			if i < 0 {
				newDigits := make([]int, 0, len+1)
				newDigits = append(newDigits, 1)
				newDigits = append(newDigits, digits...)
				digits = newDigits
				forward = false

			} else if digits[i] == 9 {
				digits[i] = 0
				forward = true
			} else {
				digits[i] += 1
				forward = false
			}
		}
	}
	return digits
}

// 引用类型：切片
// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 慢指针，表示当前唯一元素的最后位置
	slow := 0

	// 快指针遍历整个数组
	for fast := 1; fast < len(nums); fast++ {
		// 当发现新元素时
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast] // 将新元素放到正确位置
		}
	}

	// 返回唯一元素的个数（slow+1）
	return slow + 1
}

// 合并区间
func merge(intervals [][]int) [][]int {
	newIts := make([][]int, 0)
	if len(intervals) == 1 {
		newIts = append(newIts, intervals[0])
	}
	if len(intervals) > 1 {

		sort.Slice(intervals, func(i, j int) bool {
			return intervals[i][0] < intervals[j][0]
		})

		newIts = append(newIts, intervals[0])

		for i := 1; i < len(intervals); i++ {
			last := len(newIts) - 1

			if newIts[last][1] < intervals[i][0] {
				newIts = append(newIts, intervals[i]) //add
			} else {
				if newIts[last][0] >= intervals[i][0] {
					if newIts[last][1] <= intervals[i][1] {
						newIts = newIts[:last]                //delete
						newIts = append(newIts, intervals[i]) //add
					} else {
						pair := []int{intervals[i][0], newIts[last][1]}
						newIts = newIts[:last]        //delete
						newIts = append(newIts, pair) //add
					}
				} else {
					if newIts[last][1] >= intervals[i][1] {
						continue
					} else {
						pair := []int{newIts[last][0], intervals[i][1]}
						newIts = newIts[:last]        //delete
						newIts = append(newIts, pair) //add
					}
				}
			}
		}
	}
	return newIts

}

// 两数之和
func twoSum(nums []int, target int) []int {

	if len(nums) == 2 {
		return []int{0, 1}
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 1}
}

func main() {

	fmt.Println(merge([][]int{{1, 4}, {0, 4}}))
}
