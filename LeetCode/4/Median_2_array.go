package main

func main() {
	nums1 := []int{2,3,4,5,6,7}
	nums2 := []int{2,4,6,7,9,10}
	println(find_median(nums1, nums2))
	println(findMedianSortedArrays(nums1, nums2))
}

func find_median(nums1, nums2 []int) float64 {
	// ensure the small is on the first
	if len(nums1) > len(nums2) {
		return find_median(nums2, nums1)
	}

	low, high, nums1Mid, nums2Mid, k := 0, len(nums1), 0, 0, (len(nums1) + len(nums2)+1) >> 1
	for low <= high {
		nums1Mid = low + (high-low)>>1 // 分界限右侧是 mid，分界线左侧是 mid - 1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] { // nums1 中的分界线划多了，要向左边移动
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] { // nums1 中的分界线划少了，要向右边移动
			low = nums1Mid + 1
		} else {
			// 找到合适的划分了，需要输出最终结果了
			// 分为奇数偶数 2 种情况
			break
		}
	}
	// 找到左边的最大值和右边的最大值      记住基准是  numsMid-1 | numsMid
	maxLeft, minRight := 0, 0
	if nums1Mid == 0 {
		maxLeft = nums2[nums2Mid - 1]
	} else if nums2Mid == 0 {
		maxLeft = nums1[nums1Mid - 1]
	} else {
		maxLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}

	if (len(nums1) + len(nums2)) & 1 == 1 {
		return float64(maxLeft)
	}

	if nums1Mid == len(nums1) {
		minRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		minRight = nums1[nums1Mid]
	} else {
		minRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}

	return float64(minRight + maxLeft)/2   //有一个float即可使结果为float 但是float和int不能相加
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 假设 nums1 的长度小
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
		// nums1:  ……………… nums1[nums1Mid-1] | nums1[nums1Mid] ……………………
		// nums2:  ……………… nums2[nums2Mid-1] | nums2[nums2Mid] ……………………
		nums1Mid = low + (high-low)>>1 // 分界限右侧是 mid，分界线左侧是 mid - 1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] { // nums1 中的分界线划多了，要向左边移动
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] { // nums1 中的分界线划少了，要向右边移动
			low = nums1Mid + 1
		} else {
			// 找到合适的划分了，需要输出最终结果了
			// 分为奇数偶数 2 种情况
			break
		}
	}
	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft+midRight) / 2
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}