package piscine

func Capitalize(s string) string {

	result := ""
	for i,ele := range s {
		if (ele >= 'a' && ele <= 'z')&& s[i-1] <= 48 {
			result+=string(ele - 32)
		}else {
			result+=string(ele)
		}
	}
	return result
}