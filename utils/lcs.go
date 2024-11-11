package utils

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func LcsOutput(str1, str2 string, b [][]int) string {

	m := len(str1)
	n := len(str2)
	str := make([]byte, 0)
	for i, j := m, n; ; {
		if i == 0 || j == 0 {
			break
		}
		if b[i][j] == 1 {
			//str[k-1] = byte(str1[i-1])
			str = append([]byte{str1[i-1]}, str...)
			i--
			j--
			continue
		}
		if b[i][j] == 2 {
			i--
			continue
		}
		if b[i][j] == 3 {
			j--
			continue
		}

	}

	return string(str)
}
func LoogestChildStr(str1, str2 string) (int, string) {
	m := len(str1)
	n := len(str2)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	b := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		b[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i++ {
		dp[0][i] = 0
		b[0][i] = 0
	}

	for i := 0; i < n+1; i++ {
		dp[i][0] = 0
		b[i][0] = 0
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < m+1; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				b[i][j] = 1
			}
			if str1[i-1] != str2[j-1] {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				if max(dp[i][j-1], dp[i-1][j]) == dp[i][j-1] {
					b[i][j] = 3
				} else {
					b[i][j] = 2
				}
			}
		}
	}
	res := LcsOutput(str1, str2, b)

	return len(res), res
}
