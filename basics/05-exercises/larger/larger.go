package larger

func Check(x int, y int) int {
	output := 0
	larger := 0
	smaller := 0

	if x <= y {
		larger = y
		smaller = x
	} else {
		larger = x
		smaller = y
	}

	output = larger % smaller
	return output
}
