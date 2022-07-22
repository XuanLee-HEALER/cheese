package tools

func MapTo(os []any, f func(any) any) []any {
	res := make([]any, 0)
	for _, o := range os {
		res = append(res, f(o))
	}
	return res
}

func MapToString(os []any) []string {
	res := make([]string, 0)
	for _, o := range os {
		res = append(res, o.(string))
	}
	return res
}

func Filter(os []any, f func(any) bool) []any {
	cp := make([]any, 0)
	for _, o := range os {
		if f(o) {
			cp = append(cp, o)
		}
	}
	return cp
}
