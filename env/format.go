package env

import (
	"fmt"
)

// Format is an enum containing all available formats
type Format int

const (
	_ Format = iota

	// DefaultFormat output envar like <NAME>=<VALUE>
	DefaultFormat
)

// Formatter is used to format an env var
type Formatter interface {
	Format(ev Var) string
}

type defaultFormatter struct{}

// DefaultFormatter creates a new Defaultformatter
func DefaultFormatter() Formatter {
	return &defaultFormatter{}
}

// Format creates the string representation of our env var
func (f *defaultFormatter) Format(ev Var) string {
	return fmt.Sprintf("'%s'", ev.String())
}

type bashExportFormatter struct{}

func (f *bashExportFormatter) Format(ev Var) string {
	return fmt.Sprintf("export '%s'", ev.String())
}

// BashExportFormatter return an instance of the bashExportFormatter
// This exporter prefix the output with the bash 'export' statement
func BashExportFormatter() Formatter {
	return &bashExportFormatter{}
}