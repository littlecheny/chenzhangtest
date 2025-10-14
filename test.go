package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	ans := math.MaxInt
	value := 0
	cacluNumSquare(int(math.Sqrt(float64(n))), n, value, &ans)
	return ans
}

func cacluNumSquare(a, b, value int, ans *int) {
	if a == 0 {
		if b == 0 {
			if value < *ans {
				*ans = value
				return
			}
		} else {
			return
		}
	}

	if b >= a*a {
		cacluNumSquare(a, b-a*a, value+1, ans)
		cacluNumSquare(a-1, b, value, ans)
	} else {
		cacluNumSquare(a-1, b, value, ans)
	}
}

// 删除未使用的 main 函数以避免 U1000 警告
func main() {
	n := 4
	fmt.Println(numSquares(n))
}
