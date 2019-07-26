/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */
package main

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	re := make([][]int, 0)
	sort.Sort(heightAndRank(people))
	for _, v := range people {
		insert(&re, v, v[1])
	}
	return re
}

func insert(re *[][]int, h []int, k int) {
	//第一种 会因为第二个append产生新的底层数组，浪费内存，影响gc
	// *re = append((*re)[:k], append([][]int{h}, (*re)[k:]...)...)
	//第二种 直接copy覆盖，不会产生新的底层数组
	*re = append(*re, []int{})
	copy((*re)[k+1:], (*re)[k:])
	(*re)[k] = h
}

type heightAndRank [][]int

func (r heightAndRank) Len() int {
	return len(r)
}

func (r heightAndRank) Less(i, j int) bool {
	// 对h进行判断 按照降序排序
	if r[i][0] > r[j][0] {
		return true
	}
	if r[i][0] < r[j][0] {
		return false
	}
	// 对j进行判断 按照升序进行排序
	return r[i][1] < r[j][1]
}

func (r heightAndRank) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
