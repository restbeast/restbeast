package lib

import (
	"github.com/Masterminds/semver"
	"github.com/brianvoe/gofakeit/v5"
	"github.com/zclconf/go-cty/cty"
	"math/big"
	"regexp"
	"strings"
	"testing"
)

func Test_gofakeitNameImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitNameImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitNameImpl() empty string")
		}
	})
}

func Test_gofakeitNamePrefixImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitNamePrefixImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitNamePrefixImpl() empty string")
		}
	})
}

func Test_gofakeitNameSuffixImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitNameSuffixImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitNameSuffixImpl() empty string")
		}
	})
}

func Test_gofakeitFirstNameImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitFirstNameImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitFirstNameImpl() empty string")
		}
	})
}

func Test_gofakeitLastNameImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitLastNameImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitLastNameImpl() empty string")
		}
	})
}

func Test_gofakeitGenderImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitGenderImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitGenderImpl() empty string")
		}
	})
}

func Test_gofakeitSSNImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitSSNImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitSSNImpl() empty string")
		}
	})
}

func Test_gofakeitEmailImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitEmailImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitEmailImpl() empty string")
		}
	})
}

func Test_gofakeitPhoneImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitPhoneImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitPhoneImpl() empty string")
		}
	})
}

func Test_gofakeitPhoneFormattedImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitPhoneFormattedImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitPhoneFormattedImpl() empty string")
		}
	})
}

func Test_gofakeitUsernameImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitUsernameImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitUsernameImpl() empty string")
		}
	})
}

func Test_gofakeitPasswordImpl(t *testing.T) {
	argList := []cty.Value{
		cty.BoolVal(true),
		cty.BoolVal(true),
		cty.BoolVal(true),
		cty.BoolVal(true),
		cty.BoolVal(true),
		cty.NumberIntVal(10),
	}

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitPasswordImpl(argList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitPasswordImpl() empty string")
		}
	})
}

func Test_gofakeitCityImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitCityImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitCityImpl() empty string")
		}
	})
}

func Test_gofakeitCountryImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitCountryImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitCountryImpl() empty string")
		}
	})
}

func Test_gofakeitCountryAbrImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitCountryAbrImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitCountryAbrImpl() empty string")
		}
	})
}

func Test_gofakeitStateImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitStateImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitStateImpl() empty string")
		}
	})
}

func Test_gofakeitStateAbrImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitStateAbrImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitStateAbrImpl() empty string")
		}
	})
}

func Test_gofakeitStreetImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitStreetImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitStreetImpl() empty string")
		}
	})
}

func Test_gofakeitStreetNameImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitStreetNameImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitStreetNameImpl() empty string")
		}
	})
}

func Test_gofakeitStreetNumberImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitStreetNumberImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitStreetNumberImpl() empty string")
		}
	})
}

func Test_gofakeitStreetPrefixImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitStreetPrefixImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitStreetPrefixImpl() empty string")
		}
	})
}

func Test_gofakeitStreetSuffixImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitStreetSuffixImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitStreetSuffixImpl() empty string")
		}
	})
}

func Test_gofakeitZipImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitZipImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitZipImpl() empty string")
		}
	})
}

func Test_gofakeitLatitudeImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		min := big.NewFloat(-90.00)
		max := big.NewFloat(90.00)
		got, _ := gofakeitLatitudeImpl(emptyArgList, cty.String)

		if got.AsBigFloat().Cmp(min) < 1 && got.AsBigFloat().Cmp(max) > 0 {
			t.Errorf("gofakeitLatitudeImpl() empty string")
		}
	})
}

func Test_gofakeitLongitudeImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		min := big.NewFloat(0.00)
		max := big.NewFloat(180.00)
		got, _ := gofakeitLongitudeImpl(emptyArgList, cty.String)

		if got.AsBigFloat().Cmp(min) < 1 && got.AsBigFloat().Cmp(max) > 0 {
			t.Errorf("gofakeitLongitudeImpl() empty string")
		}
	})
}

func Test_gofakeitLatitudeInRangeImpl(t *testing.T) {
	min := big.NewFloat(-5.00)
	max := big.NewFloat(10.00)

	argList := []cty.Value{
		cty.NumberIntVal(-5),
		cty.NumberIntVal(10),
	}

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitLatitudeInRangeImpl(argList, cty.String)

		if got.AsBigFloat().Cmp(min) < 1 && got.AsBigFloat().Cmp(max) > 0 {
			t.Errorf("gofakeitLatitudeInRangeImpl() empty string")
		}
	})
}

func Test_gofakeitLongitudeInRangeImpl(t *testing.T) {
	min := big.NewFloat(5.00)
	max := big.NewFloat(10.00)

	argList := []cty.Value{
		cty.NumberIntVal(5),
		cty.NumberIntVal(10),
	}

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitLongitudeInRangeImpl(argList, cty.String)

		if got.AsBigFloat().Cmp(min) < 1 && got.AsBigFloat().Cmp(max) > 0 {
			t.Errorf("gofakeitLongitudeInRangeImpl() empty string")
		}
	})
}

func Test_gofakeitGamertagImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitGamertagImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitGamertagImpl() empty string")
		}
	})
}

func Test_gofakeitBeerAlcoholImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitBeerAlcoholImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitBeerAlcoholImpl() empty string")
		}
	})
}

func Test_gofakeitBeerBlgImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitBeerBlgImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitBeerBlgImpl() empty string")
		}
	})
}

func Test_gofakeitBeerHopImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitBeerHopImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitBeerHopImpl() empty string")
		}
	})
}

func Test_gofakeitBeerIbuImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitBeerIbuImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitBeerIbuImpl() empty string")
		}
	})
}

func Test_gofakeitBeerMaltImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitBeerMaltImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitBeerMaltImpl() empty string")
		}
	})
}

func Test_gofakeitBeerNameImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitBeerNameImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitBeerNameImpl() empty string")
		}
	})
}

func Test_gofakeitBeerStyleImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitBeerStyleImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitBeerStyleImpl() empty string")
		}
	})
}

func Test_gofakeitBeerYeastImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitBeerYeastImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitBeerYeastImpl() empty string")
		}
	})
}

func Test_gofakeitCarMakerImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitCarMakerImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitCarMakerImpl() empty string")
		}
	})
}

func Test_gofakeitCarModelImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitCarModelImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitCarModelImpl() empty string")
		}
	})
}

func Test_gofakeitCarTypeImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitCarTypeImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitCarTypeImpl() empty string")
		}
	})
}

func Test_gofakeitCarFuelTypeImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitCarFuelTypeImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitCarFuelTypeImpl() empty string")
		}
	})
}

func Test_gofakeitCarTransmissionTypeImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitCarTransmissionTypeImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitCarTransmissionTypeImpl() empty string")
		}
	})
}

func Test_gofakeitNounImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitNounImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitNounImpl() empty string")
		}
	})
}

func Test_gofakeitVerbImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitVerbImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitVerbImpl() empty string")
		}
	})
}

func Test_gofakeitAdverbImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitAdverbImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitAdverbImpl() empty string")
		}
	})
}

func Test_gofakeitPrepositionImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitPrepositionImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitPrepositionImpl() empty string")
		}
	})
}

func Test_gofakeitAdjectiveImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitAdjectiveImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitAdjectiveImpl() empty string")
		}
	})
}

func Test_gofakeitWordImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitWordImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitPregofakeitWordImplring")
		}
	})
}

func Test_gofakeitSentenceImpl(t *testing.T) {
	argList := []cty.Value{cty.NumberIntVal(5)}

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitSentenceImpl(argList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitSentenceImpl() empty string")
		}
	})
}

func Test_gofakeitParagraphImpl(t *testing.T) {
	argList := []cty.Value{
		cty.NumberIntVal(1),
		cty.NumberIntVal(10),
		cty.NumberIntVal(50),
		cty.StringVal(" "),
	}

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitParagraphImpl(argList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitParagraphImpl() empty string")
		}
	})
}

func Test_gofakeitLoremIpsumWordImpl(t *testing.T) {
	argList := []cty.Value{
		cty.NumberIntVal(1),
		cty.NumberIntVal(10),
		cty.NumberIntVal(50),
		cty.StringVal(" "),
	}

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitLoremIpsumWordImpl(argList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitLoremIpsumWordImpl() empty string")
		}
	})
}

func Test_gofakeitLoremIpsumSentenceImpl(t *testing.T) {
	argList := []cty.Value{
		cty.NumberIntVal(1),
		cty.NumberIntVal(10),
		cty.NumberIntVal(50),
		cty.StringVal(" "),
	}

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitLoremIpsumSentenceImpl(argList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitLoremIpsumSentenceImpl() empty string")
		}
	})
}

func Test_gofakeitLoremIpsumParagraphImpl(t *testing.T) {
	argList := []cty.Value{
		cty.NumberIntVal(1),
		cty.NumberIntVal(10),
		cty.NumberIntVal(50),
		cty.StringVal(" "),
	}

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitLoremIpsumParagraphImpl(argList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitLoremIpsumParagraphImpl() empty string")
		}
	})
}

func Test_gofakeitQuestionImpl(t *testing.T) {
	argList := []cty.Value{
		cty.NumberIntVal(1),
		cty.NumberIntVal(10),
		cty.NumberIntVal(50),
		cty.StringVal(" "),
	}

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitQuestionImpl(argList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitQuestionImpl() empty string")
		}
	})
}

func Test_gofakeitQuoteImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitQuoteImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitQuoteImpl() empty string")
		}
	})
}

func Test_gofakeitPhraseImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitPhraseImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitPhraseImpl() empty string")
		}
	})
}

func Test_gofakeitFruitImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitFruitImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitFruitImpl() empty string")
		}
	})
}

func Test_gofakeitVegetableImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitVegetableImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitVegetableImpl() empty string")
		}
	})
}

func Test_gofakeitBreakfastImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitBreakfastImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitBreakfastImpl() empty string")
		}
	})
}

func Test_gofakeitLunchImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitLunchImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitLunchImpl() empty string")
		}
	})
}

func Test_gofakeitDinnerImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitDinnerImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitDinnerImpl() empty string")
		}
	})
}

func Test_gofakeitSnackImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitSnackImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitSnackImpl() empty string")
		}
	})
}

func Test_gofakeitDessertImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitDessertImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitDessertImpl() empty string")
		}
	})
}

func Test_gofakeitBoolImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitBoolImpl(emptyArgList, cty.String)

		// This will panic if value isn't boolean
		got.True()
	})
}

func Test_gofakeitUUIDImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitUUIDImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitUUIDImpl() empty string")
		}
	})
}

func Test_gofakeitColorImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitColorImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitColorImpl() empty string")
		}
	})
}

func Test_gofakeitHexColorImpl(t *testing.T) {
	validHexColor := regexp.MustCompile(`^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`)
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitHexColorImpl(emptyArgList, cty.String)

		if !validHexColor.MatchString(got.AsString()) {
			t.Errorf("gofakeitHexColorImpl() invalid hex color")
		}
	})
}

func Test_gofakeitRGBColorImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitRGBColorImpl(emptyArgList, cty.String)

		for _, c := range got.AsValueSlice() {
			if c.AsBigFloat().Cmp(big.NewFloat(0)) == -1 || c.AsBigFloat().Cmp(big.NewFloat(256)) == 1 {
				t.Errorf("gofakeitRGBColorImpl() out of bounds")
			}
		}
	})
}

func Test_gofakeitSafeColorImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitSafeColorImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitSafeColorImpl() empty string")
		}
	})
}

func Test_gofakeitURLImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitURLImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitURLImpl() empty string")
		}
	})
}

func Test_gofakeitImageURLImpl(t *testing.T) {
	argList := []cty.Value{
		cty.NumberIntVal(200),
		cty.NumberIntVal(200),
	}

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitImageURLImpl(argList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitImageURLImpl() empty string")
		}
	})
}

func Test_gofakeitDomainNameImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitDomainNameImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitDomainNameImpl() empty string")
		}
	})
}

func Test_gofakeitDomainSuffixImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitDomainSuffixImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitDomainSuffixImpl() empty string")
		}
	})
}

func Test_gofakeitIPv4AddressImpl(t *testing.T) {
	var emptyArgList []cty.Value

	validIpv4 := regexp.MustCompile(`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`)

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitIPv4AddressImpl(emptyArgList, cty.String)

		if !validIpv4.MatchString(got.AsString()) {
			t.Errorf("gofakeitIPv4AddressImpl() empty string")
		}
	})
}

func Test_gofakeitIPv6AddressImpl(t *testing.T) {
	var emptyArgList []cty.Value

	validIpv6 := regexp.MustCompile(`^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$`)

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitIPv6AddressImpl(emptyArgList, cty.String)

		if !validIpv6.MatchString(got.AsString()) {
			t.Errorf("gofakeitIPv6AddressImpl() empty string")
		}
	})
}

func Test_gofakeitHTTPStatusCodeImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		min := big.NewFloat(100)
		max := big.NewFloat(600.00)

		got, _ := gofakeitHTTPStatusCodeImpl(emptyArgList, cty.String)
		if got.AsBigFloat().Cmp(min) < 0 && got.AsBigFloat().Cmp(max) > 0 {
			t.Errorf("gofakeitHTTPStatusCodeImpl() invalid https status code")
		}
	})
}

func Test_gofakeitHTTPSimpleStatusCodeImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		min := big.NewFloat(100)
		max := big.NewFloat(600.00)

		got, _ := gofakeitHTTPSimpleStatusCodeImpl(emptyArgList, cty.String)
		if got.AsBigFloat().Cmp(min) < 0 && got.AsBigFloat().Cmp(max) > 0 {
			t.Errorf("gofakeitHTTPSimpleStatusCodeImpl() invalid https status code")
		}
	})
}

func Test_gofakeitLogLevelImpl(t *testing.T) {
	argList := []cty.Value{
		cty.StringVal("something"),
	}

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitLogLevelImpl(argList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitLogLevelImpl() empty string")
		}
	})
}

func Test_gofakeitHTTPMethodImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitHTTPMethodImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitHTTPMethodImpl() empty string")
		}
	})
}

func Test_gofakeitUserAgentImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitUserAgentImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitUserAgentImpl() empty string")
		}
	})
}

func Test_gofakeitChromeUserAgentImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitChromeUserAgentImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitChromeUserAgentImpl() empty string")
		}
	})
}

func Test_gofakeitFirefoxUserAgentImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitFirefoxUserAgentImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitFirefoxUserAgentImpl() empty string")
		}
	})
}

func Test_gofakeitOperaUserAgentImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitOperaUserAgentImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitOperaUserAgentImpl() empty string")
		}
	})
}

func Test_gofakeitSafariUserAgentImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitSafariUserAgentImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitSafariUserAgentImpl() empty string")
		}
	})
}

func Test_gofakeitDateImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitDateImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitDateImpl() empty string")
		}
	})
}

func Test_gofakeitDateRangeImpl(t *testing.T) {
	type args struct {
		args    []cty.Value
		retType cty.Type
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"error", args{[]cty.Value{cty.StringVal("1941-02-27"), cty.StringVal("1941-02-27")}, cty.String}, true},
		{"success", args{[]cty.Value{cty.StringVal("1941-02-27T00:00:00+00:00"), cty.StringVal("2018-08-26T00:00:00+00:00")}, cty.String}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := gofakeitDateRangeImpl(tt.args.args, tt.args.retType)

			if (err != nil) != tt.wantErr {
				t.Errorf("gofakeitDateRangeImpl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_gofakeitNanoSecondImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitNanoSecondImpl(emptyArgList, cty.String)

		if got.AsBigFloat().Cmp(big.NewFloat(0)) < 1 {
			t.Errorf("gofakeitNanoSecondImpl() returns zero")
		}
	})
}

func Test_gofakeitSecondImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitSecondImpl(emptyArgList, cty.String)

		if got.AsBigFloat().Cmp(big.NewFloat(0)) < 1 {
			t.Errorf("gofakeitSecondImpl() returns zero")
		}
	})
}

func Test_gofakeitMinuteImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitMinuteImpl(emptyArgList, cty.String)

		if got.AsBigFloat().Cmp(big.NewFloat(0)) < 1 {
			t.Errorf("gofakeitMinuteImpl() returns zero")
		}
	})
}

func Test_gofakeitHourImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitHourImpl(emptyArgList, cty.String)

		if got.AsBigFloat().Cmp(big.NewFloat(0)) < 0 {
			t.Errorf("gofakeitHourImpl() returns zero")
		}
	})
}

func Test_gofakeitMonthImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitMonthImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitDateImpl() empty string")
		}
	})
}

func Test_gofakeitDayImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitDayImpl(emptyArgList, cty.String)

		if got.AsBigFloat().Cmp(big.NewFloat(0)) < 1 {
			t.Errorf("gofakeitDayImpl() returns zero")
		}
	})
}

func Test_gofakeitWeekDayImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitWeekDayImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitDateImpl() empty string")
		}
	})
}

func Test_gofakeitYearImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitYearImpl(emptyArgList, cty.String)

		if got.AsBigFloat().Cmp(big.NewFloat(0)) < 1 {
			t.Errorf("gofakeitYearImpl() returns zero")
		}
	})
}

func Test_gofakeitTimeZoneImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitTimeZoneImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitTimeZoneImpl() empty string")
		}
	})
}

func Test_gofakeitTimeZoneAbvImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitTimeZoneAbvImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitTimeZoneAbvImpl() empty string")
		}
	})
}

func Test_gofakeitTimeZoneFullImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitTimeZoneFullImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitTimeZoneFullImpl() empty string")
		}
	})
}

func Test_gofakeitTimeZoneOffsetImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		min := big.NewFloat(-12)
		max := big.NewFloat(14)

		got, _ := gofakeitTimeZoneOffsetImpl(emptyArgList, cty.Number)

		if got.AsBigFloat().Cmp(min) < 1 && got.AsBigFloat().Cmp(max) > 0 {
			t.Errorf("gofakeitTimeZoneOffsetImpl() invalid result")
		}
	})
}

func Test_gofakeitTimeZoneRegionImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitTimeZoneRegionImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitTimeZoneRegionImpl() empty string")
		}
	})
}

func Test_gofakeitPriceImpl(t *testing.T) {
	argList := []cty.Value{
		cty.NumberIntVal(0),
		cty.NumberIntVal(100),
	}

	t.Run("success", func(t *testing.T) {
		min := big.NewFloat(0.00)
		max := big.NewFloat(90.00)
		got, _ := gofakeitPriceImpl(argList, cty.String)

		if got.AsBigFloat().Cmp(min) < 1 && got.AsBigFloat().Cmp(max) > 0 {
			t.Errorf("gofakeitPriceImpl() invalid result")
		}
	})
}

func Test_gofakeitCreditCardCvvImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitCreditCardCvvImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitCreditCardCvvImpl() empty string")
		}
	})
}

func Test_gofakeitCreditCardExpImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitCreditCardExpImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitCreditCardExpImpl() empty string")
		}
	})
}

// Possible type values are visa, mastercard, american-express, diners-club, discover, jcb, unionpay, maestro, elo, hiper, hipercard
//  Some regexs
//	Amex Card: ^3[47][0-9]{13}$
//	BCGlobal: ^(6541|6556)[0-9]{12}$
//	Carte Blanche Card: ^389[0-9]{11}$
//	Diners Club Card: ^3(?:0[0-5]|[68][0-9])[0-9]{11}$
//	Discover Card: ^65[4-9][0-9]{13}|64[4-9][0-9]{13}|6011[0-9]{12}|(622(?:12[6-9]|1[3-9][0-9]|[2-8][0-9][0-9]|9[01][0-9]|92[0-5])[0-9]{10})$
//	Insta Payment Card: ^63[7-9][0-9]{13}$
//	JCB Card: ^(?:2131|1800|35\d{3})\d{11}$
//	KoreanLocalCard: ^9[0-9]{15}$
//	Laser Card: ^(6304|6706|6709|6771)[0-9]{12,15}$
//	Maestro Card: ^(5018|5020|5038|6304|6759|6761|6763)[0-9]{8,15}$
//	Mastercard: ^(5[1-5][0-9]{14}|2(22[1-9][0-9]{12}|2[3-9][0-9]{13}|[3-6][0-9]{14}|7[0-1][0-9]{13}|720[0-9]{12}))$
//	Solo Card: ^(6334|6767)[0-9]{12}|(6334|6767)[0-9]{14}|(6334|6767)[0-9]{15}$
//	Switch Card: ^(4903|4905|4911|4936|6333|6759)[0-9]{12}|(4903|4905|4911|4936|6333|6759)[0-9]{14}|(4903|4905|4911|4936|6333|6759)[0-9]{15}|564182[0-9]{10}|564182[0-9]{12}|564182[0-9]{13}|633110[0-9]{10}|633110[0-9]{12}|633110[0-9]{13}$
//	Union Pay Card: ^(62[0-9]{14,17})$
//	Visa Card: ^4[0-9]{12}(?:[0-9]{3})?$
//	Visa Master Card: ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14})$
//  Elo card: ^(?:401178|401179|431274|438935|451416|457393|457631|457632|504175|627780|636297|636368|655000|655001|651652|651653|651654|650485|650486|650487|650488|506699|5067[0-6][0-9]|50677[0-8]|509\d{3})\d{10}$
//  Hibercard: ^(606282\d{10}(\d{3})?)|(3841\d{15})$
func Test_gofakeitCreditCardNumberImpl(t *testing.T) {
	type args struct {
		args    []cty.Value
		retType cty.Type
	}
	tests := []struct {
		skip    bool
		name    string
		args    args
		wantErr bool
		regex   string
	}{
		{false, "visa", args{[]cty.Value{cty.ListVal([]cty.Value{cty.StringVal("visa")})}, cty.String}, false, "^4[0-9]{12}(?:[0-9]{3})?"},
		{false, "mastercard", args{[]cty.Value{cty.ListVal([]cty.Value{cty.StringVal("mastercard")})}, cty.String}, false, "^(5[1-5][0-9]{14}|2(22[1-9][0-9]{12}|2[3-9][0-9]{13}|[3-6][0-9]{14}|7[0-1][0-9]{13}|720[0-9]{12}))$"},
		{false, "american-express", args{[]cty.Value{cty.ListVal([]cty.Value{cty.StringVal("american-express")})}, cty.String}, false, "^3[47][0-9]{13}$"},
		{true, "diners-club", args{[]cty.Value{cty.ListVal([]cty.Value{cty.StringVal("diners-club")})}, cty.String}, false, "^3(?:0[0-5]|[68][0-9])[0-9]{11}$"},
		{false, "discover", args{[]cty.Value{cty.ListVal([]cty.Value{cty.StringVal("discover")})}, cty.String}, false, "^65[4-9][0-9]{13}|64[4-9][0-9]{13}|6011[0-9]{12}|(622(?:12[6-9]|1[3-9][0-9]|[2-8][0-9][0-9]|9[01][0-9]|92[0-5])[0-9]{10})$"},
		{true, "jcb", args{[]cty.Value{cty.ListVal([]cty.Value{cty.StringVal("jcb")})}, cty.String}, false, "^(?:2131|1800|35\\d{3})\\d{11}$"},
		{true, "unionpay", args{[]cty.Value{cty.ListVal([]cty.Value{cty.StringVal("unionpay")})}, cty.String}, false, "^(62[0-9]{14,17})$"},
		{false, "maestro", args{[]cty.Value{cty.ListVal([]cty.Value{cty.StringVal("maestro")})}, cty.String}, false, "^(50|5[6-9]|6[0-9])[0-9]{10,17}$"},
		{true, "elo", args{[]cty.Value{cty.ListVal([]cty.Value{cty.StringVal("elo")})}, cty.String}, false, "^(?:401178|401179|431274|438935|451416|457393|457631|457632|504175|627780|636297|636368|655000|655001|651652|651653|651654|650485|650486|650487|650488|506699|5067[0-6][0-9]|50677[0-8]|509\\d{3})\\d{10}$"},
		{false, "hipercard", args{[]cty.Value{cty.ListVal([]cty.Value{cty.StringVal("hipercard")})}, cty.String}, false, "^(606282\\d{10}(\\d{3})?)|(3841\\d{15})$"},
		{false, "visa with binlist", args{[]cty.Value{cty.ListVal([]cty.Value{cty.StringVal("visa")}), cty.ListVal([]cty.Value{cty.StringVal("400115")})}, cty.String}, false, "^400115[0-9]{7}(?:[0-9]{3})?"},
		{false, "visa with binlist with gaps", args{[]cty.Value{cty.ListVal([]cty.Value{cty.StringVal("visa")}), cty.ListVal([]cty.Value{cty.StringVal("400115")}), cty.BoolVal(true)}, cty.String}, false, "^4[0-9]{12}(?:[0-9]{3})?"},
		{false, "visa with binlist without gaps", args{[]cty.Value{cty.ListVal([]cty.Value{cty.StringVal("visa")}), cty.ListVal([]cty.Value{cty.StringVal("400115")}), cty.BoolVal(false)}, cty.String}, false, "^4[0-9]{12}(?:[0-9]{3})?"},
		{false, "unknown", args{[]cty.Value{cty.ListVal([]cty.Value{cty.StringVal("unknown")})}, cty.String}, true, "^4[0-9]{12}(?:[0-9]{3})?"},
		{false, "empty list", args{[]cty.Value{cty.ListValEmpty(cty.String)}, cty.String}, true, "^4[0-9]{12}(?:[0-9]{3})?"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skip {
				t.Skipf("skipping gofakeitCreditCardNumberImpl() %s", tt.name)
			}

			got, err := gofakeitCreditCardNumberImpl(tt.args.args, tt.args.retType)

			if (err != nil) != tt.wantErr {
				t.Errorf("gofakeitCreditCardNumberImpl() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if tt.wantErr == false {
				valid := regexp.MustCompile(tt.regex)
				gaps := false

				if len(tt.args.args) > 2 {
					gaps = tt.args.args[2].True()
				}

				var asString string
				if gaps {
					asString = strings.ReplaceAll(got.AsString(), " ", "")
				} else {
					asString = got.AsString()
				}

				if !valid.MatchString(asString) {
					t.Errorf("gofakeitCreditCardNumberImpl() produced invalid %s card number %s", tt.name, got.AsString())
				}
			}
		})
	}
}

func Test_gofakeitCreditCardTypeImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitCreditCardTypeImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitCreditCardTypeImpl() empty string")
		}
	})
}

func Test_gofakeitCurrencyLongImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitCurrencyLongImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitCurrencyLongImpl() empty string")
		}
	})
}

func Test_gofakeitCurrencyShortImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitCurrencyShortImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitCurrencyShortImpl() empty string")
		}
	})
}

func Test_gofakeitAchRoutingImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitAchRoutingImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitAchRoutingImpl() empty string")
		}
	})
}

func Test_gofakeitAchAccountImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitAchAccountImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitAchAccountImpl() empty string")
		}
	})
}

func Test_gofakeitBitcoinAddressImpl(t *testing.T) {
	t.Skipf("skipping this until %s gets resolved", "https://github.com/brianvoe/gofakeit/issues/116")

	var emptyArgList []cty.Value

	valid := regexp.MustCompile("(?i)^[13][a-km-zA-HJ-NP-Z0-9]{26,33}$")

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitBitcoinAddressImpl(emptyArgList, cty.String)

		if !valid.MatchString(got.AsString()) {
			t.Errorf("gofakeitBitcoinAddressImpl() invalid bitcoin address")
		}
	})
}

func Test_gofakeitBitcoinPrivateKeyImpl(t *testing.T) {
	t.Skipf("skipping this until %s gets resolved", "https://github.com/brianvoe/gofakeit/issues/116")

	var emptyArgList []cty.Value

	valid := regexp.MustCompile("^5[HJK][1-9A-Za-z][^OIl]{49}")

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitBitcoinPrivateKeyImpl(emptyArgList, cty.String)

		if !valid.MatchString(got.AsString()) {
			t.Errorf("gofakeitBitcoinPrivateKeyImpl() invalid bitcoin private key")
		}
	})
}

func Test_gofakeitBSImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitBSImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitBSImpl() empty string")
		}
	})
}

func Test_gofakeitBuzzWordImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitBuzzWordImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitBuzzWordImpl() empty string")
		}
	})
}

func Test_gofakeitCompanyImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitCompanyImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitCompanyImpl() empty string")
		}
	})
}

func Test_gofakeitCompanySuffixImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitCompanySuffixImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitCompanySuffixImpl() empty string")
		}
	})
}

func Test_gofakeitJobDescriptorImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitJobDescriptorImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitJobDescriptorImpl() empty string")
		}
	})
}

func Test_gofakeitJobLevelImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitJobLevelImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitJobLevelImpl() empty string")
		}
	})
}

func Test_gofakeitJobTitleImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitJobTitleImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitJobTitleImpl() empty string")
		}
	})
}

func Test_gofakeitHackerAbbreviationImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitHackerAbbreviationImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitHackerAbbreviationImpl() empty string")
		}
	})
}

func Test_gofakeitHackerAdjectiveImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitHackerAdjectiveImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitHackerAdjectiveImpl() empty string")
		}
	})
}

func Test_gofakeitHackeringVerbImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitHackeringVerbImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitHackeringVerbImpl() empty string")
		}
	})
}

func Test_gofakeitHackerNounImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitHackerNounImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitHackerNounImpl() empty string")
		}
	})
}

func Test_gofakeitHackerPhraseImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitHackerPhraseImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitHackerPhraseImpl() empty string")
		}
	})
}

func Test_gofakeitHackerVerbImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitHackerVerbImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitHackerVerbImpl() empty string")
		}
	})
}

func Test_gofakeitHipsterWordImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitHipsterWordImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitHipsterWordImpl() empty string")
		}
	})
}

func Test_gofakeitHipsterSentenceImpl(t *testing.T) {
	argList := []cty.Value{cty.NumberIntVal(10)}

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitHipsterSentenceImpl(argList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitHipsterSentenceImpl() empty string")
		}
	})
}

func Test_gofakeitHipsterParagraphImpl(t *testing.T) {
	argList := []cty.Value{
		cty.NumberIntVal(10),
		cty.NumberIntVal(10),
		cty.NumberIntVal(10),
		cty.StringVal(" "),
	}

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitHipsterParagraphImpl(argList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitHipsterParagraphImpl() empty string")
		}
	})
}

func Test_gofakeitAppNameImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitAppNameImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitAppNameImpl() empty string")
		}
	})
}

func Test_gofakeitAppVersionImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitAppVersionImpl(emptyArgList, cty.String)

		_, err := semver.NewVersion(got.AsString())

		if err != nil {
			t.Errorf("gofakeitAppVersionImpl() results invalid semantic versions")
		}
	})
}

func Test_gofakeitAppAuthorImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitAppAuthorImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitAppAuthorImpl() empty string")
		}
	})
}

func Test_gofakeitPetNameImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitPetNameImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitPetNameImpl() empty string")
		}
	})
}

func Test_gofakeitAnimalImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitAnimalImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitAnimalImpl() empty string")
		}
	})
}

func Test_gofakeitAnimalTypeImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitAnimalTypeImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitAnimalTypeImpl() empty string")
		}
	})
}

func Test_gofakeitFarmAnimalImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitFarmAnimalImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitFarmAnimalImpl() empty string")
		}
	})
}

func Test_gofakeitCatImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitCatImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitCatImpl() empty string")
		}
	})
}

func Test_gofakeitDogImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitDogImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitDogImpl() empty string")
		}
	})
}

func Test_gofakeitEmojiImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitEmojiImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitEmojiImpl() empty string")
		}
	})
}

func Test_gofakeitEmojiDescriptionImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitEmojiDescriptionImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitEmojiDescriptionImpl() empty string")
		}
	})
}

func Test_gofakeitEmojiCategoryImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitEmojiCategoryImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitEmojiCategoryImpl() empty string")
		}
	})
}

func Test_gofakeitEmojiAliasImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitEmojiAliasImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitEmojiAliasImpl() empty string")
		}
	})
}

func Test_gofakeitEmojiTagImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitEmojiTagImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitEmojiTagImpl() empty string")
		}
	})
}

func Test_gofakeitLanguageImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitLanguageImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitLanguageImpl() empty string")
		}
	})
}

func Test_gofakeitLanguageAbbreviationImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitLanguageAbbreviationImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitLanguageAbbreviationImpl() empty string")
		}
	})
}

func Test_gofakeitProgrammingLanguageImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitProgrammingLanguageImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitProgrammingLanguageImpl() empty string")
		}
	})
}

func Test_gofakeitProgrammingLanguageBestImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitProgrammingLanguageBestImpl(emptyArgList, cty.String)

		if len(got.AsString()) == 0 {
			t.Errorf("gofakeitProgrammingLanguageBestImpl() empty string")
		}
	})
}

func Test_gofakeitNumberImpl(t *testing.T) {
	argList := []cty.Value{
		cty.NumberIntVal(0),
		cty.NumberIntVal(10),
	}

	t.Run("success", func(t *testing.T) {
		min := big.NewFloat(-90.00)
		max := big.NewFloat(90.00)

		got, _ := gofakeitNumberImpl(argList, cty.Number)

		if got.AsBigFloat().Cmp(min) < 1 && got.AsBigFloat().Cmp(max) > 0 {
			t.Errorf("gofakeitNumberImpl() invalid result")
		}
	})
}

func Test_gofakeitInt8Impl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitInt8Impl(emptyArgList, cty.Number)

		if len(got.AsBigFloat().String()) == 0 {
			t.Errorf("gofakeitInt8Impl() invalid result")
		}
	})
}

func Test_gofakeitInt16Impl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitInt16Impl(emptyArgList, cty.Number)

		if len(got.AsBigFloat().String()) == 0 {
			t.Errorf("gofakeitInt16Impl() invalid result")
		}
	})
}

func Test_gofakeitInt32Impl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitInt32Impl(emptyArgList, cty.Number)

		if len(got.AsBigFloat().String()) == 0 {
			t.Errorf("gofakeitInt32Impl() invalid result")
		}
	})
}

func Test_gofakeitInt64Impl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitInt64Impl(emptyArgList, cty.Number)

		if len(got.AsBigFloat().String()) == 0 {
			t.Errorf("gofakeitInt64Impl() invalid result")
		}
	})
}

func Test_gofakeitUint8Impl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitUint8Impl(emptyArgList, cty.Number)

		if len(got.AsBigFloat().String()) == 0 {
			t.Errorf("gofakeitUint8Impl() invalid result")
		}
	})
}

func Test_gofakeitUint16Impl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitUint16Impl(emptyArgList, cty.Number)

		if len(got.AsBigFloat().String()) == 0 {
			t.Errorf("gofakeitUint16Impl() invalid result")
		}
	})
}

func Test_gofakeitUint32Impl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitUint32Impl(emptyArgList, cty.Number)

		if len(got.AsBigFloat().String()) == 0 {
			t.Errorf("gofakeitUint32Impl() invalid result")
		}
	})
}

func Test_gofakeitUint64Impl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitUint64Impl(emptyArgList, cty.Number)

		if len(got.AsBigFloat().String()) == 0 {
			t.Errorf("gofakeitUint64Impl() invalid result")
		}
	})
}

func Test_gofakeitFloat32Impl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitFloat32Impl(emptyArgList, cty.Number)

		if len(got.AsBigFloat().String()) == 0 {
			t.Errorf("gofakeitFloat32Impl() invalid result")
		}
	})
}

func Test_gofakeitFloat32RangeImpl(t *testing.T) {
	argList := []cty.Value{
		cty.NumberIntVal(0),
		cty.NumberIntVal(100),
	}

	t.Run("success", func(t *testing.T) {
		min := big.NewFloat(0.00)
		max := big.NewFloat(100.00)

		got, _ := gofakeitFloat32RangeImpl(argList, cty.Number)

		if got.AsBigFloat().Cmp(min) < 1 && got.AsBigFloat().Cmp(max) > 0 {
			t.Errorf("gofakeitFloat32RangeImpl() invalid result")
		}
	})
}

func Test_gofakeitFloat64Impl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitFloat64Impl(emptyArgList, cty.Number)

		if len(got.AsBigFloat().String()) == 0 {
			t.Errorf("gofakeitFloat64Impl() invalid result")
		}
	})
}

func Test_gofakeitFloat64RangeImpl(t *testing.T) {
	argList := []cty.Value{
		cty.NumberIntVal(0),
		cty.NumberIntVal(100),
	}

	t.Run("success", func(t *testing.T) {
		min := big.NewFloat(0.00)
		max := big.NewFloat(100.00)

		got, _ := gofakeitFloat64RangeImpl(argList, cty.Number)

		if got.AsBigFloat().Cmp(min) < 1 && got.AsBigFloat().Cmp(max) > 0 {
			t.Errorf("gofakeitFloat64RangeImpl() invalid result")
		}
	})
}

func Test_gofakeitShuffleIntsImpl(t *testing.T) {
	deck := []cty.Value{
		cty.NumberIntVal(1),
		cty.NumberIntVal(2),
		cty.NumberIntVal(3),
		cty.NumberIntVal(4),
	}

	argList := []cty.Value{
		cty.ListVal(deck),
	}

	t.Run("same length", func(t *testing.T) {
		got, _ := gofakeitShuffleIntsImpl(argList, cty.Number)

		valueSlice := got.AsValueSlice()

		if len(valueSlice) != len(deck) {
			t.Errorf("gofakeitShuffleIntsImpl() invalid result deck")
		}
	})

	t.Run("shuffled", func(t *testing.T) {
		got, _ := gofakeitShuffleIntsImpl(argList, cty.Number)
		valueSlice := got.AsValueSlice()

		equal := true
		for i := range valueSlice {
			if valueSlice[i].AsBigFloat().Cmp(deck[i].AsBigFloat()) != 0 {
				equal = false
			}
		}

		if equal {
			t.Errorf("gofakeitShuffleIntsImpl() invalid result deck")
		}
	})
}

func Test_gofakeitRandomIntImpl(t *testing.T) {
	deck := []cty.Value{
		cty.NumberIntVal(1),
		cty.NumberIntVal(2),
		cty.NumberIntVal(3),
		cty.NumberIntVal(4),
	}

	argList := []cty.Value{
		cty.ListVal(deck),
	}

	t.Run("same length", func(t *testing.T) {
		min := big.NewFloat(1)
		max := big.NewFloat(4)

		got, _ := gofakeitRandomIntImpl(argList, cty.Number)

		if got.AsBigFloat().Cmp(min) < 1 && got.AsBigFloat().Cmp(max) > -1 {
			t.Errorf("gofakeitRandomIntImpl() invalid result")
		}
	})
}

func Test_gofakeitDigitImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitDigitImpl(emptyArgList, cty.String)

		if len(got.AsString()) != 1 {
			t.Errorf("gofakeitDigitImpl() invalid result")
		}
	})
}

func Test_gofakeitLetterImpl(t *testing.T) {
	var emptyArgList []cty.Value

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitLetterImpl(emptyArgList, cty.String)

		if len(got.AsString()) != 1 {
			t.Errorf("gofakeitLetterImpl() invalid result")
		}
	})
}

func Test_gofakeitLexifyImpl(t *testing.T) {
	argList := []cty.Value{
		cty.StringVal("erm?ohno??"),
	}

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitLexifyImpl(argList, cty.String)

		if len(got.AsString()) != 10 {
			t.Errorf("gofakeitLexifyImpl() invalid result")
		}
	})
}

func Test_gofakeitNumerifyImpl(t *testing.T) {
	argList := []cty.Value{
		cty.StringVal("erm#ohno##"),
	}

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitNumerifyImpl(argList, cty.String)

		if len(got.AsString()) != 10 {
			t.Errorf("gofakeitNumerifyImpl() invalid result")
		}
	})
}

func Test_gofakeitShuffleStringsImpl(t *testing.T) {
	var deck []cty.Value

	for i := 0; i < 1000; i++ {
		deck = append(deck, cty.StringVal(gofakeit.Letter()))
	}

	argList := []cty.Value{
		cty.ListVal(deck),
	}

	t.Run("same length", func(t *testing.T) {
		got, _ := gofakeitShuffleStringsImpl(argList, cty.Number)

		valueSlice := got.AsValueSlice()

		if len(valueSlice) != len(deck) {
			t.Errorf("gofakeitShuffleStringsImpl() invalid result deck")
		}
	})

	t.Run("shuffled", func(t *testing.T) {
		got, _ := gofakeitShuffleStringsImpl(argList, cty.Number)
		valueSlice := got.AsValueSlice()

		equal := true
		for i := range valueSlice {
			if valueSlice[i].AsString() != deck[i].AsString() {
				equal = false
			}
		}

		if equal {
			t.Errorf("gofakeitShuffleStringsImpl() invalid result deck")
		}
	})
}

func Test_gofakeitRandomStringImpl(t *testing.T) {
	deck := []cty.Value{
		cty.StringVal("a"),
		cty.StringVal("b"),
		cty.StringVal("c"),
		cty.StringVal("d"),
		cty.StringVal("e"),
	}

	argList := []cty.Value{
		cty.ListVal(deck),
	}

	t.Run("success", func(t *testing.T) {
		got, _ := gofakeitRandomStringImpl(argList, cty.Number)

		found := false
		for i := range deck {
			if deck[i].AsString() == got.AsString() {
				found = true
			}
		}

		if !found {
			t.Errorf("gofakeitRandomStringImpl() invalid result deck")
		}
	})
}
