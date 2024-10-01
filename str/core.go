package str

import (
	"reflect"
	"strconv"
	"strings"
)

// ToString 尝试将值转换为字符串，支持数字和字符串类型
func ToString(v interface{}) string {
	switch t := v.(type) {
	case string:
		return t
	case int, int8, int16, int32, int64:
		return strconv.FormatInt(reflect.ValueOf(t).Int(), 10)
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatUint(reflect.ValueOf(t).Uint(), 10)
	default:
		return "" // 或返回错误
	}
}

// SliceToString 切片转字符串
func SliceToString[T any](arr interface{}, symbol string) string {
	if sa, ok := arr.([]T); ok {
		sm := make([]string, len(sa))
		for i, v := range sa {
			sm[i] = ToString(v)
		}
		return strings.Join(sm, symbol)
	}
	return ""
}
