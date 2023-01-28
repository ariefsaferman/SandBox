package concat

import (
	"bytes"
	"fmt"
	"strings"
)

func ConcatStrings(strs ...string) string {
	builder := strings.Builder{}

	for _, str := range strs {
		builder.WriteString(str)
	}

	return builder.String()
}

func ConcatBytesBuffer(strs ...string) string {
	var buf bytes.Buffer

	for _, str := range strs {
		buf.WriteString(str)
	}

	return buf.String()
}

func ConcatFmtSprintf(strs ...string) string {
	var result string

	for _, str := range strs {
		result = fmt.Sprintf("%s%s", result, str)
	}

	return result
}
