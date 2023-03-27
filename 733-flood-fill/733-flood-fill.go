package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	fill(image, sr, sc, image[sr][sc], newColor)
	return image
}

func fill(image [][]int, r, c, oldColor, newColor int) {
	if r < 0 || r >= len(image) || c < 0 || c >= len(image[0]) || image[r][c] != oldColor {
		return
	}
	image[r][c] = newColor
	fill(image, r-1, c, oldColor, newColor)
	fill(image, r+1, c, oldColor, newColor)
	fill(image, r, c-1, oldColor, newColor)
	fill(image, r, c+1, oldColor, newColor)
}

func main() {
	// image := [][]int{
	// 	{1, 1, 1},
	// 	{1, 1, 0},
	// 	{1, 0, 1},
	// }
	// sr := 1
	// sc := 1
	// newColor := 2
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	sr := 1
	sc := 1
	newColor := 2
	floodFill(image, sr, sc, newColor)
	fmt.Println(image)
}
