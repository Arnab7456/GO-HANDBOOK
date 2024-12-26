package main

func searchMatrix(matrix [][]int, target int) bool {
    // if len(matrix) ==0 || len(matrix[0] ==0){
	// 	return false;
	// }
	rows, cols := len(matrix), len(matrix[0])
	row := 0
	col := cols-1

	for row <rows && col >=0 {
		curr := matrix[row][col]

		if curr == target {
			return true
		}else if curr > target {
			col--
		}else {
			row ++;
		}
	}
	return false;
}