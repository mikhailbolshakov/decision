package kit

const (
	CurRUB = "RUB"
	CurUSD = "USD"
	CurRSD = "RSD"
	CurEUR = "EUR"
	CurKZT = "KZT"
)

var (
	currenciesByISO = map[string]*Currency{
		CurRUB: {
			NameEng:        "Ruble",
			TranslationKey: "currencies.rub",
			IsoCode:        CurRUB,
			Number:         "643",
			Symbol:         "₽",
			Unit:           100,
		},
		CurRSD: {
			NameEng:        "Dinar",
			TranslationKey: "currencies.rsd",
			IsoCode:        CurRSD,
			Number:         "941",
			Symbol:         "Дин.",
			Unit:           100,
		},
		CurUSD: {
			NameEng:        "US Dollar",
			TranslationKey: "currencies.usd",
			IsoCode:        CurUSD,
			Number:         "840",
			Symbol:         "$",
			Unit:           1,
		},
		CurEUR: {
			NameEng:        "Euro",
			TranslationKey: "currencies.eur",
			IsoCode:        CurEUR,
			Number:         "978",
			Symbol:         "€",
			Unit:           1,
		},
		CurKZT: {
			NameEng:        "Tenge",
			TranslationKey: "currencies.kzt",
			IsoCode:        CurKZT,
			Number:         "398",
			Symbol:         "₸",
			Unit:           500,
		},
	}
)
