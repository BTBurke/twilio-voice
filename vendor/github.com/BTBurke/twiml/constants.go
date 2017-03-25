package twiml

// Constants for language
const (
	Man                = "man"
	Woman              = "woman"
	Alice              = "alice"
	English            = "en"
	French             = "fr"
	Spanish            = "es"
	German             = "de"
	DanishDenmark      = "da-DK"
	GermanGermany      = "de-DE"
	EnglishAustralia   = "en-AU"
	EnglishCanada      = "en-CA"
	EnglishUK          = "en-UK"
	EnglishIndia       = "en-IN"
	EnglishUSA         = "en-US"
	SpanishCatalan     = "ca-ES"
	SpanishSpain       = "es-ES"
	SpanishMexico      = "es-MX"
	FinishFinland      = "fi-FI"
	FrenchCanada       = "fr-CA"
	FrenchFrance       = "fr-FR"
	ItalianItaly       = "it-IT"
	JapaneseJapan      = "ja-JP"
	KoreanKorea        = "ko-KR"
	NorwegianNorway    = "nb-NO"
	DutchNetherlands   = "nl-NL"
	PolishPoland       = "pl-PL"
	PortugueseBrazil   = "pt-BR"
	PortuguesePortugal = "pt-PT"
	RussianRussia      = "ru-RU"
	SwedishSweden      = "sv-SE"
	ChineseMandarin    = "zh-CH"
	ChineseCantonese   = "zh-HK"
	ChineseTaiwanese   = "zh-TW"
)

// Call status
const (
	Queued     = "queued"
	Ringing    = "ringing"
	InProgress = "in-progress"
	Completed  = "completed"
	Busy       = "busy"
	Failed     = "failed"
	NoAnswer   = "no-answer"
	Canceled   = "canceled"
)

// Call directions
const (
	OutboundAPI  = "outbound-api"
	Inbound      = "inbound"
	OutboundDial = "outbound-dial"
)

// Trim options
const (
	TrimSilence = "trim-silence"
	DoNotTrim   = "do-not-trim"
)
