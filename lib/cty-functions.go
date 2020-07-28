package lib

import (
	"github.com/zclconf/go-cty/cty/function"
	"github.com/zclconf/go-cty/cty/function/stdlib"
)

func getCtyFunctions() map[string]function.Function {
	return map[string]function.Function {
		// General Functions
		"equal": stdlib.EqualFunc,
		"notEqual": stdlib.NotEqualFunc,
		"coalesce": stdlib.CoalesceFunc,
		"not": stdlib.NotFunc,
		"and": stdlib.AndFunc,
		"or": stdlib.OrFunc,
		// String Functions
		"upper": stdlib.UpperFunc,
		"lower": stdlib.LowerFunc,
		"reverse": stdlib.ReverseFunc,
		"strlen": stdlib.StrlenFunc,
		"substr": stdlib.SubstrFunc,
		"join": stdlib.JoinFunc,
		"sort": stdlib.SortFunc,
		"split": stdlib.SplitFunc,
		"chomp": stdlib.ChompFunc,
		"indent": stdlib.IndentFunc,
		"title": stdlib.TitleFunc,
		"trimSpace": stdlib.TrimSpaceFunc,
		"trim": stdlib.TrimFunc,
		"trimPrefix": stdlib.TrimPrefixFunc,
		"trimSuffix": stdlib.TrimSuffixFunc,
		"format": stdlib.FormatFunc,
		"formatList": stdlib.FormatListFunc,
		"replace": stdlib.ReplaceFunc,
		// Number Functions
		"absolute": stdlib.AbsoluteFunc,
		"add": stdlib.AddFunc,
		"subtract": stdlib.SubtractFunc,
		"multiply": stdlib.MultiplyFunc,
		"divide": stdlib.DivideFunc,
		"modulo": stdlib.ModuloFunc,
		"greaterThan": stdlib.GreaterThanFunc,
		"greaterThanOrEqualTo": stdlib.GreaterThanOrEqualToFunc,
		"lessThan": stdlib.LessThanFunc,
		"lessThanOrEqualTo": stdlib.LessThanOrEqualToFunc,
		"negate": stdlib.NegateFunc,
		"min": stdlib.MinFunc,
		"max": stdlib.MaxFunc,
		"int": stdlib.IntFunc,
		"ceil": stdlib.CeilFunc,
		"floor": stdlib.FloorFunc,
		"log": stdlib.LogFunc,
		"pow": stdlib.PowFunc,
		"signum": stdlib.SignumFunc,
		"parseInt": stdlib.ParseIntFunc,
		// Collection Functions
		"hasIndex": stdlib.HasIndexFunc,
		"index": stdlib.IndexFunc,
		"length": stdlib.LengthFunc,
		"element": stdlib.ElementFunc,
		"coalesceList": stdlib.CoalesceListFunc,
		"compact": stdlib.CompactFunc,
		"contains": stdlib.ContainsFunc,
		"distinct": stdlib.DistinctFunc,
		"chunklist": stdlib.ChunklistFunc,
		"flatten": stdlib.FlattenFunc,
		"keys": stdlib.KeysFunc,
		"lookup": stdlib.LookupFunc,
		"merge": stdlib.MergeFunc,
		"reverseList": stdlib.ReverseListFunc,
		"setProduct": stdlib.SetProductFunc,
		"slice": stdlib.SliceFunc,
		"values": stdlib.ValuesFunc,
		"zipmap": stdlib.ZipmapFunc,
		"concat": stdlib.ConcatFunc,
		"range": stdlib.RangeFunc,
		// Encoding Functions
		"csvDecode": stdlib.CSVDecodeFunc,
		"jsonEncode": stdlib.JSONEncodeFunc,
		"jsonDecode": stdlib.JSONDecodeFunc,
		// Datetime Functions
		"formatDate": stdlib.FormatDateFunc,
		"timeAdd": stdlib.TimeAddFunc,
		// Regex Functions
		"regex": stdlib.RegexFunc,
		"regexAll": stdlib.RegexAllFunc,
		"regexReplace": stdlib.RegexReplaceFunc,
		// Gofakeit Functions
		// Person functions
		"gofakeitName": gofakeitName(),
		"gofakeitNamePrefix": gofakeitNamePrefix(),
		"gofakeitNameSuffix": gofakeitNameSuffix(),
		"gofakeitFirstName": gofakeitFirstName(),
		"gofakeitLastName": gofakeitLastName(),
		"gofakeitGender": gofakeitGender(),
		"gofakeitSSN": gofakeitSSN(),
		"gofakeitEmail": gofakeitEmail(),
		"gofakeitPhone": gofakeitPhone(),
		"gofakeitPhoneFormatted": gofakeitPhoneFormatted(),
		// Auth functions
		"gofakeitUsername": gofakeitUsername(),
		"gofakeitPassword": gofakeitPassword(),
		// Address functions
		"gofakeitCity": gofakeitCity(),
		"gofakeitCountry": gofakeitCountry(),
		"gofakeitCountryAbr": gofakeitCountryAbr(),
		"gofakeitState": gofakeitState(),
		"gofakeitStateAbr": gofakeitStateAbr(),
		"gofakeitStreet": gofakeitStreet(),
		"gofakeitStreetName": gofakeitStreetName(),
		"gofakeitStreetNumber": gofakeitStreetNumber(),
		"gofakeitStreetPrefix": gofakeitStreetPrefix(),
		"gofakeitStreetSuffix": gofakeitStreetSuffix(),
		"gofakeitZip": gofakeitZip(),
		"gofakeitLatitude": gofakeitLatitude(),
		"gofakeitLatitudeInRange": gofakeitLatitudeInRange(),
		"gofakeitLongitude": gofakeitLongitude(),
		"gofakeitLongitudeInRange": gofakeitLongitudeInRange(),
		// Gamer functions
		"gofakeitGamertag": gofakeitGamertag(),
		// Booze functions
		"gofakeitBeerAlcohol": gofakeitBeerAlcohol(),
		"gofakeitBeerBlg": gofakeitBeerBlg(),
		"gofakeitBeerHop": gofakeitBeerHop(),
		"gofakeitBeerIbu": gofakeitBeerIbu(),
		"gofakeitBeerMalt": gofakeitBeerMalt(),
		"gofakeitBeerName": gofakeitBeerName(),
		"gofakeitBeerStyle": gofakeitBeerStyle(),
		"gofakeitBeerYeast": gofakeitBeerYeast(),
		// Car functions
		"gofakeitCarMaker": gofakeitCarMaker(),
		"gofakeitCarModel": gofakeitCarModel(),
		"gofakeitCarType": gofakeitCarType(),
		"gofakeitCarFuelType": gofakeitCarFuelType(),
		"gofakeitTransmissionGearType": gofakeitCarTransmissionType(),
		// Word functions
		"gofakeitNoun": gofakeitNoun(),
		"gofakeitVerb": gofakeitVerb(),
		"gofakeitAdverb": gofakeitAdverb(),
		"gofakeitPreposition": gofakeitPreposition(),
		"gofakeitAdjective": gofakeitAdjective(),
		"gofakeitWord": gofakeitWord(),
		"gofakeitSentence": gofakeitSentence(),
		"gofakeitParagraph": gofakeitParagraph(),
		"gofakeitLoremIpsumWord": gofakeitLoremIpsumWord(),
		"gofakeitLoremIpsumSentence": gofakeitLoremIpsumSentence(),
		"gofakeitLoremIpsumParagraph": gofakeitLoremIpsumParagraph(),
		"gofakeitQuestion": gofakeitQuestion(),
		"gofakeitQuote": gofakeitQuote(),
		"gofakeitPhrase": gofakeitPhrase(),
		// Food functions
		"gofakeitFruit": gofakeitFruit(),
		"gofakeitVegetable": gofakeitVegetable(),
		"gofakeitBreakfast": gofakeitBreakfast(),
		"gofakeitLunch": gofakeitLunch(),
		"gofakeitDinner": gofakeitDinner(),
		"gofakeitSnack": gofakeitSnack(),
		"gofakeitDessert": gofakeitDessert(),
		// Misc functions
		"gofakeitBool": gofakeitBool(),
		"gofakeitUUID": gofakeitUUID(),
		// Color functions
		"gofakeitColor": gofakeitColor(),
		"gofakeitHexColor": gofakeitHexColor(),
		"gofakeitRGBColor": gofakeitRGBColor(),
		"gofakeitSafeColor": gofakeitSafeColor(),
		// Internet functions
		"gofakeitURL": gofakeitURL(),
		"gofakeitImageURL": gofakeitImageURL(),
		"gofakeitDomainName": gofakeitDomainName(),
		"gofakeitDomainSuffix": gofakeitDomainSuffix(),
		"gofakeitIPv4Address": gofakeitIPv4Address(),
		"gofakeitIPv6Address": gofakeitIPv6Address(),
		"gofakeitHTTPStatusCode": gofakeitHTTPStatusCode(),
		"gofakeitHTTPSimpleStatusCode": gofakeitHTTPSimpleStatusCode(),
		"gofakeitLogLevel": gofakeitLogLevel(),
		"gofakeitHTTPMethod": gofakeitHTTPMethod(),
		"gofakeitUserAgent": gofakeitUserAgent(),
		"gofakeitChromeUserAgent": gofakeitChromeUserAgent(),
		"gofakeitFirefoxUserAgent": gofakeitFirefoxUserAgent(),
		"gofakeitOperaUserAgent": gofakeitOperaUserAgent(),
		"gofakeitSafariUserAgent": gofakeitSafariUserAgent(),
		// Date/Time functions
		"gofakeitDate": gofakeitDate(),
		"gofakeitDateRange": gofakeitDateRange(),
		"gofakeitNanoSecond": gofakeitNanoSecond(),
		"gofakeitSecond": gofakeitSecond(),
		"gofakeitMinute": gofakeitMinute(),
		"gofakeitHour": gofakeitHour(),
		"gofakeitMonth": gofakeitMonth(),
		"gofakeitDay": gofakeitDay(),
		"gofakeitWeekDay": gofakeitWeekDay(),
		"gofakeitYear": gofakeitYear(),
		"gofakeitTimeZone": gofakeitTimeZone(),
		"gofakeitTimeZoneAbv": gofakeitTimeZoneAbv(),
		"gofakeitTimeZoneFull": gofakeitTimeZoneFull(),
		"gofakeitTimeZoneOffset": gofakeitTimeZoneOffset(),
		"gofakeitTimeZoneRegion": gofakeitTimeZoneRegion(),
		"gofakeitPrice": gofakeitPrice(),
		// Payment functions
		"gofakeitCreditCardCvv": gofakeitCreditCardCvv(),
		"gofakeitCreditCardExp": gofakeitCreditCardExp(),
		"gofakeitCreditCardNumber": gofakeitCreditCardNumber(),
		"gofakeitCreditCardType": gofakeitCreditCardType(),
		"gofakeitCurrencyLong": gofakeitCurrencyLong(),
		"gofakeitCurrencyShort": gofakeitCurrencyShort(),
		"gofakeitAchRouting": gofakeitAchRouting(),
		"gofakeitAchAccount": gofakeitAchAccount(),
		"gofakeitBitcoinAddress": gofakeitBitcoinAddress(),
		"gofakeitBitcoinPrivateKey": gofakeitBitcoinPrivateKey(),
		// Company functions
		"gofakeitBS": gofakeitBS(),
		"gofakeitBuzzWord": gofakeitBuzzWord(),
		"gofakeitCompany": gofakeitCompany(),
		"gofakeitCompanySuffix": gofakeitCompanySuffix(),
		"gofakeitJobDescriptor": gofakeitJobDescriptor(),
		"gofakeitJobLevel": gofakeitJobLevel(),
		"gofakeitJobTitle": gofakeitJobTitle(),
		// Hacker functions
		"gofakeitHackerAbbreviation": gofakeitHackerAbbreviation(),
		"gofakeitHackerAdjective": gofakeitHackerAdjective(),
		"gofakeitHackeringVerb": gofakeitHackeringVerb(),
		"gofakeitHackerNoun": gofakeitHackerNoun(),
		"gofakeitHackerPhrase": gofakeitHackerPhrase(),
		"gofakeitHackerVerb": gofakeitHackerVerb(),
		// Hipster functions
		"gofakeitHipsterWord": gofakeitHipsterWord(),
		"gofakeitHipsterSentence": gofakeitHipsterSentence(),
		"gofakeitHipsterParagraph": gofakeitHipsterParagraph(),
		// App functions
		"gofakeitAppName": gofakeitAppName(),
		"gofakeitAppVersion": gofakeitAppVersion(),
		"gofakeitAppAuthor": gofakeitAppAuthor(),
		// Animal functions
		"gofakeitPetName": gofakeitPetName(),
		"gofakeitAnimal": gofakeitAnimal(),
		"gofakeitAnimalType": gofakeitAnimalType(),
		"gofakeitFarmAnimal": gofakeitFarmAnimal(),
		"gofakeitCat": gofakeitCat(),
		"gofakeitDog": gofakeitDog(),
		// Emoji functions
		"gofakeitEmoji": gofakeitEmoji(),
		"gofakeitEmojiDescription": gofakeitEmojiDescription(),
		"gofakeitEmojiCategory": gofakeitEmojiCategory(),
		"gofakeitEmojiAlias": gofakeitEmojiAlias(),
		"gofakeitEmojiTag": gofakeitEmojiTag(),
		// Language functions
		"gofakeitLanguage": gofakeitLanguage(),
		"gofakeitLanguageAbbreviation": gofakeitLanguageAbbreviation(),
		"gofakeitProgrammingLanguage": gofakeitProgrammingLanguage(),
		"gofakeitProgrammingLanguageBest": gofakeitProgrammingLanguageBest(),
		// Number functions
		"gofakeitNumber": gofakeitNumber(),
		"gofakeitInt8": gofakeitInt8(),
		"gofakeitInt16": gofakeitInt16(),
		"gofakeitInt32": gofakeitInt32(),
		"gofakeitInt64": gofakeitInt64(),
		"gofakeitUint8": gofakeitUint8(),
		"gofakeitUint16": gofakeitUint16(),
		"gofakeitUint32": gofakeitUint32(),
		"gofakeitUint64": gofakeitUint64(),
		"gofakeitFloat32": gofakeitFloat32(),
		"gofakeitFloat32Range": gofakeitFloat32Range(),
		"gofakeitFloat64": gofakeitFloat64(),
		"gofakeitFloat64Range": gofakeitFloat64Range(),
		"gofakeitShuffleInts": gofakeitShuffleInts(),
		"gofakeitRandomInt": gofakeitRandomInt(),
		"gofakeitDigit": gofakeitDigit(),
		// String functions
		"gofakeitLetter": gofakeitLetter(),
		"gofakeitLexify": gofakeitLexify(),
		"gofakeitNumerify": gofakeitNumerify(),
		"gofakeitShuffleStrings": gofakeitShuffleStrings(),
		"gofakeitRandomString": gofakeitRandomString(),
	}
}