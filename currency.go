package currency

import (
	"fmt"
	"strings"
)

var formatMapping = []struct {
	SymbolOnLeft                bool
	SpaceBetweenAmountAndSymbol bool
	Format                      string
}{
	{true, false, "{{symbol}}{{value}}"},
	{true, true, "{{symbol}} {{value}}"},
	{false, false, "{{value}}{{symbol}}"},
	{false, true, "{{value}} {{symbol}}"},
}
func formatNumber(amount float64, decimalDigits int, decimalSep, thousandSep string) string {
    // Simple number formatting
    s := fmt.Sprintf(fmt.Sprintf("%%.%df", decimalDigits), amount)
    parts := strings.Split(s, ".")
    intPart := parts[0]
    decPart := ""
    if len(parts) > 1 {
        decPart = parts[1]
    }

    // Add thousands separator
    var result strings.Builder
    for i, c := range intPart {
        if i != 0 && (len(intPart)-i)%3 == 0 {
            result.WriteString(thousandSep)
        }
        result.WriteRune(c)
    }

    if decimalDigits > 0 {
        result.WriteString(decimalSep)
        result.WriteString(decPart)
    }

    return result.String()
}

func CurrencyFormat(amount float64, currency string) string {
    c, exists := CURRENCIES[currency]
    if !exists {
        return fmt.Sprintf("%.2f", amount)
    }

    formattedNumber := formatNumber(amount, c.DecimalDigits, c.DecimalSeparator, c.ThousandsSeparator)

    var format string
    for _, fm := range formatMapping {
        if fm.SymbolOnLeft == c.SymbolOnLeft && fm.SpaceBetweenAmountAndSymbol == c.SpaceBetweenAmountAndSymbol {
            format = fm.Format
            break
        }
    }

    formattedString := strings.Replace(format, "{{symbol}}", c.Symbol, 1)
    return strings.Replace(formattedString, "{{value}}", formattedNumber, 1)
}

func ConvertCurrency(fromCurrency, toCurrency string, value float64, exchangeRates map[string]float64) float64 {
	quote1 := exchangeRates[fromCurrency]
	quote2 := exchangeRates[toCurrency]
	return value / quote1 * quote2
}
