package manipulate

import (
	"errors"
	"regexp"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/guregu/null.v4"
)

var (
	Sanitizer  = bluemonday.StrictPolicy()
	spaceRegex = regexp.MustCompile(`\s+`)
	nidRegex   = regexp.MustCompile(`^[A-Z]{3}\d{7}$`)
)

var ErrNIDInvalid = errors.New("nid_invalid")

func Sanitize(s *string) {
	*s = strings.TrimSpace(Sanitizer.Sanitize(*s))
}

func SanitizeNull(s *null.String) {
	s.String = strings.TrimSpace(Sanitizer.Sanitize(s.String))
}

func SanitizeIBAN(s *null.String) {
	s.String = spaceRegex.ReplaceAllString(strings.TrimSpace(strings.ToUpper(s.String)), "")
}

func SanitizeVAT(s *null.String) {
	s.String = spaceRegex.ReplaceAllString(strings.TrimSpace(strings.ToUpper(s.String)), "")
}

func SanitizeNID(s *string) error {
	*s = strings.TrimSpace(strings.ToUpper(*s))

	if !nidRegex.MatchString(*s) {
		return ErrNIDInvalid
	}

	return nil
}
