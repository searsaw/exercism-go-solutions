package secret

var words = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(n int) []string {
	if n > 0 {
		var list []int

		if n&16 == 0 {
			list = []int{0, 1, 2, 3}
		} else {
			list = []int{3, 2, 1, 0}
		}

		var handshake []string

		for _, ele := range list {
			if n&(1<<uint(ele)) != 0 {
				handshake = append(handshake, words[ele])
			}
		}

		if len(handshake) > 0 {
			return handshake
		}
	}

	return nil
}
