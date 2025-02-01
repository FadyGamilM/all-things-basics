package main

func main() {
	inp := [][]int{
		{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60},
	}
	searchMatrix(inp, 3)
}

func searchMatrix(matrix [][]int, target int) bool {
	data := []int{}
	for i := range matrix {
		row := matrix[i]
		data = append(data, row...)
	}

	l, h := 0, len(data)-1
	m := (l + h) / 2

	for l <= h {
		if target == data[m] {
			return true
		} else if target > data[m] {
			l = m + 1
			m = (l + h) / 2
		} else {
			h = m - 1
			m = (l + h) / 2
		}
	}
	return false
}
