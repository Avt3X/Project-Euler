package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

var mat [15][15] int

func problem18() {
	open, _ := os.Open("triangle.txt")
	scanner := bufio.NewScanner(open)
	j := 0
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		var arr [15] int
		k := 0
		for i := range split {
			n, _ := strconv.Atoi(split[i])
			arr[k] = n
			k++
		}
		mat[j] = arr
		j++
	}
}

func solve() int {
	for i := 13; i >= 0; i-- {
		for j := 0; j < i+1; j++ {
			mat[i][j] += int(math.Max(float64(mat[i+1][j+1]), float64(mat[i+1][j])))
		}
	}

	res := 0
	for i := range mat[0] {
		if res < mat[0][i] {
			res = mat[0][i]
		}
	}

	return res
}
