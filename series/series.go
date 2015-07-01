package slice

func All(n int, s string) []string {
	if n > len(s) {
		return nil
	}

	last_index := len(s) - n
	series := make([]string, last_index+1)

	for i := 0; i <= last_index; i++ {
		series[i] = s[i : i+n]
	}

	return series
}

func Frist(n int, s string) string {
	return s[0:n]
}

func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "The number per series is larger than the string given.", false
	}

	return s[0:n], true
}
