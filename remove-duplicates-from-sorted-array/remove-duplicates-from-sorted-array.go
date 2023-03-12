package removeDuplicates

func removeDuplicates(nums []int) int {
    var resultIndex int = 1
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] != nums[i + 1] {
            nums[resultIndex] = nums[i + 1]
            resultIndex++
        }
    }
    return resultIndex
}
