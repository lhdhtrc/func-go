package object

import (
	"reflect"
	"slices"
)

// FilterChangeValue 过滤发生改变的值
func FilterChangeValue[O any, N any](old O, new N, ignore []string) map[string]interface{} {
	ro := reflect.ValueOf(old)
	rn := reflect.ValueOf(new)

	if ro.Kind() != reflect.Struct && (ro.Kind() != reflect.Ptr || ro.Elem().Kind() != reflect.Struct) {
		return nil
	}
	if rn.Kind() != reflect.Struct && (rn.Kind() != reflect.Ptr || rn.Elem().Kind() != reflect.Struct) {
		return nil
	}

	if ro.Kind() == reflect.Ptr {
		ro = ro.Elem()
	}
	if rn.Kind() == reflect.Ptr {
		rn = rn.Elem()
	}

	variant := make(map[string]interface{})

	for i := 0; i < rn.NumField(); i++ {
		k := rn.Type().Field(i).Name

		if slices.Contains(ignore, k) || !ro.FieldByName(k).IsValid() {
			continue
		}

		nv := rn.Field(i).Interface()
		ov := ro.FieldByName(k).Interface()

		if !reflect.DeepEqual(ov, nv) {
			variant[k] = nv
		}
	}

	return variant
}
