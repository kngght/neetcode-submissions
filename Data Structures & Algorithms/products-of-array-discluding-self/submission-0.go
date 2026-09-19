func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	pref := 1
	for i := range res{ 
		res[i] = pref
		pref *= nums[i]
	}
	suf := 1
	for i := len(res)-1; i >= 0; i-- { 
		res[i] *= suf 
		suf *= nums[i]
	}
	return res
}
