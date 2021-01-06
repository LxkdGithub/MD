package bmmatch

// BF算法是 O(m*n)的暴力破解


// RKMatch returns int
func RKMatch(main string, ptn string) int {
	if main == "" || ptn == "" {
		return -1
	}
	var m, n int = len(main), len(ptn)
	if m < n {
		return RKMatch(ptn, main)
	}

	// Hash的方式
	var ptnHash = 0
	for i := 0; i < n; i++ {
		ptnHash += int(ptn[i])
	}

	var mainHash = 0
	for i := 0; i <= m-n; i++ {
		if i == 0 {
			for l := 0; l < n; l++ {
				mainHash += int(main[l])
			}
		} else {
			mainHash = mainHash - int(main[i-1]) + int(main[i+n-1])
		}

		if mainHash == ptnHash {
			var k = i
			var j = 0
			for j < n {
				if int(main[k]) != int(ptn[j]) {
					break
				}
				k++
				j++
			}
			if j == n {
				return i
			}
		}
	}
	return -1
}


//参考链接  https://segmentfault.com/a/1190000022490177 代码较难理解
func BMMacth() {

}



// 效率上 应该是BM > KMP > RK >= BF(暴力匹配)