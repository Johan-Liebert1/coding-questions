package main

func moveZeroes2(nums []int) {
	lastFoundNonZeroAt := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastFoundNonZeroAt] = nums[i]
			lastFoundNonZeroAt++
		}
	}

	for ; lastFoundNonZeroAt < len(nums); lastFoundNonZeroAt++ {
		nums[lastFoundNonZeroAt] = 0
	}
}

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}

	left, right := 0, 1

	for right < len(nums) {
		for left < right && nums[left] != 0 {
			left++
		}

		if nums[right] != 0 && nums[left] == 0 {
			// WE can definitely swap these as left and right will always be one next to each other
			// so we do maintain the relative ordering of non zeroes
			nums[right], nums[left] = nums[left], nums[right]
		}

		right++
	}
}

func main3() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}
