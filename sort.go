package gox

// 排序相关类
type sortByFun[T any] struct {
	data []T
	fun  func(s1 T, s2 T) bool
}

func (s sortByFun[T]) Len() int       { return len(s.data) }
func (s *sortByFun[T]) Swap(i, j int) { s.data[i], s.data[j] = s.data[j], s.data[i] }
func (s *sortByFun[T]) Less(i, j int) bool {
	return s.fun(s.data[i], s.data[j])
}
