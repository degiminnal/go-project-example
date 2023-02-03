package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(strings.IndexByte("hello", 'H'))

	n := 10
	res := make([]int, n)
	fmt.Println(res[0])

	fmt.Println("创建一维数组：")
	nums := []int{0, 1, 0, 1}
	fmt.Println(nums)

	fmt.Println("创二维数组：")
	var arr [][]int
	for x := 0; x < 3; x++ {
		tmp := make([]int, 10)
		arr = append(arr, tmp)
	}
	fmt.Println(arr)
	arr1 := make([][10]int, 3)
	fmt.Println(arr1)
	arr1 = append(arr1, [10]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	fmt.Println(arr1)
	arr1 = append(arr1, arr1[1])
	fmt.Println(arr1)

	fmt.Println("用const数值初始化array:")
	const cst_a = 10
	const cst_b int = 10
	arr3 := make([][cst_a]int, 2)
	fmt.Println(cst_a, cst_b)
	fmt.Println(arr3)
	// 不能用变量初始化常量
	// aa := 12
	// const cst_aa int = aa

	fmt.Println("min, max函数：")
	// 只能两个元素min
	fmt.Println(math.Min(1, 2))

	// 数组的遍历
	fmt.Println("数组遍历：")
	groups := []int{1, 2, 3, 4, 5}
	for i, g := range groups {
		fmt.Println(i, g)
	}

	// 求余
	fmt.Println("求余\n3%2=", 3%2)

	s := "123"
	fmt.Println("字符‘1’：", s[0])
	fmt.Println("字符‘1’==49:", s[0] == 49)
	fmt.Println("字符‘1’=='1':", s[0] == '1')

}
