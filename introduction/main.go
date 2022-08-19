// @program:     go-algorithm
// @file:        main.go
// @author:      XY
// @create:      2022-08-19 14:30
// @description: 基础数据结构

package main

import (
	"fmt"
	"sort"
)

//通过切片模拟栈
func stackStruct() {
	stack := make([]int, 0)
	//压栈
	stack = append(stack, 10)
	fmt.Printf("%v\n", stack)
	//弹栈
	v := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Printf("%v\n", v)
	fmt.Printf("%v\n", stack)
}

//模拟队列
func queueStruct() {
	queue := make([]int, 0)
	//入队
	queue = append(queue, 10)
	queue = append(queue, 20)
	fmt.Printf("入队：%v\n", queue)
	//出队
	v := queue[0]
	queue = queue[1:]
	fmt.Printf("出队：%v\n", v)
	fmt.Printf("出队后：%v\n", queue)

}

//字典
func mapStruct() {
	m := make(map[string]int)
	m["张三"] = 1
	m["李四"] = 2
	m["王五"] = 3
	//删除前
	fmt.Printf("删除前:%v\n", m)
	//删除key
	delete(m, "张三")
	fmt.Printf("删除后:%v\n", m)
}

//排序
func sortFunc() {
	intArr := []int{1, 3, 2, 5}
	sort.Ints(intArr)
	fmt.Printf("[]int排序后:%v\n", intArr)
	strArr := []string{"c", "b", "a", "e"}
	sort.Strings(strArr)
	fmt.Printf("[]string排序后:%v\n", strArr)
	//自定义排序
	objArr := []int{1, 3, 2, 5}
	sort.Slice(objArr, func(i, j int) bool {
		return objArr[i] > objArr[j]
	})
	fmt.Printf("Slice排序后:%v\n", objArr)

}
func main() {
	//stackStruct()
	//queueStruct()
	//mapStruct()
	sortFunc()
}
