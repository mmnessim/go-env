# go-env

A simple Go package for loading and accessing environment variables from `.env` files.

## Features

- Reads key-value pairs from a `.env` file
- Supports custom file paths
- Handles duplicate keys (last value wins)
- Provides easy access to environment variables in Go code
- No other dependencies

## Usage

### Installation

Install the package using:

```sh
go get github.com/mmnessim/go-env
```

### Example `.env` file

```
TEST="PASS"
TEST_2="Pass 2"
TEST_3=blorg
TEST_4="/.\!@#$%^&*()_+<>?:\"
```

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/mmnessim/go-env/env"
)

func main() {
    e, err := env.New() // Loads .env by default
    if err != nil {
        panic(err)
    }
    fmt.Println(e.Get("TEST"))    // Output: PASS
    fmt.Println(e.Get("TEST_2"))  // Output: Pass 2
    fmt.Println(e.Get("TEST_3"))  // Output: blorg
    fmt.Println(e.Get("TEST_4"))  // Output: /.\!@#$%^&*()_+<>?:\
}
```

### Load a custom file

```go
e, err := env.New(".env.local")
```

## Testing

Run tests with:

```sh
go test
```
