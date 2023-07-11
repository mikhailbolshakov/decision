package kit

var (
	countries = []*Country{
		{
			NameEng:        "Russia",
			TranslationKey: "countries.rus",
			Code:           "643",
			Alfa2:          "RU",
			Alfa3:          "RUS",
			TimeZones:      []string{TzP12, TzP11, TzP10, TzP9, TzP8, TzP7, TzP6, TzP5, TzP4, TzP3, TzP2},
			PhoneCodes:     []string{"7"},
			Currencies:     GetCurrencies(CurRUB),
		},
		{
			NameEng:        "Serbia",
			TranslationKey: "countries.srb",
			Code:           "688",
			Alfa2:          "RS",
			Alfa3:          "SRB",
			TimeZones:      []string{TzP1, TzP2},
			PhoneCodes:     []string{"381"},
			Currencies:     GetCurrencies(CurRSD, CurEUR),
		},
		{
			NameEng:        "USA",
			TranslationKey: "countries.usa",
			Code:           "840",
			Alfa2:          "US",
			Alfa3:          "USA",
			TimeZones:      []string{TzM10, TzM9, TzM8, TzM7, TzM6, TzM5, TzM4},
			PhoneCodes:     []string{"1"},
			Currencies:     GetCurrencies(CurUSD),
		},
		{
			NameEng:        "Slovenia",
			TranslationKey: "countries.svn",
			Code:           "705",
			Alfa2:          "SI",
			Alfa3:          "SVN",
			TimeZones:      []string{TzP1, TzP2},
			PhoneCodes:     []string{"386"},
			Currencies:     GetCurrencies(CurEUR),
		},
		{
			NameEng:        "Kazakhstan",
			TranslationKey: "countries.kaz",
			Code:           "398",
			Alfa2:          "KZ",
			Alfa3:          "KAZ",
			TimeZones:      []string{TzP5, TzP6},
			PhoneCodes:     []string{"7", "997"},
			Currencies:     GetCurrencies(CurKZT),
		},
	}
)
