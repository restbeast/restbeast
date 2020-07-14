package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"github.com/zclconf/go-cty/cty/function/stdlib"
	"github.com/zclconf/go-cty/cty/gocty"
	ctyjson "github.com/zclconf/go-cty/cty/json"
	"os"
	"reflect"
	"regexp"
)

func getObjSpec() hcldec.ObjectSpec {
	return hcldec.ObjectSpec{
		"method": &hcldec.AttrSpec{
			Name:     "method",
			Required: true,
			Type:     cty.String,
		},
		"url": &hcldec.AttrSpec{
			Name:     "url",
			Required: true,
			Type:     cty.String,
		},
		"headers": &hcldec.AttrSpec{
			Name:     "headers",
			Required: false,
			Type:     cty.Map(cty.String),
		},
		"body": &hcldec.AttrSpec{
			Name:     "body",
			Required: false,
			Type:     cty.DynamicPseudoType,
		},
		"depends_on": &hcldec.AttrSpec{
			Name:     "depends_on",
			Required: false,
			Type:     cty.List(cty.String),
		},
	}
}

func getEvalContext(variables map[string]cty.Value, envVars cty.Value, requestAsVars map[string]cty.Value) hcl.EvalContext {
	return hcl.EvalContext{
		Variables: map[string]cty.Value{
			"var": cty.ObjectVal(variables),
			"env": envVars,
			"request": cty.ObjectVal(requestAsVars),
		},
		Functions: map[string]function.Function{
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
		},
	}
}

func findRequest(name string, rawRequests []*RequestCfg) (err error, request RequestCfg) {
	for _, r := range rawRequests {
		if name == r.Name {
			return nil, *r
		}
	}

	return errors.New("request not found"), RequestCfg{}
}

func getPossibleDependencies(diags hcl.Diagnostics) (dependencies []string) {
	if len(diags) != 0 {
		diagMessageRegex := regexp.MustCompile(`This object does not have an attribute named "(?P<name>[\w\d-_]+)"`)

		for _, diag := range diags {
			if diag.Summary == "Unsupported attribute" {
				findString := diagMessageRegex.FindStringSubmatch(diag.Detail)

				if len(findString) > 1 {
					dependencies = append(dependencies, findString[1])
				}
			}
		}

		if len(dependencies) == 0 {
			for _, diag := range diags {
				fmt.Printf("- %s\n", diag)
			}

			os.Exit(0)
		}
	}

	return dependencies
}

func getUniqueDependencies(intSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func processDependencies(dependencies []string, variables map[string]cty.Value, envVars cty.Value, version string, rawRequests RequestCfgs) (requestAsVars map[string]cty.Value) {
	requestAsVars = make(map[string]cty.Value)

	for _, dep := range getUniqueDependencies(dependencies) {
		request := parseRequest(dep, variables, envVars, version, rawRequests)
		response := DoRequest(request, version)

		var decoded interface{}
		err := json.Unmarshal(response.Body, &decoded)

		if err != nil {
			fmt.Println("Error: error decoding json response body", err)
			os.Exit(0)
		}

		requestAsVars[dep] = walkThrough(reflect.ValueOf(decoded))
	}

	return requestAsVars
}

func getRequest(cfg cty.Value) Request {
	body, _ := json.MarshalIndent(ctyjson.SimpleJSONValue{cfg.GetAttr("body")}, "", "  ")
	var headers map[string]string
	_ = gocty.FromCtyValue(cfg.GetAttr("headers"), &headers)

	request := Request{
		Method:  cfg.GetAttr("method").AsString(),
		Url:     cfg.GetAttr("url").AsString(),
		Headers: headers,
		Body:    string(body),
	}

	return request
}

func parseRequest(name string, variables map[string]cty.Value, envVars cty.Value, version string, rawRequests RequestCfgs) Request {
  err, request := findRequest(name, rawRequests)

  if err != nil {
  	fmt.Println("Error: Request not found")
  	os.Exit(1)
	}

	requestAsVars := map[string]cty.Value{}
	evalContext := getEvalContext(variables, envVars, requestAsVars)
	spec := getObjSpec()

	cfg, diags := hcldec.Decode(request.Body, spec, &evalContext)
	dependencies := getPossibleDependencies(diags)

	if len(dependencies) > 0 {
		requestAsVars := processDependencies(dependencies, variables, envVars, version, rawRequests)
		evalContext = getEvalContext(variables, envVars, requestAsVars)

		cfg, diags = hcldec.Decode(request.Body, spec, &evalContext)

		if len(diags) > 0 {
			for _, diag := range diags {
				fmt.Printf("- %s\n", diag)
			}

			os.Exit(0)
		}
	}

	return getRequest(cfg)
}
