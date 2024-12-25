package utils

import "strings"

func ProcessCity(city string) (string, string) {
	var first string
	var second string
	if strings.Index(city, "特别行政区") != -1 {
		first = strings.Replace(city, "特别行政区", "", 1)
		second = strings.Replace(city, "特别行政区", "", 1)
	} else if strings.Index(city, "·") != -1 {
		temp := strings.Split(city, "·")
		if temp[0] == "广西壮族自治区" {
			first = "广西"
			second = strings.Replace(temp[1], "市", "", 1)
		} else if temp[0] == "内蒙古自治区" {
			first = "内蒙古"
			second = strings.Replace(temp[1], "市", "", 1)
		} else {
			first = strings.Replace(temp[0], "省", "", 1)
			second = strings.Replace(temp[1], "市", "", 1)
		}

	} else if strings.Index(city, "市") != -1 {
		first = strings.Replace(city, "市", "", 1)
		second = strings.Replace(city, "市", "", 1)
	}
	return first, second
}
