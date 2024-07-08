package repeat

func Repeat(c rune, timesToRepeat int) string {
	var result string
	for i := 0; i < timesToRepeat; i++ {
		result += string(c)
	}
	return result
}
