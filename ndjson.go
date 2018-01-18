package ndjson

import (
	"bufio"
	"bytes"
	"io"
)

// ToNewlineDelimitedJSON converts json to ndjson
func ToNewlineDelimitedJSON(json []byte) {
	// todo
}

// ToJSONbuffer converts some newline-delimited JSON to valid JSON buffer
func ToJSONbuffer(reader io.Reader) bytes.Buffer {
	var buffer bytes.Buffer

	r := bufio.NewReader(reader)

	buffer.Write([]byte("["))
	for {
		bytes, err := r.ReadBytes(byte('\n'))
		buffer.Write(bytes)
		if err == io.EOF || string(bytes) == "" {
			break
		}
		buffer.Write([]byte(","))
	}
	buffer.Write([]byte("]"))
	buffer.Write([]byte("\n"))

	return buffer
}

// ToJSON converts some newline-delimited JSON to valid JSON string
func ToJSON(reader io.Reader) string {
	b := ToJSONbuffer(reader)
	return b.String()
}
