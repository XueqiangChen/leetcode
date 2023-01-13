package main

// public int maxAreaOfIsland(int[][] grid) {
//    int res = 0;
//    for (int r = 0; r < grid.length; r++) {
//        for (int c = 0; c < grid[0].length; c++) {
//            if (grid[r][c] == 1) {
//                int a = area(grid, r, c);
//                res = Math.max(res, a);
//            }
//        }
//    }
//    return res;
//}
//
//int area(int[][] grid, int r, int c) {
//    if (!inArea(grid, r, c)) {
//        return 0;
//    }
//    if (grid[r][c] != 1) {
//        return 0;
//    }
//    grid[r][c] = 2;
//
//    return 1
//        + area(grid, r - 1, c)
//        + area(grid, r + 1, c)
//        + area(grid, r, c - 1)
//        + area(grid, r, c + 1);
//}
//
//boolean inArea(int[][] grid, int r, int c) {
//    return 0 <= r && r < grid.length
//        	&& 0 <= c && c < grid[0].length;
//}

func numIslands(grid [][]byte) int {
	var dfs func(grid [][]byte, row int, col int)
	dfs = func(grid [][]byte, row int, col int) {
		if row >= len(grid) || col >= len(grid[0]) || row < 0 || col < 0 {
			return
		}

		if grid[row][col] != '1' {
			return
		}

		grid[row][col] = '2'
		dfs(grid, row-1, col)
		dfs(grid, row+1, col)
		dfs(grid, row, col-1)
		dfs(grid, row, col+1)
	}

	res := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '1' {
				dfs(grid, r, c)
				res++
			}
		}
	}
	return res
}
