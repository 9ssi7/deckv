# Deckv - Disposable Email Checker as a Key-Value Store for Golang

[![Go Reference](https://pkg.go.dev/badge/github.com/9ssi7/deckv.svg)](https://pkg.go.dev/github.com/9ssi7/deckv)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Deckv is a simple and efficient blocklist implementation for Go applications with support for multiple storage backends.

## Features

- Simple and easy-to-use API
- Multiple storage backend support
  - In-memory storage (default)
  - Redis storage
- Configurable through simple text files
- Context support for all operations
- Thread-safe operations

## Installation

```bash
go get github.com/9ssi7/deckv
```

## Usage

See the [examples](examples) directory for more detailed usage examples.


## Quick Start

### Blocklist Configuration File

You can create a blocklist configuration file with the following content:

```
0-mail.com
1-mail.com
```

or real and full domain names, use this file [disposable-email-domains config file](https://github.com/disposable-email-domains/disposable-email-domains/blob/main/disposable_email_blocklist.conf) as a reference.

### Using In-Memory Storage

Create a `blocklist.conf` file with the following content:

```
0-mail.com
1-mail.com
```

```go
client := deckv.New(deckv.WithConfFilePath("./blocklist.conf"))
err := client.Load(context.Background())
if err != nil {
    panic(err)
}
ok, err := client.Check(context.Background(), "0-mail.com")
if err != nil {
    panic(err)
}
fmt.Println(ok)
```

### Using Redis Storage

Create a `blocklist.conf` file with the following content:

```
0-mail.com
1-mail.com
```

You can use Redis storage by setting the `WithStorage` option.

```go
storage := deckvredis.New(context.Background(), deckvredis.Config{
    Host:     "localhost",
    Port:     "6379",
    Password: "",
    DB:       0,
})
client := deckv.New(deckv.WithConfFilePath("./blocklist.conf"), deckv.WithStorage(storage))
err := client.Load(context.Background())
if err != nil {
    panic(err)
}
ok, err := client.Check(context.Background(), "0-mail.com")
if err != nil {
    panic(err)
}
fmt.Println(ok)
```

## Configuration File Format

The blocklist configuration file is a simple text file with one entry per line:

```
0-mail.com
1-mail.com
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.