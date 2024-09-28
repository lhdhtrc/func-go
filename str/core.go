package str

import (
	"strconv"
	"strings"
)

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
