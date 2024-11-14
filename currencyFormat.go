package currency


type Currency struct {
	Code                        string
	Symbol                      string
	ThousandsSeparator          string
	DecimalSeparator            string
	SymbolOnLeft                bool
	SpaceBetweenAmountAndSymbol bool
	DecimalDigits               int
}

var CURRENCIES = map[string]Currency{
	"GBP": {"GBP", "£", ",", ".", true, false, 2},
	"THB": {"THB", "฿", ",", ".", true, false, 2},
	"IDR": {"IDR", "Rp", ".", ",", true, false, 0},
	"USD": {"USD", "$", ",", ".", true, false, 2},
	"VND": {"VND", "₫", ".", ".", false, true, 0},
	"MYR": {"MYR", "RM", ",", ".", true, false, 2},
	"AED": {"AED", "د.إ.‏", ",", ".", true, true, 2},
	"AFN": {"AFN", "؋", ",", ".", true, false, 2},
	"ALL": {"ALL", "Lek", ".", ",", false, false, 2},
	"AMD": {"AMD", "֏", ",", ".", false, true, 2},
	"ANG": {"ANG", "ƒ", ",", ".", true, false, 2},
	"AOA": {"AOA", "Kz", ",", ".", true, false, 2},
	"ARS": {"ARS", "$", ".", ",", true, true, 2},
	"AUD": {"AUD", "$", ",", ".", true, false, 2},
	"AWG": {"AWG", "ƒ", ",", ".", true, false, 2},
	"AZN": {"AZN", "₼", " ", ",", false, true, 2},
	"BAM": {"BAM", "КМ", ".", ",", false, true, 2},
	"BBD": {"BBD", "$", ",", ".", true, false, 2},
	"BDT": {"BDT", "৳", ",", ".", true, true, 0},
	"BGN": {"BGN", "лв.", " ", ",", false, true, 2},
	"BHD": {"BHD", "د.ب.‏", ",", ".", true, true, 3},
	"BIF": {"BIF", "FBu", ",", ".", false, false, 0},
	"BMD": {"BMD", "$", ",", ".", true, false, 2},
	"BND": {"BND", "$", ".", ",", true, false, 0},
	"BOB": {"BOB", "Bs", ".", ",", true, true, 2},
	"BRL": {"BRL", "R$", ".", ",", true, true, 2},
	"BSD": {"BSD", "$", ",", ".", true, false, 2},
	"BTC": {"BTC", "Ƀ", ",", ".", false, false, 8},
	"BTN": {"BTN", "Nu.", ",", ".", true, true, 1},
	"BWP": {"BWP", "P", ",", ".", true, false, 2},
	"BYR": {"BYR", "р.", " ", ",", false, true, 2},
	"BYN": {"BYN", "р.", " ", ",", false, true, 2},
	"BZD": {"BZD", "BZ$", ",", ".", true, false, 2},
	"CAD": {"CAD", "$", ",", ".", true, false, 2},
	"CDF": {"CDF", "FC", ",", ".", false, false, 2},
	"CHF": {"CHF", "CHF", "'", ".", true, true, 2},
	"CLP": {"CLP", "$", ".", ",", true, true, 0},
	"CNY": {"CNY", "¥", ",", ".", true, false, 2},
	"COP": {"COP", "$", ".", ",", true, true, 2},
	"CRC": {"CRC", "₡", ".", ",", true, false, 2},
	"CUC": {"CUC", "CUC", ",", ".", true, false, 2},
	"CUP": {"CUP", "$MN", ",", ".", true, false, 2},
	"CVE": {"CVE", "$", ",", ".", true, false, 2},
	"CZK": {"CZK", "Kč", " ", ",", false, true, 2},
	"DJF": {"DJF", "Fdj", ",", ".", false, false, 0},
	"DKK": {"DKK", "kr.", "", ",", false, true, 2},
	"DOP": {"DOP", "RD$", ",", ".", true, false, 2},
	"DZD": {"DZD", "د.ج.‏", ",", ".", true, true, 2},
	"EGP": {"EGP", "ج.م.‏", ",", ".", true, true, 2},
	"ERN": {"ERN", "Nfk", ",", ".", false, false, 2},
	"ETB": {"ETB", "ETB", ",", ".", true, false, 2},
	"EUR": {"EUR", "€", " ", ",", false, true, 2},
	"FJD": {"FJD", "$", ",", ".", true, false, 2},
	"FKP": {"FKP", "£", ",", ".", true, false, 2},
	"GEL": {"GEL", "GEL", " ", ",", false, true, 2},
	"GHS": {"GHS", "₵", ",", ".", true, false, 2},
	"GIP": {"GIP", "£", ",", ".", true, false, 2},
	"GMD": {"GMD", "D", ",", ".", false, false, 2},
	"GNF": {"GNF", "FG", ",", ".", false, false, 0},
	"GTQ": {"GTQ", "Q", ",", ".", true, false, 2},
	"GYD": {"GYD", "$", ",", ".", true, false, 2},
	"HKD": {"HKD", "HK$", ",", ".", true, false, 2},
	"HNL": {"HNL", "L.", ",", ".", true, true, 2},
	"HRK": {"HRK", "kn", ".", ",", false, true, 2},
	"HTG": {"HTG", "G", ",", ".", true, false, 2},
	"HUF": {"HUF", "Ft", " ", ",", false, true, 2},
	"ILS": {"ILS", "₪", ",", ".", true, true, 2},
	"INR": {"INR", "₹", ",", ".", true, false, 2},
	"IQD": {"IQD", "د.ع.‏", ",", ".", true, true, 2},
	"IRR": {"IRR", "﷼", ",", "/", true, true, 2},
	"ISK": {"ISK", "kr.", ".", ",", false, true, 0},
	"JMD": {"JMD", "J$", ",", ".", true, false, 2},
	"JOD": {"JOD", "د.ا.‏", ",", ".", true, true, 3},
	"JPY": {"JPY", "¥", ",", ".", true, false, 0},
	"KES": {"KES", "KSh", ",", ".", true, false, 2},
	"KGS": {"KGS", "сом", " ", "-", false, true, 2},
	"KHR": {"KHR", "៛", ",", ".", false, false, 0},
	"KMF": {"KMF", "CF", ",", ".", false, false, 2},
	"KPW": {"KPW", "₩", ",", ".", true, false, 0},
	"KRW": {"KRW", "₩", ",", ".", true, false, 0},
	"KWD": {"KWD", "د.ك.‏", ",", ".", true, true, 3},
	"KYD": {"KYD", "$", ",", ".", true, false, 2},
	"KZT": {"KZT", "₸", " ", "-", true, false, 2},
	"LAK": {"LAK", "₭", ",", ".", false, false, 0},
	"LBP": {"LBP", "ل.ل.‏", ",", ".", true, true, 2},
	"LKR": {"LKR", "₨", ",", ".", true, true, 0},
	"LRD": {"LRD", "$", ",", ".", true, false, 2},
	"LSL": {"LSL", "M", ",", ".", false, false, 2},
	"LYD": {"LYD", "د.ل.‏", ",", ".", true, false, 3},
	"MAD": {"MAD", "د.م.‏", ",", ".", true, true, 2},
	"MDL": {"MDL", "lei", ",", ".", false, true, 2},
	"MGA": {"MGA", "Ar", ",", ".", true, false, 0},
	"MKD": {"MKD", "ден.", ".", ",", false, true, 2},
	"MMK": {"MMK", "K", ",", ".", true, false, 2},
	"MNT": {"MNT", "₮", " ", ",", true, false, 2},
	"MOP": {"MOP", "MOP$", ",", ".", true, false, 2},
	"MRO": {"MRO", "UM", ",", ".", false, false, 2},
	"MTL": {"MTL", "₤", ",", ".", true, false, 2},
	"MUR": {"MUR", "₨", ",", ".", true, false, 2},
	"MVR": {"MVR", "MVR", ",", ".", false, true, 1},
	"MWK": {"MWK", "MK", ",", ".", true, false, 2},
	"MXN": {"MXN", "$", ",", ".", true, false, 2},
	"MZN": {"MZN", "MT", ",", ".", true, false, 0},
	"NAD": {"NAD", "$", ",", ".", true, false, 2},
	"NGN": {"NGN", "₦", ",", ".", true, false, 2},
	"NIO": {"NIO", "C$", ",", ".", true, true, 2},
	"NOK": {"NOK", "kr", " ", ",", true, true, 2},
	"NPR": {"NPR", "₨", ",", ".", true, false, 2},
	"NZD": {"NZD", "$", ",", ".", true, false, 2},
	"OMR": {"OMR", "﷼", ",", ".", true, true, 3},
	"PAB": {"PAB", "B/.", ",", ".", true, true, 2},
	"PEN": {"PEN", "S/.", ",", ".", true, true, 2},
	"PGK": {"PGK", "K", ",", ".", true, false, 2},
	"PHP": {"PHP", "₱", ",", ".", true, false, 2},
	"PKR": {"PKR", "₨", ",", ".", true, false, 2},
	"PLN": {"PLN", "zł", " ", ",", false, true, 2},
	"PYG": {"PYG", "₲", ".", ",", true, true, 2},
	"QAR": {"QAR", "﷼", ",", ".", true, true, 2},
	"RON": {"RON", "L", ".", ",", false, true, 2},
	"RSD": {"RSD", "Дин.", ".", ",", false, true, 2},
	"RUB": {"RUB", "₽", " ", ",", false, true, 2},
	"RWF": {"RWF", "RWF", " ", ",", true, true, 2},
	"SAR": {"SAR", "﷼", ",", ".", true, true, 2},
	"SBD": {"SBD", "$", ",", ".", true, false, 2},
	"SCR": {"SCR", "₨", ",", ".", true, false, 2},
	"SDD": {"SDD", "LSd", ",", ".", false, false, 2},
	"SDG": {"SDG", "£‏", ",", ".", true, false, 2},
	"SEK": {"SEK", "kr", ".", ",", false, true, 2},
	"SGD": {"SGD", "$", ",", ".", true, false, 2},
	"SHP": {"SHP", "£", ",", ".", true, false, 2},
	"SLL": {"SLL", "Le", ",", ".", true, false, 2},
	"SOS": {"SOS", "S", ",", ".", true, false, 2},
	"SRD": {"SRD", "$", ",", ".", true, false, 2},
	"STD": {"STD", "Db", ",", ".", true, false, 2},
	"SVC": {"SVC", "₡", ",", ".", true, false, 2},
	"SYP": {"SYP", "£", ",", ".", true, true, 2},
	"SZL": {"SZL", "E", ",", ".", true, false, 2},
	"TJS": {"TJS", "TJS", " ", ";", false, true, 2},
	"TMT": {"TMT", "m", " ", ",", false, false, 0},
	"TND": {"TND", "د.ت.‏", ",", ".", true, true, 3},
	"TOP": {"TOP", "T$", ",", ".", true, false, 2},
	"TRY": {"TRY", "₺", ".", ",", true, false, 2},
	"TTD": {"TTD", "TT$", ",", ".", true, false, 2},
	"TVD": {"TVD", "$", ",", ".", true, false, 2},
	"TWD": {"TWD", "NT$", ",", ".", true, false, 2},
	"TZS": {"TZS", "TSh", ",", ".", true, false, 2},
	"UAH": {"UAH", "₴", " ", ",", false, false, 2},
	"UGX": {"UGX", "USh", ",", ".", true, false, 2},
	"UYU": {"UYU", "$U", ".", ",", true, true, 2},
	"UZS": {"UZS", "сўм", " ", ",", false, true, 2},
	"VEB": {"VEB", "Bs.", ",", ".", true, false, 2},
	"VEF": {"VEF", "Bs. F.", ".", ",", true, true, 2},
	"VUV": {"VUV", "VT", ",", ".", false, false, 0},
	"WST": {"WST", "WS$", ",", ".", true, false, 2},
	"XAF": {"XAF", "F", ",", ".", false, false, 2},
	"XCD": {"XCD", "$", ",", ".", true, false, 2},
	"XBT": {"XBT", "Ƀ", ",", ".", false, false, 2},
	"XOF": {"XOF", "F", " ", ",", false, false, 2},
	"XPF": {"XPF", "F", ",", ".", false, false, 2},
	"YER": {"YER", "﷼", ",", ".", true, true, 2},
	"ZAR": {"ZAR", "R", " ", ",", true, false, 2},
	"ZMW": {"ZMW", "ZK", ",", ".", true, false, 2},
	"WON": {"WON", "₩", ",", ".", true, false, 2},
	// Add other currencies here...
}
