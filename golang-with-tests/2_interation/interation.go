package interation

func Repeat(character string, count int) string {
	var repeated string
	for range count {
		repeated += character
	}

	return repeated
}
