package format

//Currency Symbols
const (
	EURSymbol = "€"
	GBPSymbol = "£"
)

func EURCentsToCurrency(n int) string {
	return CentsToCurrency(n, EURSymbol, ",", false)
}

func EURCentsToCurrencyInternational(n int) string {
	return CentsToCurrency(n, EURSymbol, ".", true)
}

func EURCentsToCurrencyInternationalRounded(n int) string {
	return CentsToCurrencyRounded(n, EURSymbol, ".", true)
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

// CentsToCurrencyRounded returns the formated currency amount, without the fractions after the separator
func CentsToCurrencyRounded(n int, symbol string, separator string, preSymbol bool) string {
	left, _ := splitHundreths(n)
	//Use the flag passed as an argument to determine whether to put the currency
	//symbol at the end of the beginning
	if preSymbol {
		return symbol + left
	}
	return left + symbol
}
