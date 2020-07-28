# RestBeast Terminal Client
In a nut shell; A terminal API client, testing tool in terminal and in CI, easy load testing tool.

This open source terminal client is part of restbeast project which aims to simplify api development, api testing, service health checks and load testing by putting them together under one roof.

## Features

#### API Development tool
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

#### Functions
A variety of functions are available. 
See [built-in go-cty functions](https://gitlab.com/restbeast/cli/-/blob/master/docs/functions.md) and [gofakeit functions](https://gitlab.com/restbeast/cli/-/blob/master/docs/gofakeit-functions.md) 

#### Execute requests in various environments
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

#### Secrets

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

#### Randomize data in request bodies
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

#### Chaining requests

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
    "X-Amzn-Trace-Id": request.post-example.headers.X-Amzn-Trace-Id
  }
  body = {
    firstName = "Mr. ${upper(request.post-example.json.firstName)}"
    lastName = "Mr. ${upper(request.post-example.json.lastName)}"
  }
}
```
 
```shell script
restbeast r patch-example | jq
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

#### Testing / Assertion (Not Implemented Yet)
```hcl
request "new-user" {
  method = "POST"
  url = "https://${env.url}/users"
  headers = {
    "content-type" = "application/json"
  }
  body = {
    firstName = gofakeitFirstName()
    lastName = gofakeitLastName()
  }
  expect = {
    status = 201
    body = {
      id = assertUuidV4()
      name = assertNotNil()
    }
  }
}
```

#### Attack request
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

## Install

Get the latest build from [gitlab release page](https://gitlab.com/restbeast/cli/-/releases)

Decompress file
```shell script
tar zxvf restbeast.tar.gz
```

Set permissions
```shell script
chmod +x restbeast
```

Move the executable file to a location in $PATH 
```shell script
sudo mv restbeast /usr/local/bin/
```

### Help

`restbeast -h` or `restbeast {command} -h`

## FAQ and troubleshooting

## License

GNU General Public License v3.0 - see [LICENSE](LICENSE) for more details