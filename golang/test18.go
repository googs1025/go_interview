package main

import "fmt"

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

func main() {
	nums := []int{10, 4, 2, 7, 5, 6, 0}
	fmt.Println(selectSort(nums))
}
