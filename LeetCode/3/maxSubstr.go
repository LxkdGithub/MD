package main

import "time"

func main() {
	start := time.Now().UnixNano()
	println(start)
	println(findMaxSubStr("abcabcbb"))
	println(time.Now().UnixNano() - start)
}

//滑动窗口法
func findMaxSubStr(s string) int {
	if len(s) == 0 {
		return 0
	}

	// 记录每个字符出现频率
	var freq [50]int
	var result, left, right = 0, 0, -1 // 为什么要从-1开始因为代码从right+1开始
	for left < len(s) {
		if right+1<len(s) && freq[s[right+1] - 'a'] == 0 {
			freq[s[right+1] - 'a']++
			right++
		} else {
			if right+1>=len(s) {
				return result
			} else {
				// 这里有是我没想到到的 直接一次次右移而不是查找冲突位置
				freq[s[left] - 'a']-- //最左边的记录清楚
				left++					//右移
			}
		}
		if right - left + 1 > result {
			result = right - left + 1
		}
	}
	return result
}
