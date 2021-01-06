package main

import "fmt"

func test1() {
	b := &[][]int{}
	a := []int{1,2,3}
	fmt.Printf("%p", &a)
	*b = append(*b, a)
	a[0] = 2
	a[1] = 3
	a[2] = 4
	fmt.Printf("%p", &a)
	*b = append(*b, a)
	println("")
	fmt.Println(*b)
}

func test2() {
	a := [5]int{1,2,3,4,5}
	b := a[:4]
	c := a[1:]
	copy(b, c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func test3() {
	a := []int{1,2,3}
	b := [][]int{}
	b = append(b, a)
	fmt.Println(a)
	fmt.Println(b)
	a = append(a, 1)
	fmt.Println(a)
	fmt.Println(b)

}

func main() {
	test3()
}








