func calPoints(operations []string) int {
	set := make([]int, 0, 512)
	sum := 0
	for _, v := range operations {
		switch {
		case v == "D":
			set = append(set, set[len(set)-1]*2)
		case v == "+":
			set = append(set, set[len(set)-1]+set[len(set)-2])
		case v == "C":
			set = set[:len(set)-1]
		default:
			vInt, _ := strconv.Atoi(v)
			set = append(set, vInt)
		}
	}

	for _, v := range set {
		sum += v
	}

	return sum
}