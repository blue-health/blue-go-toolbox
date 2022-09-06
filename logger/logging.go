package logger

import (
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/logging"
	"github.com/blue-health/blue-go-toolbox/authn"
	"github.com/go-playground/validator"
)

type (
	Logger struct{ l Cloud }

	Fields map[string]interface{}

	Cloud interface {
		Log(e logging.Entry)
	}
)

func New(l Cloud) Logger { return Logger{l: l} }

func (l Logger) LogResponse(r *http.Request, w http.ResponseWriter, s int, m string) {
	var v logging.Severity

	switch {
	case s < http.StatusBadRequest:
		v = logging.Info
	case s >= http.StatusBadRequest && s < http.StatusInternalServerError:
		v = logging.Warning
	default:
		v = logging.Error
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)

	_ = json.NewEncoder(w).Encode(apiError{Error: apiMsg{Message: http.StatusText(s)}})

	l.l.Log(logging.Entry{
		Labels:   getLabels(r),
		Payload:  m,
		Severity: v,
		HTTPRequest: &logging.HTTPRequest{
			Request:  r,
			Status:   s,
			RemoteIP: r.Header.Get("X-Forwarded-For"),
		},
	})
}

func (l Logger) LogServiceError(r *http.Request, w http.ResponseWriter, e error) {
	var (
		u = unwrap(e)
		m = e.Error()
		s = http.StatusInternalServerError
		f []apiField
	)

	if v, ok := errorMap[e.Error()]; ok {
		s = v
	}

	switch g := u.(type) {
	case ValidationError:
		e = g

		if g.Root != nil {
			n := unwrap(g.Root)
			m = n.Error()

			if v, ok := errorMap[n.Error()]; ok {
				s = v
			}
		}

		for i := range g.Details {
			f = append(f, apiField{Name: toSnakeCase(g.Details[i].StructNamespace())})
		}

	case validator.ValidationErrors:
		e = g
		s = http.StatusBadRequest
		m = "bad_request"

		for i := range g {
			f = append(f, apiField{Name: toSnakeCase(g[i].StructNamespace())})
		}
	}

	var v logging.Severity

	switch {
	case s < http.StatusBadRequest:
		v = logging.Info
	case s >= http.StatusBadRequest && s < http.StatusInternalServerError:
		v = logging.Warning
	default:
		v = logging.Error
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)

	_ = json.NewEncoder(w).Encode(apiError{Error: apiMsg{Message: m, Fields: f}})

	l.l.Log(logging.Entry{
		Labels:   getLabels(r),
		Payload:  e.Error(),
		Severity: v,
		HTTPRequest: &logging.HTTPRequest{
			Request:  r,
			Status:   s,
			RemoteIP: r.Header.Get("X-Forwarded-For"),
		},
	})
}

func (l Logger) Log(v logging.Severity, m string, f Fields) {
	s := make(map[string]string, len(f))
	for k, v := range f {
		s[k] = fmt.Sprintf("%+v", v)
	}

	l.l.Log(logging.Entry{
		Labels:   s,
		Payload:  m,
		Severity: v,
	})
}

func getLabels(r *http.Request) map[string]string {
	m := make(map[string]string, 0)

	if i, ok := authn.GetIdentityID(r.Context()); ok {
		m["identity_id"] = i.String()
	}

	return m
}
