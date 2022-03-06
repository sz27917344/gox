package gox

import (
	"reflect"
)

// MapStruct 对结构体进行映射
func MapStruct[T any](src interface{}, dst *T) *T {
	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	srcType, srcValue := reflect.TypeOf(src), reflect.ValueOf(src)
	// dst必须结构体指针类型
	if dstType.Kind() != reflect.Ptr || dstType.Elem().Kind() != reflect.Struct {
		panic(NewCode(COPY_FAILED))
	}
	// src必须为结构体或者结构体指针
	if srcType.Kind() == reflect.Ptr {
		srcType, srcValue = srcType.Elem(), srcValue.Elem()
	}
	if srcType.Kind() != reflect.Struct {
		panic(NewCode(COPY_FAILED))
	}
	// 取具体内容
	dstType, dstValue = dstType.Elem(), dstValue.Elem()
	// 属性个数
	propertyNums := dstType.NumField()
	for i := 0; i < propertyNums; i++ {
		// 属性
		property := dstType.Field(i)
		// 待填充属性值
		propertyValue := srcValue.FieldByName(property.Name)
		// 无效，说明src没有这个属性 || 属性同名但类型不同
		if !propertyValue.IsValid() || property.Type != propertyValue.Type() {
			continue
		}
		if dstValue.Field(i).CanSet() {
			dstValue.Field(i).Set(propertyValue)
		}
	}
	return dst
}

// MapSlice 对切片类型进行映射
func MapSlice[S any, T any](srcSlice []S, dst T) []*T {
	// 创建数组
	dstSlice := make([]*T, len(srcSlice))
	// 对数组中的元素进行遍历复制
	for i, src := range srcSlice {
		dstSlice[i] = MapStruct(src, new(T))
	}
	return dstSlice
}

// MapPage 对分页结果进行类型映射
func MapPage[S any, T any](srcPage PageInfo[S], dst T) (targetPage PageInfo[T]) {
	// 设置最大值
	targetPage.Total = srcPage.Total
	// 对列表进行映射
	targetPage.List = MapSlice(srcPage.List, dst)
	return
}
