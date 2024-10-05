//go:generate go run gen.go
//go:generate go fmt ./...
//go:generate go vet ./...
//go:generate go test ./...

package currency

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Currency is a type that by its methods represents currency in ISO 4217 format.
type Currency string

// AlphabeticCode returns the alphabetic code of the currency (e.g., "USD", "EUR")
//
// Is based on another ISO standard, ISO 3166, which lists the codes for country names.
// The first two letters of the three-letter ISO 4217 code match the country name code.
// Where possible, the third letter corresponds to the first letter of the currency name.
//
// If the currency is invalid, it returns empty string.
func (c Currency) AlphabeticCode() string { return currencies[c].alphabetic }

// NumericCode returns the three-digit numeric code of the currency (e.g., "840" for USD).
//
// Is useful when currency codes need to be understood in countries
// that do not use Latin scripts and for computerized systems.
// Where possible, the three-digit numeric code is the same as the numeric country code.
//
// If the currency is invalid, it returns empty string.
func (c Currency) NumericCode() string { return currencies[c].numeric }

// MinorUnits returns the number of digits after the decimal separator.
//
// If the currency is invalid, it returns 0 ("0").
func (c Currency) MinorUnits() MinorUnits { return currencies[c].units }

// Name returns the name of the currency (e.g., "US Dollar").
//
// If the currency is invalid, it returns "".
func (c Currency) Name() string { return currencies[c].name }

// CountryNames returns a list of locations listed for this currency.
//
// If the currency is invalid, it returns nil.
func (c Currency) CountryNames() []string { return currencies[c].countries }

// Alphabetic switches the currency to alphabetic representation.
//
// If the currency is invalid, the value will not change.
func (c *Currency) Alphabetic() {
	tmp, ok := currencies[*c]
	if !ok {
		return
	}
	*c = Currency(tmp.alphabetic)
}

// Numeric switches the currency to numeric representation.
//
// If the currency is invalid, the value will not change.
func (c *Currency) Numeric() {
	tmp, ok := currencies[*c]
	if !ok {
		return
	}
	*c = Currency(tmp.numeric)
}

func (c *Currency) UnmarshalJSON(data []byte) error {
	ccy, err := Identify(strings.Trim(string(data), `"`))
	if err == nil {
		*c = ccy
	}
	return err
}

// MinorUnits interface represents the number of digits after the decimal separator.
type MinorUnits interface {
	// Int returns the integer representation of the minor units.
	Int() int
	// String returns the string representation of the minor units.
	String() string
	// Power returns the power of 10 corresponding to the minor units of the currency.
	// It is used to convert an integer representation of units into a decimal amount.
	// For example, if the minor unit is 2 (as in USD or EUR), Power returns 100 (10^2).
	Power() float64
}

type units uint8

func (u units) Int() int {
	if u == 0xff {
		return 0
	}
	return int(u)
}

func (u units) String() string {
	if u == 0xff {
		return NotApplicable
	}
	return strconv.Itoa(int(u))
}

func (u units) Power() float64 { return math.Pow10(u.Int()) }

type stringNumber interface {
	~string | ~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Identify returns the currency based on the provided alphabetic code (e.g. "USD") or numeric code (e.g. "840").
// The numeric code can be a string or any integer type.
//
// It returns an error if the code is unknown.
func Identify[T stringNumber](code T) (Currency, error) {
	var c Currency
	switch v := any(code).(type) {
	case string:
		c = normalize(v)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		c = Currency(fmt.Sprintf("%03d", v))
	}
	if _, ok := currencies[c]; !ok {
		return "", fmt.Errorf("unknown currency code: %v", code)
	}
	return c, nil
}

func normalize(s string) Currency {
	if s = strings.TrimSpace(s); len(s) != 3 {
		return ""
	}
	return Currency(strings.ToUpper(s))
}
