[![Actions Status](https://github.com/restbeast/restbeast/workflows/ci/badge.svg)](https://github.com/restbeast/restbeast/actions)
[![codecov](https://codecov.io/gh/restbeast/restbeast/branch/master/graph/badge.svg)](https://codecov.io/gh/restbeast/restbeast)

# RestBeast Terminal Client
In a nut shell;
- A terminal API client
- Testing tool either in terminal or in CI
- Easy load testing tool.

This open source terminal client aims to simplify api development, api testing, service health checks and load testing by putting them together under one roof.

## Features

* [API Request tool](#api-request-tool)
* [Built-in Functions](#built-in-functions)
* [External (Custom) Functions](#external-functions)
* [Multi Environment Configurations](#multi-environment-configurations)
* [Secrets](#secrets)
* [Randomize Data](#randomize-data)
* [Chaining requests](#chaining-requests)
* [Attack request](#attack-request)
* [Fixing version](#fixing-version)
* [Authorization](#authorization)
* [Testing and Assertions](#testing-and-assertions)

### API Request Tool

Having ability to do API requests as you develop a restful service and being able to share configuration required with a team is a requirement for API development.
Best way to share restbeast configuration is to commit your `hcl` files to vcs.     

Example of a simple request
```hcl
request get-example {
  method = "GET"
  url = "http://httpbin.org/get"
}
```

Executing the given example and piping output to `jq`
```shell script
restbeast r get-example | jq
```

A more complex example for adding a user
```hcl
request post-example {
  method = "POST"
  url = "http://httpbin.org/post"
  headers = {
    "content-type" = "application/json"
  }
  body = {
    firstName = "John",
    lastName = "Doe"
  }
}
```

```shell script
restbeast r post-example | jq
```

### Built-in Functions
A variety of functions are available. 
See [built-in go-cty functions](https://github.com/restbeast/restbeast/blob/master/docs/functions.md) and [gofakeit functions](https://github.com/restbeast/restbeast/blob/master/docs/gofakeit-functions.md) 

### External Functions
It's possible to define external programs or scripts as functions. 

Possible argument types are `string`, `list`, `map`, `number`.

Given contents of `test-function.js` as follows.
```javascript
process.stdout.write(process.argv[2].toUpperCase());
```
```hcl
external-function testFunction {
  interpreter = "node"
  script = "test-function.js"
  args = ["string"]
}

request function-example {
  method = "GET"
  url = "localhost/${testFunction("hello")}"
}
```

It's possible to use external functions with `variable` blocks. 

```javascript
process.stdout.write(
  JSON.stringify(
    {
      key1: process.argv[2].toUpperCase(),
      key2: process.argv[2].toLowerCase()
    }
  )
);
```
```hcl
external-function testFunction {
  interpreter = "node"
  script = "test-function.js"
  args = ["string", "string"]
}

variable test {
  value = jsonDecode(testFunction("value1", "value2"))
}

request function-example {
  method = "POST"
  url = "localhost/"

  body = {
    key1 = var.test.key1,
    key2 = var.test.key2
  }
}
```

### Multi Environment Configurations
Environment variables and related secrets can be changed with just a simple env flag.

```hcl
env prod {
  variables = {
    url = "http://httpbin.org"
    apiKey = "oh-my-secret"
  }
}

env local {
  default = true
  variables = {
    url = "http://localhost"
    apiKey = "not-so-important-secret"
  }
}

request env-example {
  method = "GET"
  url = "${env.url}/get"
  headers = {
    "x-api-key" = env.apiKey
  }
}
```

Execute with `-e`, useful for testing against various environments or testing in CI pipelines
```shell script
restbeast request env-example --env prod
```

### Secrets

```hcl
env local {
  default = true

  secrets from_shell_env {
    type = "env-var"
    paths = {
      apikey = "APIKEY"
    }
  }

  variables = {
    url = "http://localhost:8080"
    apiKey = secret.from_shell_env.apikey
  }
}

request secret-example {
  method = "GET"
  url = "${env.url}/get"
  headers = {
    "x-api-key" = env.apiKey
  }
}
```

```shell script
restbeast_var_APIKEY="very-secret-key" restbeast r secret-example --env local
```

##### Secrets from environment variables
```hcl
env test {
  secrets env {
    type = "env-var"
    paths = {
      val1 = "VAL1"
      val2 = "VAL2"
    }
  }  
}
```

Prefix environment variables with `restbeast_var_`
```shell script
restbeast_var_VAL1=secret1 restbeast_var_VAL2=secret2 restbeast r xxx --env test
```

### Randomize Data
Leverage `https://github.com/brianvoe/gofakeit` library in your requests.

An example with randomized user data
```hcl
request random-example {
  method = "POST"
  url = "https://httpbin.org/post"
  headers = {
    "content-type" = "application/json"
  }
  body = {
    firstName = gofakeitFirstName()
    lastName = gofakeitLastName()
  }
}
```

### Chaining requests

When `patch-example` request executed, it will do a `post-example` request first and use it's response as a depencency in `patch-example` request.

```hcl
request post-example {
  method = "POST"
  url = "https://httpbin.org/post"
  headers = {
    "content-type" = "application/json"
  }
  body = {
    firstName = gofakeitFirstName()
    lastName = gofakeitLastName()
  }
}

request patch-example {
  method = "PATCH"
  url = "https://httpbin.org/patch"
  headers = {
    "content-type" = "application/json"
    "X-Amzn-Trace-Id": request.post-example.body.headers.X-Amzn-Trace-Id
  }
  body = {
    firstName = "Mr. ${upper(request.post-example.body.json.firstName)}"
    lastName = "Mr. ${upper(request.post-example.body.json.lastName)}"
    contentLengthOfFirstResponse = request.post-example.headers["content-length"][0]
  }
}
```
 
```shell
$ restbeast request patch-example | jq
```

Example response will containt the first and the second `X-Amzn-Trace-Id` also 
it will include upper cased values of firstName and lastName generated in `post-example` request

```json
{
  "args": {},
  "data": "{\n  \"firstName\": \"Mr. JUNIOR\",\n  \"lastName\": \"Mr. RUSSEL\"\n}",
  "files": {},
  "form": {},
  "headers": {
    "Accept-Encoding": "gzip",
    "Content-Length": "59",
    "Content-Type": "application/json",
    "Host": "httpbin.org",
    "User-Agent": "RestBeast-v0.4.0",
    "X-Amzn-Trace-Id": "Self=1-5f15a0f8-cd3f68948e10488e520dcb5a;Root=1-5f15a0f8-26cd99544146b6540841f272"
  },
  "json": {
    "firstName": "Mr. JUNIOR",
    "lastName": "Mr. RUSSEL"
  },
  "origin": "213.127.104.191",
  "url": "https://httpbin.org/patch"
}
```

### Attack request
Keep targeted server busy. This command will execute given request `c` times in given `p` period.
Request count has to be equal or higher than 1 request per second.

```shell script
restbeast ar test-request-name -c 60 -p 60s
```

Example output
```text
Status 200 response: %78 (47)
Status 400 response: %15 (9)
Status 500 response: %6 (4)
95 Percentile: 1.091473938s
99 Percentile: 1.100081803s
AverageTime: 585.411933ms
```

### Fixing version
It's possible to fix your restbeast configuration to a specific version. See [here](https://github.com/restbeast/restbeast/blob/master/docs/semver.md) for extra comparison options.

```hcl
version = "~0.9"
```

### Authorization
It's possible to handle basic or bearer authorization through `auth` block.

Basic auth example
```hcl
request example {
  auth {
    basic {
      username = "a-username"
      password = request.some-other-request.body.password
    }
  }
}
```

Bearer auth example
```hcl
request example {
  auth {
    bearer {
      token = request.sign-in.body.jwt
    }
  }
}
```

### Testing and Assertions
```hcl
request get-example {
  method = "GET"
  url = "http://httpbin.org/get"
}
test get-example {
  valid-origin = assertIpv4(request.get-example.body.origin)
  valid-host = assertEqual(request.get-example.body.headers.Host, "httpbin.org")
}
```
Running this example with `test` command will provide the output below. With exit code as `0`. Failed tests will result exit code `1` and fail descriptions. 
```
$ restbeast test get-example
PASS: valid-host
PASS: valid-origin

2 passes, 0 failures, 2 total.
```
Please see [assertions](https://github.com/restbeast/restbeast/blob/master/docs/assertions.md). 

Just run the test command without any arguments to run all tests at once 
```
$ restbeast test
````

## Install

### Install From Binary
Get the latest build from [github release page](https://github.com/restbeast/restbeast/releases/latest).   
Decompress the file and Move the executable file to a location in $PATH
```shell script
tar zxvf restbeast-v0.10.0-linux-amd64.tar.gz
sudo mv restbeast /usr/local/bin/
```

### Compile From Source
Install `go >= 1.14` [](https://golang.org/doc/install)

Get the latest source from [github release page](https://github.com/restbeast/restbeast/releases/latest) and unzip
```shell script
unzip v0.10.0.zip
cd v0.10.0
```
Or clone from gitlab
```shell script
git clone https://github.com/restbeast/restbeast.git restbeast-cli
cd restbeast-cli
```

```shell script
make
sudo make install
```

### Help

`restbeast -h` or `restbeast {command} -h`

## License

GNU General Public License v3.0 - see [LICENSE](LICENSE) for more details
