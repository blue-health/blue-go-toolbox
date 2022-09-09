package logger

import (
	"errors"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Root    error
	Rewrite [2]string
	Details validator.ValidationErrors
}

var errorMap = map[string]int{}

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func RegisterErrors(m map[error]int) {
	for k, v := range m {
		errorMap[k.Error()] = v
	}
}

func FailValidation(root, details error) *ValidationError {
	var v ValidationError

	v.Root = root

	if e, ok := details.(validator.ValidationErrors); ok {
		v.Details = e
	}

	return &v
}

func (e *ValidationError) Error() string {
	var sb strings.Builder

	if e.Root != nil {
		sb.WriteString(e.Root.Error())
	}

	if len(e.Details) > 0 {
		sb.WriteString(": ")
		sb.WriteString(e.Details.Error())
	}

	return sb.String()
}

func (e *ValidationError) RewritePrefix(original, updated string) {
	e.Rewrite = [2]string{original, updated}
}

func (e *ValidationError) RewritesPrefix() bool {
	return e.Rewrite != [2]string{}
}

func unwrap(err error) error {
	for errors.Unwrap(err) != nil {
		err = errors.Unwrap(err)
	}

	return err
}

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}
