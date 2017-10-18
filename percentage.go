package format

import "fmt"

//Percentage2dp formats a float as a 2dp percentag with % sign
func Percentage2dp(n float64) string {
	return fmt.Sprintf("%.2f%%", n)
}

//Percentage1dp formats a float as a 1dp percentag with % sign
func Percentage1dp(n float64) string {
	return fmt.Sprintf("%.1f%%", n)
}
