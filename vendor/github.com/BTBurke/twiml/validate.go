package twiml

import "regexp"

//type ValidationFunc func(...interface{}) bool

func Validate(vf ...bool) bool {
	for _, f := range vf {
		if !f {
			return false
		}
	}
	return true
}

func OneOf(field string, options ...string) bool {
	for _, w := range options {
		if field == w {
			return true
		}
	}
	return false
}

func IntBetween(field int, high int, low int) bool {
	if (field <= high) && (field >= low) {
		return true
	}
	return false
}

func Required(field string) bool {
	if len(field) > 0 {
		return true
	}
	return false
}

func OneOfOpt(field string, options ...string) bool {
	if field == "" {
		return true
	}
	return OneOf(field, options...)
}

func AllowedMethod(field string) bool {
	// optional field always set with default (typically POST)
	if field == "" {
		return true
	}
	if (field != "GET") && (field != "POST") {
		return false
	}
	return true
}

func Numeric(field string) bool {
	matched, err := regexp.MatchString("[0-9]+", field)
	if err != nil {
		return false
	}
	return matched
}

func NumericOpt(field string) bool {
	if field == "" {
		return true
	}
	return Numeric(field)
}

func AllowedLanguage(speaker string, language string) bool {
	switch speaker {
	case Man, Woman:
		return OneOfOpt(language, English, French, German, Spanish, EnglishUK)
	case Alice:
		return OneOfOpt(language,
			DanishDenmark,
			GermanGermany,
			EnglishAustralia,
			EnglishCanada,
			EnglishUK,
			EnglishIndia,
			EnglishUSA,
			SpanishCatalan,
			SpanishSpain,
			SpanishMexico,
			FinishFinland,
			FrenchCanada,
			FrenchFrance,
			ItalianItaly,
			JapaneseJapan,
			KoreanKorea,
			NorwegianNorway,
			DutchNetherlands,
			PolishPoland,
			PortugueseBrazil,
			PortuguesePortugal,
			RussianRussia,
			SwedishSweden,
			ChineseMandarin,
			ChineseCantonese,
			ChineseTaiwanese,
		)
	default:
		return OneOfOpt(language, English, French, German, Spanish, EnglishUK)
	}
}

func AllowedCallbackEvent(events string) bool {
	if events == "" {
		return true
	}
	var validEvents = regexp.MustCompile(`^(initiated\s?|ringing\s?|answered\s?|completed\s?)+$`)
	return validEvents.MatchString(events)
}
