package solution

func FindLeft(arr []int, find int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)>>1
		if arr[mid] == find {
			if mid == 0 || arr[mid-1] != find {
				return mid
			} else {
				right = mid - 1
			}
		} else if arr[mid] < find {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func FindRight(arr []int, find int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)>>1
		if arr[mid] == find {
			if mid == len(arr)-1 || arr[mid+1] != find {
				return mid
			} else {
				left = mid + 1
			}
		} else if arr[mid] < find {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
