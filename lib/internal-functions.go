package lib

import (
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"github.com/zclconf/go-cty/cty/function/stdlib"
)

var defaultFunctions = map[string]function.Function{
	// General Functions
	"equal":    stdlib.EqualFunc,
	"notEqual": stdlib.NotEqualFunc,
	"coalesce": stdlib.CoalesceFunc,
	"not":      stdlib.NotFunc,
	"and":      stdlib.AndFunc,
	"or":       stdlib.OrFunc,
	// String Functions
	"upper":      stdlib.UpperFunc,
	"lower":      stdlib.LowerFunc,
	"reverse":    stdlib.ReverseFunc,
	"strlen":     stdlib.StrlenFunc,
	"substr":     stdlib.SubstrFunc,
	"join":       stdlib.JoinFunc,
	"sort":       stdlib.SortFunc,
	"split":      stdlib.SplitFunc,
	"chomp":      stdlib.ChompFunc,
	"indent":     stdlib.IndentFunc,
	"title":      stdlib.TitleFunc,
	"trimSpace":  stdlib.TrimSpaceFunc,
	"trim":       stdlib.TrimFunc,
	"trimPrefix": stdlib.TrimPrefixFunc,
	"trimSuffix": stdlib.TrimSuffixFunc,
	"format":     stdlib.FormatFunc,
	"formatList": stdlib.FormatListFunc,
	"replace":    stdlib.ReplaceFunc,
	// Number Functions
	"absolute":             stdlib.AbsoluteFunc,
	"add":                  stdlib.AddFunc,
	"subtract":             stdlib.SubtractFunc,
	"multiply":             stdlib.MultiplyFunc,
	"divide":               stdlib.DivideFunc,
	"modulo":               stdlib.ModuloFunc,
	"greaterThan":          stdlib.GreaterThanFunc,
	"greaterThanOrEqualTo": stdlib.GreaterThanOrEqualToFunc,
	"lessThan":             stdlib.LessThanFunc,
	"lessThanOrEqualTo":    stdlib.LessThanOrEqualToFunc,
	"negate":               stdlib.NegateFunc,
	"min":                  stdlib.MinFunc,
	"max":                  stdlib.MaxFunc,
	"int":                  stdlib.IntFunc,
	"ceil":                 stdlib.CeilFunc,
	"floor":                stdlib.FloorFunc,
	"log":                  stdlib.LogFunc,
	"pow":                  stdlib.PowFunc,
	"signum":               stdlib.SignumFunc,
	"parseInt":             stdlib.ParseIntFunc,
	// Collection Functions
	"hasIndex":     stdlib.HasIndexFunc,
	"index":        stdlib.IndexFunc,
	"length":       stdlib.LengthFunc,
	"element":      stdlib.ElementFunc,
	"coalesceList": stdlib.CoalesceListFunc,
	"compact":      stdlib.CompactFunc,
	"contains":     stdlib.ContainsFunc,
	"distinct":     stdlib.DistinctFunc,
	"chunklist":    stdlib.ChunklistFunc,
	"flatten":      stdlib.FlattenFunc,
	"keys":         stdlib.KeysFunc,
	"lookup":       stdlib.LookupFunc,
	"merge":        stdlib.MergeFunc,
	"reverseList":  stdlib.ReverseListFunc,
	"setProduct":   stdlib.SetProductFunc,
	"slice":        stdlib.SliceFunc,
	"values":       stdlib.ValuesFunc,
	"zipmap":       stdlib.ZipmapFunc,
	"concat":       stdlib.ConcatFunc,
	"range":        stdlib.RangeFunc,
	// Encoding Functions
	"csvDecode":  stdlib.CSVDecodeFunc,
	"jsonEncode": stdlib.JSONEncodeFunc,
	"jsonDecode": stdlib.JSONDecodeFunc,
	// Datetime Functions
	"formatDate": stdlib.FormatDateFunc,
	"timeAdd":    stdlib.TimeAddFunc,
	// Regex Functions
	"regex":        stdlib.RegexFunc,
	"regexAll":     stdlib.RegexAllFunc,
	"regexReplace": stdlib.RegexReplaceFunc,
}

type GofakeitFunc struct {
	Params []function.Parameter
	Type   function.TypeFunc
	Impl   function.ImplFunc
}

var gofakeitFunctionList = map[string]GofakeitFunc{
	"gofakeitName": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitNameImpl,
	},
	"gofakeitNamePrefix": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitNamePrefixImpl,
	},
	"gofakeitNameSuffix": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitNameSuffixImpl,
	},
	"gofakeitFirstName": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitFirstNameImpl,
	},
	"gofakeitLastName": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitLastNameImpl,
	},
	"gofakeitGender": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitGenderImpl,
	},
	"gofakeitSSN": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitSSNImpl,
	},
	"gofakeitEmail": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitEmailImpl,
	},
	"gofakeitPhone": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitPhoneImpl,
	},
	"gofakeitPhoneFormatted": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitPhoneFormattedImpl,
	},
	"gofakeitUsername": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitUsernameImpl,
	},
	"gofakeitPassword": {
		Params: []function.Parameter{
			{
				Name:             "lower",
				Type:             cty.Bool,
				AllowDynamicType: false,
			},
			{
				Name:             "upper",
				Type:             cty.Bool,
				AllowDynamicType: false,
			},
			{
				Name:             "numeric",
				Type:             cty.Bool,
				AllowDynamicType: false,
			},
			{
				Name:             "special",
				Type:             cty.Bool,
				AllowDynamicType: false,
			},
			{
				Name:             "space",
				Type:             cty.Bool,
				AllowDynamicType: false,
			},
			{
				Name:             "length",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.String), Impl: gofakeitPasswordImpl,
	},
	"gofakeitCity": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitCityImpl,
	},
	"gofakeitCountry": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitCountryImpl,
	},
	"gofakeitCountryAbr": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitCountryAbrImpl,
	},
	"gofakeitState": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitStateImpl,
	},
	"gofakeitStateAbr": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitStateAbrImpl,
	},
	"gofakeitStreet": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitStreetImpl,
	},
	"gofakeitStreetName": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitStreetNameImpl,
	},
	"gofakeitStreetNumber": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitStreetNumberImpl,
	},
	"gofakeitStreetPrefix": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitStreetPrefixImpl,
	},
	"gofakeitStreetSuffix": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitStreetSuffixImpl,
	},
	"gofakeitZip": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitZipImpl,
	},
	"gofakeitLatitude": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitLatitudeImpl,
	},
	"gofakeitLatitudeInRange": {
		Params: []function.Parameter{
			{
				Name:             "min",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
			{
				Name:             "max",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.String), Impl: gofakeitLatitudeInRangeImpl,
	},
	"gofakeitLongitude": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitLongitudeImpl,
	},
	"gofakeitLongitudeInRange": {
		Params: []function.Parameter{
			{
				Name:             "min",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
			{
				Name:             "max",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.String), Impl: gofakeitLongitudeInRangeImpl,
	},
	"gofakeitGamertag": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitGamertagImpl,
	},
	"gofakeitBeerAlcohol": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitBeerAlcoholImpl,
	},
	"gofakeitBeerBlg": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitBeerBlgImpl,
	},
	"gofakeitBeerHop": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitBeerHopImpl,
	},
	"gofakeitBeerIbu": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitBeerIbuImpl,
	},
	"gofakeitBeerMalt": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitBeerMaltImpl,
	},
	"gofakeitBeerName": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitBeerNameImpl,
	},
	"gofakeitBeerStyle": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitBeerStyleImpl,
	},
	"gofakeitBeerYeast": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitBeerYeastImpl,
	},
	"gofakeitCarMaker": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitCarMakerImpl,
	},
	"gofakeitCarModel": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitCarModelImpl,
	},
	"gofakeitCarType": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitCarTypeImpl,
	},
	"gofakeitCarFuelType": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitCarFuelTypeImpl,
	},
	"gofakeitNoun": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitNounImpl,
	},
	"gofakeitVerb": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitVerbImpl,
	},
	"gofakeitAdverb": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitAdverbImpl,
	},
	"gofakeitPreposition": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitPrepositionImpl,
	},
	"gofakeitAdjective": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitAdjectiveImpl,
	},
	"gofakeitWord": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitWordImpl,
	},
	"gofakeitSentence": {
		Params: []function.Parameter{
			{
				Name:             "wordCount",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.String), Impl: gofakeitSentenceImpl,
	},
	"gofakeitParagraph": {
		Params: []function.Parameter{
			{
				Name:             "paragraphCount",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
			{
				Name:             "sentenceCount",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
			{
				Name:             "wordCount",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
			{
				Name:             "separator",
				Type:             cty.String,
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.String), Impl: gofakeitParagraphImpl,
	},
	"gofakeitLoremIpsumWord": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitLoremIpsumWordImpl,
	},
	"gofakeitLoremIpsumSentence": {
		Params: []function.Parameter{
			{
				Name:             "wordCount",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.String), Impl: gofakeitLoremIpsumSentenceImpl,
	},
	"gofakeitLoremIpsumParagraph": {
		Params: []function.Parameter{
			{
				Name:             "paragraphCount",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
			{
				Name:             "sentenceCount",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
			{
				Name:             "wordCount",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
			{
				Name:             "separator",
				Type:             cty.String,
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.String), Impl: gofakeitLoremIpsumParagraphImpl,
	},
	"gofakeitQuestion": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitQuestionImpl,
	},
	"gofakeitQuote": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitQuoteImpl,
	},
	"gofakeitPhrase": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitPhraseImpl,
	},
	"gofakeitFruit": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitFruitImpl,
	},
	"gofakeitVegetable": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitVegetableImpl,
	},
	"gofakeitBreakfast": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitBreakfastImpl,
	},
	"gofakeitLunch": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitLunchImpl,
	},
	"gofakeitDinner": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitDinnerImpl,
	},
	"gofakeitSnack": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitSnackImpl,
	},
	"gofakeitDessert": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitDessertImpl,
	},
	"gofakeitBool": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitBoolImpl,
	},
	"gofakeitUUID": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitUUIDImpl,
	},
	"gofakeitColor": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitColorImpl,
	},
	"gofakeitHexColor": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitHexColorImpl,
	},
	"gofakeitRGBColor": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitRGBColorImpl,
	},
	"gofakeitSafeColor": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitSafeColorImpl,
	},
	"gofakeitURL": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitURLImpl,
	},
	"gofakeitImageURL": {
		Params: []function.Parameter{
			{
				Name:             "width",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
			{
				Name:             "height",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.String), Impl: gofakeitImageURLImpl,
	},
	"gofakeitDomainName": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitDomainNameImpl,
	},
	"gofakeitDomainSuffix": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitDomainSuffixImpl,
	},
	"gofakeitIPv4Address": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitIPv4AddressImpl,
	},
	"gofakeitIPv6Address": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitIPv6AddressImpl,
	},
	"gofakeitHTTPStatusCode": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitHTTPStatusCodeImpl,
	},
	"gofakeitHTTPSimpleStatusCode": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitHTTPSimpleStatusCodeImpl,
	},
	"gofakeitLogLevel": {
		Params: []function.Parameter{
			{
				Name:             "logType",
				Type:             cty.String,
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.String), Impl: gofakeitLogLevelImpl,
	},
	"gofakeitHTTPMethod": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitHTTPMethodImpl,
	},
	"gofakeitUserAgent": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitUserAgentImpl,
	},
	"gofakeitChromeUserAgent": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitChromeUserAgentImpl,
	},
	"gofakeitFirefoxUserAgent": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitFirefoxUserAgentImpl,
	},
	"gofakeitOperaUserAgent": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitOperaUserAgentImpl,
	},
	"gofakeitSafariUserAgent": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitSafariUserAgentImpl,
	},
	"gofakeitDate": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitDateImpl,
	},
	"gofakeitDateRange": {
		Params: []function.Parameter{
			{
				Name:             "separator",
				Type:             cty.String,
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.String), Impl: gofakeitDateRangeImpl,
	},
	"gofakeitNanoSecond": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number),
		Impl: gofakeitNanoSecondImpl,
	},
	"gofakeitSecond": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number),
		Impl: gofakeitSecondImpl,
	},
	"gofakeitMinute": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number),
		Impl: gofakeitMinuteImpl,
	},
	"gofakeitHour": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number),
		Impl: gofakeitHourImpl,
	},
	"gofakeitMonth": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitMonthImpl,
	},
	"gofakeitDay": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number),
		Impl: gofakeitDayImpl,
	},
	"gofakeitWeekDay": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number),
		Impl: gofakeitWeekDayImpl,
	},
	"gofakeitYear": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number),
		Impl: gofakeitYearImpl,
	},
	"gofakeitTimeZone": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitTimeZoneImpl,
	},
	"gofakeitTimeZoneAbv": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitTimeZoneAbvImpl,
	},
	"gofakeitTimeZoneFull": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitTimeZoneFullImpl,
	},
	"gofakeitTimeZoneOffset": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitTimeZoneOffsetImpl,
	},
	"gofakeitTimeZoneRegion": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitTimeZoneRegionImpl,
	},
	"gofakeitPrice": {
		Params: []function.Parameter{
			{
				Name:             "min",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
			{
				Name:             "max",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitPriceImpl,
	},
	"gofakeitCreditCardCvv": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitCreditCardCvvImpl,
	},
	"gofakeitCreditCardExp": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitCreditCardExpImpl,
	},
	"gofakeitCreditCardNumber": {
		Params: []function.Parameter{
			{
				Name:             "types",
				Type:             cty.List(cty.String),
				AllowDynamicType: false,
			},
			{
				Name:             "bins",
				Type:             cty.List(cty.String),
				AllowDynamicType: false,
			},
			{
				Name:             "gaps",
				Type:             cty.Bool,
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.String), Impl: gofakeitCreditCardNumberImpl,
	},
	"gofakeitCreditCardType": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitCreditCardTypeImpl,
	},
	"gofakeitCurrencyLong": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitCurrencyLongImpl,
	},
	"gofakeitCurrencyShort": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitCurrencyShortImpl,
	},
	"gofakeitAchRouting": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitAchRoutingImpl,
	},
	"gofakeitAchAccount": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitAchAccountImpl,
	},
	"gofakeitBitcoinAddress": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitBitcoinAddressImpl,
	},
	"gofakeitBitcoinPrivateKey": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitBitcoinPrivateKeyImpl,
	},
	"gofakeitBS": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitBSImpl,
	},
	"gofakeitBuzzWord": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitBuzzWordImpl,
	},
	"gofakeitCompany": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitCompanyImpl,
	},
	"gofakeitCompanySuffix": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitCompanySuffixImpl,
	},
	"gofakeitJobDescriptor": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitJobDescriptorImpl,
	},
	"gofakeitJobLevel": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitJobLevelImpl,
	},
	"gofakeitJobTitle": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitJobTitleImpl,
	},
	"gofakeitHackerAbbreviation": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitHackerAbbreviationImpl,
	},
	"gofakeitHackerAdjective": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitHackerAdjectiveImpl,
	},
	"gofakeitHackeringVerb": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitHackeringVerbImpl,
	},
	"gofakeitHackerNoun": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitHackerNounImpl,
	},
	"gofakeitHackerPhrase": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitHackerPhraseImpl,
	},
	"gofakeitHackerVerb": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitHackerVerbImpl,
	},
	"gofakeitHipsterWord": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitHipsterWordImpl,
	},
	"gofakeitHipsterSentence": {
		Params: []function.Parameter{
			{
				Name:             "wordCount",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.String), Impl: gofakeitHipsterSentenceImpl,
	},
	"gofakeitHipsterParagraph": {
		Params: []function.Parameter{
			{
				Name:             "paragraphCount",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
			{
				Name:             "sentenceCount",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
			{
				Name:             "wordCount",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
			{
				Name:             "separator",
				Type:             cty.String,
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.String), Impl: gofakeitHipsterParagraphImpl,
	},
	"gofakeitAppName": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitAppNameImpl,
	},
	"gofakeitAppVersion": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitAppVersionImpl,
	},
	"gofakeitAppAuthor": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitAppAuthorImpl,
	},
	"gofakeitPetName": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitPetNameImpl,
	},
	"gofakeitAnimal": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitAnimalImpl,
	},
	"gofakeitAnimalType": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitAnimalTypeImpl,
	},
	"gofakeitFarmAnimal": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitFarmAnimalImpl,
	},
	"gofakeitCat": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitCatImpl,
	},
	"gofakeitDog": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitDogImpl,
	},
	"gofakeitEmoji": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitEmojiImpl,
	},
	"gofakeitEmojiDescription": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitEmojiDescriptionImpl,
	},
	"gofakeitEmojiCategory": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitEmojiCategoryImpl,
	},
	"gofakeitEmojiAlias": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitEmojiAliasImpl,
	},
	"gofakeitEmojiTag": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitEmojiTagImpl,
	},
	"gofakeitLanguage": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitLanguageImpl,
	},
	"gofakeitLanguageAbbreviation": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitLanguageAbbreviationImpl,
	},
	"gofakeitProgrammingLanguage": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitProgrammingLanguageImpl,
	},
	"gofakeitProgrammingLanguageBest": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitProgrammingLanguageBestImpl,
	},
	"gofakeitNumber": {
		Params: []function.Parameter{
			{
				Name:             "min",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
			{
				Name:             "max",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitNumberImpl,
	},
	"gofakeitInt8": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number),
		Impl: gofakeitInt8Impl,
	},
	"gofakeitInt16": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number),
		Impl: gofakeitInt16Impl,
	},
	"gofakeitInt32": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number),
		Impl: gofakeitInt32Impl,
	},
	"gofakeitInt64": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number),
		Impl: gofakeitInt64Impl,
	},
	"gofakeitUint8": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number),
		Impl: gofakeitUint8Impl,
	},
	"gofakeitUint16": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number),
		Impl: gofakeitUint16Impl,
	},
	"gofakeitUint32": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number),
		Impl: gofakeitUint32Impl,
	},
	"gofakeitUint64": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number),
		Impl: gofakeitUint64Impl,
	},
	"gofakeitFloat32": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number),
		Impl: gofakeitFloat32Impl,
	},
	"gofakeitFloat32Range": {
		Params: []function.Parameter{
			{
				Name:             "min",
				Type:             cty.Number,
				AllowDynamicType: false,
			}, {
				Name:             "max",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitFloat32RangeImpl,
	},
	"gofakeitFloat64": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number),
		Impl: gofakeitFloat64Impl,
	},
	"gofakeitFloat64Range": {
		Params: []function.Parameter{
			{
				Name:             "min",
				Type:             cty.Number,
				AllowDynamicType: false,
			}, {
				Name:             "max",
				Type:             cty.Number,
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitFloat64RangeImpl,
	},
	"gofakeitShuffleInts": {
		Params: []function.Parameter{
			{
				Name:             "deck",
				Type:             cty.List(cty.Number),
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.List(cty.Number)), Impl: gofakeitShuffleIntsImpl,
	},
	"gofakeitRandomInt": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number),
		Impl: gofakeitRandomIntImpl,
	},
	"gofakeitDigit": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitDigitImpl,
	},
	"gofakeitLetter": {
		Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String),
		Impl: gofakeitLetterImpl,
	},
	"gofakeitLexify": {
		Params: []function.Parameter{
			{
				Name:             "str",
				Type:             cty.String,
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.String), Impl: gofakeitLexifyImpl,
	},
	"gofakeitNumerify": {
		Params: []function.Parameter{
			{
				Name:             "str",
				Type:             cty.String,
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.String), Impl: gofakeitNumerifyImpl,
	},
	"gofakeitShuffleStrings": {
		Params: []function.Parameter{
			{
				Name:             "deck",
				Type:             cty.List(cty.String),
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.List(cty.String)), Impl: gofakeitShuffleStringsImpl,
	},
	"gofakeitRandomString": {
		Params: []function.Parameter{
			{
				Name:             "list",
				Type:             cty.List(cty.String),
				AllowDynamicType: false,
			},
		}, Type: function.StaticReturnType(cty.String), Impl: gofakeitRandomStringImpl,
	},
}

var restbeastFunctionList = map[string]function.Function{
	"readfile": function.New(
		&function.Spec{
			Params: []function.Parameter{
				{
					Name:             "file",
					Type:             cty.String,
					AllowNull:        false,
					AllowUnknown:     false,
					AllowDynamicType: false,
					AllowMarked:      false,
				},
			},
			Type: function.StaticReturnType(cty.String),
			Impl: restbeastReadFileImpl,
		},
	),
	"read_file": function.New(
		&function.Spec{
			Params: []function.Parameter{
				{
					Name:             "file",
					Type:             cty.String,
					AllowNull:        false,
					AllowUnknown:     false,
					AllowDynamicType: false,
					AllowMarked:      false,
				},
			},
			Type: function.StaticReturnType(cty.String),
			Impl: restbeastReadFileImpl,
		},
	),
	"read_file_part": function.New(
		&function.Spec{
			Params: []function.Parameter{
				{
					Name:             "file",
					Type:             cty.String,
					AllowNull:        false,
					AllowUnknown:     false,
					AllowDynamicType: false,
					AllowMarked:      false,
				},
				{
					Name:             "offset",
					Type:             cty.Number,
					AllowNull:        false,
					AllowUnknown:     false,
					AllowDynamicType: false,
					AllowMarked:      false,
				},
				{
					Name:             "length",
					Type:             cty.Number,
					AllowNull:        false,
					AllowUnknown:     false,
					AllowDynamicType: false,
					AllowMarked:      false,
				},
			},
			Type: function.StaticReturnType(cty.String),
			Impl: restbeastFilePartImpl,
		},
	),
	"fill_null": function.New(
		&function.Spec{
			Params: []function.Parameter{
				{
					Name:             "probability",
					Type:             cty.Number,
					AllowNull:        false,
					AllowUnknown:     false,
					AllowDynamicType: false,
					AllowMarked:      false,
				},
				{
					Name:             "nonNullValue",
					Type:             cty.DynamicPseudoType,
					AllowNull:        false,
					AllowUnknown:     false,
					AllowDynamicType: true,
					AllowMarked:      false,
				},
			},
			Type: function.StaticReturnType(cty.DynamicPseudoType),
			Impl: restbeastFillNullImpl,
		},
	),
	"env_var": function.New(
		&function.Spec{
			Params: []function.Parameter{
				{
					Name:             "key",
					Type:             cty.String,
					AllowNull:        false,
					AllowUnknown:     false,
					AllowDynamicType: false,
					AllowMarked:      false,
				},
			},
			Type: function.StaticReturnType(cty.String),
			Impl: restbeastEnvVarImpl,
		},
	),
	"env_var_with_default": function.New(
		&function.Spec{
			Params: []function.Parameter{
				{
					Name:             "key",
					Type:             cty.String,
					AllowNull:        false,
					AllowUnknown:     false,
					AllowDynamicType: false,
					AllowMarked:      false,
				},
				{
					Name:             "default",
					Type:             cty.String,
					AllowNull:        false,
					AllowUnknown:     false,
					AllowDynamicType: false,
					AllowMarked:      false,
				},
			},
			Type: function.StaticReturnType(cty.String),
			Impl: restbeastEnvVarWithDefaultImpl,
		},
	),
	"unixTimestamp": function.New(
		&function.Spec{
			Params: []function.Parameter{
				{
					Name:             "timestamp",
					Type:             cty.String,
					AllowNull:        false,
					AllowUnknown:     false,
					AllowDynamicType: false,
					AllowMarked:      false,
				},
			},
			Type: function.StaticReturnType(cty.Number),
			Impl: restbeastUnixTimestampImpl,
		},
	),
	"now": function.New(
		&function.Spec{
			Type: function.StaticReturnType(cty.String),
			Impl: restbeastNowImpl,
		},
	),
}

func getCtyFunctions() map[string]function.Function {
	allFunctions := make(map[string]function.Function)

	for k, v := range gofakeitFunctionList {
		allFunctions[k] = function.New(
			&function.Spec{
				Params: v.Params,
				Type:   v.Type,
				Impl:   v.Impl,
			},
		)
	}

	for k, v := range defaultFunctions {
		allFunctions[k] = v
	}

	for k, v := range restbeastFunctionList {
		allFunctions[k] = v
	}

	return allFunctions
}
