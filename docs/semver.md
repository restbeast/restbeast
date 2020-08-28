# Versioning

#### Hyphen Range Comparisons

There are multiple methods to handle ranges and the first is hyphens ranges. These look like:

    1.2 - 1.4.5 which is equivalent to >= 1.2 <= 1.4.5
    2.3.4 - 4.5 which is equivalent to >= 2.3.4 <= 4.5

#### Wildcards In Comparisons

The x, X, and * characters can be used as a wildcard character. This works for all comparison operators. When used on the = operator it falls back to the patch level comparison (see tilde below). For example,

    1.2.x is equivalent to >= 1.2.0, < 1.3.0
    >= 1.2.x is equivalent to >= 1.2.0
    <= 2.x is equivalent to < 3
    * is equivalent to >= 0.0.0

#### Tilde Range Comparisons (Patch)

The tilde (~) comparison operator is for patch level ranges when a minor version is specified and major level changes when the minor number is missing. For example,

    ~1.2.3 is equivalent to >= 1.2.3, < 1.3.0
    ~1 is equivalent to >= 1, < 2
    ~2.3 is equivalent to >= 2.3, < 2.4
    ~1.2.x is equivalent to >= 1.2.0, < 1.3.0
    ~1.x is equivalent to >= 1, < 2

#### Caret Range Comparisons (Major)

The caret (^) comparison operator is for major level changes once a stable (1.0.0) release has occurred. Prior to a 1.0.0 release the minor versions acts as the API stability level. This is useful when comparisons of API versions as a major change is API breaking. For example,

    ^1.2.3 is equivalent to >= 1.2.3, < 2.0.0
    ^1.2.x is equivalent to >= 1.2.0, < 2.0.0
    ^2.3 is equivalent to >= 2.3, < 3
    ^2.x is equivalent to >= 2.0.0, < 3
    ^0.2.3 is equivalent to >=0.2.3 <0.3.0
    ^0.2 is equivalent to >=0.2.0 <0.3.0
    ^0.0.3 is equivalent to >=0.0.3 <0.0.4
    ^0.0 is equivalent to >=0.0.0 <0.1.0
    ^0 is equivalent to >=0.0.0 <1.0.0


Taken from https://github.com/Masterminds/semver
