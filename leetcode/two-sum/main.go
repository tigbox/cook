package twosum

import (
	"sort"
)

// 暴力法：遍历当前的数组，每个元素又向后遍历到最后一个，时间复杂度是O(n^2)
func TwoSum(list []int, target int) []int {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i]+list[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 哈希表法：遍历当前的数组，当前元素值作为key写入到哈希表，再从哈希表中找出和sum的差值，如果存在就return。时间复杂度是n
func TwoSumHash(list []int, target int) []int {
	dictionary := make(map[int]int)
	for currentIndex, value := range list {
		dictionary[value] = currentIndex
		if anotherIndex, ok := dictionary[target-value]; ok {
			// 一共就两个index， currentIndex是后找到的，那么anotherIndex就是先找到的
			return []int{anotherIndex, currentIndex}
		}

	}
	return []int{}
}

// 哈希表法：遍历当前的数组，当前的元素值作为key写入到哈希表，再次哈希表中找出和sum的差值，如果存在就把两个index打入到temp中
// 考虑到temp中的元素就是索引，并且希望索引按照从小到大的顺序， 所以这里搞了个sort排序对temp进行整理
// 时间复杂度是n
func TwoSumHash1(list []int, target int) []int {
	dic := make(map[int]int)
	for currentIndex, value := range list {
		dic[value] = currentIndex
		if anotherIndex, ok := dic[target-value]; ok {
			temp := []int{currentIndex, anotherIndex}
			sort.Ints(temp)
			return temp
		}
	}
	return []int{}
}
