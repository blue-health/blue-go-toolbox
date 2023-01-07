package manipulate

import (
	"errors"
	"path"
	"regexp"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/guregu/null.v4"
)

type (
	Sanitizer struct {
		policy *bluemonday.Policy
	}

	Kind int
)

const (
	Strict Kind = iota + 1
	UGC
)

var (
	ugcPolicy    = bluemonday.UGCPolicy()
	strictPolicy = bluemonday.StrictPolicy()
	spaceRegex   = regexp.MustCompile(`\s+`)
	nidRegex     = regexp.MustCompile(`^[A-Z]{3}\d{7}$`)
)

var ErrNIDInvalid = errors.New("nid_invalid")

func NewSanitizer(kind Kind) *Sanitizer {
	var policy *bluemonday.Policy

	switch kind {
	case UGC:
		policy = ugcPolicy
	default:
		policy = strictPolicy
	}

	return &Sanitizer{policy: policy}
}

func (p *Sanitizer) Sanitize(s *string) {
	*s = strings.TrimSpace(p.policy.Sanitize(*s))
}

func (p *Sanitizer) SanitizeNull(s *null.String) {
	s.String = strings.TrimSpace(p.policy.Sanitize(s.String))
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

func SanitizeFileName(s string) string {
	return cleanString(path.Clean(path.Base(s)), illegalName)
}

var (
	separators       = regexp.MustCompile(`[ &_=+:]`)
	dashes           = regexp.MustCompile(`[\-]+`)
	illegalName      = regexp.MustCompile(`[^[:alnum:]-.]`)
	transliterations = map[rune]string{
		'À': "A",
		'Á': "A",
		'Â': "A",
		'Ã': "A",
		'Ä': "A",
		'Å': "AA",
		'Æ': "AE",
		'Ç': "C",
		'È': "E",
		'É': "E",
		'Ê': "E",
		'Ë': "E",
		'Ì': "I",
		'Í': "I",
		'Î': "I",
		'Ï': "I",
		'Ð': "D",
		'Ł': "L",
		'Ñ': "N",
		'Ò': "O",
		'Ó': "O",
		'Ô': "O",
		'Õ': "O",
		'Ö': "OE",
		'Ø': "OE",
		'Œ': "OE",
		'Ù': "U",
		'Ú': "U",
		'Ü': "UE",
		'Û': "U",
		'Ý': "Y",
		'Þ': "TH",
		'ẞ': "SS",
		'à': "a",
		'á': "a",
		'â': "a",
		'ã': "a",
		'ä': "ae",
		'å': "aa",
		'æ': "ae",
		'ç': "c",
		'è': "e",
		'é': "e",
		'ê': "e",
		'ë': "e",
		'ì': "i",
		'í': "i",
		'î': "i",
		'ï': "i",
		'ð': "d",
		'ł': "l",
		'ñ': "n",
		'ń': "n",
		'ò': "o",
		'ó': "o",
		'ô': "o",
		'õ': "o",
		'ō': "o",
		'ö': "oe",
		'ø': "oe",
		'œ': "oe",
		'ś': "s",
		'ù': "u",
		'ú': "u",
		'û': "u",
		'ū': "u",
		'ü': "ue",
		'ý': "y",
		'ÿ': "y",
		'ż': "z",
		'þ': "th",
		'ß': "ss",
	}
)

func cleanString(s string, r *regexp.Regexp) string {
	s = strings.Trim(s, " ")
	s = accents(s)
	s = separators.ReplaceAllString(s, "-")
	s = r.ReplaceAllString(s, "")
	s = dashes.ReplaceAllString(s, "-")

	return s
}

func accents(s string) string {
	var b strings.Builder

	b.Grow(len(s))

	for _, c := range s {
		if val, ok := transliterations[c]; ok {
			b.WriteString(val)
		} else {
			b.WriteRune(c)
		}
	}

	return b.String()
}
