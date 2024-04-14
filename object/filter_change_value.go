package object

import (
	"github.com/lhdhtrc/func-go/array"
	"reflect"
)

// FilterChangeValue 过滤发生改变的值
func FilterChangeValue[O any, N any](old O, new N, ignore *[]string) map[string]interface{} {
	variant := make(map[string]interface{})

	ro := reflect.ValueOf(old).Elem()
	rn := reflect.ValueOf(new).Elem()
	for i := 0; i < rn.NumField(); i++ {
		k := rn.Type().Field(i).Name
		v := rn.Field(i)

		if array.Include(*ignore, k) || v.IsZero() || ro.FieldByName(k).Interface() == v.Interface() {
			continue
		}

		variant[k] = v.Interface()
	}

	return variant
}
