package helpers

func InSliceString(list []string, str string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func MergeSliceString(strings ...[]string) (merged []string) {
	for _, str := range strings {
		if str == nil {
			continue
		}
		merged = append(merged, str...)
	}

	return merged
}
