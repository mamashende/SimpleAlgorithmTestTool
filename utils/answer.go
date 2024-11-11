package utils

// lcs 计算两个字符串的最长公共子序列的长度

func Lcs(str1, str2 string) (int, string) {
	m := len(str1)
	n := len(str2)
	c := make([][]int, m+1)
	b := make([][]int, m+1)
	for i := range c {
		c[i] = make([]int, n+1)
		b[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				c[i][j] = c[i-1][j-1] + 1
				b[i][j] = 1
			} else if c[i-1][j] >= c[i][j-1] {
				c[i][j] = c[i-1][j]
				b[i][j] = 2
			} else {
				c[i][j] = c[i][j-1]
				b[i][j] = 3
			}
		}
	}
	res := LcsOutput(str1, str2, b)

	return len(res), res
}
