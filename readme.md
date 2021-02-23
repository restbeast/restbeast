[![Actions Status](https://github.com/restbeast/restbeast/workflows/ci/badge.svg)](https://github.com/restbeast/restbeast/actions)
[![codecov](https://codecov.io/gh/restbeast/restbeast/branch/master/graph/badge.svg)](https://codecov.io/gh/restbeast/restbeast)

# RestBeast Terminal Client
In a nut shell;
- A terminal API client
- Testing tool either in terminal or in CI
- Easy load testing tool.

This open source terminal client aims to simplify api development, api testing, service health checks and load testing by putting them together under one roof.

## Visit [https://restbeast.com/docs/v1.0/](https://restbeast.com/docs/v1.0/) website for the documentation, and examples

## Installation

### Homebrew

Go [here](https://brew.sh/) to install brew, then; 

```shell
brew tap restbeast/brew
brew install restbeast
```

### Binary
Get the latest build from [github release page](https://github.com/restbeast/restbeast/releases/latest).     
Decompress the file and move the executable file to a location in $PATH

```shell
tar zxvf restbeast-v1.0.0-linux-amd64.tar.gz
sudo mv restbeast /usr/local/bin/
```

### Compile From Source
Install `go >= 1.16` [go docs](https://golang.org/doc/install)

Get the latest source from [github release page](https://github.com/restbeast/restbeast/releases/latest) and unzip
```shell
unzip restbeast-v1.0.0.zip
cd restbeast-v1.0.0
```

Or clone from github
```shell
git clone https://github.com/restbeast/restbeast.git
cd restbeast
```

```shell
make VERSION=v1.0.0
sudo make install
```

## License

GNU General Public License v3.0 - see [LICENSE](LICENSE) for more details
