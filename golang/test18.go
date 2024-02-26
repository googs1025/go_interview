package main

import (
	"container/heap"
	"fmt"
)

// 冒泡排序
func bubbleSort(nums []int) []int {

	for i := 0; i < len(nums); i++ {

		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}

	}
	return nums
}

// 选择排序
func selectSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		min := i

		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
			nums[i], nums[min] = nums[min], nums[i]
		}
	}
	return nums
}

// MinHeap 定义一个自定义的最小堆类型
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 使用堆排序对切片进行排序
func heapSort(nums []int) []int {
	h := &MinHeap{}

	// 初始化堆
	heap.Init(h)

	// 将切片元素插入堆中
	for _, num := range nums {
		heap.Push(h, num)
	}

	// 从堆中弹出元素并存入有序切片
	sortedNums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sortedNums[i] = heap.Pop(h).(int)
	}

	return sortedNums
}

func main() {
	nums := []int{10, 4, 2, 7, 5, 6, 0}
	fmt.Println(selectSort(nums))
	sortedNums := heapSort(nums)
	fmt.Println(sortedNums) // 输出: [0 2 4 5 6 7 10]
}
