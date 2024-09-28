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

// StringToUintSlice 字符串转数字切片
func StringToUintSlice(s string) []uint {
	parts := strings.Split(s, ",")
	var result []uint

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		num, err := strconv.ParseUint(part, 10, strconv.IntSize)
		if err != nil {
			continue
		}
		result = append(result, uint(num))
	}

	return result
}

func UintSliceToString(slice []uint) string {
	str := make([]string, len(slice))
	for i, num := range slice {
		str[i] = strconv.Itoa(int(num))
	}
	return strings.Join(str, ",")
}
