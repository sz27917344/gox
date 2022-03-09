package xmvc

import (
	"fmt"
	"github.com/sz27917344/gox/xerr"
	"github.com/sz27917344/gox/xres"
	"reflect"
	"strings"
)

func MvcValidate(bean interface{}) {
	object := reflect.ValueOf(bean)
	myref := object.Elem()
	typeOfType := myref.Type()
	for i := 0; i < myref.NumField(); i++ {
		field := myref.Field(i)
		mvc := typeOfType.Field(i).Tag.Get("mvc")
		if mvc == "" {
			switch field.Type().Kind() {
			case reflect.Int64:
				if field.Int() == 0 {
					panic(xerr.NewCodeMessage(xres.ParamError, typeOfType.Field(i).Name+"不能为空"))
				}
				break
			case reflect.String:
				if strings.TrimSpace(field.String()) == "" {
					panic(xerr.NewCodeMessage(xres.ParamError, typeOfType.Field(i).Name+"不能为空"))
				}
			}
			fmt.Printf("%d. %s %s = %v \n", i, typeOfType.Field(i).Name, field.Type(), field.Interface())
		}

	}
}