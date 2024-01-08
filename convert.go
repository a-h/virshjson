package virshjson

import (
	"bufio"
	"errors"
	"io"
	"regexp"
	"strings"
)

var ErrMalformedHeader = errors.New("malformed input header")
var ErrMalformedSeparator = errors.New("malformed input separator")
var ErrMalformedBody = errors.New("malformed input body")

type field struct {
	name  string
	start int
}

var fieldsRegexp = regexp.MustCompile(`.+?(\s{2,}|$)`)

func getFields(s string) (fields []field) {
	matches := fieldsRegexp.FindAllStringIndex(s, -1)
	for _, match := range matches {
		start, end := match[0], match[1]
		fields = append(fields, field{
			name:  strings.TrimSpace(s[start:end]),
			start: start,
		})
	}
	return fields
}

var separatorRegexp = regexp.MustCompile(`^\-+$`)

func Convert(input io.Reader) ([]map[string]any, error) {
	scanner := bufio.NewScanner(input)
	// Read headers.
	scanner.Scan()
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	fields := getFields(scanner.Text())
	if len(fields) == 0 {
		return nil, ErrMalformedHeader
	}

	// Read the line of hyphens.
	scanner.Scan()
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	if !separatorRegexp.MatchString(scanner.Text()) {
		return nil, ErrMalformedSeparator
	}

	// Read lines until an empty one.
	data := make([]map[string]any, 0)
	for scanner.Scan() {
		if scanner.Err() != nil {
			return nil, scanner.Err()
		}
		if len(scanner.Text()) == 0 {
			continue
		}
		value := make(map[string]any)
		for i, field := range fields {
			end := len(scanner.Text())
			if i < len(fields)-1 {
				end = fields[i+1].start
			}
			value[field.name] = strings.TrimSpace(scanner.Text()[field.start:end])
		}
		data = append(data, value)
	}

	return data, nil
}
