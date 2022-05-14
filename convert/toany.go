package convert

func ToAny[T any](vals ...T) []any {

	r := make([]any, len(vals))

	for i, v := range vals {
		r[i] = v
	}

	return r
}
