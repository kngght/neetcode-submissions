func carFleet(target int, position []int, speed []int) int {
	type car struct{pos, spd int}
	cars := make([]car, len(position))	
	for i := 0; i < len(position); i++ {
		cars[i] = car{pos: position[i], spd: speed[i]} 
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	fleets := 0
	time := 0.0

	for _, c := range cars {
		t := float64(target - c.pos) / float64(c.spd)
		if t > time {
			fleets++
			time = t
		}
	}

	return fleets
}
