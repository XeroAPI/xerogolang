package helpers

import (
	"bytes"
	"fmt"
	"io"
)

//ReaderToString converts an io.ReadCloser to a string
func ReaderToString(reader io.ReadCloser) string {
	if reader == nil {
		return ""
	}
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(reader)
	if err != nil {
		return ""
	}
	newString := buf.String()
	_, err = fmt.Printf(newString)
	if err != nil {
		return ""
	}
	return newString
}
