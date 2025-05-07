package utils

import "strings"

// HiddenPhone 隐藏手机号中间4位
func HiddenPhone(phone string) string {
	if phone == "" || len(phone) != 11 {
		return ""
	}
	slice := strings.Split(phone, "")
	return strings.Join(slice[0:3], "") + "****" + strings.Join(slice[7:], "")
}
