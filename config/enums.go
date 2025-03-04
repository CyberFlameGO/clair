package config

import (
	"encoding"
	"fmt"
	"strings"
)

//go:generate stringer -type Mode,LogLevel -linecomment -output enums_string.go

// A Mode is an operating mode recognized by Clair.
//
// This is not directly settable by serializing into a Config object.
type Mode int

// Clair modes, with their string representations as the comments.
const (
	ComboMode    Mode = iota // combo
	IndexerMode              // indexer
	MatcherMode              // matcher
	NotifierMode             // notifier
)

// ParseMode returns a mode for the given string.
//
// The passed string is case-insensitive.
func ParseMode(s string) (Mode, error) {
	for i, lim := 0, len(_Mode_index); i < lim; i++ {
		m := Mode(i)
		if strings.EqualFold(s, m.String()) {
			return m, nil
		}
	}
	return Mode(-1), fmt.Errorf(`unknown mode %q`, s)
}

// A LogLevel is a log level recognized by Clair.
//
// The zero value is "info".
type LogLevel int

// The recognized log levels, with their string representations as the comments.
//
// NB "Fatal" and "Panic" are not used in clair or claircore, and will result in
// almost no logging.
const (
	InfoLog       LogLevel = iota // info
	DebugColorLog                 // debug-color
	DebugLog                      // debug
	WarnLog                       // warn
	ErrorLog                      // error
	FatalLog                      // fatal
	PanicLog                      // panic
)

// ParseLogLevel returns the log lever for the given string.
//
// The passed string is case-insensitive.
func ParseLogLevel(s string) (LogLevel, error) {
	for i, lim := 0, len(_LogLevel_index); i < lim; i++ {
		l := LogLevel(i)
		if strings.EqualFold(s, l.String()) {
			return l, nil
		}
	}
	return LogLevel(-1), fmt.Errorf(`unknown log level %q`, s)
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (l *LogLevel) UnmarshalText(b []byte) (err error) {
	*l, err = ParseLogLevel(string(b))
	if err != nil {
		return err
	}
	return nil
}

// Assert LogLevel implements TextUnmarshaler.
var _ encoding.TextUnmarshaler = (*LogLevel)(nil)
