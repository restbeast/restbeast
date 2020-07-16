# RestBeast Terminal Client
This open source terminal client is part of restbeast project which aims to simplify api development, api testing, service health checks and load testing by putting them together under one roof.

## Features

#### API Development tool
Having ability to do API requests as you develop a restful service and being able to share configuration required with a team is a requirement for API development.
Best way to share restbeast configuration is to commit your `hcl` files to vcs.     

Example of a simple request
```hcl
request "get-users" {
  method = "GET"
  url = "https://${env.url}/users"
}
```

Executing the given example and piping output to `jq`
```shell script
restbeast r get-users | jq
```

A more complex example for adding a user
```hcl
request "new-user" {
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

#### Functions
A variety of functions are available. 
See [built-in go-cty functions](https://gitlab.com/restbeast/cli/-/blob/master/docs/functions.md) and [gofakeit functions](https://gitlab.com/restbeast/cli/-/blob/master/docs/gofakeit-functions.md) 

#### Execute requests in various environments
Environment variables and related secrets can be changed with just a simple env flag.

```hcl
env local {
  default = true
  variables = {
    url = "localhost"
    apiKey = "oh-my-secret"
  }
}

request "get-users" {
  method = "GET"
  url = "https://${env.url}/users"
  headers = {
    "x-api-key" = env.apiKey
  }
}
```

Execute with `-e`, useful for testing against various environments or testing in CI pipelines
```shell script
restbeast request get-users --env dev
```

#### Secrets

```hcl
env local {
  secrets from_shell_env {
    type = "env-var"
    paths = {
      apikey = "APIKEY"
    }
  }

  // SSM secret engine isn't implemented yet.
  secrets from_aws_ssm {
    type = "aws-ssm"
    paths = {
      apikey = "/path/to/apikey"
    }
    provider = "aws.dev-acc-eu-west-1"
  }

  // SSM secret engine isn't implemented yet.
  secrets from_aws_ssm {
    type = "aws-ssm"
    paths = {
      apikey = "/path/to/apikey"
    }
    provider = "aws.dev-acc-eu-west-1"
  }

  variables = {
    url = "localhost"
    apiKey = secret.from_shell_env.apikey
    anOtherKey = secret.aws-ssm.apikey
  }
}
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
}
```

#### Chaining requests

When `update-user` request executed, it will do a `new-user` request first and use it's response as a depencency in `update-user` request.

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
}

request "update-user" {
  method = "PATCH"
  url = "https://${env.url}/users/${request.new-user.id}"
  headers = {
    "content-type" = "application/json"
  }
  body = {
    firstName = "Mr. ${upper(request.new-user.firstName)}"
    lastName = gofakeitLastName()
  }
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

#### Secrets
```hcl
env local {
  default = true
  variables = {
    url = "localhost"
    apiKey = "${secret.apiKey}"
  }
  secretEngine = "env-vars"
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
