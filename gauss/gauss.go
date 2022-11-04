package gauss

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
)

func swap(mat [][]int, x int, y int, n int) {
	for i := 0; i < n; i++ {
		var tmp int = mat[x][i]
		mat[x][i] = mat[y][i]
		mat[y][i] = tmp
	}
}

func gauss(m [][]int, n int, w http.ResponseWriter) {
	var issng int = felm(m, n)
	if issng == -1 {
		belm(m, n, w)
	} else {
		if m[issng][n-1] == 0 {
			io.WriteString(w, "Infinite solutions")

		} else {
			io.WriteString(w, "No solutions")
		}
	}
}

func felm(m [][]int, n int) int {
	for k := 0; k < n; k++ {
		var im int = k
		var mval int = int(math.Abs(float64(m[im][k])))
		for j := k + 1; j < n; j++ {
			if float64(mval) < math.Abs(float64(m[j][k])) {
				mval = m[j][k]
				im = j
			}
		}

		if mval == 0 {
			return k
		}
		swap(m, im, k, 0)
		for i := k + 1; i < n; i++ {
			fct := m[i][k] / m[k][k]
			for j := 0; j < n; j++ {
				m[i][j] -= fct * m[k][j]
			}
			m[i][k] = 0
		}
	}
	return -1
}

func belm(m [][]int, n int, w http.ResponseWriter) {
	x := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		x[i] = float64(m[i][n])
		for j := i + 1; j < n; j++ {
			x[i] -= float64(m[i][j]) * float64(x[j])
		}
		x[i] = x[i] / float64(m[i][i])
	}
	io.WriteString(w, "Solution is:\n ")
	var mp=make(map[int]string)
	mp[0]="x"
	mp[1]="y"
	mp[2]="z"
	mp[3]="w"

	for i, a := range x {
		io.WriteString(w, fmt.Sprintf("Value of %s is : %.6f\n",mp[i], a))
	}
}

func Build(n int, w http.ResponseWriter) {
	io.WriteString(w, "Matrix: \n")
	var size int
	size = n
	mat := make([][]int, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size+1; j++ {
			mat[i] = append(mat[i], rand.Intn(10)-rand.Intn(1))
		}
	}
	for i := 0; i < size; i++ {
		for j := 0; j <= size; j++ {
			io.WriteString(w, fmt.Sprintf("%v ", mat[i][j]))
		}
		io.WriteString(w, "\n")
	}
	gauss(mat, size, w)

}
