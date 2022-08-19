// @program:     go-algorithm
// @file:        strStr.go
// @author:      XY
// @create:      2022-08-19 11:25
// @description: 给定一个  haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从 0 开始)。如果不存在，则返回  -1。

package main

import "fmt"

//思路：核心点遍历给定字符串字符，判断以当前字符开头字符串是否等于目标字符串
//技巧：判断 haystack的某个字符是否等于needle第一个字符
//如果不等，则进入跳出内部循环
//如果相等，则继续内部循环判断
//判断完成之后，校验一下是否判断完，用 len(needle) == j判断，如果内部循环跑完，则相等
//为true，表示存在
func strStr(haystack, needle string) int {
	//如果需要判断的字符串长度为0
	if len(needle) == 0 {
		return 0
	}
	var i, j int
	for i = 0; i < len(haystack)-len(needle)+1; i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if len(needle) == j {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("abcde", "cd"))
}
