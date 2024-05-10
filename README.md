[![CircleCI](https://circleci.com/gh/alexfalkowski/servicectl.svg?style=svg)](https://circleci.com/gh/alexfalkowski/servicectl)

# Service Control

A tool for go-service and go-service-templates.

## Crypto

This command allows us to rotate different keys.

### AES

```sh
❯ ./servicectl aes --help
AES crypto.

Usage:
  servicectl aes [flags]

Flags:
  -h, --help            help for aes
  -o, --output string   output config location (format kind:location, default env:AES_CONFIG_FILE) (default "env:AES_CONFIG_FILE")
  -r, --rotate          rotate keys

Global Flags:
  -i, --input string   input config location (format kind:location, default env:CONFIG_FILE) (default "env:CONFIG_FILE")
```

## Development

If you would like to contribute, here is how you can get started.

### Structure

The project follows the structure in [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

### Dependencies

Please make sure that you have the following installed:
- [Ruby](.ruby-version)
- [Golang](go.mod)

### Style

This project favours the [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

### Setup

Check out [CI](.circleci/config.yml).

### Changes

To see what has changed, please have a look at `CHANGELOG.md`
