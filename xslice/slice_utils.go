package xslice

// IsNotBlank 切片非空
func IsNotBlank(slice []interface{}) bool {
	return slice != nil && len(slice) > 0
}

// IsBlank 切片为空
func IsBlank(slice []interface{}) bool {
	return slice == nil || len(slice) == 0
}
