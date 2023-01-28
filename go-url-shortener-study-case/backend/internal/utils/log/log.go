package log

import (
	"errors"
	"log"
	"strconv"
	"strings"

	cerrors "url-shortener/internal/utils/errors"
)

/*
Logs error from outermost to innermost
Still could be improved, made just to make debugging this app easier
*/
func LogAppErr(ew *cerrors.ErrorWrapper) {
	var s strings.Builder

	if ew == nil {
		return
	}

	setMessage(&s, ew)
	s.WriteByte('\n')

	for ew.Err != nil && errors.As(ew.Err, &ew) {
		setMessage(&s, ew)
		s.WriteByte('\n')
	}

	if ew.Err != nil && !errors.As(ew.Err, &ew) {
		s.WriteString("error: ")
		s.WriteString(ew.Err.Error())
		s.WriteByte('\n')
	}

	log.Println(strings.TrimSpace(s.String()))
}

func setMessage(s *strings.Builder, ew *cerrors.ErrorWrapper) {
	s.WriteString("error: ")
	s.WriteString(ew.Message)
	s.WriteString("\n file: ")
	s.WriteString(ew.Filename)
	s.WriteString(":")
	s.WriteString(strconv.Itoa(ew.LineNumber))
}
