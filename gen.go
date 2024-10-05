//go:build ignore

// Generator for currency-related data.

package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/gromey/currency"
)

func main() {
	if err := gen(); err != nil {
		log.Fatal(err)
	}
}

const (
	rawURL = "https://www.six-group.com/dam/download/financial-information/data-center/iso-currrency/lists/list-one.xml"

	readme        = "README.md"
	constFileName = "constants.go"
	constTemplate = `package currency

const NotApplicable string = "N.A." // No digits after the decimal separator.

const ({{range .Currencies}}
	{{.AlphabeticCode}}, _{{.NumericCode}} Currency = "{{.AlphabeticCode}}", "{{.NumericCode}}" // {{.Name}}{{end}}
)
`

	sourcesFileName = "sources.go"
	sourcesTemplate = `package currency

// internal currency type represents currencies in ISO 4217 format.
type currency struct {
    alphabetic string
    numeric    string
    units      units
    name       string
    countries  []string
}

var (
{{- range .Currencies}}
    {{.AlphabeticCodeLow}} = currency{alphabetic: "{{.AlphabeticCode}}", numeric: "{{.NumericCode}}", units: 0x{{.MinorUnits}}, name: "{{.Name}}", countries: {{.CountryName}}}{{end}}
)

var currencies = map[Currency]currency{
{{- range $i, $el := .Currencies}}{{if lf $i 5}}
    {{end}}{{.AlphabeticCode}}: {{.AlphabeticCodeLow}}, _{{.NumericCode}}: {{.AlphabeticCodeLow}}, {{end}}
}
`
)

var (
	prefix = []byte(`published: ***`)
	suffix = []byte(`***`)
)

type minorUnits string

func (m *minorUnits) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}
	if strings.TrimSpace(v) != currency.NotApplicable {
		n, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*m = minorUnits(fmt.Sprintf("%02d", n))
	} else {
		*m = "ff"
	}
	return nil
}

type currencyRow struct {
	AlphabeticCode    string     `xml:"Ccy"`
	AlphabeticCodeLow string     `xml:"-"`
	NumericCode       string     `xml:"CcyNbr"`
	MinorUnits        minorUnits `xml:"CcyMnrUnts"`
	Name              string     `xml:"CcyNm"`
	CountryName       string     `xml:"CtryNm"`
}

type currencyTable struct {
	CurrencyRows []currencyRow `xml:"CcyNtry"`
}

type iso4217 struct {
	Published     string        `xml:"Pblshd,attr"`
	CurrencyTable currencyTable `xml:"CcyTbl"`
}

type data struct {
	Counter    int
	Currencies []currencyRow
}

type isoResult struct {
	alphabeticCode string
	numericCode    string
	minorUnits     minorUnits
	name           string
	countryNames   []string
}

func gen() error {
	iso := new(iso4217)

	if err := iso.get(); err != nil {
		return err
	}

	return iso.makeFiles()
}

func (iso *iso4217) get() error {
	response, err := http.Get(rawURL)
	if err != nil {
		return err
	}
	defer func() { _ = response.Body.Close() }()

	var body []byte
	if body, err = io.ReadAll(response.Body); err != nil {
		return err
	}

	return xml.Unmarshal(body, iso)
}

func (iso *iso4217) makeFiles() error {
	mapISO := make(map[string]isoResult, len(iso.CurrencyTable.CurrencyRows))

	for _, ccyRow := range iso.CurrencyTable.CurrencyRows {
		if strings.TrimSpace(ccyRow.AlphabeticCode) == "" {
			continue
		}

		if v, ok := mapISO[strings.TrimSpace(ccyRow.AlphabeticCode)]; ok {
			v.countryNames = append(v.countryNames, strings.TrimSpace(ccyRow.CountryName))
			mapISO[strings.TrimSpace(ccyRow.AlphabeticCode)] = v
			continue
		}

		isoResultRow := isoResult{
			alphabeticCode: strings.TrimSpace(ccyRow.AlphabeticCode),
			numericCode:    ccyRow.NumericCode,
			minorUnits:     ccyRow.MinorUnits,
			name:           strings.TrimSpace(ccyRow.Name),
			countryNames:   []string{strings.TrimSpace(ccyRow.CountryName)},
		}

		mapISO[ccyRow.AlphabeticCode] = isoResultRow
	}

	currencies := make([]currencyRow, len(mapISO))

	i := 0
	for _, v := range mapISO {
		currencies[i] = currencyRow{
			AlphabeticCode:    v.alphabeticCode,
			AlphabeticCodeLow: strings.ToLower(v.alphabeticCode),
			NumericCode:       v.numericCode,
			MinorUnits:        v.minorUnits,
			Name:              v.name,
			CountryName:       fmt.Sprintf("%#v", v.countryNames),
		}
		i++
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].AlphabeticCode < currencies[j].AlphabeticCode
	})

	result := data{
		Counter:    len(currencies) + 1,
		Currencies: currencies,
	}

	if err := createFileByTemplate(constTemplate, constFileName, result); err != nil {
		return err
	}

	if err := createFileByTemplate(sourcesTemplate, sourcesFileName, result); err != nil {
		return err
	}

	return updateReadme(iso.Published)
}

var funcMap = template.FuncMap{
	"lf": func(i, n int) bool { return (i)%n == 0 },
	"ii": func(i int) int { return i + 1 },
}

func createFileByTemplate(tempData, filename string, data any) error {
	temp, err := template.New(filename).Funcs(funcMap).Parse(tempData)
	if err != nil {
		return err
	}

	var file *os.File
	if file, err = os.Create(filename); err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	return temp.Execute(file, data)
}

func updateReadme(published string) error {
	content, err := os.ReadFile(readme)
	if err != nil {
		return err
	}

	i := bytes.Index(content, prefix)
	if i == -1 {
		return fmt.Errorf("could not find prefix in README.md file")
	}
	i += len(prefix)

	n := bytes.Index(content[i:], suffix)
	if n == -1 {
		return fmt.Errorf("could not find suffix in README.md file")
	}
	n += i

	var file *os.File
	if file, err = os.Create(readme); err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	if _, err = file.Write(content[:i]); err != nil {
		return err
	}
	if _, err = file.Write([]byte(published)); err != nil {
		return err
	}
	if _, err = file.Write(content[n:]); err != nil {
		return err
	}

	return nil
}
