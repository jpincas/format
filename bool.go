package format

//BoolToYesOrNo formats True/False as "Yes" or "No"
func BoolToYesOrNo(tf bool) string {

	if tf {
		return "Yes"
	}
	return "No"
}
