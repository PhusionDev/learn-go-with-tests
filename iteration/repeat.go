package iteration

/*
Repeat takes a character and count and then returns a string with that
character repeated count number of times
*/
func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
