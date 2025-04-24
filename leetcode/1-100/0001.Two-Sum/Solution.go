package Solution

import "sort"

func TwoSum(list []int, target int) []int {
	for i, v := range list {
		for j := i + 1; j < len(list); j++ {
			if v+list[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func TwoSumMap(list []int, target int) []int {
	m := make(map[int]int)
	for index, value := range list {
		if anotherIndex, ok := m[target-value]; ok {
			return []int{anotherIndex, index}
		} else {
			m[value] = index
		}
	}
	return []int{}
}

func TwoSumMap1(list []int, target int) []int {
	m := make(map[int]int)
	for index, value := range list {
		if anotherIndex, ok := m[target-value]; ok {
			temp := []int{index, anotherIndex}
			sort.Ints(temp)
			return temp
		} else {
			m[value] = index

		}
	}
	return []int{}
}
