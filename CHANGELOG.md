1.3.0 (unreleased)
- Fix gofakeitPrice return value
- Use environment variables directly in hcl [#63](https://github.com/restbeast/restbeast/issues/63)

1.2.0
 - Fix crash on empty slice
 - Repeat request by param and argument [#60](https://github.com/restbeast/restbeast/issues/60) [#61](https://github.com/restbeast/restbeast/issues/61)

1.1.0
 - Improved logging
 - Add fill_null function [#48](https://github.com/restbeast/restbeast/issues/48)
 - Fix headers printout order [#49](https://github.com/restbeast/restbeast/issues/49)

1.0.0
 - Ability to construct body with content/type [#12](https://github.com/restbeast/restbeast/issues/12)
 - Cookie support [#16](https://github.com/restbeast/restbeast/issues/16)
 - Apple Silicon M1 support for macOS
 - Enhance json decode error message
 - New assertion function `assertNonEmptyString` [#35](https://github.com/restbeast/restbeast/issues/35)
 - List command [#30](https://github.com/restbeast/restbeast/issues/30)
 - Respect response content-type [#34](https://github.com/restbeast/restbeast/issues/34)

0.13.0
- Request parameters [#13](https://github.com/restbeast/restbeast/issues/13)

0.12.4
- Re tweak agent header
- Fix bug with emtpy response body
- Add missing types to json decoder
- Fix test count output
- Fix test dependent requests mixing into each other

0.12.3
- Fix issues gofakeit date related functions
- Fix agent header formatting

0.12.2
- Fix test all command issue when handling non unsupported attribute diagnostics  

0.12.1
- Fix test command issue when handling hcl diags 

0.12.0
- Test command to run all tests at once [#21](https://github.com/restbeast/restbeast/issues/21)

0.11.0
- Test command and assertions [#7](https://github.com/restbeast/restbeast/issues/7)

0.10.0
- Bytes received and sent output [#14](https://github.com/restbeast/restbeast/issues/14)
- Bug report command `restbeast bug` [#8](https://github.com/restbeast/restbeast/issues/8)
- Fix URL length [#3](https://github.com/restbeast/restbeast/issues/3) 

0.9.0
- Introducing Auth block
- Basic auth
- Bearer auth

0.8.0
- Introduce dynamic variables

0.7.1
- Fixes non-empty body issue

0.7.0
- Improved stability
- Changed access structure to request variables. Now headers and status are accessible through `request` object.
- Version fixing #13
- Fixes on external script parser
