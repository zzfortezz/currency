package currency

import "testing"

func TestCurrencyFormat(t *testing.T) {
    tests := []struct {
        amount   float64
        currency string
        expected string
    }{
        {1234.56, "USD", "$1,234.56"},
        {1234.56, "EUR", "1.234,56 €"},
        {1234.56, "JPY", "1234.56"}, // Assuming JPY not defined
    }

    for _, test := range tests {
        result := CurrencyFormat(test.amount, test.currency)
        if result != test.expected {
            t.Errorf("CurrencyFormat(%f, %s) = %s; want %s", test.amount, test.currency, result, test.expected)
        }
    }
}
