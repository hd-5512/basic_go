package leetcode

/**
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。

 思路从 i=0 j=0 开始 查看行与列中 没有重复
 然后把数字放到 对应 0-8的 数组里  0-8 =  i/3 + j/3*3  行 加 列的坐标位置 0  3

 */

var b = [][]byte{
	{'8','3','.','.','7','.','.','.','.'},
	{'6','.','.','1','9','5','.','.','.'},
	{'.','9','8','.','.','.','.','6','.'},
	{'8','.','.','.','6','.','.','.','3'},
	{'4','.','.','8','.','3','.','.','1'},
	{'7','.','.','.','2','.','.','.','6'},
	{'.','6','.','.','.','.','2','8','.'},
	{'.','.','.','4','1','9','.','.','5'},
	{'.','.','.','.','8','.','.','7','9'},
}

var c = [][]byte{
	{'.','.','5','.','.','.','6','3','.'},
	{'.','.','.','.','.','.','.','.','.'},
	{'5','.','.','.','.','.','.','9','.'},
	{'.','.','.','5','.','.','.','.','.'},
	{'4','.','3','.','.','.','.','.','1'},
	{'.','.','.','7','.','.','.','.','.'},
	{'.','.','.','5','.','.','.','.','.'},
	{'.','.','.','.','.','.','.','.','.'},
	{'.','.','.','.','.','.','.','.','.'},
}


func isValidSudoku(board [][]byte) bool {

	board = c

	line  := map[int] []string {}
	col   := map[int] []string {}
	block := map[int] []string {}


	for i:=0;i<9;i++ {
		line  = map[int] []string {}
		for j:=0;j<9;j++ {
			cur := string(board[i][j])
			if cur == "." {
				continue
			}

			//行中是否重复
			if _, ok := line[i]; ok {
				//存在
				for key := range line[i] {
					if cur == line[i][key] {
						return false  // 存在重复元素，标识为false
					}
				}
			}

			line[i] = append(line[i],cur)

			//列中是否重复
			if _, ok := col[j]; ok {
				for key := range col[j] {
					if cur == col[j][key] {
						return false // 存在重复元素，标识为false
					}
				}
			}
			col[j] = append(col[j],cur)


			//9格内是否重复
			board_pos := i/3*3 + j/3
			if _, ok := block[board_pos]; ok {
				for key := range block[board_pos] {
					if cur == block[board_pos][key] {
						return false // 存在重复元素，标识为false
					}
				}
			}
			block[board_pos] = append(block[board_pos],cur)
		}
	}
	return true
}