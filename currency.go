package format

//Currency Symbols
const (
	EURSymbol = ""
	GBPSymbol = "£"
)

//EURCentsToCurrency formats as follows: E5.67
func EURCentsToCurrency(n int) string {
	return CentsToCurrency(n, EURSymbol, ",", false)
}

//EURCentsToCurrencyInternational formats as follows: 5,67E
func EURCentsToCurrencyInternational(n int) string {
	return CentsToCurrency(n, EURSymbol, ".", true)
}

//GBPPenceToCurrency formats as follows: £5.67
func GBPPenceToCurrency(n int) string {
	return CentsToCurrency(n, GBPSymbol, ".", true)
}

//CentsToCurrency formats an integer number of cents (or pence etc) to a string
//with currency symbol and seperator
func CentsToCurrency(n int, symbol string, separator string, preSymbol bool) string {

	left, right := splitHundreths(n)
	amount := left + separator + right
	//Use the flag passed as an argument to determine whether to put the currency
	//symbol at the end of the beginning
	if preSymbol {
		return symbol + amount
	}
	return amount + symbol
}
