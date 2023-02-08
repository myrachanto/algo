package main

import (
	"fmt"
	"log"
)

func init(){
	log.SetPrefix("Matrix multiplication ...")

}
func main(){
	mat1 := [][]int{
		{1,2,4},
		{5,6,4},
		{1,2,1},
	}
	mat2 := [][]int{
		{1,3,2},
		{5,2,4},
		{3,2,1},
	}
	results := MatrixMulti(mat1, mat2)
	fmt.Println(results)

}
func MatrixMulti(mat1, mat2 [][]int)[][]int{
	results := make([][]int, len(mat1))
	for i := 0; i < len(mat1); i++{
		results[i] = make([]int, len(mat1))
		for j := 0; j < len(mat2); j++{
			for k := 0; k <len(mat2); k++{
				results[i][j] += mat1[i][k] * mat2[k][j]
			}
		}
	}
	return results
}