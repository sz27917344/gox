package gox

import (
	"sort"
)

// Map 数组映射
func Map[S, T any](s []S, transform func(S) T) []T {
	rs := make([]T, len(s))
	for i, v := range s {
		rs[i] = transform(v)
	}
	return rs
}

// ForEach 循环
func ForEach[T any](s []*T, fn func(*T)) {
	for _, v := range s {
		fn(v)
	}
}

// Peek 提取数据进行处理
func Peek[T any](s []T, fn func(T)) []T {
	for _, v := range s {
		fn(v)
	}
	return s
}

// Filter 数组过滤
func Filter[T any](s []T, isPass func(T) bool) []T {
	rs := make([]T, 0)
	for _, v := range s {
		if isPass(v) {
			rs = append(rs, v)
		}
	}
	return rs
}

// Sort 排序
func Sort[T any](s []T, sortFunc func(s1, s2 T) bool) []T {
	sortStruct := sortByFun[T]{data: s, fun: sortFunc}
	sort.Sort(&sortStruct)
	return sortStruct.data
}

// Any 匹配
func Any[T any](s []T, anyFunc func(s T) bool) bool {
	for _, v := range s {
		if anyFunc(v) {
			return true
		}
	}
	return false
}

// None 匹配
func None[T any](s []T, noneFunc func(s T) bool) bool {
	for _, v := range s {
		if !noneFunc(v) {
			return true
		}
	}
	return false
}

// First 取符合条件的第一个值
func First[T any](s []T, fn func(s T) bool) *T {
	for _, v := range s {
		if fn(v) {
			return &v
		}
	}
	return nil
}

type mapKey interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64 | complex64 | complex128 | string | bool
}

// CollectToMap 转换成Map
func CollectToMap[T any, K mapKey, V any](s []T, fnKey func(s T) K, fnValue func(s T) V) map[K]V {
	m := make(map[K]V)
	for _, t := range s {
		m[fnKey(t)] = fnValue(t)
	}
	return m
}
