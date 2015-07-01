package pascal

func Triangle(n int) [][]int {
	var triangle = make([][]int, n)
	triangle[0] = []int{1}

	for i := 1; i < n; i++ {
		previous := triangle[i-1]
		new_row_length := i + 1
		new_row := make([]int, new_row_length)
		var middle_index int

		if new_row_length%2 == 0 {
			// Row has even number length
			middle_index = (new_row_length / 2) - 1
		} else {
			// Row has odd number length
			middle_index = new_row_length / 2
		}

		// Middle out!
		for j := middle_index; j > 0; j-- {
			a := previous[j-1]
			b := previous[j]
			new_row[j], new_row[(new_row_length-1)-j] = a+b, a+b
		}

		new_row[0], new_row[new_row_length-1] = 1, 1

		triangle[i] = new_row
	}

	return triangle
}
