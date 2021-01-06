package main

func main() {
	a := "ADOBECODEBANC"
	b := "ABC"
	println(minWindow(a, b))
}



func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	var tFreq, sFreq [256]int
	result, left, right, finalLeft, finalRight, minW, count := "", 0, -1, -1, -1, len(s)+1, 0


	// tFreq因为t定而是定的一开始计算好
	// sFreq是left和right之间的部分频率，用来和tFerq对比来知道有没有完全包含，所有字符的sFerq[x] == tFreq[x]就是完全
	// 包含了，这使就要执行大的else来尝试右移left看看，每次移动都要看sFreq[x]是否和tFreq相等(因为右移right意味着
	// sFreq[x] >= tFreq[x]),
	for i := 0; i < len(t); i++ {
		tFreq[t[i]-'a']++
	}

	for left < len(s) {
		if right+1 < len(s) && count < len(t) {
			sFreq[s[right+1]-'a']++
			if sFreq[s[right+1]-'a'] <= tFreq[s[right+1]-'a'] {
				count++
			}
			right++
		} else {
			// 完全包含后尝试右移left，但一旦sFreq[x] < tFreq[x] 也就是count不足就要停止，继续右移right(上边代码)
			if right-left+1 < minW && count == len(t) {
				minW = right - left + 1
				finalLeft = left
				finalRight = right
			}
			if sFreq[s[left]-'a'] == tFreq[s[left]-'a'] {  //大的话不用减，继续执行大else(还可以后移)，小的话执行上边的大if
				count--
			}
			sFreq[s[left]-'a']--
			left++
		}
	}
	if finalLeft != -1 {
		for i := finalLeft; i < finalRight+1; i++ {
			result += string(s[i])
		}
	}
	return result
}