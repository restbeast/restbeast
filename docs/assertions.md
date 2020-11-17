# Assertions

**assertEqual**(val1 *Dynamic*, val2 *Dynamic*)  
Equal determines whether the two given values are equal

**assertNotEqual**(val1 *Dynamic*, val2 *Dynamic*)  
Equal determines whether the two given values are not equal

**assertGreaterThan**(val1 *Numeric*, val2 *Numeric*)  
GreaterThan passes in case of  val1 > val2

**assertGreaterThanOrEqualTo**(val1 *Numeric*, val2 *Numeric*)  
assertGreaterThanOrEqualTo passes in case of  val1 >= val2

**assertLessThan**(val1 *Numeric*, val2 *Numeric*)  
assertLessThan passes in case of  val1 < val2

**assertLessThanOrEqualTo**(val1 *Numeric*, val2 *Numeric*)  
assertLessThanOrEqualTo passes in case of  val1 <= val2

**assertTrue**(val1 *Bool*)  
assertTrue passes if val1 is True

**assertFalse**(val1 *Bool*)  
assertFalse passes if val1 is False

**assertEmail**(val *String*)  
assertEmail passes if val1 is a valid email

**assertUUIDv4**(val *String*)  
assertUUIDv4 passes if val1 is a valid UUIDv4

**assertIpv4**(val *String*)  
assertIpv4 passes if val1 is a valid ip v4

**assertRegex**(regex *String*, val *String*)  
assertRegex passes if val1 passes regex given regex rule. 
