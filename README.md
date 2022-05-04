# yaml2json

Convert YAML to JSON

## Usage

```console
$ yaml2json input.yaml # Output JSON to standard output
```

## Install

* [GitHub Releases](https://github.com/suzuki-shunsuke/yaml2json/releases): Download asset from GitHub Releases and install the executable binary in `PATH`
* [Homebrew](https://brew.sh/): `$ brew install suzuki-shunsuke/yaml2json/yaml2json`

## How does it work?

This is very simple.

1. Read YAML with [gopkg.in/yaml.v2](https://pkg.go.dev/gopkg.in/yaml.v2)
1. Convert `map[interface{}]interface{}` to `map[string]interface{}` with [suzuki-shunsuke/go-convmap](https://github.com/suzuki-shunsuke/go-convmap)
1. Output JSON to the Standard Output with [encoding/json](https://pkg.go.dev/encoding/json)

## Why do we reinvent a wheel?

We know there are already some tools to convert YAML to JSON,
but we develop this tool, because we want a single executable binary for macOS, Linux / AMD64, ARM64.
Unfortunately, we can't find a single executable binary for Linux ARM64.

## LICENSE

[MIT](LICENSE)
