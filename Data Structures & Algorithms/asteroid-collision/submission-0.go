func asteroidCollision(asteroids []int) []int {
    stack := []int{}

    for _, v := range asteroids {
        alive := true

        for alive && v < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
            top := stack[len(stack)-1]
            if top < -v {
                stack = stack[:len(stack)-1]
            } else if top == -v {
                stack = stack[:len(stack)-1]
                alive = false
            } else {
                alive = false
            }
        }

        if alive {
            stack = append(stack, v)
        }
    }

    return stack
}