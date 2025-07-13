package iteration

func Repetead(char string, time int) string {
	var repetitions string
	for index := 0; index < time; index++ {
		repetitions += char
	}
	return repetitions
}
