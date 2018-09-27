package listops

type binFunc func(int, int) int

type predFunc func(int) bool

type unaryFunc func(int) int

type IntList []int

func (i IntList) Foldr(f binFunc, initial int) int {
	result := initial

	for j := len(i) - 1; j >= 0; j-- {
		result = f(i[j], result)
	}

	return result
}

func (i IntList) Foldl(f binFunc, initial int) int {
	result := initial

	for j := range i {
		result = f(result, i[j])
	}

	return result
}

func (i IntList) Filter(f predFunc) IntList {
	var filtered IntList = []int{}

	for _, v := range i {
		if f(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

func (i IntList) Length() int {
	return len(i)
}

func (i IntList) Map(f unaryFunc) IntList {
	var mapped IntList = []int{}

	for _, v := range i {
		mapped = append(mapped, f(v))
	}

	return mapped
}

func (i IntList) Reverse() IntList {
	var reversed IntList = []int{}

	for j := len(i) - 1; j >= 0; j-- {
		reversed = append(reversed, i[j])
	}

	return reversed
}

func (i IntList) Append(other IntList) IntList {
	var appended IntList = []int{}

	appended = append(appended, i...)
	appended = append(appended, other...)

	return appended
}

func (i IntList) Concat(others []IntList) IntList {
	var concatenated IntList = []int{}

	concatenated = concatenated.Append(i)

	for _, other := range others {
		concatenated = concatenated.Append(other)
	}

	return concatenated
}
