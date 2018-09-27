package fb

func Setstrings(first string, second string) []string {
	feed := make([]string, 10)
	feed[0] = first
	feed[1] = second
	return feed
}
