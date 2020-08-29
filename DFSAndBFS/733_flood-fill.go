package DFSAndBFS

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if len(image) == 0 {
		return nil
	}
	startColor := image[sr][sc]
	dfs5(sr, sc, startColor, newColor, image)
	return image
}

func dfs5(startX, startY, startColor, newColor int, image [][]int) {
	if startX < 0 || startX >= len(image) || startY < 0 || startY >= len(image[0]) {
		return
	}
	if image[startX][startY] == newColor || image[startX][startY] != startColor {
		return
	}
	image[startX][startY] = newColor
	dfs5(startX-1, startY, startColor, newColor, image)
	dfs5(startX+1, startY, startColor, newColor, image)
	dfs5(startX, startY-1, startColor, newColor, image)
	dfs5(startX, startY+1, startColor, newColor, image)
}
