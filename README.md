# currency

![https://img.shields.io/github/v/tag/gromey/currency](https://img.shields.io/github/v/tag/gromey/currency)
![https://img.shields.io/github/license/gromey/currency](https://img.shields.io/github/license/gromey/currency)

`currency` library of currencies based on the [ISO 4217 standard](https://www.iso.org/iso-4217-currency-codes.html)
published: ***2024-06-25***

## Installation

`currency` can be installed like any other Go library through `go get`:

```console
go get github.com/gromey/currency@latest
```

## Getting Started

```go
package main

import (
	"fmt"

	"github.com/gromey/currency"
)

func main() {
	jpy := currency.JPY

	fmt.Printf("AlphabeticCode: %s, NumericCode: %s, MinorUnits.Int: %d, MinorUnits.String: %s, Name: %s, CountryNames: %v\n",
		jpy.AlphabeticCode(), jpy.NumericCode(), jpy.MinorUnits().Int(), jpy.MinorUnits().String(), jpy.Name(), jpy.CountryNames())
	// Output: AlphabeticCode: JPY, NumericCode: 392, MinorUnits.Int: 0, MinorUnits.String: 0, Name: Yen, CountryNames: [JAPAN]

	ccy, err := currency.Identify("XAU")
	if err != nil {
		panic(err)
	}

	fmt.Printf("AlphabeticCode: %s, NumericCode: %s, MinorUnits.Int: %d, MinorUnits.String: %s, Name: %s, CountryNames: %v\n",
		ccy.AlphabeticCode(), ccy.NumericCode(), ccy.MinorUnits().Int(), ccy.MinorUnits().String(), ccy.Name(), ccy.CountryNames())
	// Output: AlphabeticCode: XAU, NumericCode: 959, MinorUnits.Int: 0, MinorUnits.String: N.A., Name: Gold, CountryNames: [ZZ08_Gold]

	ccy, err = currency.Identify(8)
	if err != nil {
		panic(err)
	}

	fmt.Printf("AlphabeticCode: %s, NumericCode: %s, MinorUnits.Int: %d, MinorUnits.String: %s, Name: %s, CountryNames: %v\n",
		ccy.AlphabeticCode(), ccy.NumericCode(), ccy.MinorUnits().Int(), ccy.MinorUnits().String(), ccy.Name(), ccy.CountryNames())
	// Output: AlphabeticCode: ALL, NumericCode: 008, MinorUnits.Int: 2, MinorUnits.String: 2, Name: Lek, CountryNames: [ALBANIA]
}
```

### WARNING!
**To simplify the use of the library, currency methods do not return errors.  
It is not recommended to create currency manually, like `currency.Currency("usd")`.  
Use the identify function instead: `currency.Identify("usd")`.**
