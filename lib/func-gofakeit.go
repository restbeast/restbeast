package lib

import (
	. "fmt"
	"github.com/brianvoe/gofakeit/v5"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	"time"
)

func gofakeitNameImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Name()), nil
}

func gofakeitNamePrefixImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.NamePrefix()), nil
}

func gofakeitNameSuffixImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.NamePrefix()), nil
}

func gofakeitFirstNameImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.FirstName()), nil
}

func gofakeitLastNameImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.LastName()), nil
}

func gofakeitGenderImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Gender()), nil
}

func gofakeitSSNImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.SSN()), nil
}

func gofakeitEmailImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Email()), nil
}

func gofakeitPhoneImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Phone()), nil
}

func gofakeitPhoneFormattedImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.PhoneFormatted()), nil
}

func gofakeitUsernameImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Username()), nil
}

func gofakeitPasswordImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)

	var size int
	_ = gocty.FromCtyValue(args[5], &size)
	return cty.StringVal(gofakeit.Password(args[0].True(), args[1].True(), args[2].True(), args[3].True(), args[4].True(), size)), nil
}

func gofakeitCityImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.City()), nil
}

func gofakeitCountryImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Country()), nil
}

func gofakeitCountryAbrImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.CountryAbr()), nil
}

func gofakeitStateImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.State()), nil
}

func gofakeitStateAbrImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.StateAbr()), nil
}

func gofakeitStreetImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Street()), nil
}

func gofakeitStreetNameImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.StreetName()), nil
}

func gofakeitStreetNumberImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.StreetNumber()), nil
}

func gofakeitStreetPrefixImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.StreetPrefix()), nil
}

func gofakeitStreetSuffixImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.StreetSuffix()), nil
}
func gofakeitZipImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Zip()), nil
}

func gofakeitLatitudeImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.NumberFloatVal(gofakeit.Latitude()), nil
}

func gofakeitLatitudeInRangeImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	min, _ := args[0].AsBigFloat().Float64()
	max, _ := args[1].AsBigFloat().Float64()
	inRange, err := gofakeit.LatitudeInRange(min, max)

	return cty.NumberFloatVal(inRange), err
}

func gofakeitLongitudeImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.NumberFloatVal(gofakeit.Longitude()), nil
}

func gofakeitLongitudeInRangeImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	min, _ := args[0].AsBigFloat().Float64()
	max, _ := args[1].AsBigFloat().Float64()
	inRange, err := gofakeit.LongitudeInRange(min, max)

	return cty.NumberFloatVal(inRange), err
}

func gofakeitGamertagImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Gamertag()), nil
}

func gofakeitBeerAlcoholImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.BeerAlcohol()), nil
}

func gofakeitBeerBlgImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.BeerBlg()), nil
}

func gofakeitBeerHopImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.BeerHop()), nil
}

func gofakeitBeerIbuImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.BeerIbu()), nil
}

func gofakeitBeerMaltImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.BeerMalt()), nil
}

func gofakeitBeerNameImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.BeerName()), nil
}

func gofakeitBeerStyleImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.BeerStyle()), nil
}

func gofakeitBeerYeastImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.BeerYeast()), nil
}

func gofakeitCarMakerImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.CarMaker()), nil
}

func gofakeitCarModelImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.CarModel()), nil
}

func gofakeitCarTypeImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.CarType()), nil
}

func gofakeitCarFuelTypeImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.CarFuelType()), nil
}

func gofakeitCarTransmissionTypeImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.CarTransmissionType()), nil
}

func gofakeitNounImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Noun()), nil
}

func gofakeitVerbImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Verb()), nil
}

func gofakeitAdverbImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Adverb()), nil
}

func gofakeitPrepositionImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Preposition()), nil
}

func gofakeitAdjectiveImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Adjective()), nil
}

func gofakeitWordImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Word()), nil
}

func gofakeitSentenceImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	var wordCount int
	_ = gocty.FromCtyValue(args[0], &wordCount)

	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Sentence(wordCount)), nil
}

func gofakeitParagraphImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	var paragraphCount, sentenceCount, wordCount int
	_ = gocty.FromCtyValue(args[0], &paragraphCount)
	_ = gocty.FromCtyValue(args[1], &sentenceCount)
	_ = gocty.FromCtyValue(args[2], &wordCount)

	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Paragraph(paragraphCount, sentenceCount, wordCount, args[3].AsString())), nil
}

func gofakeitLoremIpsumWordImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.LoremIpsumWord()), nil
}

func gofakeitLoremIpsumSentenceImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	var wordCount int
	_ = gocty.FromCtyValue(args[0], &wordCount)

	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.LoremIpsumSentence(wordCount)), nil
}

func gofakeitLoremIpsumParagraphImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	var paragraphCount, sentenceCount, wordCount int
	_ = gocty.FromCtyValue(args[0], &paragraphCount)
	_ = gocty.FromCtyValue(args[1], &sentenceCount)
	_ = gocty.FromCtyValue(args[2], &wordCount)

	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.LoremIpsumParagraph(paragraphCount, sentenceCount, wordCount, args[3].AsString())), nil
}

func gofakeitQuestionImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Question()), nil
}

func gofakeitQuoteImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Quote()), nil
}

func gofakeitPhraseImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Phrase()), nil
}

func gofakeitFruitImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Fruit()), nil
}

func gofakeitVegetableImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Vegetable()), nil
}

func gofakeitBreakfastImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Breakfast()), nil
}

func gofakeitLunchImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Lunch()), nil
}

func gofakeitDinnerImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Dinner()), nil
}

func gofakeitSnackImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Snack()), nil
}

func gofakeitDessertImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Dessert()), nil
}

func gofakeitBoolImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.BoolVal(gofakeit.Bool()), nil
}

func gofakeitUUIDImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.UUID()), nil
}

func gofakeitColorImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Color()), nil
}

func gofakeitHexColorImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.HexColor()), nil
}

func gofakeitRGBColorImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	rgbcolor, err := gocty.ToCtyValue(gofakeit.RGBColor(), cty.List(cty.Number))
	return rgbcolor, err
}

func gofakeitSafeColorImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.SafeColor()), nil
}

func gofakeitURLImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.URL()), nil
}

func gofakeitImageURLImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	var width, height int
	_ = gocty.FromCtyValue(args[0], &width)
	_ = gocty.FromCtyValue(args[1], &height)

	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.ImageURL(width, height)), nil
}

func gofakeitDomainNameImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.DomainName()), nil
}

func gofakeitDomainSuffixImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.DomainSuffix()), nil
}

func gofakeitIPv4AddressImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.IPv4Address()), nil
}

func gofakeitIPv6AddressImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.IPv6Address()), nil
}

func gofakeitHTTPStatusCodeImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	value, err := gocty.ToCtyValue(gofakeit.HTTPStatusCode(), cty.Number)

	return value, err
}

func gofakeitHTTPSimpleStatusCodeImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	value, err := gocty.ToCtyValue(gofakeit.HTTPStatusCodeSimple(), cty.Number)

	return value, err
}

func gofakeitLogLevelImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.LogLevel(args[0].AsString())), nil
}

func gofakeitHTTPMethodImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.HTTPMethod()), nil
}

func gofakeitUserAgentImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.UserAgent()), nil
}

func gofakeitChromeUserAgentImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.ChromeUserAgent()), nil
}

func gofakeitFirefoxUserAgentImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.FirefoxUserAgent()), nil
}

func gofakeitOperaUserAgentImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.OperaUserAgent()), nil
}

func gofakeitSafariUserAgentImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.SafariUserAgent()), nil
}

func gofakeitDateImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)

	return cty.StringVal(gofakeit.Date().String()), nil
}

func gofakeitDateRangeImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)

	start, err := time.Parse(time.RFC3339, args[0].AsString())
	if err != nil {
		return cty.Value{}, err
	}

	end, err := time.Parse(time.RFC3339, args[1].AsString())
	if err != nil {
		return cty.Value{}, err
	}

	return cty.StringVal(gofakeit.DateRange(start, end).String()), err
}

func gofakeitNanoSecondImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)

	return gocty.ToCtyValue(gofakeit.NanoSecond(), cty.Number)
}

func gofakeitSecondImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)

	return gocty.ToCtyValue(gofakeit.Second(), cty.Number)
}

func gofakeitMinuteImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)

	return gocty.ToCtyValue(gofakeit.Minute(), cty.Number)
}

func gofakeitHourImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)

	return gocty.ToCtyValue(gofakeit.Hour(), cty.Number)
}

func gofakeitMonthImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Month()), nil
}

func gofakeitDayImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)

	return gocty.ToCtyValue(gofakeit.Day(), cty.Number)
}

func gofakeitWeekDayImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.WeekDay()), nil
}

func gofakeitYearImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)

	return gocty.ToCtyValue(gofakeit.Year(), cty.Number)
}

func gofakeitTimeZoneImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.TimeZone()), nil
}

func gofakeitTimeZoneAbvImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.TimeZoneAbv()), nil
}

func gofakeitTimeZoneFullImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.TimeZoneFull()), nil
}

func gofakeitTimeZoneOffsetImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return gocty.ToCtyValue(gofakeit.TimeZoneOffset(), cty.Number)
}

func gofakeitTimeZoneRegionImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.TimeZoneRegion()), nil
}

func gofakeitPriceImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	min, _ := args[0].AsBigFloat().Float64()
	max, _ := args[1].AsBigFloat().Float64()

	gofakeit.Seed(0)
	return cty.NumberFloatVal(gofakeit.Price(min, max)), nil
}

func gofakeitCreditCardCvvImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.CreditCardCvv()), nil
}

func gofakeitCreditCardExpImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.CreditCardExp()), nil
}

func gofakeitCreditCardNumberImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	var types, bins []string
	gaps := false

	err := gocty.FromCtyValue(args[0], &types)
	if err != nil {
		return cty.StringVal(""), err
	}

	if len(types) == 0 {
		return cty.StringVal(""), Errorf("at least 1 credit card type needs to be specified")
	}

	validTypes := []string{"visa", "mastercard", "american-express", "diners-club", "discover", "jcb", "unionpay", "maestro", "elo", "hiper", "hipercard"}
	for _, t := range types {
		if !sliceContains(validTypes, t) {
			return cty.StringVal(""), Errorf("given type(%s) isn't allowed. Pick one from visa, mastercard, american-express, diners-club, discover, jcb, unionpay, maestro, elo, hiper, hipercard", t)
		}
	}

	if len(args) > 1 {
		err = gocty.FromCtyValue(args[1], &bins)

		if err != nil {
			return cty.StringVal(""), err
		}
	}

	if len(args) > 2 {
		gaps = args[2].True() == true
	}

	cco := gofakeit.CreditCardOptions{
		Types: types,
		Bins:  bins,
		Gaps:  gaps,
	}

	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.CreditCardNumber(&cco)), nil
}

func gofakeitCreditCardTypeImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.CreditCardType()), nil
}

func gofakeitCurrencyLongImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.CurrencyLong()), nil
}

func gofakeitCurrencyShortImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.CurrencyShort()), nil
}

func gofakeitAchRoutingImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.AchRouting()), nil
}

func gofakeitAchAccountImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.AchAccount()), nil
}

func gofakeitBitcoinAddressImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.BitcoinAddress()), nil
}

func gofakeitBitcoinPrivateKeyImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.BitcoinPrivateKey()), nil
}

func gofakeitBSImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.BS()), nil
}

func gofakeitBuzzWordImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.BuzzWord()), nil
}

func gofakeitCompanyImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Company()), nil
}

func gofakeitCompanySuffixImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.CompanySuffix()), nil
}

func gofakeitJobDescriptorImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.JobDescriptor()), nil
}

func gofakeitJobLevelImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.JobLevel()), nil
}

func gofakeitJobTitleImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.JobTitle()), nil
}

func gofakeitHackerAbbreviationImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.HackerAbbreviation()), nil
}

func gofakeitHackerAdjectiveImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.HackerAdjective()), nil
}

func gofakeitHackeringVerbImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.HackeringVerb()), nil
}

func gofakeitHackerNounImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.HackerNoun()), nil
}

func gofakeitHackerPhraseImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.HackerPhrase()), nil
}

func gofakeitHackerVerbImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.HackerVerb()), nil
}

func gofakeitHipsterWordImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.HipsterWord()), nil
}

func gofakeitHipsterSentenceImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	var wordCount int
	_ = gocty.FromCtyValue(args[0], &wordCount)

	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.HipsterSentence(wordCount)), nil
}

func gofakeitHipsterParagraphImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	var paragraphCount, sentenceCount, wordCount int
	_ = gocty.FromCtyValue(args[0], &paragraphCount)
	_ = gocty.FromCtyValue(args[1], &sentenceCount)
	_ = gocty.FromCtyValue(args[2], &wordCount)

	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.HipsterParagraph(paragraphCount, sentenceCount, wordCount, args[3].AsString())), nil
}

func gofakeitAppNameImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.AppName()), nil
}

func gofakeitAppVersionImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.AppVersion()), nil
}

func gofakeitAppAuthorImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.AppAuthor()), nil
}

func gofakeitPetNameImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.PetName()), nil
}

func gofakeitAnimalImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Animal()), nil
}

func gofakeitAnimalTypeImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.AnimalType()), nil
}

func gofakeitFarmAnimalImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.FarmAnimal()), nil
}

func gofakeitCatImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Cat()), nil
}

func gofakeitDogImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Dog()), nil
}

func gofakeitEmojiImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Emoji()), nil
}

func gofakeitEmojiDescriptionImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.EmojiDescription()), nil
}

func gofakeitEmojiCategoryImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.EmojiCategory()), nil
}

func gofakeitEmojiAliasImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.EmojiAlias()), nil
}

func gofakeitEmojiTagImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.EmojiTag()), nil
}

func gofakeitLanguageImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Language()), nil
}

func gofakeitLanguageAbbreviationImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.LanguageAbbreviation()), nil
}

func gofakeitProgrammingLanguageImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.ProgrammingLanguage()), nil
}

func gofakeitProgrammingLanguageBestImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.ProgrammingLanguageBest()), nil
}

func gofakeitNumberImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	var min, max int
	err := gocty.FromCtyValue(args[0], &min)

	if err != nil {
		return cty.NumberIntVal(0), err
	}

	err = gocty.FromCtyValue(args[1], &max)

	if err != nil {
		return cty.NumberIntVal(0), err
	}

	return cty.NumberIntVal(int64(gofakeit.Number(min, max))), nil
}

func gofakeitInt8Impl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.NumberIntVal(int64(gofakeit.Int8())), nil
}

func gofakeitInt16Impl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.NumberIntVal(int64(gofakeit.Int16())), nil
}

func gofakeitInt32Impl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.NumberIntVal(int64(gofakeit.Int32())), nil
}

func gofakeitInt64Impl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.NumberIntVal(gofakeit.Int64()), nil
}

func gofakeitUint8Impl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.NumberUIntVal(uint64(gofakeit.Uint8())), nil
}

func gofakeitUint16Impl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.NumberUIntVal(uint64(gofakeit.Uint16())), nil
}

func gofakeitUint32Impl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.NumberUIntVal(uint64(gofakeit.Uint32())), nil
}

func gofakeitUint64Impl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.NumberUIntVal(gofakeit.Uint64()), nil
}

func gofakeitFloat32Impl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.NumberFloatVal(float64(gofakeit.Float32())), nil
}

func gofakeitFloat32RangeImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	var min, max float32
	err := gocty.FromCtyValue(args[0], &min)

	if err != nil {
		return cty.NumberIntVal(0), err
	}

	err = gocty.FromCtyValue(args[1], &max)

	if err != nil {
		return cty.NumberIntVal(0), err
	}

	gofakeit.Seed(0)
	return cty.NumberFloatVal(float64(gofakeit.Float32Range(min, max))), nil
}
func gofakeitFloat64Impl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.NumberFloatVal(gofakeit.Float64()), nil
}

func gofakeitFloat64RangeImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	var min, max float64
	err := gocty.FromCtyValue(args[0], &min)

	if err != nil {
		return cty.NumberIntVal(0), err
	}

	err = gocty.FromCtyValue(args[1], &max)

	if err != nil {
		return cty.NumberIntVal(0), err
	}

	gofakeit.Seed(0)
	return cty.NumberFloatVal(gofakeit.Float64Range(min, max)), nil
}

func gofakeitShuffleIntsImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	var deck []int
	err := gocty.FromCtyValue(args[0], &deck)

	if err != nil {
		return cty.Value{}, err
	}

	gofakeit.Seed(0)
	gofakeit.ShuffleInts(deck)

	return gocty.ToCtyValue(deck, cty.List(cty.Number))
}

func gofakeitRandomIntImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	var deck []int
	err := gocty.FromCtyValue(args[0], &deck)

	if err != nil {
		return cty.Value{}, err
	}

	gofakeit.Seed(0)
	return cty.NumberIntVal(int64(gofakeit.RandomInt(deck))), nil
}

func gofakeitDigitImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Digit()), nil
}

func gofakeitLetterImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Letter()), nil
}

func gofakeitLexifyImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Lexify(args[0].AsString())), nil
}

func gofakeitNumerifyImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)
	return cty.StringVal(gofakeit.Numerify(args[0].AsString())), nil
}

func gofakeitShuffleStringsImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	var deck []string
	err := gocty.FromCtyValue(args[0], &deck)

	if err != nil {
		return cty.Value{}, err
	}

	gofakeit.Seed(0)
	gofakeit.ShuffleStrings(deck)

	return gocty.ToCtyValue(deck, cty.List(cty.String))
}

func gofakeitRandomStringImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	gofakeit.Seed(0)

	var list []string
	err := gocty.FromCtyValue(args[0], &list)

	return cty.StringVal(gofakeit.RandomString(list)), err
}
