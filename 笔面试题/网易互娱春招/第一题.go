package main

import "fmt"

/**
 * 接收两个表示9进制数的字符串，返回表示它们相加后的9进制数的字符串
 * @param num1 string字符串 第一个加数
 * @param num2 string字符串 第二个加数
 * @return string字符串
 */
func add(num1 string, num2 string) string {
	// write code here
	n1, n2 := find(num1), find(num2)
	result1 := ""
	if len(num1)-1-n1 > len(num2)-1-n2 {
		result1 = sum1(num2[n2+1:], num1[n1+1:])
	} else {
		result1 = sum1(num1[n1+1:], num2[n2+1:])
	}
	result2 := ""
	if n1 > n2 {
		result2 = sum2(num2[:n2], num1[:n1])
	} else {
		result2 = sum2(num1[:n1], num2[:n2])
	}
	if len(result1) > len(num1)-1-n1 && len(result1) > len(num2)-1-n2 {
		result2 = sum2(result2, result1[:1])
	}
	result := result2 + result1[1:]
	return result
}

func sum2(num1, num2 string) string {
	result := ""
	tmp := byte('0')
	for i := len(num1) - 1; i >= 0; i-- {
		result = fmt.Sprintf("%d%s", (num1[i]+num2[i]+tmp-'0'-'0'-'0')%9, result)
		tmp = (num1[i]+num2[i]+tmp-'0'-'0'-'0')/9 + '0'
	}
	for i := len(num2) - len(num1); i >= 0; i-- {
		result = fmt.Sprintf("%d%s", (num2[i]+tmp-'0'-'0')%9, result)
		tmp = (num2[i]+tmp-'0'-'0')/9 + '0'
	}
	for tmp > '9' {
		result = fmt.Sprintf("%d%s", (tmp-'0')%9, result)
		tmp = (tmp-'0')/9 + '0'
	}
	return result
}

func sum1(num1, num2 string) string {
	result := num2[len(num1)-1:]
	tmp := byte('0')
	for i := len(num1) - 1; i >= 0; i-- {
		result = fmt.Sprintf("%d%s", (num1[i]+num2[i]+tmp-'0'-'0'-'0')%9, result)
		tmp = (num1[i]+num2[i]+tmp-'0'-'0'-'0')/9 + '0'
	}
	for tmp > '9' {
		result = fmt.Sprintf("%d%s", (tmp-'0')%9, result)
		tmp = (tmp-'0')/9 + '0'
	}
	return result
}

func find(str string) int {
	for i := 0; i < len(str); i++ {
		if str[i] == '.' {
			return i
		}
	}
	return len(str)
}
