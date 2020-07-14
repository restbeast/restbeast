https://github.com/brianvoe/gofakeit#person

# Functions

Restbeast implements almost all of the [gofakeit](https://github.com/brianvoe/gofakeit#functions) functions

## Person

**gofakeitName**(): *String*  

**gofakeitNamePrefix**(): *String*  

**gofakeitNameSuffix**(): *String*  

**gofakeitFirstName**(): *String*  

**gofakeitLastName**(): *String*  

**gofakeitGender**(): *String*  

**gofakeitSSN**(): *String*  

**gofakeitEmail**(): *String*  

**gofakeitPhone**(): *String*  

**gofakeitPhoneFormatted**(): *String*  

## Auth

**gofakeitUsername**(): *String*

**gofakeitPassword**(lower *Bool*, upper *Bool*, numeric *Bool*, special *Bool*, space *Bool*, length *Number*): *String*

## Address

**gofakeitCity**(): *String*

**gofakeitCountry**(): *String*

**gofakeitCountryAbr**(): *String*

**gofakeitState**(): *String*

**gofakeitStateAbr**(): *String*

**gofakeitStreet**(): *String*

**gofakeitStreetName**(): *String*

**gofakeitStreetNumber**(): *String*

**gofakeitStreetPrefix**(): *String*

**gofakeitStreetSuffix**(): *String*

**gofakeitZip**(): *String*

**gofakeitLatitude**(): *Number*

**gofakeitLatitudeInRange**(min *Float*, max *Float*): *Number*

**gofakeitLongitude**(): *Number*

**gofakeitLongitudeInRange**(min *Float*, max *Float*): *Number*

## Gamer

**gofakeitGamertag**(): *String*

**gofakeitBeerAlcohol**(): *string*

**gofakeitBeerBlg**(): *string*

**gofakeitBeerHop**(): *string*

**gofakeitBeerIbu**(): *string*

**gofakeitBeerMalt**(): *string*

**gofakeitBeerName**(): *string*

**gofakeitBeerStyle**(): *string*

**gofakeitBeerYeast**(): *string*

## Car

**gofakeitCarMaker**(): *string*

**gofakeitCarModel**(): *string*

**gofakeitCarType**(): *string*

**gofakeitCarFuelType**(): *string*

**gofakeitCarTransmissionType**(): *string*

## Word 

**gofakeitNoun**(): *string*

**gofakeitVerb**(): *string*

**gofakeitAdverb**(): *string*

**gofakeitPreposition**(): *string*

**gofakeitAdjective**(): *string*

**gofakeitWord**(): *string*

**gofakeitSentence**(wordCount int): *string*

**gofakeitParagraph**(paragraphCount *int*, sentenceCount *int*, wordCount *int*, separator *string*): *string*

**gofakeitLoremIpsumWord**(): *string*

**gofakeitLoremIpsumSentence**(wordCount int): *string*

**gofakeitLoremIpsumParagraph**(paragraphCount *int*, sentenceCount *int*, wordCount *int*, separator *string*): *string*

**gofakeitQuestion**(): *string*

**gofakeitQuote**(): *string*

**gofakeitPhrase**(): *string*

## Food

**gofakeitFruit**(): *string*

**gofakeitVegetable**(): *string*

**gofakeitBreakfast**(): *string*

**gofakeitLunch**(): *string*

**gofakeitDinner**(): *string*

**gofakeitSnack**(): *string*

**gofakeitDessert**(): *string*

## Misc

**gofakeitBool**(): *bool*

**gofakeitUUID**(): *string*

## Color

**gofakeitColor**(): *string*

**gofakeitHexColor**(): *string*

**gofakeitRGBColor**(): *[]int*

**gofakeitSafeColor**(): *string*

## Internet

**gofakeitURL**(): *string*

**gofakeitImageURL**(width *int*, height *int*): *string*

**gofakeitDomainName**(): *string*

**gofakeitDomainSuffix**(): *string*

**gofakeitIPv4Address**(): *string*

**gofakeitIPv6Address**(): *string*

**gofakeitStatusCode**(): *string*

**gofakeitSimpleStatusCode**(): *int*

**gofakeitLogLevel**(logType *string*): *string*

**gofakeitHTTPMethod**(): *string*

**gofakeitUserAgent**(): *string*

**gofakeitChromeUserAgent**(): *string*

**gofakeitFirefoxUserAgent**(): *string*

**gofakeitOperaUserAgent**(): *string*

**gofakeitSafariUserAgent**(): *string*

**gofakeitDate**(): *string*

**gofakeitDateRange**(start *string*, end *string*) *string*  
Start and end dates are in ISO-8601 format.  
Example: 2020-07-15T00:01:20+00:00

**gofakeitNanoSecond**(): *int*

**gofakeitSecond**(): *int*

**gofakeitMinute**(): *int*

**gofakeitHour**(): *int*

**gofakeitMonth**(): *string*

**gofakeitDay**(): *int*

**gofakeitWeekDay**(): *string*

**gofakeitYear**(): *int*

**gofakeitTimeZone**(): *string*

**gofakeitTimeZoneAbv**(): *string*

**gofakeitTimeZoneFull**(): *string*

**gofakeitTimeZoneOffset**(): *float32*

**gofakeitTimeZoneRegion**(): *string*

**gofakeitPrice**(min *float*, max *float*): *float*

## Payment

**gofakeitCreditCardCvv**(): *string*

**gofakeitCreditCardExp**(): *string*

**gofakeitCreditCardNumber**(types *[]string*, [bins *[]string*, \[gaps *bool*\]]) *string*  
Possible type values are visa, mastercard, american-express, diners-club, discover, jcb, unionpay, maestro, elo, hiper, hipercard

**gofakeitCreditCardType**(): *string*

**gofakeitCurrency**(): **CurrencyInfo*

**gofakeitCurrencyLong**(): *string*

**gofakeitCurrencyShort**(): *string*

**gofakeitAchRouting**(): *string*

**gofakeitAchAccount**(): *string*

**gofakeitBitcoinAddress**(): *string*

**gofakeitBitcoinPrivateKey**(): *string*

## Company

**gofakeitBS**(): *string*

**gofakeitBuzzWord**(): *string*

**gofakeitCompany**(): *string*

**gofakeitCompanySuffix**(): *string*

**gofakeitJob**(): **JobInfo*

**gofakeitJobDescriptor**(): *string*

**gofakeitJobLevel**(): *string*

**gofakeitJobTitle**(): *string*

## Hacker

**gofakeitHackerAbbreviation**(): *string*

**gofakeitHackerAdjective**(): *string*

**gofakeitHackeringVerb**(): *string*

**gofakeitHackerNoun**(): *string*

**gofakeitHackerPhrase**(): *string*

**gofakeitHackerVerb**(): *string*

## Hipster

**gofakeitHipsterWord**(): *string*

**gofakeitHipsterSentence**(wordCount int): *string*

**gofakeitHipsterParagraph**(paragraphCount *int*, sentenceCount *int*, wordCount *int*, separator *string*): *string*

## App

**gofakeitAppName**(): *string*

**gofakeitAppVersion**(): *string*

**gofakeitAppAuthor**(): *string*

## Animal

**gofakeitPetName**(): *string*

**gofakeitAnimal**(): *string*

**gofakeitAnimalType**(): *string*

**gofakeitFarmAnimal**(): *string*

**gofakeitCat**(): *string*

**gofakeitDog**(): *string*

## Emoji

**gofakeitEmoji**(): *string*

**gofakeitEmojiDescription**(): *string*

**gofakeitEmojiCategory**(): *string*

**gofakeitEmojiAlias**(): *string*

**gofakeitEmojiTag**(): *string*

## Language

**gofakeitLanguage**(): *string*

**gofakeitLanguageAbbreviation**(): *string*

**gofakeitProgrammingLanguage**(): *string*

**gofakeitProgrammingLanguageBest**(): *string*

## Number

**gofakeitNumber**(min *int*, max *int*): *int*

**gofakeitInt8**(): *int8*

**gofakeitInt16**(): *int16*

**gofakeitInt32**(): *int32*

**gofakeitInt64**(): *int64*

**gofakeitUint8**(): *uint8*

**gofakeitUint16**(): *uint16*

**gofakeitUint32**(): *uint32*

**gofakeitUint64**(): *uint64*

**gofakeitFloat32**(): *float32*

**gofakeitFloat32Range**(min *float32*, max *float32*): *float32*

**gofakeitFloat64**(): *float64*

**gofakeitFloat64Range**(min *float64*, max *float64*): *float64*

**gofakeitShuffleInts**(deck *[]int*): *[]int* 

**RandomInt**(deck *[]int*) *int*

**gofakeitDigit**(): *string*

## String

**gofakeitLetter**(): *string*

**gofakeitLexify**(str string): *string*

**gofakeitNumerify**(str string): *string*

**gofakeitShuffleStrings**(deck *[]string*): *[]string*
 
**gofakeitRandomString**(deck *[]string*) *string*
