# go-query-string

![Project status](https://img.shields.io/badge/version-1.0.1-green.svg)
![License](https://img.shields.io/dub/l/vibe-d.svg)

`go-query-string` allows for easy encodes and decodes structs to query param

Features:

- Encodes structs to query string
- Decodes query string to structs

## Installation

To install `go-query-string`

```shell
go get -u github.com/teerapon19/go-query-string
```

Import the package into code:

```go
import "github.com/teerapon19/go-query-string"
```

## Usages

Encode

```go
package main

import "github.com/teerapon19/go-query-string"

type QueryParams struct {
    ID string
}

func main() {
    queryParams, err := query.Marshal(QueryParams{
        ID: "1234567890",
    })
    if err != nil {
        log.Fatal(err)
    }

    url := fmt.Sprintf("https://go-query-string-test.com?%s", queryParams)
    // url => https://go-query-string-test.com?id=1234567890
}
```

Decode

```go
package main

import "github.com/teerapon19/go-query-string"

type QueryParams struct {
    ID string
}

func main() {

    var queryParams QueryParams

    err := query.Unmarshal("id=1234567890", &queryParams)
    if err != nil {
        log.Fatal(err)
    }

    // queryParams => {ID:1234567890}
}
```

## Struct Tag

- `query:"-"` to ignore encode and decode
- `query:"name:type"` to use field name to be key if leave it's empty the name will use field name change to snake case by default
- `query:"<name>"` to change name of key query
