// 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
//每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。

package _240_search_a_2d_matrix_2

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	// 从左下角启动开始遍历
	i := m - 1
	j := 0
	for i >= 0 && j <= n-1 {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}
	return false
}
