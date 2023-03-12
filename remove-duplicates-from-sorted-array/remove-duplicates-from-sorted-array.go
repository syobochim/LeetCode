package removeDuplicates

import "fmt"

func removeDuplicates(nums []int) int {
	var resultIndex int = 1
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			fmt.Printf("i: %d, j: %d, resultIndex: %d \n", i, j, resultIndex)
			if nums[i] != nums[j] {
				nums[resultIndex] = nums[j]
				resultIndex++
				i = j
			}
		}
	}
	return resultIndex
}
