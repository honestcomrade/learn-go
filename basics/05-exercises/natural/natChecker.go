package natural

func NatChecker(x int) bool {
	found := false
	if x%3 == 0 {
		found = true
	} else if x%5 == 0 {
		found = true
	} else {
		found = false
	}
	return found
}
