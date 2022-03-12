package xmvc

import (
	"github.com/sz27917344/gox/xerr"
	"github.com/sz27917344/gox/xres"
	"reflect"
)

// MvcValidate 对参数进行校验
func MvcValidate(bean interface{}) {
	value := reflect.ValueOf(bean)
	// 对指针或非指针进行兼容
	value = reflect.Indirect(value)
	typeOfType := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		mvc := typeOfType.Field(i).Tag.Get("mvc")
		// 存在mvc标记时
		if len(mvc) == 0 {
			// 字段值为空
			if isBlank(field) {
				panic(xerr.NewCodeMessage(xres.ParamError, typeOfType.Field(i).Name+"不能为空"))
			}
		}
	}
}

// 判断值是否非空
func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}
