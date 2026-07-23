// func findDuplicate(nums []int) int {
// 	set := make(map[int]struct{}, len(nums)/2)
// 	for _, v := range nums {
// 		if _, ok := set[v]; ok {
// 			return v
// 		}
// 		set[v] = struct{}{}
// 	}
// 	return 0
// }

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	target := 0
	for slow != target {
		target = nums[target]
		slow = nums[slow]
	}
	return target
}