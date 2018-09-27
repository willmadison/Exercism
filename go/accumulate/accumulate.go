package accumulate

const testVersion = 1

func Accumulate(values []string, converter func(string) string) []string {
	result := []string{}

	for _, value := range values {
		result = append(result, converter(value))
	}

	return result
}
