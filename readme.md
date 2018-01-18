# go-ndjson

A set of golang utils to work with newline delimited json (ndjson)

## installation

`go get github.com/kandros/go-ndjson`

## usage

Reading from a file

```
package main

import (
	"os"

	ndjson "github.com/kandros/go-ndjson"
)

func main() {
	f, err := os.Open("db/logstore.log")
	if err != nil {
		panic(err)
	}

	s :=ndjson.ToJSON(f)
    // or as raw buffer b := ndjson.ToJSONToJSONbuffer(f)
    fmt.Print(s)
}
```
