package main

import "fmt"

func main() {
	res := combine(4, 3)
	fmt.Println(res)
}


// 此题深度优先搜索 可以当做范例
func combine(n, k int) [][]int {
	// 这只是个入口函数
	if n <= 0 || k <= 0 {
		return [][]int{}
	}
	tmp, res := []int{}, [][]int{}
	generateCombniations(n, k, 1, tmp, &res)
	return res
}


// 这里思路很重要 不要死记，for循环主要判断开始和结束位置
// 先append递归完了之后在删掉继续循环

func generateCombniations(n, k, start int,c []int, res *[][]int) {
	// 上边是终止条件
	if len(c) == k {				//这里切片res append 切片c会有什么结果
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		//*res = append(*res, c)
		return
	}


	for i:=start;i<=n-(k-len(c))+1;i++ {
		c = append(c, i)
		generateCombniations(n, k, i+1, c, res)
		c = c[:len(c)-1]
	}


}
