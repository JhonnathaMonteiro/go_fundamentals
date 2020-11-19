package iteration

// Repeat return a repetition of character 5 times
func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += character
	}

	return repeated
}
