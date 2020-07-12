# Functions

Restbeast implements almost all of the [go-cty](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib) functions

## General Function

**equal**(val1 *Dynamic*, val2 *Dynamic*): *Bool*  
Equal determines whether the two given values are equal, returning a bool value. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Equal)

**notEqual**(val1 *Dynamic*, val2 *Dynamic*): *Bool*  
NotEqual is the opposite of Equal. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#NotEqual)

**coalesce**(val1 *Dynamic*, val2 *Dynamic*, ...): *Dynamic*  
Coalesce returns the first of the given arguments that is not null. If all arguments are null, an error is produced. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Coalesce) 

**not**(param *Bool*): *Bool*  
Not returns the logical complement of the given boolean value. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Not)

**and**(param1 *Bool*, param2 *Bool*): *Bool*  
And returns true if and only if both of the given boolean values are true. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#And)

**or**(param1 *Bool*, param2 *Bool*): *Bool*  
Or returns true if either of the given boolean values are true. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Or)

## String Functions

**upper**(str *String*): *String*  
Upper is a Function that converts a given string to uppercase. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Upper)

**lower**(str *String*): *String*  
Lower is a Function that converts a given string to lowercase. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Lower)

**reverse**(str *String*): *String*  
Reverse is a Function that reverses the order of the characters in the given string. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Reverse)

**strlen**(str *String*): *Number*  
Strlen is a Function that returns the length of the given string in characters. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Strlen)

**substr**(str *String*, offset *Number*, length *Number*): *String*  
Substr is a Function that extracts a sequence of characters from another string and creates a new string. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Substr)

**join**(separator *String*, list1 *List*, list2 *List*...): *String*  
Join concatenates together the string elements of one or more lists with a given separator. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Join)

**sort**(list *List*): *List*  
Sort re-orders the elements of a given list of strings so that they are in ascending lexicographical order. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Sort) 

**chomp**(str *String*): *String*  
Chomp removes newline characters at the end of a string. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Chomp)

**indent**(str *String*): *String*  
Indent adds a given number of spaces to the beginnings of all but the first line in a given multi-line string. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Indent)

**title**(str *String*): *String*  
Title converts the first letter of each word in the given string to uppercase. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Title)

**trimSpace**(str *String*): *String*  
Trim removes the specified characters from the start and end of the given string. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#TrimSpace)

**trim**(str *String*): *String*  
Trim removes the specified characters from the start and end of the given string. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Trim)

**trimPrefix**(str *String*): *String*  
TrimPrefix removes the specified prefix from the start of the given string. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#TrimPrefix)
 
**trimSuffix**(str *String*): *String*  
TrimSuffix removes the specified suffix from the end of the given string. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#TrimSuffix)

**format**(format *String*, val1 *Dynamic*, val2 *Dynamic*, ...): *String*  
Format produces a string representation of zero or more values using a format string similar to the "printf" function in C. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Format)

It supports the following "verbs":

|   |   |
|---|---|
|%%|Literal percent sign, consuming no value|
|%v|A default formatting of the value based on type, as described below.|
|%#v|JSON serialization of the value|
|%t|Converts to boolean and then produces "true" or "false"|
|%b|Converts to number, requires integer, produces binary representation|
|%d|Converts to number, requires integer, produces decimal representation|
|%o|Converts to number, requires integer, produces octal representation|
|%x|Converts to number, requires integer, produces hexadecimal representation with lowercase letters|
|%X|Like %x but with uppercase letters|
|%e|Converts to number, produces scientific notation like -1.234456e+78|
|%E|Like %e but with an uppercase "E" representing the exponent|
|%f|Converts to number, produces decimal representation with fractional part but no exponent, like 123.456|
|%g|%e for large exponents or %f otherwise|
|%G|%E for large exponents or %f otherwise|
|%s|Converts to string and produces the string's characters|
|%q|Converts to string and produces JSON-quoted string representation, like %v.|

The default format selections made by %v are:

|   |   |
|---|---|
|string|%s|
|number|%g|
|bool|%t|
|other|%#v|

Null values produce the literal keyword "null" for %v and %#v, and produce an error otherwise.

Width is specified by an optional decimal number immediately preceding the verb letter. If absent, the width is whatever is necessary to represent the value. Precision is specified after the (optional) width by a period followed by a decimal number. If no period is present, a default precision is used. A period with no following number is invalid. For examples:

|   |   |
|---|---|
|%f|default width, default precision|
|%9f|width 9, default precision|
|%.2f|default width, precision 2|
|%9.2f|width 9, precision 2|

Width and precision are measured in unicode characters (grapheme clusters).

For most values, width is the minimum number of characters to output, padding the formatted form with spaces if necessary.

For strings, precision limits the length of the input to be formatted (not the size of the output), truncating if necessary.

For numbers, width sets the minimum width of the field and precision sets the number of places after the decimal, if appropriate, except that for %g/%G precision sets the total number of significant digits.

The following additional symbols can be used immediately after the percent introducer as flags:

|   |   |
|---|---|
| |(a space) leave a space where the sign would be if number is positive|
|+|Include a sign for a number even if it is positive (numeric only)|
|-|Pad with spaces on the left rather than the right|
|0|Pad with zeros rather than spaces.|

Flag characters are ignored for verbs that do not support them.

By default, % sequences consume successive arguments starting with the first. Introducing a [n] sequence immediately before the verb letter, where n is a decimal integer, explicitly chooses a particular value argument by its one-based index. Subsequent calls without an explicit index will then proceed with n+1, n+2, etc.

An error is produced if the format string calls for an impossible conversion or accesses more values than are given. An error is produced also for an unsupported format verb. 

**formatList**(format *String*, val1 *Dynamic*, val2 *Dynamic*, ...): *String*  
FormatList applies the same formatting behavior as Format, but accepts a mixture of list and non-list values as arguments. Any list arguments passed must have the same length, which dictates the length of the resulting list. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#FormatList)

Any non-list arguments are used repeatedly for each iteration over the list arguments. The list arguments are iterated in order by key, so corresponding items are formatted together. 

**replace**(str *String*, substr *String*, replace *String*): *String*  
Replace searches a given string for another given substring, and replaces all occurrences with a given replacement string. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Replace) 

## Number Functions

**absolute**(num *Number*): *Number*  
Absolute returns the magnitude of the given number, without its sign. That is, it turns negative values into positive values. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Absolute)
 
**add**(num1 *Number*, num2 *Number*): *Number*  
Add returns the sum of the two given numbers. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Add) 
 
**subtract**(num1 *Number*, num2 *Number*): *Number*  
Subtract returns the difference between the two given numbers. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Subtract)

**multiply**(num1 *Number*, num2 *Number*): *Number*  
Multiply returns the product of the two given numbers. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Multiply)
 
**divide**(num1 *Number*, num2 *Number*): *Number*  
Divide returns a divided by b, where both a and b are numbers. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Divide)

**modulo**(num1 *Number*, num2 *Number*): *Number*  
Modulo returns the remainder of a divided by b under integer division, where both a and b are numbers. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Modulo)

**greaterThan***(num1 *Number*, num2 *Number*): *Number*  
GreaterThan returns true if a is less than b. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#GreaterThan)

**greaterThanOrEqualTo***(num1 *Number*, num2 *Number*): *Number*  
GreaterThanOrEqualTo returns true if a is less than b. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#GreaterThanOrEqualTo)
   
**lessThan***(num1 *Number*, num2 *Number*): *Number*  
LessThan returns true if a is less than b. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#LessThan)
   
**lessThanOrEqualTo***(num1 *Number*, num2 *Number*): *Number*  
LessThanOrEqualTo returns true if a is less than b. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#LessThanOrEqualTo)

**negate**(num *Number*): *Number*  
Negate returns the given number multipled by -1. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Negate)

**min**(num1 *Number*, num2 *Number*, ...): *Number*  
Min returns the minimum number from the given numbers. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Min)
  
**max**(num1 *Number*, num2 *Number*, ...): *Number*  
Max returns the maximum number from the given numbers. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Max)
  
**int**(num *Number*): *Number*  
Int removes the fractional component of the given number returning an integer representing the whole number component, rounding towards zero. For example, -1.5 becomes -1. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Int)

If an infinity is passed to Int, an error is returned.

**ceil**(num *Number*): *Number*  
Ceil returns the closest whole number greater than or equal to the given value. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Ceil)
   
**floor**(num *Number*): *Number*  
Floor returns the closest whole number lesser than or equal to the given value. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Floor)
  
**log**(num *Number*, base *Number*): *Number*  
Log returns returns the logarithm of a given number in a given base. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Log)
  
**pow**(num *Number*, power *Number*): *Number*  
Pow returns the logarithm of a given number in a given base. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Pow)
  
**signum**(num *Number*): *Number*  
Signum determines the sign of a number, returning a number between -1 and 1 to represent the sign. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Signum) 
 
**parseInt**(num *Number*, base *Number*): *Number*  
ParseInt parses a string argument and returns an integer of the specified base. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#ParseInt)

## Collection Functions

**hasIndex**(collection *Collection*, key *Dynamic*): *Bool*  
HasIndex determines whether the given collection can be indexed with the given key. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#HasIndex) 

**index**(collection *Collection*, key *Dynamic*): *Dynamic*  
Index returns an element from the given collection using the given key, or returns an error if there is no element for the given key. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Index)

**length**(collection *Collection*): *Number*  
Length returns the number of elements in the given collection. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Length)

**element**(list *List*, index *Number*): *Dynamic*  
Element returns a single element from a given list at the given index. If index is greater than the length of the list then it is wrapped modulo the list length. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Element) 

**coalesceList**(list1 *List*, list2 *List*, ...): *List*  
CoalesceList takes any number of list arguments and returns the first one that isn't empty. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#CoalesceList)

**compact**(list *List*): *List*  
Compact takes a list of strings and returns a new list with any empty string elements removed.  [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Compact) 

**contains**(list *List*, value *Dynamic*): *Bool*    
Contains determines whether a given list contains a given single value as one of its elements. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Contains)

**distinct**(list *List*): *List*  
Distinct takes a list and returns a new list with any duplicate elements removed. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Distinct) 

**chunklist**(list *List*, size *Number*): *List*  
Chunklist splits a single list into fixed-size chunks, returning a list of lists. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Chunklist)

**flatten**(list *List*): *List*  
Flatten takes a list and replaces any elements that are lists with a flattened sequence of the list contents. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Flatten)

**keys**(collection *Map*): *List*  
Keys takes a map and returns a sorted list of the map keys. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Keys)

**lookup**(map *Map*, key *Dynamic*, defaultValue *Dynamic*): *Dynamic*  
Lookup performs a dynamic lookup into a map. There are two required arguments, map and key, plus an optional default, which is a value to return if no key is found in map. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Lookup) 

**merge**(map1 *Map*, map2 *Map*, ...): *Map*  
Merge takes an arbitrary number of maps and returns a single map that contains a merged set of elements from all of the maps. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Merge)

If more than one given map defines the same key then the one that is later in the argument sequence takes precedence. 

**reverseList**(list *List*): *List*  
ReverseList takes a sequence and produces a new sequence of the same length with all of the same elements as the given sequence but in reverse order. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#ReverseList) 
 
**setProduct**(set1 *Set*, set2 *Set*, ...): *Set*  
SetProduct computes the Cartesian product of sets or sequences. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#SetProduct) 

**slice**(list *List*, start *Number*, end *Number*): *List*  
Slice extracts some consecutive elements from within a list. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Slice)

**values**(map *Map*): *List*  
Values returns a list of the map values, in the order of the sorted keys. This function only works on flat maps. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Values)

**zipmap**(keys *List*, values *List*): *Map*  
Zipmap constructs a map from a list of keys and a corresponding list of values. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#ZipMap)

**concat**(list1 *List*, list2 *List*, ...): *List*  
Concat takes one or more sequences (lists or tuples) and returns the single sequence that results from concatenating them together in order.

If all of the given sequences are lists of the same element type then the result is a list of that type. Otherwise, the result is a of a tuple type constructed from the given sequence types. 

**range**(val1 *Number*, val2 *Number*, ...): *List*  
 Range creates a list of numbers by starting from the given starting value, then adding the given step value until the result is greater than or equal to the given stopping value. Each intermediate result becomes an element in the resulting list.

When all three parameters are set, the order is (start, end, step). If only two parameters are set, they are the start and end respectively and step defaults to 1. If only one argument is set, it gives the end value with start defaulting to 0 and step defaulting to 1.

Because the resulting list must be fully buffered in memory, there is an artificial cap of 1024 elements, after which this function will return an error to avoid consuming unbounded amounts of memory. The Range function is primarily intended for creating small lists of indices to iterate over, so there should be no reason to generate huge lists with it. 

## Encoding Functions

**csvDecode**(csv *String*): *Dynamic*  
CSVDecode parses the given CSV (RFC 4180) string and, if it is valid, returns a list of objects representing the rows. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#CSVDecode)

The result is always a list of some object type. The first row of the input is used to determine the object attributes, and subsequent rows determine the values of those attributes.

**jsonEncode**(val *Dynamic*): *String*  
JSONEncode returns a JSON serialization of the given value. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#JSONEncode)
  
**jsonDecode**(str *String*): *Dynamic*  
JSONDecode parses the given JSON string and, if it is valid, returns the value it represents. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#JSONDecode)

Note that applying JSONDecode to the result of JSONEncode may not produce an identically-typed result, since JSON encoding is lossy for cty Types. The resulting value will consist only of primitive types, object types, and tuple types. 

## DateTime Functions

**formatDate**(format *String*, timestamp *String*): *String*  
FormatDate reformats a timestamp given in [RFC 3339](http://tools.ietf.org/html/rfc3339) syntax into another time syntax defined by a given format string.

The format string uses letter mnemonics to represent portions of the timestamp, with repetition signifying length variants of each portion. Single quote characters ' can be used to quote sequences of literal letters that should not be interpreted as formatting mnemonics.

The full set of supported mnemonic sequences is listed below:

|   |   |
|---|---|
|YY|Year modulo 100 zero-padded to two digits, like "06".|
|YYYY|Four (or more) digit year, like "2006".|
|M|Month number, like "1" for January.|
|MM|Month number zero-padded to two digits, like "01".|
|MMM|English month name abbreviated to three letters, like "Jan".|
|MMMM|English month name unabbreviated, like "January".|
|D|Day of month number, like "2".|
|DD|Day of month number zero-padded to two digits, like "02".|
|EEE|English day of week name abbreviated to three letters, like "Mon".|
|EEEE|English day of week name unabbreviated, like "Monday".|
|h|24-hour number, like "2".|
|hh|24-hour number zero-padded to two digits, like "02".|
|H|12-hour number, like "2".|
|HH|12-hour number zero-padded to two digits, like "02".|
|AA|Hour AM/PM marker in uppercase, like "AM".|
|aa|Hour AM/PM marker in lowercase, like "am".|
|m|Minute within hour, like "5".|
|mm|Minute within hour zero-padded to two digits, like "05".|
|s|Second within minute, like "9".|
|ss|Second within minute zero-padded to two digits, like "09".|
|ZZZZ|Timezone offset with just sign and digit, like "-0800".|
|ZZZZZ|Timezone offset with colon separating hours and minutes, like "-08:00".|
|Z|Like ZZZZZ but with a special case "Z" for UTC.|
|ZZZ|Like ZZZZ but with a special case "UTC" for UTC.|

The format syntax is optimized mainly for generating machine-oriented timestamps rather than human-oriented timestamps; the English language portions of the output reflect the use of English names in a number of machine-readable date formatting standards. For presentation to humans, a locale-aware time formatter (not included in this package) is a better choice.

The format syntax is not compatible with that of any other language, but is optimized so that patterns for common standard date formats can be recognized quickly even by a reader unfamiliar with the format syntax. 

**timeAdd**(timestamp *String*, duration *String*): *String*    
TimeAdd adds a duration to a timestamp, returning a new timestamp.

In the HCL language, timestamps are conventionally represented as strings using RFC 3339 "Date and Time format" syntax. Timeadd requires the timestamp argument to be a string conforming to this syntax.

`duration` is a string representation of a time difference, consisting of sequences of number and unit pairs, like `"1.5h"` or `1h30m`. The accepted units are `ns`, `us` (or `µs`), `"ms"`, `"s"`, `"m"`, and `"h"`. The first number may be negative to indicate a negative duration, like `"-2h5m"`.

The result is a string, also in [RFC 3339](http://tools.ietf.org/html/rfc3339) format, representing the result of adding the given direction to the given timestamp.

## Regex Functions

**regex**(pattern *String*, str *String*): *String*  
Regex is a function that extracts one or more substrings from a given string by applying a regular expression pattern, describing the first match. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#Regex) 

The return type depends on the composition of the capture groups (if any) in the pattern:

- If there are no capture groups at all, the result is a single string
  representing the entire matched pattern.
- If all of the capture groups are named, the result is an object whose
  keys are the named groups and whose values are their sub-matches, or
  null if a particular sub-group was inside another group that didn't
  match.
- If none of the capture groups are named, the result is a tuple whose
  elements are the sub-groups in order and whose values are their
  sub-matches, or null if a particular sub-group was inside another group
  that didn't match.
- It is invalid to use both named and un-named capture groups together in
  the same pattern.

If the pattern doesn't match, this function returns an error. To test for a match, call RegexAll and check if the length of the result is greater than zero. 

**regexAll**(pattern *String*, str *String*): *String*  
RegexAll is similar to Regex but it finds all of the non-overlapping matches in the given string and returns a list of them. [docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#RegexAll)

The result type is always a list, whose element type is deduced from the pattern in the same way as the return type for Regex is decided.

If the pattern doesn't match at all, this function returns an empty list. 

**regexReplace**(str *String*, pattern *String*, replace *String*): *String*
[docs](https://godoc.org/github.com/zclconf/go-cty/cty/function/stdlib#RegexReplace)
