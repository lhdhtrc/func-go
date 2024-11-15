package array

// DiffStringSlice 返回两个字符串切片之间的差异：新增的和删除的元素
func DiffStringSlice(original, updated []string) ([]string, []string) {
	originalMap := make(map[string]bool)
	updatedMap := make(map[string]bool)
	var added []string
	var removed []string

	// 构建原切片的映射
	for _, item := range original {
		originalMap[item] = true
	}

	// 构建新切片的映射，并找出差异
	for _, item := range updated {
		updatedMap[item] = true
		if !originalMap[item] {
			// 如果新切片中的元素不在原切片中，则视为新增
			added = append(added, item)
		} else {
			// 标记为已存在，以便后续检查删除
			delete(originalMap, item)
		}
	}

	// 找出原切片中删除的元素
	for item := range originalMap {
		removed = append(removed, item)
	}

	return added, removed
}
