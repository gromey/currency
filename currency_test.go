package currency

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

func equal(t *testing.T, exp, got interface{}) {
	if !reflect.DeepEqual(exp, got) {
		t.Fatalf("Not equal:\nexp: %v\ngot: %v", exp, got)
	}
}

type testCS struct {
	currency Currency
	expect   string
}

type testCC struct {
	currency Currency
	expect   Currency
}

func TestCurrency_AlphabeticCode(t *testing.T) {
	var tests = []testCS{
		{
			currency: ALL,
			expect:   "ALL",
		},
		{
			currency: _008,
			expect:   "ALL",
		},
		{
			currency: "",
		},
		{
			currency: "AAA",
		},
		{
			currency: "000",
		},
	}
	for _, tt := range tests {
		iso := tt.currency
		equal(t, tt.expect, iso.AlphabeticCode())
	}
}

func TestCurrency_NumericCode(t *testing.T) {
	var tests = []testCS{
		{
			currency: ALL,
			expect:   "008",
		},
		{
			currency: _008,
			expect:   "008",
		},
		{
			currency: "",
		},
		{
			currency: "AAA",
		},
		{
			currency: "000",
		},
	}
	for _, tt := range tests {
		iso := tt.currency
		equal(t, tt.expect, iso.NumericCode())
	}
}

func TestCurrency_MinorUnits(t *testing.T) {
	var tests = []struct {
		currency     Currency
		expectInt    int
		expectString string
		expectPower  float64
	}{
		{
			currency:     ALL,
			expectInt:    2,
			expectString: "2",
			expectPower:  100,
		},
		{
			currency:     _008,
			expectInt:    2,
			expectString: "2",
			expectPower:  100,
		},
		{
			currency:     XAU,
			expectString: NotApplicable,
			expectPower:  1,
		},
		{
			currency:     _959,
			expectString: NotApplicable,
			expectPower:  1,
		},
		{
			currency:     "",
			expectString: "0",
			expectPower:  1,
		},
		{
			currency:     "AAA",
			expectString: "0",
			expectPower:  1,
		},
		{
			currency:     "000",
			expectString: "0",
			expectPower:  1,
		},
	}
	for _, tt := range tests {
		iso := tt.currency
		equal(t, tt.expectInt, iso.MinorUnits().Int())
		equal(t, tt.expectString, iso.MinorUnits().String())
		equal(t, tt.expectPower, iso.MinorUnits().Power())
	}
}

func TestCurrency_Name(t *testing.T) {
	var tests = []testCS{
		{
			currency: ALL,
			expect:   "Lek",
		},
		{
			currency: _008,
			expect:   "Lek",
		},
		{
			currency: "",
		},
		{
			currency: "AAA",
		},
		{
			currency: "000",
		},
	}
	for _, tt := range tests {
		iso := tt.currency
		equal(t, tt.expect, iso.Name())
	}
}

func TestCurrency_Countries(t *testing.T) {
	var tests = []struct {
		currency Currency
		expect   []string
	}{
		{
			currency: ALL,
			expect:   []string{"ALBANIA"},
		},
		{
			currency: _008,
			expect:   []string{"ALBANIA"},
		},
		{
			currency: "",
		},
		{
			currency: "AAA",
		},
		{
			currency: "000",
		},
	}
	for _, tt := range tests {
		iso := tt.currency
		equal(t, tt.expect, iso.CountryNames())
	}
}

func TestCurrency_Alphabetic(t *testing.T) {
	var tests = []testCC{
		{
			currency: _008,
			expect:   ALL,
		},
		{
			currency: ALL,
			expect:   ALL,
		},
		{
			currency: "",
			expect:   "",
		},
		{
			currency: "AAA",
			expect:   "AAA",
		},
		{
			currency: "000",
			expect:   "000",
		},
	}
	for _, tt := range tests {
		tt.currency.Alphabetic()
		equal(t, tt.expect, tt.currency)
	}
}

func TestCurrency_Numeric(t *testing.T) {
	var tests = []testCC{
		{
			currency: ALL,
			expect:   _008,
		},
		{
			currency: _008,
			expect:   _008,
		},
		{
			currency: "",
			expect:   "",
		},
		{
			currency: "AAA",
			expect:   "AAA",
		},
		{
			currency: "000",
			expect:   "000",
		},
	}
	for _, tt := range tests {
		tt.currency.Numeric()
		equal(t, tt.expect, tt.currency)
	}
}

func TestCurrency_UnmarshalJSON(t *testing.T) {
	type testStruct struct {
		Currency Currency
	}

	var tests = []struct {
		data        []byte
		testStruct  testStruct
		expect      Currency
		expectError error
	}{
		{
			data:   []byte("{\"currency\":\"ALL\"}"),
			expect: ALL,
		},
		{
			data:   []byte("{\"currency\":\"all\"}"),
			expect: ALL,
		},
		{
			data:   []byte("{\"currency\":\"008\"}"),
			expect: _008,
		},
		{
			data:   []byte("{\"currency\":\"  008  \"}"),
			expect: _008,
		},
		{
			data:        []byte("{\"currency\":\"AAA\"}"),
			testStruct:  testStruct{Currency: ALL},
			expect:      ALL,
			expectError: errors.New("unknown currency code: AAA"),
		},
		{
			data:        []byte("{\"currency\":8}"),
			testStruct:  testStruct{Currency: ALL},
			expect:      ALL,
			expectError: errors.New("unknown currency code: 8"),
		},
		{
			data:        []byte("{\"currency\":-9}"),
			testStruct:  testStruct{Currency: _008},
			expect:      _008,
			expectError: errors.New("unknown currency code: -9"),
		},
	}
	for _, tt := range tests {
		err := json.Unmarshal(tt.data, &tt.testStruct)
		equal(t, tt.expectError, err)
		equal(t, tt.expect, tt.testStruct.Currency)
	}
}

func TestIdentify(t *testing.T) {
	iso, err := Identify("ALL")
	equal(t, nil, err)
	equal(t, ALL, iso)

	iso, err = Identify("all")
	equal(t, nil, err)
	equal(t, ALL, iso)

	iso, err = Identify("008")
	equal(t, nil, err)
	equal(t, _008, iso)

	iso, err = Identify(8)
	equal(t, nil, err)
	equal(t, _008, iso)

	iso, err = Identify(uint(8))
	equal(t, nil, err)
	equal(t, _008, iso)

	_, err = Identify("ER0")
	equal(t, errors.New("unknown currency code: ER0"), err)

	_, err = Identify("ERU")
	equal(t, errors.New("unknown currency code: ERU"), err)

	_, err = Identify("32")
	equal(t, errors.New("unknown currency code: 32"), err)

	_, err = Identify(0)
	equal(t, errors.New("unknown currency code: 0"), err)
}
