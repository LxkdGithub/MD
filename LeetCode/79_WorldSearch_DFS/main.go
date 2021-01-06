package main

func main() {
	board := [][]byte{
		{'A','B','C','E'},
		{'S','F','C','S'},
		{'A','D','E','E'},
	}
	world := "ABCCED"
	println(WorldSearch(board, world))
	world = "SEE"
	println(WorldSearch(board, world))
	world = "ABCB"
	println(WorldSearch(board, world))
}

func WorldSearch(board [][]byte, world string) bool {
	if len(world) == 0 {
		return true
	}
	if len(board) == 0 {
		return false
	}

	visited := make([][]bool, len(board))
	for i:=0;i<len(visited);i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	// for 循环注意，这道题目是二维循环太大，放在主函数里
	for i:=0;i<len(board);i++ {
		for j:=0;j<len(board[0]);j++ {
			if DFSSearch(board, world, i, j, visited, 1) == true {
				return true
			}
		}
	}
	return false

}

func isInBoard(x, y int, board [][]byte) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}


func DFSSearch(board [][]byte, world string, x, y int, visited [][]bool, count int) bool {
	if count == len(world) {
		return board[x][y] == world[count-1]
	}
	if world[count - 1] == board[x][y] {
		visited[x][y] = true
		for i:=-1;i<=1;i++ {
			for j := -1;j<=1;j++ {
				new_x := x + i
				new_y := y + j
				if isInBoard(new_x, new_y, board) {
					if !visited[new_x][new_y] &&
						DFSSearch(board, world, new_x, new_y, visited, count+1) {
						return true
					}
				}
			}
		}
		visited[x][y] = false // 这里很重要，没有考虑到，不会影响到后边
	}
	return false
}
