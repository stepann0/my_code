package main

func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])
	visited := map[int]bool{}

	for i := range grid {
		for j := range grid[i] {
			cell := grid[i][j]
			id := i*cols + j
			if cell == byte('1') {
				if _, ok := visited[id]; !ok {
					visited[id] = true
				}
			}
		}
	}
}

func main() {

}
