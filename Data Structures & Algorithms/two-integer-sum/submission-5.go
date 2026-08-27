func twoSum(nums []int, target int) []int {
	set := make(map[int]int, len(nums))
	for i, v := range nums {
		set[v] = i
	}
	for i, v := range nums {
		diff := target - v
		if j, ok := set[diff]; ok && j != i {
			return []int{i, j}
		}
	}
	return []int{}
}
