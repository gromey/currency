package currency

// internal currency type represents currencies in ISO 4217 format.
type currency struct {
	alphabetic string
	numeric    string
	units      units
	name       string
	countries  []string
}

var (
	aed = currency{alphabetic: "AED", numeric: "784", units: 0x02, name: "UAE Dirham", countries: []string{"UNITED ARAB EMIRATES (THE)"}}
	afn = currency{alphabetic: "AFN", numeric: "971", units: 0x02, name: "Afghani", countries: []string{"AFGHANISTAN"}}
	all = currency{alphabetic: "ALL", numeric: "008", units: 0x02, name: "Lek", countries: []string{"ALBANIA"}}
	amd = currency{alphabetic: "AMD", numeric: "051", units: 0x02, name: "Armenian Dram", countries: []string{"ARMENIA"}}
	ang = currency{alphabetic: "ANG", numeric: "532", units: 0x02, name: "Netherlands Antillean Guilder", countries: []string{"CURAÇAO", "SINT MAARTEN (DUTCH PART)"}}
	aoa = currency{alphabetic: "AOA", numeric: "973", units: 0x02, name: "Kwanza", countries: []string{"ANGOLA"}}
	ars = currency{alphabetic: "ARS", numeric: "032", units: 0x02, name: "Argentine Peso", countries: []string{"ARGENTINA"}}
	aud = currency{alphabetic: "AUD", numeric: "036", units: 0x02, name: "Australian Dollar", countries: []string{"AUSTRALIA", "CHRISTMAS ISLAND", "COCOS (KEELING) ISLANDS (THE)", "HEARD ISLAND AND McDONALD ISLANDS", "KIRIBATI", "NAURU", "NORFOLK ISLAND", "TUVALU"}}
	awg = currency{alphabetic: "AWG", numeric: "533", units: 0x02, name: "Aruban Florin", countries: []string{"ARUBA"}}
	azn = currency{alphabetic: "AZN", numeric: "944", units: 0x02, name: "Azerbaijan Manat", countries: []string{"AZERBAIJAN"}}
	bam = currency{alphabetic: "BAM", numeric: "977", units: 0x02, name: "Convertible Mark", countries: []string{"BOSNIA AND HERZEGOVINA"}}
	bbd = currency{alphabetic: "BBD", numeric: "052", units: 0x02, name: "Barbados Dollar", countries: []string{"BARBADOS"}}
	bdt = currency{alphabetic: "BDT", numeric: "050", units: 0x02, name: "Taka", countries: []string{"BANGLADESH"}}
	bgn = currency{alphabetic: "BGN", numeric: "975", units: 0x02, name: "Bulgarian Lev", countries: []string{"BULGARIA"}}
	bhd = currency{alphabetic: "BHD", numeric: "048", units: 0x03, name: "Bahraini Dinar", countries: []string{"BAHRAIN"}}
	bif = currency{alphabetic: "BIF", numeric: "108", units: 0x00, name: "Burundi Franc", countries: []string{"BURUNDI"}}
	bmd = currency{alphabetic: "BMD", numeric: "060", units: 0x02, name: "Bermudian Dollar", countries: []string{"BERMUDA"}}
	bnd = currency{alphabetic: "BND", numeric: "096", units: 0x02, name: "Brunei Dollar", countries: []string{"BRUNEI DARUSSALAM"}}
	bob = currency{alphabetic: "BOB", numeric: "068", units: 0x02, name: "Boliviano", countries: []string{"BOLIVIA (PLURINATIONAL STATE OF)"}}
	bov = currency{alphabetic: "BOV", numeric: "984", units: 0x02, name: "Mvdol", countries: []string{"BOLIVIA (PLURINATIONAL STATE OF)"}}
	brl = currency{alphabetic: "BRL", numeric: "986", units: 0x02, name: "Brazilian Real", countries: []string{"BRAZIL"}}
	bsd = currency{alphabetic: "BSD", numeric: "044", units: 0x02, name: "Bahamian Dollar", countries: []string{"BAHAMAS (THE)"}}
	btn = currency{alphabetic: "BTN", numeric: "064", units: 0x02, name: "Ngultrum", countries: []string{"BHUTAN"}}
	bwp = currency{alphabetic: "BWP", numeric: "072", units: 0x02, name: "Pula", countries: []string{"BOTSWANA"}}
	byn = currency{alphabetic: "BYN", numeric: "933", units: 0x02, name: "Belarusian Ruble", countries: []string{"BELARUS"}}
	bzd = currency{alphabetic: "BZD", numeric: "084", units: 0x02, name: "Belize Dollar", countries: []string{"BELIZE"}}
	cad = currency{alphabetic: "CAD", numeric: "124", units: 0x02, name: "Canadian Dollar", countries: []string{"CANADA"}}
	cdf = currency{alphabetic: "CDF", numeric: "976", units: 0x02, name: "Congolese Franc", countries: []string{"CONGO (THE DEMOCRATIC REPUBLIC OF THE)"}}
	che = currency{alphabetic: "CHE", numeric: "947", units: 0x02, name: "WIR Euro", countries: []string{"SWITZERLAND"}}
	chf = currency{alphabetic: "CHF", numeric: "756", units: 0x02, name: "Swiss Franc", countries: []string{"LIECHTENSTEIN", "SWITZERLAND"}}
	chw = currency{alphabetic: "CHW", numeric: "948", units: 0x02, name: "WIR Franc", countries: []string{"SWITZERLAND"}}
	clf = currency{alphabetic: "CLF", numeric: "990", units: 0x04, name: "Unidad de Fomento", countries: []string{"CHILE"}}
	clp = currency{alphabetic: "CLP", numeric: "152", units: 0x00, name: "Chilean Peso", countries: []string{"CHILE"}}
	cny = currency{alphabetic: "CNY", numeric: "156", units: 0x02, name: "Yuan Renminbi", countries: []string{"CHINA"}}
	cop = currency{alphabetic: "COP", numeric: "170", units: 0x02, name: "Colombian Peso", countries: []string{"COLOMBIA"}}
	cou = currency{alphabetic: "COU", numeric: "970", units: 0x02, name: "Unidad de Valor Real", countries: []string{"COLOMBIA"}}
	crc = currency{alphabetic: "CRC", numeric: "188", units: 0x02, name: "Costa Rican Colon", countries: []string{"COSTA RICA"}}
	cuc = currency{alphabetic: "CUC", numeric: "931", units: 0x02, name: "Peso Convertible", countries: []string{"CUBA"}}
	cup = currency{alphabetic: "CUP", numeric: "192", units: 0x02, name: "Cuban Peso", countries: []string{"CUBA"}}
	cve = currency{alphabetic: "CVE", numeric: "132", units: 0x02, name: "Cabo Verde Escudo", countries: []string{"CABO VERDE"}}
	czk = currency{alphabetic: "CZK", numeric: "203", units: 0x02, name: "Czech Koruna", countries: []string{"CZECHIA"}}
	djf = currency{alphabetic: "DJF", numeric: "262", units: 0x00, name: "Djibouti Franc", countries: []string{"DJIBOUTI"}}
	dkk = currency{alphabetic: "DKK", numeric: "208", units: 0x02, name: "Danish Krone", countries: []string{"DENMARK", "FAROE ISLANDS (THE)", "GREENLAND"}}
	dop = currency{alphabetic: "DOP", numeric: "214", units: 0x02, name: "Dominican Peso", countries: []string{"DOMINICAN REPUBLIC (THE)"}}
	dzd = currency{alphabetic: "DZD", numeric: "012", units: 0x02, name: "Algerian Dinar", countries: []string{"ALGERIA"}}
	egp = currency{alphabetic: "EGP", numeric: "818", units: 0x02, name: "Egyptian Pound", countries: []string{"EGYPT"}}
	ern = currency{alphabetic: "ERN", numeric: "232", units: 0x02, name: "Nakfa", countries: []string{"ERITREA"}}
	etb = currency{alphabetic: "ETB", numeric: "230", units: 0x02, name: "Ethiopian Birr", countries: []string{"ETHIOPIA"}}
	eur = currency{alphabetic: "EUR", numeric: "978", units: 0x02, name: "Euro", countries: []string{"ÅLAND ISLANDS", "ANDORRA", "AUSTRIA", "BELGIUM", "CROATIA", "CYPRUS", "ESTONIA", "EUROPEAN UNION", "FINLAND", "FRANCE", "FRENCH GUIANA", "FRENCH SOUTHERN TERRITORIES (THE)", "GERMANY", "GREECE", "GUADELOUPE", "HOLY SEE (THE)", "IRELAND", "ITALY", "LATVIA", "LITHUANIA", "LUXEMBOURG", "MALTA", "MARTINIQUE", "MAYOTTE", "MONACO", "MONTENEGRO", "NETHERLANDS (THE)", "PORTUGAL", "RÉUNION", "SAINT BARTHÉLEMY", "SAINT MARTIN (FRENCH PART)", "SAINT PIERRE AND MIQUELON", "SAN MARINO", "SLOVAKIA", "SLOVENIA", "SPAIN"}}
	fjd = currency{alphabetic: "FJD", numeric: "242", units: 0x02, name: "Fiji Dollar", countries: []string{"FIJI"}}
	fkp = currency{alphabetic: "FKP", numeric: "238", units: 0x02, name: "Falkland Islands Pound", countries: []string{"FALKLAND ISLANDS (THE) [MALVINAS]"}}
	gbp = currency{alphabetic: "GBP", numeric: "826", units: 0x02, name: "Pound Sterling", countries: []string{"GUERNSEY", "ISLE OF MAN", "JERSEY", "UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND (THE)"}}
	gel = currency{alphabetic: "GEL", numeric: "981", units: 0x02, name: "Lari", countries: []string{"GEORGIA"}}
	ghs = currency{alphabetic: "GHS", numeric: "936", units: 0x02, name: "Ghana Cedi", countries: []string{"GHANA"}}
	gip = currency{alphabetic: "GIP", numeric: "292", units: 0x02, name: "Gibraltar Pound", countries: []string{"GIBRALTAR"}}
	gmd = currency{alphabetic: "GMD", numeric: "270", units: 0x02, name: "Dalasi", countries: []string{"GAMBIA (THE)"}}
	gnf = currency{alphabetic: "GNF", numeric: "324", units: 0x00, name: "Guinean Franc", countries: []string{"GUINEA"}}
	gtq = currency{alphabetic: "GTQ", numeric: "320", units: 0x02, name: "Quetzal", countries: []string{"GUATEMALA"}}
	gyd = currency{alphabetic: "GYD", numeric: "328", units: 0x02, name: "Guyana Dollar", countries: []string{"GUYANA"}}
	hkd = currency{alphabetic: "HKD", numeric: "344", units: 0x02, name: "Hong Kong Dollar", countries: []string{"HONG KONG"}}
	hnl = currency{alphabetic: "HNL", numeric: "340", units: 0x02, name: "Lempira", countries: []string{"HONDURAS"}}
	htg = currency{alphabetic: "HTG", numeric: "332", units: 0x02, name: "Gourde", countries: []string{"HAITI"}}
	huf = currency{alphabetic: "HUF", numeric: "348", units: 0x02, name: "Forint", countries: []string{"HUNGARY"}}
	idr = currency{alphabetic: "IDR", numeric: "360", units: 0x02, name: "Rupiah", countries: []string{"INDONESIA"}}
	ils = currency{alphabetic: "ILS", numeric: "376", units: 0x02, name: "New Israeli Sheqel", countries: []string{"ISRAEL"}}
	inr = currency{alphabetic: "INR", numeric: "356", units: 0x02, name: "Indian Rupee", countries: []string{"BHUTAN", "INDIA"}}
	iqd = currency{alphabetic: "IQD", numeric: "368", units: 0x03, name: "Iraqi Dinar", countries: []string{"IRAQ"}}
	irr = currency{alphabetic: "IRR", numeric: "364", units: 0x02, name: "Iranian Rial", countries: []string{"IRAN (ISLAMIC REPUBLIC OF)"}}
	isk = currency{alphabetic: "ISK", numeric: "352", units: 0x00, name: "Iceland Krona", countries: []string{"ICELAND"}}
	jmd = currency{alphabetic: "JMD", numeric: "388", units: 0x02, name: "Jamaican Dollar", countries: []string{"JAMAICA"}}
	jod = currency{alphabetic: "JOD", numeric: "400", units: 0x03, name: "Jordanian Dinar", countries: []string{"JORDAN"}}
	jpy = currency{alphabetic: "JPY", numeric: "392", units: 0x00, name: "Yen", countries: []string{"JAPAN"}}
	kes = currency{alphabetic: "KES", numeric: "404", units: 0x02, name: "Kenyan Shilling", countries: []string{"KENYA"}}
	kgs = currency{alphabetic: "KGS", numeric: "417", units: 0x02, name: "Som", countries: []string{"KYRGYZSTAN"}}
	khr = currency{alphabetic: "KHR", numeric: "116", units: 0x02, name: "Riel", countries: []string{"CAMBODIA"}}
	kmf = currency{alphabetic: "KMF", numeric: "174", units: 0x00, name: "Comorian Franc", countries: []string{"COMOROS (THE)"}}
	kpw = currency{alphabetic: "KPW", numeric: "408", units: 0x02, name: "North Korean Won", countries: []string{"KOREA (THE DEMOCRATIC PEOPLE’S REPUBLIC OF)"}}
	krw = currency{alphabetic: "KRW", numeric: "410", units: 0x00, name: "Won", countries: []string{"KOREA (THE REPUBLIC OF)"}}
	kwd = currency{alphabetic: "KWD", numeric: "414", units: 0x03, name: "Kuwaiti Dinar", countries: []string{"KUWAIT"}}
	kyd = currency{alphabetic: "KYD", numeric: "136", units: 0x02, name: "Cayman Islands Dollar", countries: []string{"CAYMAN ISLANDS (THE)"}}
	kzt = currency{alphabetic: "KZT", numeric: "398", units: 0x02, name: "Tenge", countries: []string{"KAZAKHSTAN"}}
	lak = currency{alphabetic: "LAK", numeric: "418", units: 0x02, name: "Lao Kip", countries: []string{"LAO PEOPLE’S DEMOCRATIC REPUBLIC (THE)"}}
	lbp = currency{alphabetic: "LBP", numeric: "422", units: 0x02, name: "Lebanese Pound", countries: []string{"LEBANON"}}
	lkr = currency{alphabetic: "LKR", numeric: "144", units: 0x02, name: "Sri Lanka Rupee", countries: []string{"SRI LANKA"}}
	lrd = currency{alphabetic: "LRD", numeric: "430", units: 0x02, name: "Liberian Dollar", countries: []string{"LIBERIA"}}
	lsl = currency{alphabetic: "LSL", numeric: "426", units: 0x02, name: "Loti", countries: []string{"LESOTHO"}}
	lyd = currency{alphabetic: "LYD", numeric: "434", units: 0x03, name: "Libyan Dinar", countries: []string{"LIBYA"}}
	mad = currency{alphabetic: "MAD", numeric: "504", units: 0x02, name: "Moroccan Dirham", countries: []string{"MOROCCO", "WESTERN SAHARA"}}
	mdl = currency{alphabetic: "MDL", numeric: "498", units: 0x02, name: "Moldovan Leu", countries: []string{"MOLDOVA (THE REPUBLIC OF)"}}
	mga = currency{alphabetic: "MGA", numeric: "969", units: 0x02, name: "Malagasy Ariary", countries: []string{"MADAGASCAR"}}
	mkd = currency{alphabetic: "MKD", numeric: "807", units: 0x02, name: "Denar", countries: []string{"NORTH MACEDONIA"}}
	mmk = currency{alphabetic: "MMK", numeric: "104", units: 0x02, name: "Kyat", countries: []string{"MYANMAR"}}
	mnt = currency{alphabetic: "MNT", numeric: "496", units: 0x02, name: "Tugrik", countries: []string{"MONGOLIA"}}
	mop = currency{alphabetic: "MOP", numeric: "446", units: 0x02, name: "Pataca", countries: []string{"MACAO"}}
	mru = currency{alphabetic: "MRU", numeric: "929", units: 0x02, name: "Ouguiya", countries: []string{"MAURITANIA"}}
	mur = currency{alphabetic: "MUR", numeric: "480", units: 0x02, name: "Mauritius Rupee", countries: []string{"MAURITIUS"}}
	mvr = currency{alphabetic: "MVR", numeric: "462", units: 0x02, name: "Rufiyaa", countries: []string{"MALDIVES"}}
	mwk = currency{alphabetic: "MWK", numeric: "454", units: 0x02, name: "Malawi Kwacha", countries: []string{"MALAWI"}}
	mxn = currency{alphabetic: "MXN", numeric: "484", units: 0x02, name: "Mexican Peso", countries: []string{"MEXICO"}}
	mxv = currency{alphabetic: "MXV", numeric: "979", units: 0x02, name: "Mexican Unidad de Inversion (UDI)", countries: []string{"MEXICO"}}
	myr = currency{alphabetic: "MYR", numeric: "458", units: 0x02, name: "Malaysian Ringgit", countries: []string{"MALAYSIA"}}
	mzn = currency{alphabetic: "MZN", numeric: "943", units: 0x02, name: "Mozambique Metical", countries: []string{"MOZAMBIQUE"}}
	nad = currency{alphabetic: "NAD", numeric: "516", units: 0x02, name: "Namibia Dollar", countries: []string{"NAMIBIA"}}
	ngn = currency{alphabetic: "NGN", numeric: "566", units: 0x02, name: "Naira", countries: []string{"NIGERIA"}}
	nio = currency{alphabetic: "NIO", numeric: "558", units: 0x02, name: "Cordoba Oro", countries: []string{"NICARAGUA"}}
	nok = currency{alphabetic: "NOK", numeric: "578", units: 0x02, name: "Norwegian Krone", countries: []string{"BOUVET ISLAND", "NORWAY", "SVALBARD AND JAN MAYEN"}}
	npr = currency{alphabetic: "NPR", numeric: "524", units: 0x02, name: "Nepalese Rupee", countries: []string{"NEPAL"}}
	nzd = currency{alphabetic: "NZD", numeric: "554", units: 0x02, name: "New Zealand Dollar", countries: []string{"COOK ISLANDS (THE)", "NEW ZEALAND", "NIUE", "PITCAIRN", "TOKELAU"}}
	omr = currency{alphabetic: "OMR", numeric: "512", units: 0x03, name: "Rial Omani", countries: []string{"OMAN"}}
	pab = currency{alphabetic: "PAB", numeric: "590", units: 0x02, name: "Balboa", countries: []string{"PANAMA"}}
	pen = currency{alphabetic: "PEN", numeric: "604", units: 0x02, name: "Sol", countries: []string{"PERU"}}
	pgk = currency{alphabetic: "PGK", numeric: "598", units: 0x02, name: "Kina", countries: []string{"PAPUA NEW GUINEA"}}
	php = currency{alphabetic: "PHP", numeric: "608", units: 0x02, name: "Philippine Peso", countries: []string{"PHILIPPINES (THE)"}}
	pkr = currency{alphabetic: "PKR", numeric: "586", units: 0x02, name: "Pakistan Rupee", countries: []string{"PAKISTAN"}}
	pln = currency{alphabetic: "PLN", numeric: "985", units: 0x02, name: "Zloty", countries: []string{"POLAND"}}
	pyg = currency{alphabetic: "PYG", numeric: "600", units: 0x00, name: "Guarani", countries: []string{"PARAGUAY"}}
	qar = currency{alphabetic: "QAR", numeric: "634", units: 0x02, name: "Qatari Rial", countries: []string{"QATAR"}}
	ron = currency{alphabetic: "RON", numeric: "946", units: 0x02, name: "Romanian Leu", countries: []string{"ROMANIA"}}
	rsd = currency{alphabetic: "RSD", numeric: "941", units: 0x02, name: "Serbian Dinar", countries: []string{"SERBIA"}}
	rub = currency{alphabetic: "RUB", numeric: "643", units: 0x02, name: "Russian Ruble", countries: []string{"RUSSIAN FEDERATION (THE)"}}
	rwf = currency{alphabetic: "RWF", numeric: "646", units: 0x00, name: "Rwanda Franc", countries: []string{"RWANDA"}}
	sar = currency{alphabetic: "SAR", numeric: "682", units: 0x02, name: "Saudi Riyal", countries: []string{"SAUDI ARABIA"}}
	sbd = currency{alphabetic: "SBD", numeric: "090", units: 0x02, name: "Solomon Islands Dollar", countries: []string{"SOLOMON ISLANDS"}}
	scr = currency{alphabetic: "SCR", numeric: "690", units: 0x02, name: "Seychelles Rupee", countries: []string{"SEYCHELLES"}}
	sdg = currency{alphabetic: "SDG", numeric: "938", units: 0x02, name: "Sudanese Pound", countries: []string{"SUDAN (THE)"}}
	sek = currency{alphabetic: "SEK", numeric: "752", units: 0x02, name: "Swedish Krona", countries: []string{"SWEDEN"}}
	sgd = currency{alphabetic: "SGD", numeric: "702", units: 0x02, name: "Singapore Dollar", countries: []string{"SINGAPORE"}}
	shp = currency{alphabetic: "SHP", numeric: "654", units: 0x02, name: "Saint Helena Pound", countries: []string{"SAINT HELENA, ASCENSION AND TRISTAN DA CUNHA"}}
	sle = currency{alphabetic: "SLE", numeric: "925", units: 0x02, name: "Leone", countries: []string{"SIERRA LEONE"}}
	sos = currency{alphabetic: "SOS", numeric: "706", units: 0x02, name: "Somali Shilling", countries: []string{"SOMALIA"}}
	srd = currency{alphabetic: "SRD", numeric: "968", units: 0x02, name: "Surinam Dollar", countries: []string{"SURINAME"}}
	ssp = currency{alphabetic: "SSP", numeric: "728", units: 0x02, name: "South Sudanese Pound", countries: []string{"SOUTH SUDAN"}}
	stn = currency{alphabetic: "STN", numeric: "930", units: 0x02, name: "Dobra", countries: []string{"SAO TOME AND PRINCIPE"}}
	svc = currency{alphabetic: "SVC", numeric: "222", units: 0x02, name: "El Salvador Colon", countries: []string{"EL SALVADOR"}}
	syp = currency{alphabetic: "SYP", numeric: "760", units: 0x02, name: "Syrian Pound", countries: []string{"SYRIAN ARAB REPUBLIC"}}
	szl = currency{alphabetic: "SZL", numeric: "748", units: 0x02, name: "Lilangeni", countries: []string{"ESWATINI"}}
	thb = currency{alphabetic: "THB", numeric: "764", units: 0x02, name: "Baht", countries: []string{"THAILAND"}}
	tjs = currency{alphabetic: "TJS", numeric: "972", units: 0x02, name: "Somoni", countries: []string{"TAJIKISTAN"}}
	tmt = currency{alphabetic: "TMT", numeric: "934", units: 0x02, name: "Turkmenistan New Manat", countries: []string{"TURKMENISTAN"}}
	tnd = currency{alphabetic: "TND", numeric: "788", units: 0x03, name: "Tunisian Dinar", countries: []string{"TUNISIA"}}
	top = currency{alphabetic: "TOP", numeric: "776", units: 0x02, name: "Pa’anga", countries: []string{"TONGA"}}
	try = currency{alphabetic: "TRY", numeric: "949", units: 0x02, name: "Turkish Lira", countries: []string{"TÜRKİYE"}}
	ttd = currency{alphabetic: "TTD", numeric: "780", units: 0x02, name: "Trinidad and Tobago Dollar", countries: []string{"TRINIDAD AND TOBAGO"}}
	twd = currency{alphabetic: "TWD", numeric: "901", units: 0x02, name: "New Taiwan Dollar", countries: []string{"TAIWAN (PROVINCE OF CHINA)"}}
	tzs = currency{alphabetic: "TZS", numeric: "834", units: 0x02, name: "Tanzanian Shilling", countries: []string{"TANZANIA, UNITED REPUBLIC OF"}}
	uah = currency{alphabetic: "UAH", numeric: "980", units: 0x02, name: "Hryvnia", countries: []string{"UKRAINE"}}
	ugx = currency{alphabetic: "UGX", numeric: "800", units: 0x00, name: "Uganda Shilling", countries: []string{"UGANDA"}}
	usd = currency{alphabetic: "USD", numeric: "840", units: 0x02, name: "US Dollar", countries: []string{"AMERICAN SAMOA", "BONAIRE, SINT EUSTATIUS AND SABA", "BRITISH INDIAN OCEAN TERRITORY (THE)", "ECUADOR", "EL SALVADOR", "GUAM", "HAITI", "MARSHALL ISLANDS (THE)", "MICRONESIA (FEDERATED STATES OF)", "NORTHERN MARIANA ISLANDS (THE)", "PALAU", "PANAMA", "PUERTO RICO", "TIMOR-LESTE", "TURKS AND CAICOS ISLANDS (THE)", "UNITED STATES MINOR OUTLYING ISLANDS (THE)", "UNITED STATES OF AMERICA (THE)", "VIRGIN ISLANDS (BRITISH)", "VIRGIN ISLANDS (U.S.)"}}
	usn = currency{alphabetic: "USN", numeric: "997", units: 0x02, name: "US Dollar (Next day)", countries: []string{"UNITED STATES OF AMERICA (THE)"}}
	uyi = currency{alphabetic: "UYI", numeric: "940", units: 0x00, name: "Uruguay Peso en Unidades Indexadas (UI)", countries: []string{"URUGUAY"}}
	uyu = currency{alphabetic: "UYU", numeric: "858", units: 0x02, name: "Peso Uruguayo", countries: []string{"URUGUAY"}}
	uyw = currency{alphabetic: "UYW", numeric: "927", units: 0x04, name: "Unidad Previsional", countries: []string{"URUGUAY"}}
	uzs = currency{alphabetic: "UZS", numeric: "860", units: 0x02, name: "Uzbekistan Sum", countries: []string{"UZBEKISTAN"}}
	ved = currency{alphabetic: "VED", numeric: "926", units: 0x02, name: "Bolívar Soberano", countries: []string{"VENEZUELA (BOLIVARIAN REPUBLIC OF)"}}
	ves = currency{alphabetic: "VES", numeric: "928", units: 0x02, name: "Bolívar Soberano", countries: []string{"VENEZUELA (BOLIVARIAN REPUBLIC OF)"}}
	vnd = currency{alphabetic: "VND", numeric: "704", units: 0x00, name: "Dong", countries: []string{"VIET NAM"}}
	vuv = currency{alphabetic: "VUV", numeric: "548", units: 0x00, name: "Vatu", countries: []string{"VANUATU"}}
	wst = currency{alphabetic: "WST", numeric: "882", units: 0x02, name: "Tala", countries: []string{"SAMOA"}}
	xaf = currency{alphabetic: "XAF", numeric: "950", units: 0x00, name: "CFA Franc BEAC", countries: []string{"CAMEROON", "CENTRAL AFRICAN REPUBLIC (THE)", "CHAD", "CONGO (THE)", "EQUATORIAL GUINEA", "GABON"}}
	xag = currency{alphabetic: "XAG", numeric: "961", units: 0xff, name: "Silver", countries: []string{"ZZ11_Silver"}}
	xau = currency{alphabetic: "XAU", numeric: "959", units: 0xff, name: "Gold", countries: []string{"ZZ08_Gold"}}
	xba = currency{alphabetic: "XBA", numeric: "955", units: 0xff, name: "Bond Markets Unit European Composite Unit (EURCO)", countries: []string{"ZZ01_Bond Markets Unit European_EURCO"}}
	xbb = currency{alphabetic: "XBB", numeric: "956", units: 0xff, name: "Bond Markets Unit European Monetary Unit (E.M.U.-6)", countries: []string{"ZZ02_Bond Markets Unit European_EMU-6"}}
	xbc = currency{alphabetic: "XBC", numeric: "957", units: 0xff, name: "Bond Markets Unit European Unit of Account 9 (E.U.A.-9)", countries: []string{"ZZ03_Bond Markets Unit European_EUA-9"}}
	xbd = currency{alphabetic: "XBD", numeric: "958", units: 0xff, name: "Bond Markets Unit European Unit of Account 17 (E.U.A.-17)", countries: []string{"ZZ04_Bond Markets Unit European_EUA-17"}}
	xcd = currency{alphabetic: "XCD", numeric: "951", units: 0x02, name: "East Caribbean Dollar", countries: []string{"ANGUILLA", "ANTIGUA AND BARBUDA", "DOMINICA", "GRENADA", "MONTSERRAT", "SAINT KITTS AND NEVIS", "SAINT LUCIA", "SAINT VINCENT AND THE GRENADINES"}}
	xdr = currency{alphabetic: "XDR", numeric: "960", units: 0xff, name: "SDR (Special Drawing Right)", countries: []string{"INTERNATIONAL MONETARY FUND (IMF)"}}
	xof = currency{alphabetic: "XOF", numeric: "952", units: 0x00, name: "CFA Franc BCEAO", countries: []string{"BENIN", "BURKINA FASO", "CÔTE D'IVOIRE", "GUINEA-BISSAU", "MALI", "NIGER (THE)", "SENEGAL", "TOGO"}}
	xpd = currency{alphabetic: "XPD", numeric: "964", units: 0xff, name: "Palladium", countries: []string{"ZZ09_Palladium"}}
	xpf = currency{alphabetic: "XPF", numeric: "953", units: 0x00, name: "CFP Franc", countries: []string{"FRENCH POLYNESIA", "NEW CALEDONIA", "WALLIS AND FUTUNA"}}
	xpt = currency{alphabetic: "XPT", numeric: "962", units: 0xff, name: "Platinum", countries: []string{"ZZ10_Platinum"}}
	xsu = currency{alphabetic: "XSU", numeric: "994", units: 0xff, name: "Sucre", countries: []string{"SISTEMA UNITARIO DE COMPENSACION REGIONAL DE PAGOS \"SUCRE\""}}
	xts = currency{alphabetic: "XTS", numeric: "963", units: 0xff, name: "Codes specifically reserved for testing purposes", countries: []string{"ZZ06_Testing_Code"}}
	xua = currency{alphabetic: "XUA", numeric: "965", units: 0xff, name: "ADB Unit of Account", countries: []string{"MEMBER COUNTRIES OF THE AFRICAN DEVELOPMENT BANK GROUP"}}
	xxx = currency{alphabetic: "XXX", numeric: "999", units: 0xff, name: "The codes assigned for transactions where no currency is involved", countries: []string{"ZZ07_No_Currency"}}
	yer = currency{alphabetic: "YER", numeric: "886", units: 0x02, name: "Yemeni Rial", countries: []string{"YEMEN"}}
	zar = currency{alphabetic: "ZAR", numeric: "710", units: 0x02, name: "Rand", countries: []string{"LESOTHO", "NAMIBIA", "SOUTH AFRICA"}}
	zmw = currency{alphabetic: "ZMW", numeric: "967", units: 0x02, name: "Zambian Kwacha", countries: []string{"ZAMBIA"}}
	zwg = currency{alphabetic: "ZWG", numeric: "924", units: 0x02, name: "Zimbabwe Gold", countries: []string{"ZIMBABWE"}}
)

var currencies = map[Currency]currency{
	AED: aed, _784: aed, AFN: afn, _971: afn, ALL: all, _008: all, AMD: amd, _051: amd, ANG: ang, _532: ang,
	AOA: aoa, _973: aoa, ARS: ars, _032: ars, AUD: aud, _036: aud, AWG: awg, _533: awg, AZN: azn, _944: azn,
	BAM: bam, _977: bam, BBD: bbd, _052: bbd, BDT: bdt, _050: bdt, BGN: bgn, _975: bgn, BHD: bhd, _048: bhd,
	BIF: bif, _108: bif, BMD: bmd, _060: bmd, BND: bnd, _096: bnd, BOB: bob, _068: bob, BOV: bov, _984: bov,
	BRL: brl, _986: brl, BSD: bsd, _044: bsd, BTN: btn, _064: btn, BWP: bwp, _072: bwp, BYN: byn, _933: byn,
	BZD: bzd, _084: bzd, CAD: cad, _124: cad, CDF: cdf, _976: cdf, CHE: che, _947: che, CHF: chf, _756: chf,
	CHW: chw, _948: chw, CLF: clf, _990: clf, CLP: clp, _152: clp, CNY: cny, _156: cny, COP: cop, _170: cop,
	COU: cou, _970: cou, CRC: crc, _188: crc, CUC: cuc, _931: cuc, CUP: cup, _192: cup, CVE: cve, _132: cve,
	CZK: czk, _203: czk, DJF: djf, _262: djf, DKK: dkk, _208: dkk, DOP: dop, _214: dop, DZD: dzd, _012: dzd,
	EGP: egp, _818: egp, ERN: ern, _232: ern, ETB: etb, _230: etb, EUR: eur, _978: eur, FJD: fjd, _242: fjd,
	FKP: fkp, _238: fkp, GBP: gbp, _826: gbp, GEL: gel, _981: gel, GHS: ghs, _936: ghs, GIP: gip, _292: gip,
	GMD: gmd, _270: gmd, GNF: gnf, _324: gnf, GTQ: gtq, _320: gtq, GYD: gyd, _328: gyd, HKD: hkd, _344: hkd,
	HNL: hnl, _340: hnl, HTG: htg, _332: htg, HUF: huf, _348: huf, IDR: idr, _360: idr, ILS: ils, _376: ils,
	INR: inr, _356: inr, IQD: iqd, _368: iqd, IRR: irr, _364: irr, ISK: isk, _352: isk, JMD: jmd, _388: jmd,
	JOD: jod, _400: jod, JPY: jpy, _392: jpy, KES: kes, _404: kes, KGS: kgs, _417: kgs, KHR: khr, _116: khr,
	KMF: kmf, _174: kmf, KPW: kpw, _408: kpw, KRW: krw, _410: krw, KWD: kwd, _414: kwd, KYD: kyd, _136: kyd,
	KZT: kzt, _398: kzt, LAK: lak, _418: lak, LBP: lbp, _422: lbp, LKR: lkr, _144: lkr, LRD: lrd, _430: lrd,
	LSL: lsl, _426: lsl, LYD: lyd, _434: lyd, MAD: mad, _504: mad, MDL: mdl, _498: mdl, MGA: mga, _969: mga,
	MKD: mkd, _807: mkd, MMK: mmk, _104: mmk, MNT: mnt, _496: mnt, MOP: mop, _446: mop, MRU: mru, _929: mru,
	MUR: mur, _480: mur, MVR: mvr, _462: mvr, MWK: mwk, _454: mwk, MXN: mxn, _484: mxn, MXV: mxv, _979: mxv,
	MYR: myr, _458: myr, MZN: mzn, _943: mzn, NAD: nad, _516: nad, NGN: ngn, _566: ngn, NIO: nio, _558: nio,
	NOK: nok, _578: nok, NPR: npr, _524: npr, NZD: nzd, _554: nzd, OMR: omr, _512: omr, PAB: pab, _590: pab,
	PEN: pen, _604: pen, PGK: pgk, _598: pgk, PHP: php, _608: php, PKR: pkr, _586: pkr, PLN: pln, _985: pln,
	PYG: pyg, _600: pyg, QAR: qar, _634: qar, RON: ron, _946: ron, RSD: rsd, _941: rsd, RUB: rub, _643: rub,
	RWF: rwf, _646: rwf, SAR: sar, _682: sar, SBD: sbd, _090: sbd, SCR: scr, _690: scr, SDG: sdg, _938: sdg,
	SEK: sek, _752: sek, SGD: sgd, _702: sgd, SHP: shp, _654: shp, SLE: sle, _925: sle, SOS: sos, _706: sos,
	SRD: srd, _968: srd, SSP: ssp, _728: ssp, STN: stn, _930: stn, SVC: svc, _222: svc, SYP: syp, _760: syp,
	SZL: szl, _748: szl, THB: thb, _764: thb, TJS: tjs, _972: tjs, TMT: tmt, _934: tmt, TND: tnd, _788: tnd,
	TOP: top, _776: top, TRY: try, _949: try, TTD: ttd, _780: ttd, TWD: twd, _901: twd, TZS: tzs, _834: tzs,
	UAH: uah, _980: uah, UGX: ugx, _800: ugx, USD: usd, _840: usd, USN: usn, _997: usn, UYI: uyi, _940: uyi,
	UYU: uyu, _858: uyu, UYW: uyw, _927: uyw, UZS: uzs, _860: uzs, VED: ved, _926: ved, VES: ves, _928: ves,
	VND: vnd, _704: vnd, VUV: vuv, _548: vuv, WST: wst, _882: wst, XAF: xaf, _950: xaf, XAG: xag, _961: xag,
	XAU: xau, _959: xau, XBA: xba, _955: xba, XBB: xbb, _956: xbb, XBC: xbc, _957: xbc, XBD: xbd, _958: xbd,
	XCD: xcd, _951: xcd, XDR: xdr, _960: xdr, XOF: xof, _952: xof, XPD: xpd, _964: xpd, XPF: xpf, _953: xpf,
	XPT: xpt, _962: xpt, XSU: xsu, _994: xsu, XTS: xts, _963: xts, XUA: xua, _965: xua, XXX: xxx, _999: xxx,
	YER: yer, _886: yer, ZAR: zar, _710: zar, ZMW: zmw, _967: zmw, ZWG: zwg, _924: zwg,
}
