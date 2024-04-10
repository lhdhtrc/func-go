package array

// Include 是否包含某个值
func Include[T string | int](array []T, val T) bool {
	for _, item := range array {
		if item == val {
			return true
		}
	}
	return false
}
