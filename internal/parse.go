package internal

import (
	"regexp"
	"strings"
)

var (
	pattern     *regexp.Regexp
	whiteSpaces = regexp.MustCompile(`[\n\s]+`)
)

func init() {
	const delimiter = `\s*`
	patternStr := strings.Join(patterns, delimiter)
	pattern = regexp.MustCompile(patternStr)
}

type RawRecord struct {
	User         string
	Host         string
	IP           string
	ID           string
	Duration     string
	LockWait     string
	RowsSent     string
	RowsExamined string
	Context      string
	Timestamp    string
	Query        string
}

func ParseRaw(s string) RawRecord {
	s = whiteSpaces.ReplaceAllString(s, " ")
	match := pattern.FindStringSubmatch(s)
	return RawRecord{
		User:         match[1],
		Host:         match[2],
		IP:           match[3],
		ID:           match[4],
		Duration:     match[5],
		LockWait:     match[6],
		RowsSent:     match[7],
		RowsExamined: match[8],
		Context:      match[9],
		Timestamp:    match[10],
		Query:        match[11],
	}
}
