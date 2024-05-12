[![CircleCI](https://circleci.com/gh/alexfalkowski/servicectl.svg?style=svg)](https://circleci.com/gh/alexfalkowski/servicectl)

# System Control

A tool for go-service and go-service-templates.

## Crypto

The ability to handle the cryptography configuration.

### AES

Handle AES configuration.

```sh
❯ ./servicectl aes --help
AES crypto.

Usage:
  servicectl aes [flags]

Flags:
  -h, --help     help for aes
  -r, --rotate   rotate key
  -v, --verify   verify key

Global Flags:
  -i, --input string    input config location (format kind:location)
  -o, --output string   output config location (format kind:location)
```

### HMAC

Handle HMAC configuration.

```sh
❯ ./servicectl hmac --help
HMAC crypto.

Usage:
  servicectl hmac [flags]

Flags:
  -h, --help     help for hmac
  -r, --rotate   rotate key
  -v, --verify   verify key

Global Flags:
  -i, --input string    input config location (format kind:location)
  -o, --output string   output config location (format kind:location)
```

### RSA

Handle RSA configuration.

```sh
❯ ./servicectl rsa --help
RSA crypto.

Usage:
  servicectl rsa [flags]

Flags:
  -h, --help     help for rsa
  -r, --rotate   rotate keys
  -v, --verify   verify keys

Global Flags:
  -i, --input string    input config location (format kind:location)
  -o, --output string   output config location (format kind:location)
```

### Ed25519

Handle Ed25519 configuration.

```sh
❯ ./servicectl ed25519 --help
Ed25519 crypto.

Usage:
  servicectl ed25519 [flags]

Flags:
  -h, --help     help for ed25519
  -r, --rotate   rotate keys
  -v, --verify   verify keys

Global Flags:
  -i, --input string    input config location (format kind:location)
  -o, --output string   output config location (format kind:location
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
