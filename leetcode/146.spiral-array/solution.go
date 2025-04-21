package main

import "fmt"

// {1, 2, 3, 4}
// {5, 6, 7, 8}
// {9, 10, 11, 12}
// {13, 14, 15, 16}

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	result := SpiralArray2(matrix)
	fmt.Println(result)
}

// version0 感觉这个废话很多， 尤其是游标的控制
func SpiralArray(matrix [][]int) []int {
	var result []int
	if len(matrix) == 0 {
		return result
	}
	count := len(matrix) * len(matrix[0])
	minHigh, minWidth := 0, 0
	maxHigh, maxWidth := len(matrix), len(matrix[0])

	var i, j int

	for len(result) < count {
		for ; len(result) < count && j < maxWidth; j++ {
			result = append(result, matrix[i][j])
		}
		minHigh++
		j = maxWidth - 1
		i++

		for ; len(result) < count && i < maxHigh; i++ {
			result = append(result, matrix[i][j])
		}
		maxWidth--
		i = maxHigh - 1
		j--

		for ; len(result) < count && j >= minWidth; j-- {
			result = append(result, matrix[i][j])
		}
		maxHigh--
		j = minWidth
		i--

		for ; len(result) < count && i >= minHigh; i-- {
			result = append(result, matrix[i][j])
		}
		minWidth++
		i++
		j++
	}

	return result
}

// version3 这个思路不错， 就判断数量
func SpiralArray2(matrix [][]int) []int {
	var result []int
	if len(matrix) == 0 {
		return result
	}
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	count := len(matrix) * len(matrix[0])
	for len(result) < count {
		for i := left; i <= right && len(result) < count; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		for i := top; i <= bottom && len(result) < count; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		for i := right; i >= left && len(result) < count; i-- {
			result = append(result, matrix[bottom][i])
		}
		bottom--

		for i := bottom; i >= top && len(result) < count; i-- {
			result = append(result, matrix[i][left])
		}
		left++
	}
	return result
}

// version2
func SpiralArray1(matrix [][]int) []int {
	var result []int
	if len(matrix) == 0 {
		return result
	}
	count := len(matrix) * len(matrix[0])
	minHigh, minWidth := 0, 0
	maxHigh, maxWidth := len(matrix), len(matrix[0])

	var i, j int

	for len(result) < count {
		for ; len(result) < count && j < maxWidth; j++ {
			result = append(result, matrix[minHigh][j])
		}
		minHigh++
		i = minHigh

		for ; len(result) < count && i < maxHigh; i++ {
			result = append(result, matrix[i][maxWidth-1])
		}
		maxWidth--
		j = maxWidth - 1

		for ; len(result) < count && j >= minWidth; j-- {
			result = append(result, matrix[maxHigh-1][j])
		}
		maxHigh--
		i = maxHigh - 1

		for ; len(result) < count && i >= minHigh; i-- {
			result = append(result, matrix[i][minWidth])
		}
		minWidth++
		j = minWidth
	}

	return result
}
