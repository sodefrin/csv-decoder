# csv-decoder

[![GoDoc](https://godoc.org/github.com/sodefrin/csv-decoder?status.svg)](https://godoc.org/github.com/sodefrin/csv-decoder)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/sodefrin/csv-decoder/master/LICENSE)

csv decoder.

## Install

```
$ go get -u github.com/sodefrin/csv-decoder
```

requirements: go1.13

## Usage

```go
package main

import (
	"fmt"
	"log"
	"strings"

	csvdecoder "github.com/sodefrin/csv-decoder"
)

type sample struct {
	Test1 string
	Test2 int
}

func main() {
	out := []*sample{}

	r := strings.NewReader("Test1,Test2\na,10\n")

	if err := csvdecoder.Decode(&out, r); err != nil {
		log.Fatal(err)
	}

	for _, v := range out {
		fmt.Printf("%s %d", v.Test1, v.Test2)
		// Output:
		// a 10
	}
}
```
