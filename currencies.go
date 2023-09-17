// Copyright S&K Software Development Ltd.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package currencies implements ISO 4217 database of currencies allowing to resolve
// currency name and precision by their numeric or 3-letter codes.
package currencies

// The currency information
type Currency struct {
	code          string
	number        int
	decimalPlaces int
	scale         int
	description   string
}

var (
	currencyAllCurrencies = []Currency{
		{code: "AED", number: 784, decimalPlaces: 2, description: "United Arab Emirates dirham"},
		{code: "AFN", number: 971, decimalPlaces: 2, description: "Afghan afghani"},
		{code: "ALL", number: 8, decimalPlaces: 2, description: "Albanian lek"},
		{code: "AMD", number: 51, decimalPlaces: 2, description: "Armenian dram"},
		{code: "ANG", number: 532, decimalPlaces: 2, description: "Netherlands Antillean guilder"},
		{code: "AOA", number: 973, decimalPlaces: 2, description: "Angolan kwanza"},
		{code: "ARS", number: 32, decimalPlaces: 2, description: "Argentine peso"},
		{code: "AUD", number: 36, decimalPlaces: 2, description: "Australian dollar"},
		{code: "AWG", number: 533, decimalPlaces: 2, description: "Aruban florin"},
		{code: "AZN", number: 944, decimalPlaces: 2, description: "Azerbaijani manat"},
		{code: "BAM", number: 977, decimalPlaces: 2, description: "Bosnia and Herzegovina convertible mark"},
		{code: "BBD", number: 52, decimalPlaces: 2, description: "Barbados dollar"},
		{code: "BDT", number: 50, decimalPlaces: 2, description: "Bangladeshi taka"},
		{code: "BGN", number: 975, decimalPlaces: 2, description: "Bulgarian lev"},
		{code: "BHD", number: 48, decimalPlaces: 3, description: "Bahraini dinar"},
		{code: "BIF", number: 108, decimalPlaces: 0, description: "Burundian franc"},
		{code: "BMD", number: 60, decimalPlaces: 2, description: "Bermudian dollar"},
		{code: "BND", number: 96, decimalPlaces: 2, description: "Brunei dollar"},
		{code: "BOB", number: 68, decimalPlaces: 2, description: "Boliviano"},
		{code: "BOV", number: 984, decimalPlaces: 2, description: "Bolivian Mvdol (funds code)"},
		{code: "BRL", number: 986, decimalPlaces: 2, description: "Brazilian real"},
		{code: "BSD", number: 44, decimalPlaces: 2, description: "Bahamian dollar"},
		{code: "BTN", number: 64, decimalPlaces: 2, description: "Bhutanese ngultrum"},
		{code: "BWP", number: 72, decimalPlaces: 2, description: "Botswana pula"},
		{code: "BYR", number: 974, decimalPlaces: 0, description: "Belarusian ruble"},
		{code: "BZD", number: 84, decimalPlaces: 2, description: "Belize dollar"},
		{code: "CAD", number: 124, decimalPlaces: 2, description: "Canadian dollar"},
		{code: "CDF", number: 976, decimalPlaces: 2, description: "Congolese franc"},
		{code: "CHE", number: 947, decimalPlaces: 2, description: "WIR Euro (complementary currency)"},
		{code: "CHF", number: 756, decimalPlaces: 2, description: "Swiss franc"},
		{code: "CHW", number: 948, decimalPlaces: 2, description: "WIR Franc (complementary currency)"},
		{code: "CLF", number: 990, decimalPlaces: 4, description: "Unidad de Fomento (funds code)"},
		{code: "CLP", number: 152, decimalPlaces: 0, description: "Chilean peso"},
		{code: "CNY", number: 156, decimalPlaces: 2, description: "Chinese yuan"},
		{code: "COP", number: 170, decimalPlaces: 2, description: "Colombian peso"},
		{code: "COU", number: 970, decimalPlaces: 2, description: "Unidad de Valor Real (UVR) (funds code)[7]"},
		{code: "CRC", number: 188, decimalPlaces: 2, description: "Costa Rican colon"},
		{code: "CUC", number: 931, decimalPlaces: 2, description: "Cuban convertible peso"},
		{code: "CUP", number: 192, decimalPlaces: 2, description: "Cuban peso"},
		{code: "CVE", number: 132, decimalPlaces: 0, description: "Cape Verde escudo"},
		{code: "CZK", number: 203, decimalPlaces: 2, description: "Czech koruna"},
		{code: "DJF", number: 262, decimalPlaces: 0, description: "Djiboutian franc"},
		{code: "DKK", number: 208, decimalPlaces: 2, description: "Danish krone"},
		{code: "DOP", number: 214, decimalPlaces: 2, description: "Dominican peso"},
		{code: "DZD", number: 12, decimalPlaces: 2, description: "Algerian dinar"},
		{code: "EGP", number: 818, decimalPlaces: 2, description: "Egyptian pound"},
		{code: "ERN", number: 232, decimalPlaces: 2, description: "Eritrean nakfa"},
		{code: "ETB", number: 230, decimalPlaces: 2, description: "Ethiopian birr"},
		{code: "EUR", number: 978, decimalPlaces: 2, description: "Euro"},
		{code: "FJD", number: 242, decimalPlaces: 2, description: "Fiji dollar"},
		{code: "FKP", number: 238, decimalPlaces: 2, description: "Falkland Islands pound"},
		{code: "GBP", number: 826, decimalPlaces: 2, description: "Pound sterling"},
		{code: "GEL", number: 981, decimalPlaces: 2, description: "Georgian lari"},
		{code: "GHS", number: 936, decimalPlaces: 2, description: "Ghanaian cedi"},
		{code: "GIP", number: 292, decimalPlaces: 2, description: "Gibraltar pound"},
		{code: "GMD", number: 270, decimalPlaces: 2, description: "Gambian dalasi"},
		{code: "GNF", number: 324, decimalPlaces: 0, description: "Guinean franc"},
		{code: "GTQ", number: 320, decimalPlaces: 2, description: "Guatemalan quetzal"},
		{code: "GYD", number: 328, decimalPlaces: 2, description: "Guyanese dollar"},
		{code: "HKD", number: 344, decimalPlaces: 2, description: "Hong Kong dollar"},
		{code: "HNL", number: 340, decimalPlaces: 2, description: "Honduran lempira"},
		{code: "HRK", number: 191, decimalPlaces: 2, description: "Croatian kuna"},
		{code: "HTG", number: 332, decimalPlaces: 2, description: "Haitian gourde"},
		{code: "HUF", number: 348, decimalPlaces: 2, description: "Hungarian forint"},
		{code: "IDR", number: 360, decimalPlaces: 2, description: "Indonesian rupiah"},
		{code: "ILS", number: 376, decimalPlaces: 2, description: "Israeli new shekel"},
		{code: "INR", number: 356, decimalPlaces: 2, description: "Indian rupee"},
		{code: "IQD", number: 368, decimalPlaces: 3, description: "Iraqi dinar"},
		{code: "IRR", number: 364, decimalPlaces: 2, description: "Iranian rial"},
		{code: "ISK", number: 352, decimalPlaces: 0, description: "Icelandic króna"},
		{code: "JMD", number: 388, decimalPlaces: 2, description: "Jamaican dollar"},
		{code: "JOD", number: 400, decimalPlaces: 3, description: "Jordanian dinar"},
		{code: "JPY", number: 392, decimalPlaces: 0, description: "Japanese yen"},
		{code: "KES", number: 404, decimalPlaces: 2, description: "Kenyan shilling"},
		{code: "KGS", number: 417, decimalPlaces: 2, description: "Kyrgyzstani som"},
		{code: "KHR", number: 116, decimalPlaces: 2, description: "Cambodian riel"},
		{code: "KMF", number: 174, decimalPlaces: 0, description: "Comoro franc"},
		{code: "KPW", number: 408, decimalPlaces: 2, description: "North Korean won"},
		{code: "KRW", number: 410, decimalPlaces: 0, description: "South Korean won"},
		{code: "KWD", number: 414, decimalPlaces: 3, description: "Kuwaiti dinar"},
		{code: "KYD", number: 136, decimalPlaces: 2, description: "Cayman Islands dollar"},
		{code: "KZT", number: 398, decimalPlaces: 2, description: "Kazakhstani tenge"},
		{code: "LAK", number: 418, decimalPlaces: 2, description: "Lao kip"},
		{code: "LBP", number: 422, decimalPlaces: 2, description: "Lebanese pound"},
		{code: "LKR", number: 144, decimalPlaces: 2, description: "Sri Lankan rupee"},
		{code: "LRD", number: 430, decimalPlaces: 2, description: "Liberian dollar"},
		{code: "LSL", number: 426, decimalPlaces: 2, description: "Lesotho loti"},
		{code: "LYD", number: 434, decimalPlaces: 3, description: "Libyan dinar"},
		{code: "MAD", number: 504, decimalPlaces: 2, description: "Moroccan dirham"},
		{code: "MDL", number: 498, decimalPlaces: 2, description: "Moldovan leu"},
		{code: "MGA", number: 969, decimalPlaces: 1, description: "Malagasy ariary"},
		{code: "MKD", number: 807, decimalPlaces: 2, description: "Macedonian denar"},
		{code: "MMK", number: 104, decimalPlaces: 2, description: "Myanmar kyat"},
		{code: "MNT", number: 496, decimalPlaces: 2, description: "Mongolian tugrik"},
		{code: "MOP", number: 446, decimalPlaces: 2, description: "Macanese pataca"},
		{code: "MRO", number: 478, decimalPlaces: 1, description: "Mauritanian ouguiya"},
		{code: "MUR", number: 480, decimalPlaces: 2, description: "Mauritian rupee"},
		{code: "MVR", number: 462, decimalPlaces: 2, description: "Maldivian rufiyaa"},
		{code: "MWK", number: 454, decimalPlaces: 2, description: "Malawian kwacha"},
		{code: "MXN", number: 484, decimalPlaces: 2, description: "Mexican peso"},
		{code: "MXV", number: 979, decimalPlaces: 2, description: "Mexican Unidad de Inversion (UDI) (funds code)"},
		{code: "MYR", number: 458, decimalPlaces: 2, description: "Malaysian ringgit"},
		{code: "MZN", number: 943, decimalPlaces: 2, description: "Mozambican metical"},
		{code: "NAD", number: 516, decimalPlaces: 2, description: "Namibian dollar"},
		{code: "NGN", number: 566, decimalPlaces: 2, description: "Nigerian naira"},
		{code: "NIO", number: 558, decimalPlaces: 2, description: "Nicaraguan córdoba"},
		{code: "NOK", number: 578, decimalPlaces: 2, description: "Norwegian krone"},
		{code: "NPR", number: 524, decimalPlaces: 2, description: "Nepalese rupee"},
		{code: "NZD", number: 554, decimalPlaces: 2, description: "New Zealand dollar"},
		{code: "OMR", number: 512, decimalPlaces: 3, description: "Omani rial"},
		{code: "PAB", number: 590, decimalPlaces: 2, description: "Panamanian balboa"},
		{code: "PEN", number: 604, decimalPlaces: 2, description: "Peruvian nuevo sol"},
		{code: "PGK", number: 598, decimalPlaces: 2, description: "Papua New Guinean kina"},
		{code: "PHP", number: 608, decimalPlaces: 2, description: "Philippine peso"},
		{code: "PKR", number: 586, decimalPlaces: 2, description: "Pakistani rupee"},
		{code: "PLN", number: 985, decimalPlaces: 2, description: "Polish złoty"},
		{code: "PYG", number: 600, decimalPlaces: 0, description: "Paraguayan guaraní"},
		{code: "QAR", number: 634, decimalPlaces: 2, description: "Qatari riyal"},
		{code: "RON", number: 946, decimalPlaces: 2, description: "Romanian leu"},
		{code: "RSD", number: 941, decimalPlaces: 2, description: "Serbian dinar"},
		{code: "RUB", number: 643, decimalPlaces: 2, description: "Russian ruble"},
		{code: "RWF", number: 646, decimalPlaces: 0, description: "Rwandan franc"},
		{code: "SAR", number: 682, decimalPlaces: 2, description: "Saudi riyal"},
		{code: "SBD", number: 90, decimalPlaces: 2, description: "Solomon Islands dollar"},
		{code: "SCR", number: 690, decimalPlaces: 2, description: "Seychelles rupee"},
		{code: "SDG", number: 938, decimalPlaces: 2, description: "Sudanese pound"},
		{code: "SEK", number: 752, decimalPlaces: 2, description: "Swedish krona/kronor"},
		{code: "SGD", number: 702, decimalPlaces: 2, description: "Singapore dollar"},
		{code: "SHP", number: 654, decimalPlaces: 2, description: "Saint Helena pound"},
		{code: "SLL", number: 694, decimalPlaces: 2, description: "Sierra Leonean leone"},
		{code: "SOS", number: 706, decimalPlaces: 2, description: "Somali shilling"},
		{code: "SRD", number: 968, decimalPlaces: 2, description: "Surinamese dollar"},
		{code: "SSP", number: 728, decimalPlaces: 2, description: "South Sudanese pound"},
		{code: "STD", number: 678, decimalPlaces: 2, description: "São Tomé and Príncipe dobra"},
		{code: "SYP", number: 760, decimalPlaces: 2, description: "Syrian pound"},
		{code: "SZL", number: 748, decimalPlaces: 2, description: "Swazi lilangeni"},
		{code: "THB", number: 764, decimalPlaces: 2, description: "Thai baht"},
		{code: "TJS", number: 972, decimalPlaces: 2, description: "Tajikistani somoni"},
		{code: "TMT", number: 934, decimalPlaces: 2, description: "Turkmenistani manat"},
		{code: "TND", number: 788, decimalPlaces: 3, description: "Tunisian dinar"},
		{code: "TOP", number: 776, decimalPlaces: 2, description: "Tongan paʻanga"},
		{code: "TRY", number: 949, decimalPlaces: 2, description: "Turkish lira"},
		{code: "TTD", number: 780, decimalPlaces: 2, description: "Trinidad and Tobago dollar"},
		{code: "TWD", number: 901, decimalPlaces: 2, description: "New Taiwan dollar"},
		{code: "TZS", number: 834, decimalPlaces: 2, description: "Tanzanian shilling"},
		{code: "UAH", number: 980, decimalPlaces: 2, description: "Ukrainian hryvnia"},
		{code: "UGX", number: 800, decimalPlaces: 0, description: "Ugandan shilling"},
		{code: "USD", number: 840, decimalPlaces: 2, description: "United States dollar"},
		{code: "USN", number: 997, decimalPlaces: 2, description: "United States dollar (next day) (funds code)"},
		{code: "USS", number: 998, decimalPlaces: 2, description: "United States dollar (same day) (funds code)[10]"},
		{code: "UYI", number: 940, decimalPlaces: 0, description: "Uruguay Peso en Unidades Indexadas (URUIURUI) (funds code)"},
		{code: "UYU", number: 858, decimalPlaces: 2, description: "Uruguayan peso"},
		{code: "UZS", number: 860, decimalPlaces: 2, description: "Uzbekistan som"},
		{code: "VEF", number: 937, decimalPlaces: 2, description: "Venezuelan bolívar"},
		{code: "VND", number: 704, decimalPlaces: 0, description: "Vietnamese dong"},
		{code: "VUV", number: 548, decimalPlaces: 0, description: "Vanuatu vatu"},
		{code: "WST", number: 882, decimalPlaces: 2, description: "Samoan tala"},
		{code: "XAF", number: 950, decimalPlaces: 0, description: "CFA franc BEAC"},
		{code: "XCD", number: 951, decimalPlaces: 2, description: "East Caribbean dollar"},
		{code: "XOF", number: 952, decimalPlaces: 0, description: "CFA franc BCEAO"},
		{code: "XPF", number: 953, decimalPlaces: 0, description: "CFP franc (franc Pacifique)"},
		{code: "YER", number: 886, decimalPlaces: 2, description: "Yemeni rial"},
		{code: "ZAR", number: 710, decimalPlaces: 2, description: "South African rand"},
		{code: "ZMW", number: 967, decimalPlaces: 2, description: "Zambian kwacha"},
	}
	currencyByCodeMap   map[string]*Currency
	currencyByNumberMap map[int]*Currency
)

func init() {
	currencyByCodeMap = make(map[string]*Currency)
	currencyByNumberMap = make(map[int]*Currency)

	scales := []int{1, 10, 100, 1000, 10000}

	for i := range currencyAllCurrencies {
		currency := &(currencyAllCurrencies[i])
		currency.scale = scales[currency.decimalPlaces]
		currencyByCodeMap[currency.code] = currency
		currencyByNumberMap[currency.number] = currency
	}
}

// GetCode returns 3-letter currency code, e.g. "USD" for US dollar
func (this *Currency) GetCode() string {
	return this.code
}

// GetNumber returns numeric currency code, e.g. 840 for US dollar
func (this *Currency) GetNumber() int {
	return this.number
}

// GetDecimalPlaces returns number of decimal places used in this currency. E.g. 2 for US dollar or 0 for Japanese Yen.
func (this *Currency) GetDecimalPlaces() int {
	return this.decimalPlaces
}

// GetScale returns a multiple that can be used to convert an amount to "amont in lowest denomitation". E.g. 100 for US dollar (to convert dollar amount to amount in cents) or 1 for Japanese Yen (as Yen is the lowest denomination).
func (this *Currency) GetScale() int {
	return this.scale
}

// GetDescription returns a plain english name of the currency. E.g. "United States dollar" for US dollar.
func (this *Currency) GetDescription() string {
	return this.description
}

// GetCurrencyByCode finds a currency by its 3-letter code, returns nil if not present in the database.
func GetCurrencyByCode(code string) *Currency {
	ret, ok := currencyByCodeMap[code]
	if !ok {
		ret = nil
	}
	return ret
}

// GetCurrencyByNumber finds a currency by its numeric code, returns nil if not present in the database.
func GetCurrencyByNumber(number int) *Currency {
	ret, ok := currencyByNumberMap[number]
	if !ok {
		ret = nil
	}
	return ret
}
