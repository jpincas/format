package format

import (
	"strconv"
	"strings"
)

//HundrethsTo2dp takes an integer corresponding to hundreths e.g, cents, and returns a 2dp formatted string
//Examples:
//7 -> 0.07
//65 -> 0.65
//456 -> 4.65 etc.
func HundrethsTo2dp(n int) string {
	//If n is exactly 0, just return "0"
	if n == 0 {
		return "0"
	}
	//Otherwise split
	left, right := splitHundreths(n)
	return left + "." + right
}

func splitHundreths(n int) (left, right string) {
	var sign string
	//Convert to string
	s := strconv.Itoa(n)
	//Left zero padding if necessary
	//Mechanics depend on whether there is a negative sign or not
	//If it is a negative number, strip the prefix first and 'cache' it for later
	if n < 0 {
		s = strings.TrimPrefix(s, "-")
		sign = "-"
	}

	switch len(s) {
	case 1:
		s = "00" + s
	case 2:
		s = "0" + s
	}

	left = sign + s[:len(s)-2]
	right = s[len(s)-2:]

	return
}
