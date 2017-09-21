package format

import "strconv"

//HundrethsTo2dp takes an integer corresponding to hundreths e.g, cents, and returns a 2dp formatted string
//Examples:
//7 -> 0.07
//65 -> 0.65
//456 -> 4.65 etc.
func HundrethsTo2dp(n int) string {

	left, right := splitHundreths(n)
	return left + "." + right
}

func splitHundreths(n int) (left, right string) {
	//Convert to string and left pad if necessary
	s := strconv.Itoa(n)
	switch len(s) {
	case 1:
		s = "00" + s
	case 2:
		s = "0" + s
	}

	left = s[:len(s)-2]
	right = s[len(s)-2:]

	return
}
