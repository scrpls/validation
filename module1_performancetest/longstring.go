package longstring

import (
	"bytes"
	"strings"
)

func ThroughConcatenation(length int) string {
	var s string
	for i := 0; i < length; i++ {
		s += "text"
	}
	return s
}

func ThroughArrangment(length int) string {
	s := []string{}
	for i := 0; i < length; i++ {
		s = append(s, "text")
	}
	return strings.Join(s, "")
}

func ThroughBuffer(length int) string {
	var buffer bytes.Buffer
	for i := 0; i < length; i++ {
		buffer.WriteString("text")
	}
	return buffer.String()
}
