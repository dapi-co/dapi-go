package types

type Currency struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

var (
	AUD = &Currency{Code: "AUD", Name: "AUSTRALIAN DOLLAR"}
	AED = &Currency{Code: "AED", Name: "UNITED ARAB EMIRATES DIRHAM"}
	BHD = &Currency{Code: "BHD", Name: "BAHRAINI DINAR"}
	CAD = &Currency{Code: "CAD", Name: "CANADIAN DOLLAR"}
	DKK = &Currency{Code: "DKK", Name: "DANISH KRONE"}
	EGP = &Currency{Code: "EGP", Name: "EGYPTIAN POUND"}
	EUR = &Currency{Code: "EUR", Name: "EURO"}
	HKD = &Currency{Code: "HKD", Name: "HONG KONG DOLLAR"}
	INR = &Currency{Code: "INR", Name: "INDIAN RUPEE"}
	JPY = &Currency{Code: "JPY", Name: "JAPANESE YEN"}
	JOD = &Currency{Code: "JOD", Name: "JORDANIAN DINAR"}
	CZK = &Currency{Code: "CZK", Name: "KORONA"}
	KWD = &Currency{Code: "KWD", Name: "KUWAITI DINAR"}
	MAD = &Currency{Code: "MAD", Name: "MOROCCAN DIRHAM"}
	NZD = &Currency{Code: "NZD", Name: "NEW ZEALAND DOLLAR"}
	NOK = &Currency{Code: "NOK", Name: "NORWEGIAN KRONE"}
	PKR = &Currency{Code: "PKR", Name: "PAKISTAN RUPEE"}
	GBP = &Currency{Code: "GBP", Name: "POUND STERLING"}
	PHP = &Currency{Code: "PHP", Name: "PHILIPPINE PESO"}
	QAR = &Currency{Code: "QAR", Name: "QATARI RIYAL"}
	OMR = &Currency{Code: "OMR", Name: "RIAL OMANI"}
	SAR = &Currency{Code: "SAR", Name: "SAUDI RIYAL"}
	SGD = &Currency{Code: "SGD", Name: "SINGAPORE DOLLAR"}
	ZAR = &Currency{Code: "ZAR", Name: "SOUTH AFRICAN RAND"}
	LKR = &Currency{Code: "LKR", Name: "SRI LANKA RUPEE"}
	SEK = &Currency{Code: "SEK", Name: "SWEDISH KRONA"}
	CHF = &Currency{Code: "CHF", Name: "SWISS FRANC"}
	TND = &Currency{Code: "TND", Name: "TUNISIAN DINARS"}
	TRY = &Currency{Code: "TRY", Name: "TURKISH LIRA"}
	USD = &Currency{Code: "USD", Name: "US DOLLAR"}
	XAU = &Currency{Code: "XAU", Name: "GOLD"}
)
