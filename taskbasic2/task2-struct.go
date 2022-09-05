package taskbasic2

type Number struct {
	N int
}

func (num *Number) Prime() []int {
	var result []int
	if num.N >= 5 {
		result = append(result, 2, 3, 5)
	} else if num.N >= 3 {
		result = append(result, 2, 3)
	} else if num.N >= 2 {
		result = append(result, 2)
	}
	for i := 0; i <= num.N; i++ {
		if i <= 1 || i%2 == 0 || i%3 == 0 || i%5 == 0 {
			continue
		} else {
			result = append(result, i)
		}
	}
	return result
}

func (num *Number) Odd() []int {
	var result []int
	for i := 0; i <= num.N; i++ {
		if i%2 == 1 {
			result = append(result, i)
		}
	}
	return result
}

func (num *Number) Even() []int {
	var result []int
	for i := 0; i <= num.N; i++ {
		if i%2 == 0 && i > 0 {
			result = append(result, i)
		}
	}
	return result
}

func (num *Number) Fibonacci() []int {
	var result []int
	if num.N >= 1 {
		result = append(result, 0, 1)
	} else if num.N >= 0 {
		result = append(result, 0)
	}
	for i := 0; i < num.N-1; i++ {
		var nextVal = result[i] + result[i+1]
		if nextVal > num.N {
			break
		}
		result = append(result, nextVal)
	}
	return result
}
