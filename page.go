package gox

type PageInfo[T any] struct {
	Total int32 `json:"total"`
	List  []*T  `json:"list"`
}

func NewPageInfo[T any](total int32, list []*T) PageInfo[T] {
	return PageInfo[T]{Total: total, List: list}
}

func (pageInfo PageInfo[T]) Size() int {
	return len(pageInfo.List)
}
