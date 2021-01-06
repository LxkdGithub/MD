package main

func main() {
	l := []int{1,8,6,2,5,4,8,3,7}
	println(maxArea(l))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}



func maxArea(height []int) int {
	left, right := 0, len(height) - 1
	result := 0
	for left < right {
		result = max(result, min(height[left], height[right]) * (right - left))
		if left >= right {
			right -= 1
		} else {
			left += 1
		}
	}
	return result
}
