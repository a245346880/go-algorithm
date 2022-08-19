// @program:     go-algorithm
// @file:        subsets.go
// @author:      XY
// @create:      2022-08-19 11:53
// @description: 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

package main

import "fmt"

//思路：这是一个典型的应用回溯法的题目，简单来说就是穷尽所有可能性
//通过不停的选择，撤销选择，来穷尽所有可能性，最后将满足条件的结果返回
func subsets(nums []int) [][]int {
	//保存最终结果
	result := make([][]int, 0) //初始化一个长度为0的二维数组
	//保存中间结果
	list := make([]int, 0)
	backtrack(nums, 0, list, &result)
	return result
}

//nums 给定的集合
//pos 下次添加到集合中的元素位置索引
//list 临时结果集合（每次需要复制保存）
//result 最终结果
func backtrack(nums []int, pos int, list []int, result *[][]int) {
	//把临时结果复制出来保存到最终结果
	ans := make([]int, len(list))
	//拷贝list到临时数组ans
	copy(ans, list)
	//将临时数组ans存入最终结果的二维数组中
	*result = append(*result, ans)
	//选择、处理结果、再撤销选择
	for i := pos; i < len(nums); i++ {
		//将从post开始，依次取nums中的数据存入list数组中
		list = append(list, nums[i])
		//递归，传入当前节点+1，再次遍历，相当于第一个大循环从第一个元素开始，挨个顺序组成数组存入result
		//第二个大循环时，采用nums[1]为起点进行顺序组合
		backtrack(nums, i+1, list, result)
		//list =
		list = list[0 : len(list)-1]
	}
}

func main() {
	//TODO 理解有点问题，后面再看
	nums := []int{1, 5, 7, 4, 9, 11}
	result := subsets(nums)
	fmt.Printf("%v\n", result)
	fmt.Printf("%d\n", len(result))
}
