func candy(ratings []int) int {
	candies := make([]int, len(ratings)) 
	for i := range candies {
		candies[i]++
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i-1] < ratings[i] {
			for candies[i-1] >= candies[i] {
				candies[i]++
			}
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i+1] < ratings[i] {
			for candies[i+1] >= candies[i] {
				candies[i]++
			}
		}
	}

	bag := 0
	for _, v := range candies {
		bag += v
	} 
	return bag
}