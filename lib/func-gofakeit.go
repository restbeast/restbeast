package lib

import (
	"errors"
	"github.com/brianvoe/gofakeit/v5"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"github.com/zclconf/go-cty/cty/gocty"
	"time"
)

func gofakeitName() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Name()), nil
		},
	})
}

func gofakeitNamePrefix() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.NamePrefix()), nil
		},
	})
}

func gofakeitNameSuffix() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.NamePrefix()), nil
		},
	})
}

func gofakeitFirstName() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.FirstName()), nil
		},
	})
}

func gofakeitLastName() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.LastName()), nil
		},
	})
}

func gofakeitGender() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Gender()), nil
		},
	})
}

func gofakeitSSN() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.SSN()), nil
		},
	})
}

func gofakeitEmail() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Email()), nil
		},
	})
}

func gofakeitPhone() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Phone()), nil
		},
	})
}

func gofakeitPhoneFormatted() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.PhoneFormatted()), nil
		},
	})
}

func gofakeitUsername() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Username()), nil
		},
	})
}

func gofakeitPassword() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "lower",
				Type: cty.Bool,
				AllowDynamicType: false,
			},
			{
				Name: "upper",
				Type: cty.Bool,
				AllowDynamicType: false,
			},
			{
				Name: "numeric",
				Type: cty.Bool,
				AllowDynamicType: false,
			},
			{
				Name: "special",
				Type: cty.Bool,
				AllowDynamicType: false,
			},
			{
				Name: "space",
				Type: cty.Bool,
				AllowDynamicType: false,
			},
			{
				Name: "length",
				Type: cty.Number,
				AllowDynamicType: false,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)

			var size int
			_ = gocty.FromCtyValue(args[5], &size)
			return cty.StringVal(gofakeit.Password(args[0].True(), args[1].True(), args[2].True(), args[3].True(), args[4].True(), size)), nil
		},
	})
}

func gofakeitCity() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.City()), nil
		},
	})
}

func gofakeitCountry() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Country()), nil
		},
	})
}

func gofakeitCountryAbr() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.CountryAbr()), nil
		},
	})
}

func gofakeitState() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.State()), nil
		},
	})
}

func gofakeitStateAbr() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.StateAbr()), nil
		},
	})
}

func gofakeitStreet() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Street()), nil
		},
	})
}

func gofakeitStreetName() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.StreetName()), nil
		},
	})
}

func gofakeitStreetNumber() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.StreetNumber()), nil
		},
	})
}

func gofakeitStreetPrefix() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.StreetPrefix()), nil
		},
	})
}

func gofakeitStreetSuffix() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.StreetSuffix()), nil
		},
	})
}

func gofakeitZip() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Zip()), nil
		},
	})
}

func gofakeitLatitude() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.NumberFloatVal(gofakeit.Latitude()), nil
		},
	})
}

func gofakeitLatitudeInRange() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "min",
				Type: cty.Number,
				AllowDynamicType: false,
			},
			{
				Name: "max",
				Type: cty.Number,
				AllowDynamicType: false,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			min, _ := args[0].AsBigFloat().Float64()
			max, _ := args[1].AsBigFloat().Float64()
			inRange, err := gofakeit.LatitudeInRange(min, max)

			return cty.NumberFloatVal(inRange), err
		},
	})
}

func gofakeitLongitude() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.NumberFloatVal(gofakeit.Longitude()), nil
		},
	})
}

func gofakeitLongitudeInRange() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "min",
				Type: cty.Number,
				AllowDynamicType: false,
			},
			{
				Name: "max",
				Type: cty.Number,
				AllowDynamicType: false,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			min, _ := args[0].AsBigFloat().Float64()
			max, _ := args[1].AsBigFloat().Float64()
			inRange, err := gofakeit.LongitudeInRange(min, max)

			return cty.NumberFloatVal(inRange), err
		},
	})
}

func gofakeitGamertag() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Gamertag()), nil
		},
	})
}

func gofakeitBeerAlcohol() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.BeerAlcohol()), nil
		},
	})
}

func gofakeitBeerBlg() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.BeerBlg()), nil
		},
	})
}

func gofakeitBeerHop() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.BeerHop()), nil
		},
	})
}

func gofakeitBeerIbu() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.BeerIbu()), nil
		},
	})
}

func gofakeitBeerMalt() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.BeerMalt()), nil
		},
	})
}

func gofakeitBeerName() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.BeerName()), nil
		},
	})
}

func gofakeitBeerStyle() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.BeerStyle()), nil
		},
	})
}

func gofakeitBeerYeast() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.BeerYeast()), nil
		},
	})
}

func gofakeitCarMaker() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.CarMaker()), nil
		},
	})
}

func gofakeitCarModel() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.CarModel()), nil
		},
	})
}

func gofakeitCarType() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.CarType()), nil
		},
	})
}

func gofakeitCarFuelType() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.CarFuelType()), nil
		},
	})
}

func gofakeitCarTransmissionType() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.CarTransmissionType()), nil
		},
	})
}

func gofakeitNoun() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Noun()), nil
		},
	})
}

func gofakeitVerb() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Verb()), nil
		},
	})
}

func gofakeitAdverb() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Adverb()), nil
		},
	})
}

func gofakeitPreposition() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Preposition()), nil
		},
	})
}

func gofakeitAdjective() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Adjective()), nil
		},
	})
}

func gofakeitWord() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Word()), nil
		},
	})
}

func gofakeitSentence() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "wordCount",
				Type: cty.Number,
				AllowDynamicType: false,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			var wordCount int
			_ = gocty.FromCtyValue(args[0], &wordCount)

			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Sentence(wordCount)), nil
		},
	})
}

func gofakeitParagraph() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "paragraphCount",
				Type: cty.Number,
				AllowDynamicType: false,
			},
			{
				Name: "sentenceCount",
				Type: cty.Number,
				AllowDynamicType: false,
			},
			{
				Name: "wordCount",
				Type: cty.Number,
				AllowDynamicType: false,
			},
			{
				Name: "separator",
				Type: cty.String,
				AllowDynamicType: false,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			var paragraphCount, sentenceCount, wordCount int
			_ = gocty.FromCtyValue(args[0], &paragraphCount)
			_ = gocty.FromCtyValue(args[1], &sentenceCount)
			_ = gocty.FromCtyValue(args[2], &wordCount)

			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Paragraph(paragraphCount, sentenceCount, wordCount, args[3].AsString())), nil
		},
	})
}

func gofakeitLoremIpsumWord() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.LoremIpsumWord()), nil
		},
	})
}

func gofakeitLoremIpsumSentence() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "wordCount",
				Type: cty.Number,
				AllowDynamicType: false,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			var wordCount int
			_ = gocty.FromCtyValue(args[0], &wordCount)

			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.LoremIpsumSentence(wordCount)), nil
		},
	})
}

func gofakeitLoremIpsumParagraph() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "paragraphCount",
				Type: cty.Number,
				AllowDynamicType: false,
			},
			{
				Name: "sentenceCount",
				Type: cty.Number,
				AllowDynamicType: false,
			},
			{
				Name: "wordCount",
				Type: cty.Number,
				AllowDynamicType: false,
			},
			{
				Name: "separator",
				Type: cty.String,
				AllowDynamicType: false,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			var paragraphCount, sentenceCount, wordCount int
			_ = gocty.FromCtyValue(args[0], &paragraphCount)
			_ = gocty.FromCtyValue(args[1], &sentenceCount)
			_ = gocty.FromCtyValue(args[2], &wordCount)

			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.LoremIpsumParagraph(paragraphCount, sentenceCount, wordCount, args[3].AsString())), nil
		},
	})
}

func gofakeitQuestion() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Question()), nil
		},
	})
}

func gofakeitQuote() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Quote()), nil
		},
	})
}

func gofakeitPhrase() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Phrase()), nil
		},
	})
}

func gofakeitFruit() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Fruit()), nil
		},
	})
}

func gofakeitVegetable() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Vegetable()), nil
		},
	})
}

func gofakeitBreakfast() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Breakfast()), nil
		},
	})
}

func gofakeitLunch() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Lunch()), nil
		},
	})
}

func gofakeitDinner() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Dinner()), nil
		},
	})
}

func gofakeitSnack() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Snack()), nil
		},
	})
}

func gofakeitDessert() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Dessert()), nil
		},
	})
}

func gofakeitBool() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Bool),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.BoolVal(gofakeit.Bool()), nil
		},
	})
}

func gofakeitUUID() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.UUID()), nil
		},
	})
}

func gofakeitColor() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Color()), nil
		},
	})
}

func gofakeitHexColor() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.HexColor()), nil
		},
	})
}

func gofakeitRGBColor() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.List(cty.Number)),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			rgbcolor, err := gocty.ToCtyValue(gofakeit.RGBColor(), cty.List(cty.Number))

			return rgbcolor, err
		},
	})
}

func gofakeitSafeColor() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.SafeColor()), nil
		},
	})
}

func gofakeitURL() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.URL()), nil
		},
	})
}

func gofakeitImageURL() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			var width, height int
			_ = gocty.FromCtyValue(args[0], &width)
			_ = gocty.FromCtyValue(args[0], &height)

			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.ImageURL(width, height)), nil
		},
	})
}

func gofakeitDomainName() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.DomainName()), nil
		},
	})
}

func gofakeitDomainSuffix() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.DomainSuffix()), nil
		},
	})
}

func gofakeitIPv4Address() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.IPv4Address()), nil
		},
	})
}

func gofakeitIPv6Address() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.IPv6Address()), nil
		},
	})
}

func gofakeitHTTPStatusCode() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			value, err := gocty.ToCtyValue(gofakeit.HTTPStatusCode(), cty.Number)

			return value, err
		},
	})
}

func gofakeitHTTPSimpleStatusCode() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			value, err := gocty.ToCtyValue(gofakeit.HTTPStatusCodeSimple(), cty.Number)

			return value, err
		},
	})
}

func gofakeitLogLevel() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "logType",
				Type: cty.String,
				AllowDynamicType: false,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.LogLevel(args[0].AsString())), nil
		},
	})
}

func gofakeitHTTPMethod() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.HTTPMethod()), nil
		},
	})
}

func gofakeitUserAgent() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.UserAgent()), nil
		},
	})
}

func gofakeitChromeUserAgent() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.ChromeUserAgent()), nil
		},
	})
}

func gofakeitFirefoxUserAgent() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.FirefoxUserAgent()), nil
		},
	})
}

func gofakeitOperaUserAgent() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.OperaUserAgent()), nil
		},
	})
}

func gofakeitSafariUserAgent() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.SafariUserAgent()), nil
		},
	})
}

func gofakeitDate() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)

			return cty.StringVal(gofakeit.Date().String()), nil
		},
	})
}

func gofakeitDateRange() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "separator",
				Type: cty.String,
				AllowDynamicType: false,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)

			start, err := time.Parse("2020-07-14T12:14:54+00:00", args[0].AsString())
			if err != nil {
				return cty.Value{}, err
			}

			end, err := time.Parse("2020-07-14T12:14:54+00:00", args[0].AsString())
			if err != nil {
				return cty.Value{}, err
			}

			return cty.StringVal(gofakeit.DateRange(start, end).String()), err
		},
	})
}

func gofakeitNanoSecond() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)

			return gocty.ToCtyValue(gofakeit.NanoSecond(), cty.Number)
		},
	})
}

func gofakeitSecond() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)

			return gocty.ToCtyValue(gofakeit.Second(), cty.Number)
		},
	})
}

func gofakeitMinute() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)

			return gocty.ToCtyValue(gofakeit.Minute(), cty.Number)
		},
	})
}

func gofakeitHour() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)

			return gocty.ToCtyValue(gofakeit.Hour(), cty.Number)
		},
	})
}

func gofakeitMonth() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Month()), nil
		},
	})
}

func gofakeitDay() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)

			return gocty.ToCtyValue(gofakeit.Day(), cty.Number)
		},
	})
}

func gofakeitWeekDay() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.WeekDay()), nil
		},
	})
}

func gofakeitYear() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)

			return gocty.ToCtyValue(gofakeit.Year(), cty.Number)

		},
	})
}

func gofakeitTimeZone() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.TimeZone()), nil
		},
	})
}

func gofakeitTimeZoneAbv() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.TimeZoneAbv()), nil
		},
	})
}

func gofakeitTimeZoneFull() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.TimeZoneFull()), nil
		},
	})
}

func gofakeitTimeZoneOffset() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return gocty.ToCtyValue(gofakeit.TimeZoneOffset(), cty.Number)
		},
	})
}

func gofakeitTimeZoneRegion() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.TimeZoneRegion()), nil
		},
	})
}

func gofakeitPrice() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "min",
				Type: cty.Number,
				AllowDynamicType: false,
			},
			{
				Name: "max",
				Type: cty.Number,
				AllowDynamicType: false,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			min, _ := args[0].AsBigFloat().Float64()
			max, _ := args[1].AsBigFloat().Float64()

			gofakeit.Seed(0)
			return cty.NumberFloatVal(gofakeit.Price(min, max)), nil
		},
	})
}

func gofakeitCreditCardCvv() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.CreditCardCvv()), nil
		},
	})
}

func gofakeitCreditCardExp() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.CreditCardExp()), nil
		},
	})
}

func gofakeitCreditCardNumber() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "types",
				Type: cty.List(cty.String),
				AllowDynamicType: false,
			},
			{
				Name: "bins",
				Type: cty.List(cty.String),
				AllowDynamicType: false,
			},
			{
				Name: "gaps",
				Type: cty.Bool,
				AllowDynamicType: false,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			var types, bins []string

			err := gocty.FromCtyValue(args[0], &types)
			if err != nil {
				return cty.StringVal(""), err
			}

			if len(types) == 0 {
				return cty.Value{}, errors.New("at least 1 credit card type needs to be specified")
			}

			validTypes := []string{"visa", "mastercard", "american-express", "diners-club", "discover", "jcb", "unionpay", "maestro", "elo", "hiper", "hipercard"}
			for _, t := range types {
				if !sliceContains(validTypes, t) {
					return cty.Value{}, errors.New("given type("+t+") isn't allowed. Pick one from visa, mastercard, american-express, diners-club, discover, jcb, unionpay, maestro, elo, hiper, hipercard")
				}
			}

			err = gocty.FromCtyValue(args[0], &bins)

			if err != nil {
				return cty.StringVal(""), err
			}

			cco := gofakeit.CreditCardOptions{
				Types: types,
				Bins:  bins,
				Gaps:  args[2].True(),
			}

			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.CreditCardNumber(&cco)), nil
		},
	})
}

func gofakeitCreditCardType() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.CreditCardType()), nil
		},
	})
}

func gofakeitCurrencyLong() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.CurrencyLong()), nil
		},
	})
}

func gofakeitCurrencyShort() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.CurrencyShort()), nil
		},
	})
}

func gofakeitAchRouting() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.AchRouting()), nil
		},
	})
}

func gofakeitAchAccount() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.AchAccount()), nil
		},
	})
}

func gofakeitBitcoinAddress() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.BitcoinAddress()), nil
		},
	})
}

func gofakeitBitcoinPrivateKey() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.BitcoinPrivateKey()), nil
		},
	})
}

func gofakeitBS() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.BS()), nil
		},
	})
}

func gofakeitBuzzWord() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.BuzzWord()), nil
		},
	})
}

func gofakeitCompany() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Company()), nil
		},
	})
}

func gofakeitCompanySuffix() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.CompanySuffix()), nil
		},
	})
}

func gofakeitJobDescriptor() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.JobDescriptor()), nil
		},
	})
}

func gofakeitJobLevel() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.JobLevel()), nil
		},
	})
}

func gofakeitJobTitle() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.JobTitle()), nil
		},
	})
}

func gofakeitHackerAbbreviation() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.HackerAbbreviation()), nil
		},
	})
}

func gofakeitHackerAdjective() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.HackerAdjective()), nil
		},
	})
}

func gofakeitHackeringVerb() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.HackeringVerb()), nil
		},
	})
}

func gofakeitHackerNoun() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.HackerNoun()), nil
		},
	})
}

func gofakeitHackerPhrase() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.HackerPhrase()), nil
		},
	})
}

func gofakeitHackerVerb() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.HackerVerb()), nil
		},
	})
}

func gofakeitHipsterWord() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.HipsterWord()), nil
		},
	})
}

func gofakeitHipsterSentence() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "wordCount",
				Type: cty.Number,
				AllowDynamicType: false,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			var wordCount int
			_ = gocty.FromCtyValue(args[0], &wordCount)

			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.HipsterSentence(wordCount)), nil
		},
	})
}

func gofakeitHipsterParagraph() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "paragraphCount",
				Type: cty.Number,
				AllowDynamicType: false,
			},
			{
				Name: "sentenceCount",
				Type: cty.Number,
				AllowDynamicType: false,
			},
			{
				Name: "wordCount",
				Type: cty.Number,
				AllowDynamicType: false,
			},
			{
				Name: "separator",
				Type: cty.String,
				AllowDynamicType: false,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			var paragraphCount, sentenceCount, wordCount int
			_ = gocty.FromCtyValue(args[0], &paragraphCount)
			_ = gocty.FromCtyValue(args[1], &sentenceCount)
			_ = gocty.FromCtyValue(args[2], &wordCount)

			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.HipsterParagraph(paragraphCount, sentenceCount, wordCount, args[3].AsString())), nil
		},
	})
}

func gofakeitAppName() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.AppName()), nil
		},
	})
}

func gofakeitAppVersion() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.AppVersion()), nil
		},
	})
}

func gofakeitAppAuthor() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.AppAuthor()), nil
		},
	})
}

func gofakeitPetName() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.PetName()), nil
		},
	})
}

func gofakeitAnimal() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Animal()), nil
		},
	})
}

func gofakeitAnimalType() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.AnimalType()), nil
		},
	})
}

func gofakeitFarmAnimal() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.FarmAnimal()), nil
		},
	})
}

func gofakeitCat() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Cat()), nil
		},
	})
}

func gofakeitDog() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Dog()), nil
		},
	})
}

func gofakeitEmoji() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Emoji()), nil
		},
	})
}

func gofakeitEmojiDescription() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.EmojiDescription()), nil
		},
	})
}

func gofakeitEmojiCategory() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.EmojiCategory()), nil
		},
	})
}

func gofakeitEmojiAlias() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.EmojiAlias()), nil
		},
	})
}

func gofakeitEmojiTag() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.EmojiTag()), nil
		},
	})
}

func gofakeitLanguage() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Language()), nil
		},
	})
}

func gofakeitLanguageAbbreviation() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.LanguageAbbreviation()), nil
		},
	})
}

func gofakeitProgrammingLanguage() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.ProgrammingLanguage()), nil
		},
	})
}

func gofakeitProgrammingLanguageBest() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.ProgrammingLanguageBest()), nil
		},
	})
}

func gofakeitNumber() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "min",
				Type: cty.Number,
				AllowDynamicType: false,
			},
			{
				Name: "max",
				Type: cty.Number,
				AllowDynamicType: false,
			},
		},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
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
		},
	})
}

func gofakeitInt8() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.NumberIntVal(int64(gofakeit.Int8())), nil
		},
	})
}

func gofakeitInt16() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.NumberIntVal(int64(gofakeit.Int16())), nil
		},
	})
}

func gofakeitInt32() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.NumberIntVal(int64(gofakeit.Int32())), nil
		},
	})
}

func gofakeitInt64() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.NumberIntVal(gofakeit.Int64()), nil
		},
	})
}

func gofakeitUint8() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.NumberUIntVal(uint64(gofakeit.Uint8())), nil
		},
	})
}

func gofakeitUint16() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.NumberUIntVal(uint64(gofakeit.Uint16())), nil
		},
	})
}

func gofakeitUint32() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.NumberUIntVal(uint64(gofakeit.Uint32())), nil
		},
	})
}

func gofakeitUint64() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.NumberUIntVal(gofakeit.Uint64()), nil
		},
	})
}

func gofakeitFloat32() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.NumberFloatVal(float64(gofakeit.Float32())), nil
		},
	})
}

func gofakeitFloat32Range() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
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
		},
	})
}

func gofakeitFloat64() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.NumberFloatVal(gofakeit.Float64()), nil
		},
	})
}

func gofakeitFloat64Range() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.Number),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
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
		},
	})
}

func gofakeitShuffleInts() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "deck",
				Type: cty.List(cty.Number),
				AllowDynamicType: false,
			},
		},
		Type: function.StaticReturnType(cty.List(cty.Number)),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			var deck []int
			err := gocty.FromCtyValue(args[0], &deck)

			if err != nil {
				return cty.Value{}, err
			}

			gofakeit.Seed(0)
			gofakeit.ShuffleInts(deck)

			return gocty.ToCtyValue(deck, cty.List(cty.Number))
		},
	})
}

func gofakeitRandomInt() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "deck",
				Type: cty.List(cty.Number),
				AllowDynamicType: false,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			var deck []int
			err := gocty.FromCtyValue(args[0], &deck)

			if err != nil {
				return cty.Value{}, err
			}

			gofakeit.Seed(0)
			return cty.NumberIntVal(int64(gofakeit.RandomInt(deck))), nil
		},
	})
}

func gofakeitDigit() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Digit()), nil
		},
	})
}

func gofakeitLetter() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Letter()), nil
		},
	})
}

func gofakeitLexify() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "str",
				Type: cty.String,
				AllowDynamicType: false,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Lexify(args[0].AsString())), nil
		},
	})
}

func gofakeitNumerify() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "str",
				Type: cty.String,
				AllowDynamicType: false,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)
			return cty.StringVal(gofakeit.Numerify(args[0].AsString())), nil
		},
	})
}

func gofakeitShuffleStrings() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "deck",
				Type: cty.List(cty.String),
				AllowDynamicType: false,
			},
		},
		Type: function.StaticReturnType(cty.List(cty.String)),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			var deck []string
			err := gocty.FromCtyValue(args[0], &deck)

			if err != nil {
				return cty.Value{}, err
			}

			gofakeit.Seed(0)
			gofakeit.ShuffleStrings(deck)

			return gocty.ToCtyValue(deck, cty.List(cty.String))
		},
	})
}

func gofakeitRandomString() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "list",
				Type: cty.List(cty.String),
				AllowDynamicType: false,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			gofakeit.Seed(0)

			var list []string
			err := gocty.FromCtyValue(args[0], &list)

			return cty.StringVal(gofakeit.RandomString(list)), err
		},
	})
}
