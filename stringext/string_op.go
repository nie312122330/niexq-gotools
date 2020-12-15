package stringext

// CutString ...
func CutString(str string, length int) string {
	if len(str) > length {
		if length > 6 {
			resultStr := str[0 : length-3]
			return resultStr + "..."
		}
		return str[0:length]
	}
	return str
}
