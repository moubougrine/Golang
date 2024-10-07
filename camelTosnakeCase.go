package piscine

func CamelToSnakeCase(s string) string {
	result := ""
	for i, ele := range s {
		if i != len(s)-1 && (ele >= 'A' && ele <= 'Z') && (s[i+1] >= 'A' && s[i+1] <= 'Z') {
			result+= s
			break
		}
		if i != len(s)-1 && (ele >= 'A' && ele <= 'Z') && i != 0 {
			if (s[i-1] >= 'a' && s[i-1] <= 'z') && (s[i+1] >= 'a' && s[i+1] <= 'z') {
				result += "_" + string(ele)
			}
		} else {
			result += string(ele)
		}
	}
	return result
}
