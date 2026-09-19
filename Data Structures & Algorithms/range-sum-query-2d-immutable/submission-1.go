type NumMatrix struct {
	Mtx [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sumMtx := make([][]int, len(matrix)+1)
	for i := range sumMtx {
		sumMtx[i] = make([]int, len(matrix[0])+1)
	}
	for i := range matrix {
		for j := range matrix[i] {
			sumMtx[i+1][j+1] = sumMtx[i][j+1] + sumMtx[i+1][j] - sumMtx[i][j] + matrix[i][j]
		}
	}
	return NumMatrix{
		Mtx: sumMtx,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.Mtx[row1][col1] + this.Mtx[row2+1][col2+1] - this.Mtx[row2+1][col1] - this.Mtx[row1][col2+1]
}

// Your NumMatrix object will be instantiated and called as such:
// obj := Constructor(matrix)
// param_1 := obj.SumRegion(row1,col1,row2,col2)
