package helpers

func StrLimit(str *string, limit int) *string {
	if nil == str {
		return nil
	}
	if limit < 0 {
		empty := ""
		return &empty
	}
	newStr := *str
	if len(newStr) > limit {
		newStr = newStr[0:limit]
		return &newStr
	}
	return &newStr
}
