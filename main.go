package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"utils/utils"
)

// 示例算法：计算数组元素之和
func sumArray(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

// 示例算法：计算字符串长度
func stringLength(s string) int {
	return len(s)
}

// 读取测试数据文件
func readTestData(filename string) ([][]interface{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][]interface{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid test data format")
		}
		group := []interface{}{parts[0], parts[1]}
		data = append(data, group)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

// 写入结果文件
func writeResult(filename string, results []interface{}, times []time.Duration, datasizes []int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i, result := range results {
		_, err = writer.WriteString(fmt.Sprintf("Test %d:\n", i+1))
		if err != nil {
			return err
		}
		_, err = writer.WriteString(fmt.Sprintf("Data size: %d\n", datasizes[i]))
		if err != nil {
			return err
		}
		_, err = writer.WriteString(fmt.Sprintf("Result: %v\n", result))
		if err != nil {
			return err
		}
		_, err = writer.WriteString(fmt.Sprintf("Execution time: %s\n", times[i]))
		if err != nil {
			return err
		}
		_, err = writer.WriteString(fmt.Sprintf("Execution time in nanoseconds: %d ns\n\n", times[i].Nanoseconds()))
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

func main() {
	// 读取测试数据
	testDataGroups, err := readTestData("testdata.txt")
	if err != nil {
		fmt.Println("Error reading test data:", err)
		return
	}

	var results []interface{}
	var times []time.Duration
	var dataSizes []int
	// 循环处理每组测试数据
	for i, testData := range testDataGroups {
		// 记录开始时间
		start := time.Now()

		// 调用需要测试的算法
		var result interface{}
		switch len(testData) {
		case 1:
			// 处理单个整数或字符串
			switch v := testData[0].(type) {
			case int:
				result = v // 示例：直接返回整数
			case string:
				result = stringLength(v) // 示例：计算字符串长度
			}
		case 2:
			// 处理一个整数和一个数组
			if arr, ok := testData[1].([]int); ok {
				result = sumArray(arr) // 示例：计算数组元素之和
			}
			if str1, ok := testData[0].(string); ok {
				if str2, ok := testData[1].(string); ok {
					result, _ = utils.LoogestChildStr(str1, str2) // 示例：计算最长公共子序列
				}
			}
		}

		// 记录结束时间并计算耗时
		elapsed := time.Since(start)

		// 输出结果至控制台
		fmt.Printf("Test %d:\n", i+1)
		fmt.Printf("Result: %v\n", result)
		fmt.Printf("Execution time: %s\n", elapsed)
		fmt.Printf("Execution time in nanoseconds: %d ns\n\n", elapsed.Nanoseconds())

		// 记录结果和时间
		results = append(results, result)
		times = append(times, elapsed)

		// 输出数据规模
		dataSize := len(fmt.Sprintf("%v", testData[1]))
		fmt.Printf("Data size: %d\n\n", dataSize)
		dataSizes = append(dataSizes, dataSize)
		// Output data size to console
		//fmt.Printf("Data size: %d\n\n", dataSize)
	}

	// 写入结果文件
	err = writeResult("result.txt", results, times, dataSizes)
	if err != nil {
		fmt.Println("Error writing result:", err)
	}

	err = writeStandardAnswers("answer.txt", testDataGroups)
	if err != nil {
		fmt.Println("Error writing standard answers:", err)
	}
}

// 使用标准答案执行测试样例并输出到answer.txt
func writeStandardAnswers(filename string, testDataGroups [][]interface{}) error {
	var answers []interface{}

	for _, testData := range testDataGroups {
		if len(testData) == 2 {
			if str1, ok := testData[0].(string); ok {
				if str2, ok := testData[1].(string); ok {
					answer, _ := utils.Lcs(str1, str2)
					answers = append(answers, answer)
				}
			}
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i, answer := range answers {
		_, err = writer.WriteString(fmt.Sprintf("Test %d:\n", i+1))
		if err != nil {
			return err
		}
		_, err = writer.WriteString(fmt.Sprintf("Answer: %v\n\n", answer))
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
