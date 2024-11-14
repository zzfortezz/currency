# Currency

A Go package for formatting currency amounts according to various international standards.

## Features

- Supports multiple currencies with customizable formatting.
- Easy to format amounts with appropriate symbols, separators, and decimal places.
- Extendable to include additional currencies.

## Installation

```bash
go get github.com/yourusername/currency
```

## Example

```go
package main

import (
    "fmt"
    "github.com/zzforetezz/currency"
)

func main() {
    amount := 1234.56
    currencyCode := "USD"
    formatted := currency.CurrencyFormat(amount, currencyCode)
    fmt.Println(formatted) // Output: $1,234.56
}
```

# Supported Currencies

- USD
- GBP
- THB
- IDR
- VND
- MYR
- AED
- AFN
- ALL
- AMD
- ANG
- AOA
- ARS
- AUD
- AWG
- AZN
- BAM
- BBD

Add more currencies as needed in the `CURRENCIES` map.

## Testing

Run the unit tests using:

```sh
# Command to run unit tests (replace with actual command if available)
