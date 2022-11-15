package iteration

func Repeat(sentence string) string {
	var result string

	for i := 0; i < 5; i++ {
		result += sentence
	}

	return result
}
