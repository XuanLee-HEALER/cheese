package tools

import "regexp"

func FindAll(str string, reg string) []string {
	re, err := regexp.Compile(reg)
	if err != nil {
		return []string{}
	}
	if matched := re.FindAllString(str, -1); len(matched) > 0 {
		return matched
	}
	return []string{}
}
