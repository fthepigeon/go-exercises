package main

func generate(numRows int) [][]int {
	triangle := [][]int{}
	triangle = append(triangle, []int{1})

	for i := 1; i < numRows; i++ {
		inner := make([]int, i+1)
		inner[0], inner[len(inner)-1] = 1, 1

		for j := 1; j < len(inner)-1; j++ {
			inner[j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
		triangle = append(triangle, inner)
	}
	return triangle
}
