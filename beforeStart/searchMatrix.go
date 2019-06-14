func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        //如果是空矩阵，肯定找不到任何元素
        return false
    }
    n := len(matrix[0])
    for _, row := range matrix {
        if row[n-1] < target {
            // 如果本行最大值小于目标值，继续检查下一行
            continue
        }
        for _, col := range row {
            if row[0] > target {
                // 如果本行最小值大于目标值，则矩阵中不可能找到
                return false
            }
            if col == target {
                return true
            }
        }
    }
    return false    
}
