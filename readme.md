# RestBeast Terminal Client
This open source terminal client is part of restbeast project which aims to simplify api development, api testing, service health checks and load testing by putting them together under one roof.

## Features

#### API Development tool
Having ability to do API requests as you develop a restful service and being able to share configuration required with a team is a requirement for API development.
Best way to share restbeast configuration is to commit your `hcl` files to vcs.     

Example of a simple request
```hcl
request "get-users" "default" {
  method = "GET"
  url = "https://${env.url}/users"
}
```
Executing the given example and piping output to `jq`
```shell script
restbeast r -l get-users | jq
```

A more complex example for adding a user
```hcl
request "new-user" "default" {
  method = "POST"
  url = "https://${env.url}/users"
  headers = {
    "content-type" = "application/json"
  }
  body = {
    firstName = "John",
    lastName = "Doe"
  }
}
```

#### Execute requests in various environments
Environment variables and related secrets can be changed with just a simple env flag.

```hcl
env local {
  default = true
  variables = {
    url = "localhost"
    apiKey = "${secret.apiKey}"
  }
  secretEngine = "env-vars"
}

env dev {
  variables = {
    url = "dev-domain.com"
    apiKey = "${secret.apiKey}"
  }
  secretEngine = "vault"
}

request "get-users" "default" {
  method = "GET"
  url = "https://${env.url}/users"
  headers = {
    "x-api-key" = env.apiKey
  }
}
```

Execute with `-e`, useful for testing against various environments or testing in CI pipelines
```shell script
  restbeast r -l get-users -e dev
```

#### Randomize data in request bodies
Leverage `https://github.com/brianvoe/gofakeit` library in your requests.

An example with randomized user data
```hcl
request "new-user" "default" {
  method = "POST"
  url = "https://${env.url}/users"
  headers = {
    "content-type" = "application/json"
  }
  body = {
    firstName = gofakeit.FirstName()
    lastName = gofakeit.LastName()
  }
}
```

#### Chaining requests

When `update-user` request executed, it will do a `new-user` request first and use it's response as a depencency in `update-user` request.
```hcl
request "new-user" "default" {
  method = "POST"
  url = "https://${env.url}/users"
  headers = {
    "content-type" = "application/json"
  }
  body = {
    firstName = gofakeit.FirstName()
    lastName = gofakeit.LastName()
  }
}

request "update-user" "default" {
  method = "PATCH"
  url = "https://${env.url}/users/${request.new-user.id}"
  headers = {
    "content-type" = "application/json"
  }
  body = {
    firstName = "Mr. ${upper(request.new-user.firstName)}"
    lastName = gofakeit.LastName()
  }
}
```

#### Testing / Assertion (TODO)
```hcl
request "new-user" "default" {
  method = "POST"
  url = "https://${env.url}/users"
  headers = {
    "content-type" = "application/json"
  }
  body = {
    firstName = gofakeit.FirstName(),
    lastName = gofakeit.LastName()
  }
  expect = {
    status = 201
    body = {
      id = assert.uuidV4()
      name = assert.notNil()
    }
  }
}
```

#### Attack request
Keep targeted server busy. This command will execute given request `c` times in given `p` period.

## Install

## Usage

### Help

`rb` or `rb -h` or `rb command -h`

### Regular api requests

`rb r get https://domain.com/get/something`

#### Saving requests

`rb r get https://domain.com/get/something -s a-request-name` 

#### Using saved requests

`rb r get -l a-request-name`

## FAQ and troubleshooting

## License

GNU General Public License v3.0 - see [LICENSE](LICENSE) for more details
