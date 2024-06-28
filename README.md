[![CircleCI](https://circleci.com/gh/alexfalkowski/servicectl.svg?style=svg)](https://circleci.com/gh/alexfalkowski/servicectl)
[![codecov](https://codecov.io/gh/alexfalkowski/servicectl/graph/badge.svg?token=7GMMUXLU6Y)](https://codecov.io/gh/alexfalkowski/servicectl)

# System Control

A tool for systems that use:
- https://github.com/alexfalkowski/go-service
- https://github.com/alexfalkowski/go-service-template
- https://github.com/alexfalkowski/go-client-template

## Commands

The below is what the cli handles:

```sh
❯ ./servicectl --help
servicectl

Usage:
  servicectl [command]

Available Commands:
  aes         AES crypto.
  completion  Generate the autocompletion script for the specified shell
  ed25519     Ed25519 crypto.
  feature     Feature flags.
  grpc        gRPC Server.
  help        Help about any command
  hmac        HMAC crypto.
  hooks       Webhooks.
  http        HTTP Server.
  pg          Postgres DB.
  redis       Redis cache.
  rsa         RSA crypto.
  token       Token.

Flags:
  -h, --help           help for servicectl
  -i, --input string   input config location (format kind:location)
  -v, --version        version for servicectl

Use "servicectl [command] --help" for more information about a command.
```

## Cache

The ability to handle the cache configuration.

### Redis

Handle redis configuration.

```sh
❯ ./servicectl redis --help
Redis cache.

Usage:
  servicectl redis [flags]

Flags:
  -h, --help     help for redis
  -v, --verify   verify connection

Global Flags:
  -i, --input string    input config location (format kind:location)
```

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
```

## Database

The ability to handle the database configuration.

### PG

Handle pg configuration.

```sh
❯ ./servicectl pg --help
Postgres DB.

Usage:
  servicectl pg [flags]

Flags:
  -h, --help     help for pg
  -v, --verify   verify connection

Global Flags:
  -i, --input string    input config location (format kind:location)
```

## Feature

The ability to handle the feature configuration.

```sh
❯ ./servicectl feature --help
Feature flags.

Usage:
  servicectl feature [flags]

Flags:
  -h, --help     help for feature
  -v, --verify   verify connection

Global Flags:
  -i, --input string    input config location (format kind:location)
```

## Hook

The ability to handle the hook configuration.

```sh
❯ ./servicectl hook --help
Webhook.

Usage:
  servicectl hook [flags]

Flags:
  -h, --help   help for hook
  -s, --sign   sign webhook

Global Flags:
  -i, --input string    input config location (format kind:location)
```
## NET

The ability to handle the net configuration.

### HTTP

Handle HTTP configuration.

```sh
❯ ./servicectl http --help
HTTP Server.

Usage:
  servicectl http [flags]

Flags:
  -h, --help     help for http
  -v, --verify   verify server

Global Flags:
  -i, --input string    input config location (format kind:location)
```

### gRPC

Handle gRPC configuration.

```sh
❯ ./servicectl grpc --help
gRPC Server.

Usage:
  servicectl grpc [flags]

Flags:
  -h, --help     help for grpc
  -v, --verify   verify server

Global Flags:
  -i, --input string    input config location (format kind:location)
```

## Security

The ability to handle the security configuration.

### Token

Handle token configuration.

```sh
❯ ./servicectl token --help
Token.

Usage:
  servicectl token [flags]

Flags:
  -h, --help     help for token
  -r, --rotate   rotate key
  -v, --verify   verify key

Global Flags:
  -i, --input string   input config location (format kind:location)
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
