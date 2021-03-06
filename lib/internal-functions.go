package lib

import (
	"fmt"
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
	"gofakeitName":           GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitNameImpl},
	"gofakeitNamePrefix":     GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitNamePrefixImpl},
	"gofakeitNameSuffix":     GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitNameSuffixImpl},
	"gofakeitFirstName":      GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitFirstNameImpl},
	"gofakeitLastName":       GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitLastNameImpl},
	"gofakeitGender":         GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitGenderImpl},
	"gofakeitSSN":            GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitSSNImpl},
	"gofakeitEmail":          GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitEmailImpl},
	"gofakeitPhone":          GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitPhoneImpl},
	"gofakeitPhoneFormatted": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitPhoneFormattedImpl},
	"gofakeitUsername":       GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitUsernameImpl},
	"gofakeitPassword": GofakeitFunc{Params: []function.Parameter{
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
	}, Type: function.StaticReturnType(cty.String), Impl: gofakeitPasswordImpl},
	"gofakeitCity":         GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitCityImpl},
	"gofakeitCountry":      GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitCountryImpl},
	"gofakeitCountryAbr":   GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitCountryAbrImpl},
	"gofakeitState":        GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitStateImpl},
	"gofakeitStateAbr":     GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitStateAbrImpl},
	"gofakeitStreet":       GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitStreetImpl},
	"gofakeitStreetName":   GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitStreetNameImpl},
	"gofakeitStreetNumber": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitStreetNumberImpl},
	"gofakeitStreetPrefix": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitStreetPrefixImpl},
	"gofakeitStreetSuffix": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitStreetSuffixImpl},
	"gofakeitZip":          GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitZipImpl},
	"gofakeitLatitude":     GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitLatitudeImpl},
	"gofakeitLatitudeInRange": GofakeitFunc{Params: []function.Parameter{
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
	}, Type: function.StaticReturnType(cty.String), Impl: gofakeitLatitudeInRangeImpl},
	"gofakeitLongitude": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitLongitudeImpl},
	"gofakeitLongitudeInRange": GofakeitFunc{Params: []function.Parameter{
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
	}, Type: function.StaticReturnType(cty.String), Impl: gofakeitLongitudeInRangeImpl},
	"gofakeitGamertag":    GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitGamertagImpl},
	"gofakeitBeerAlcohol": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitBeerAlcoholImpl},
	"gofakeitBeerBlg":     GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitBeerBlgImpl},
	"gofakeitBeerHop":     GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitBeerHopImpl},
	"gofakeitBeerIbu":     GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitBeerIbuImpl},
	"gofakeitBeerMalt":    GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitBeerMaltImpl},
	"gofakeitBeerName":    GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitBeerNameImpl},
	"gofakeitBeerStyle":   GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitBeerStyleImpl},
	"gofakeitBeerYeast":   GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitBeerYeastImpl},
	"gofakeitCarMaker":    GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitCarMakerImpl},
	"gofakeitCarModel":    GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitCarModelImpl},
	"gofakeitCarType":     GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitCarTypeImpl},
	"gofakeitCarFuelType": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitCarFuelTypeImpl},
	"gofakeitNoun":        GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitNounImpl},
	"gofakeitVerb":        GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitVerbImpl},
	"gofakeitAdverb":      GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitAdverbImpl},
	"gofakeitPreposition": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitPrepositionImpl},
	"gofakeitAdjective":   GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitAdjectiveImpl},
	"gofakeitWord":        GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitWordImpl},
	"gofakeitSentence": GofakeitFunc{Params: []function.Parameter{{
		Name:             "wordCount",
		Type:             cty.Number,
		AllowDynamicType: false,
	}}, Type: function.StaticReturnType(cty.String), Impl: gofakeitSentenceImpl},
	"gofakeitParagraph": GofakeitFunc{Params: []function.Parameter{
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
		}}, Type: function.StaticReturnType(cty.String), Impl: gofakeitParagraphImpl},
	"gofakeitLoremIpsumWord": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitLoremIpsumWordImpl},
	"gofakeitLoremIpsumSentence": GofakeitFunc{Params: []function.Parameter{{
		Name:             "wordCount",
		Type:             cty.Number,
		AllowDynamicType: false,
	}}, Type: function.StaticReturnType(cty.String), Impl: gofakeitLoremIpsumSentenceImpl},
	"gofakeitLoremIpsumParagraph": GofakeitFunc{Params: []function.Parameter{{
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
		}}, Type: function.StaticReturnType(cty.String), Impl: gofakeitLoremIpsumParagraphImpl},
	"gofakeitQuestion":  GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitQuestionImpl},
	"gofakeitQuote":     GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitQuoteImpl},
	"gofakeitPhrase":    GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitPhraseImpl},
	"gofakeitFruit":     GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitFruitImpl},
	"gofakeitVegetable": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitVegetableImpl},
	"gofakeitBreakfast": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitBreakfastImpl},
	"gofakeitLunch":     GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitLunchImpl},
	"gofakeitDinner":    GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitDinnerImpl},
	"gofakeitSnack":     GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitSnackImpl},
	"gofakeitDessert":   GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitDessertImpl},
	"gofakeitBool":      GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitBoolImpl},
	"gofakeitUUID":      GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitUUIDImpl},
	"gofakeitColor":     GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitColorImpl},
	"gofakeitHexColor":  GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitHexColorImpl},
	"gofakeitRGBColor":  GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitRGBColorImpl},
	"gofakeitSafeColor": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitSafeColorImpl},
	"gofakeitURL":       GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitURLImpl},
	"gofakeitImageURL": GofakeitFunc{Params: []function.Parameter{{
		Name:             "width",
		Type:             cty.Number,
		AllowDynamicType: false,
	},
		{
			Name:             "height",
			Type:             cty.Number,
			AllowDynamicType: false,
		}}, Type: function.StaticReturnType(cty.String), Impl: gofakeitImageURLImpl},
	"gofakeitDomainName":           GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitDomainNameImpl},
	"gofakeitDomainSuffix":         GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitDomainSuffixImpl},
	"gofakeitIPv4Address":          GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitIPv4AddressImpl},
	"gofakeitIPv6Address":          GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitIPv6AddressImpl},
	"gofakeitHTTPStatusCode":       GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitHTTPStatusCodeImpl},
	"gofakeitHTTPSimpleStatusCode": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitHTTPSimpleStatusCodeImpl},
	"gofakeitLogLevel": GofakeitFunc{Params: []function.Parameter{{
		Name:             "logType",
		Type:             cty.String,
		AllowDynamicType: false,
	}}, Type: function.StaticReturnType(cty.String), Impl: gofakeitLogLevelImpl},
	"gofakeitHTTPMethod":       GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitHTTPMethodImpl},
	"gofakeitUserAgent":        GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitUserAgentImpl},
	"gofakeitChromeUserAgent":  GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitChromeUserAgentImpl},
	"gofakeitFirefoxUserAgent": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitFirefoxUserAgentImpl},
	"gofakeitOperaUserAgent":   GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitOperaUserAgentImpl},
	"gofakeitSafariUserAgent":  GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitSafariUserAgentImpl},
	"gofakeitDate":             GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitDateImpl},
	"gofakeitDateRange": GofakeitFunc{Params: []function.Parameter{{
		Name:             "separator",
		Type:             cty.String,
		AllowDynamicType: false,
	}}, Type: function.StaticReturnType(cty.String), Impl: gofakeitDateRangeImpl},
	"gofakeitNanoSecond":     GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitNanoSecondImpl},
	"gofakeitSecond":         GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitSecondImpl},
	"gofakeitMinute":         GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitMinuteImpl},
	"gofakeitHour":           GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitHourImpl},
	"gofakeitMonth":          GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitMonthImpl},
	"gofakeitDay":            GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitDayImpl},
	"gofakeitWeekDay":        GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitWeekDayImpl},
	"gofakeitYear":           GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitYearImpl},
	"gofakeitTimeZone":       GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitTimeZoneImpl},
	"gofakeitTimeZoneAbv":    GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitTimeZoneAbvImpl},
	"gofakeitTimeZoneFull":   GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitTimeZoneFullImpl},
	"gofakeitTimeZoneOffset": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitTimeZoneOffsetImpl},
	"gofakeitTimeZoneRegion": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitTimeZoneRegionImpl},
	"gofakeitPrice": GofakeitFunc{Params: []function.Parameter{{
		Name:             "min",
		Type:             cty.Number,
		AllowDynamicType: false,
	},
		{
			Name:             "max",
			Type:             cty.Number,
			AllowDynamicType: false,
		}}, Type: function.StaticReturnType(cty.String), Impl: gofakeitPriceImpl},
	"gofakeitCreditCardCvv": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitCreditCardCvvImpl},
	"gofakeitCreditCardExp": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitCreditCardExpImpl},
	"gofakeitCreditCardNumber": GofakeitFunc{Params: []function.Parameter{{
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
		}}, Type: function.StaticReturnType(cty.String), Impl: gofakeitCreditCardNumberImpl},
	"gofakeitCreditCardType":     GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitCreditCardTypeImpl},
	"gofakeitCurrencyLong":       GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitCurrencyLongImpl},
	"gofakeitCurrencyShort":      GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitCurrencyShortImpl},
	"gofakeitAchRouting":         GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitAchRoutingImpl},
	"gofakeitAchAccount":         GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitAchAccountImpl},
	"gofakeitBitcoinAddress":     GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitBitcoinAddressImpl},
	"gofakeitBitcoinPrivateKey":  GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitBitcoinPrivateKeyImpl},
	"gofakeitBS":                 GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitBSImpl},
	"gofakeitBuzzWord":           GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitBuzzWordImpl},
	"gofakeitCompany":            GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitCompanyImpl},
	"gofakeitCompanySuffix":      GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitCompanySuffixImpl},
	"gofakeitJobDescriptor":      GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitJobDescriptorImpl},
	"gofakeitJobLevel":           GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitJobLevelImpl},
	"gofakeitJobTitle":           GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitJobTitleImpl},
	"gofakeitHackerAbbreviation": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitHackerAbbreviationImpl},
	"gofakeitHackerAdjective":    GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitHackerAdjectiveImpl},
	"gofakeitHackeringVerb":      GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitHackeringVerbImpl},
	"gofakeitHackerNoun":         GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitHackerNounImpl},
	"gofakeitHackerPhrase":       GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitHackerPhraseImpl},
	"gofakeitHackerVerb":         GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitHackerVerbImpl},
	"gofakeitHipsterWord":        GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitHipsterWordImpl},
	"gofakeitHipsterSentence": GofakeitFunc{Params: []function.Parameter{{
		Name:             "wordCount",
		Type:             cty.Number,
		AllowDynamicType: false,
	}}, Type: function.StaticReturnType(cty.String), Impl: gofakeitHipsterSentenceImpl},
	"gofakeitHipsterParagraph": GofakeitFunc{Params: []function.Parameter{{
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
		}}, Type: function.StaticReturnType(cty.String), Impl: gofakeitHipsterParagraphImpl},
	"gofakeitAppName":                 GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitAppNameImpl},
	"gofakeitAppVersion":              GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitAppVersionImpl},
	"gofakeitAppAuthor":               GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitAppAuthorImpl},
	"gofakeitPetName":                 GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitPetNameImpl},
	"gofakeitAnimal":                  GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitAnimalImpl},
	"gofakeitAnimalType":              GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitAnimalTypeImpl},
	"gofakeitFarmAnimal":              GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitFarmAnimalImpl},
	"gofakeitCat":                     GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitCatImpl},
	"gofakeitDog":                     GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitDogImpl},
	"gofakeitEmoji":                   GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitEmojiImpl},
	"gofakeitEmojiDescription":        GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitEmojiDescriptionImpl},
	"gofakeitEmojiCategory":           GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitEmojiCategoryImpl},
	"gofakeitEmojiAlias":              GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitEmojiAliasImpl},
	"gofakeitEmojiTag":                GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitEmojiTagImpl},
	"gofakeitLanguage":                GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitLanguageImpl},
	"gofakeitLanguageAbbreviation":    GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitLanguageAbbreviationImpl},
	"gofakeitProgrammingLanguage":     GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitProgrammingLanguageImpl},
	"gofakeitProgrammingLanguageBest": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitProgrammingLanguageBestImpl},
	"gofakeitNumber": GofakeitFunc{Params: []function.Parameter{{
		Name:             "min",
		Type:             cty.Number,
		AllowDynamicType: false,
	},
		{
			Name:             "max",
			Type:             cty.Number,
			AllowDynamicType: false,
		}}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitNumberImpl},
	"gofakeitInt8":    GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitInt8Impl},
	"gofakeitInt16":   GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitInt16Impl},
	"gofakeitInt32":   GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitInt32Impl},
	"gofakeitInt64":   GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitInt64Impl},
	"gofakeitUint8":   GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitUint8Impl},
	"gofakeitUint16":  GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitUint16Impl},
	"gofakeitUint32":  GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitUint32Impl},
	"gofakeitUint64":  GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitUint64Impl},
	"gofakeitFloat32": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitFloat32Impl},
	"gofakeitFloat32Range": GofakeitFunc{Params: []function.Parameter{{
		Name:             "min",
		Type:             cty.Number,
		AllowDynamicType: false,
	}, {
		Name:             "max",
		Type:             cty.Number,
		AllowDynamicType: false,
	}}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitFloat32RangeImpl},
	"gofakeitFloat64": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitFloat64Impl},
	"gofakeitFloat64Range": GofakeitFunc{Params: []function.Parameter{{
		Name:             "min",
		Type:             cty.Number,
		AllowDynamicType: false,
	}, {
		Name:             "max",
		Type:             cty.Number,
		AllowDynamicType: false,
	}}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitFloat64RangeImpl},
	"gofakeitShuffleInts": GofakeitFunc{Params: []function.Parameter{{
		Name:             "deck",
		Type:             cty.List(cty.Number),
		AllowDynamicType: false,
	}}, Type: function.StaticReturnType(cty.List(cty.Number)), Impl: gofakeitShuffleIntsImpl},
	"gofakeitRandomInt": GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.Number), Impl: gofakeitRandomIntImpl},
	"gofakeitDigit":     GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitDigitImpl},
	"gofakeitLetter":    GofakeitFunc{Params: []function.Parameter{}, Type: function.StaticReturnType(cty.String), Impl: gofakeitLetterImpl},
	"gofakeitLexify": GofakeitFunc{Params: []function.Parameter{{
		Name:             "str",
		Type:             cty.String,
		AllowDynamicType: false,
	}}, Type: function.StaticReturnType(cty.String), Impl: gofakeitLexifyImpl},
	"gofakeitNumerify": GofakeitFunc{Params: []function.Parameter{{
		Name:             "str",
		Type:             cty.String,
		AllowDynamicType: false,
	}}, Type: function.StaticReturnType(cty.String), Impl: gofakeitNumerifyImpl},
	"gofakeitShuffleStrings": GofakeitFunc{Params: []function.Parameter{{
		Name:             "deck",
		Type:             cty.List(cty.String),
		AllowDynamicType: false,
	}}, Type: function.StaticReturnType(cty.List(cty.String)), Impl: gofakeitShuffleStringsImpl},
	"gofakeitRandomString": GofakeitFunc{Params: []function.Parameter{{
		Name:             "list",
		Type:             cty.List(cty.String),
		AllowDynamicType: false,
	}}, Type: function.StaticReturnType(cty.String), Impl: gofakeitRandomStringImpl},
}

var restbeastFunctionList = map[string]function.Function{
	"readfile": function.New(&function.Spec{
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
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			if len(args) < 1 {
				return cty.StringVal(""), fmt.Errorf("Invalid argument count")
			}

			return cty.StringVal(fmt.Sprintf("###READFILE=%s###", args[0].AsString())), nil
		},
	}),
}

func getCtyFunctions() map[string]function.Function {
	allFunctions := make(map[string]function.Function)

	for k, v := range gofakeitFunctionList {
		allFunctions[k] = function.New(&function.Spec{
			Params: v.Params,
			Type:   v.Type,
			Impl:   v.Impl,
		})
	}

	for k, v := range defaultFunctions {
		allFunctions[k] = v
	}

	for k, v := range restbeastFunctionList {
		allFunctions[k] = v
	}

	return allFunctions
}
