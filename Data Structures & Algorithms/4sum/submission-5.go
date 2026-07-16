func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0, 32)

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
        if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if nums[i]+nums[len(nums)-3]+nums[len(nums)-2]+nums[len(nums)-1] < target {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			buf := twoSum(nums[j+1:], target-nums[i]-nums[j])
			for _, v := range buf {
				res = append(res, []int{nums[i], nums[j], v[0], v[1]})
			}
		}
	}
	return res
}

func twoSum(nums []int, target int) [][]int {
	left, right := 0, len(nums)-1
	buf := make([][]int, 0, 16)
	for left < right {
		sum := nums[left] + nums[right]
		switch {
		case sum == target:
			buf = append(buf, []int{nums[left], nums[right]})
			left++
			right--
			for left < right && nums[left] == nums[left-1] {
				left++
			}
			for left < right && nums[right] == nums[right+1] {
				right--
			}
		case sum < target:
			left++
		default:
			right--
		}
	}
	return buf
}