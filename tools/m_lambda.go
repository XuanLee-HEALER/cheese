package tools

func MapToStr(o any, f func(any) string) string {
	return f(o)
}

func MapToStrArr(os []any, f func(any) string) []string {
	res := []string{}
	for _, o := range os {
		res = append(res, f(o))
	}
	return res
}
