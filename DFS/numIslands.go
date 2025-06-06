// 200. Number of Islands
package dfs

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				//遞迴完是一座島
				grid = dfs(grid, i, j)
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) [][]byte {
	grid[i][j] = '0'
	if i-1 >= 0 && grid[i-1][j] == '1' { //i==0不走這個方向
		dfs(grid, i-1, j)
	}
	if i+1 < len(grid) && grid[i+1][j] == '1' { //i==len -1(最後一層)不走這個方向
		dfs(grid, i+1, j)
	}
	if j-1 >= 0 && grid[i][j-1] == '1' {
		dfs(grid, i, j-1)
	}
	if j+1 < len(grid[i]) && grid[i][j+1] == '1' {
		dfs(grid, i, j+1)
	}
	return grid
}

//Runtime 4ms Beats 70.88%
//Memory 5.64MB Beats 89.28%
