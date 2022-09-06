package logging_test

import (
	"errors"
	"net/http"
	"testing"

	"cloud.google.com/go/logging"
	blueLogging "github.com/blue-health/blue-go-toolbox/logging"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type (
	FakeWriter struct {
		buf        []byte
		statusCode int
	}

	MockCloudLogger struct{ mock.Mock }
)

func (w *FakeWriter) Header() http.Header {
	return http.Header{}
}

func (w *FakeWriter) Write(b []byte) (int, error) {
	w.buf = append(w.buf, b...)
	return len(b), nil
}

func (w *FakeWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

//nolint:gocritic // AA
func (m *MockCloudLogger) Log(e logging.Entry) {
	m.Called(e)
}

func TestLogResponse(t *testing.T) {
	testCases := []struct {
		name       string
		statusCode int
		message    string
		severity   logging.Severity
	}{
		{
			name:       "info log",
			statusCode: http.StatusOK,
			message:    "hello world",
			severity:   logging.Info,
		},
		{
			name:       "warning log",
			statusCode: http.StatusBadRequest,
			message:    "bad uuid",
			severity:   logging.Warning,
		},
		{
			name:       "error log",
			statusCode: http.StatusInternalServerError,
			message:    "oom",
			severity:   logging.Error,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			var (
				r               = &http.Request{}
				mockCloudLogger = new(MockCloudLogger)
				w               = &FakeWriter{}
			)

			mockCloudLogger.On("Log", mock.MatchedBy(func(e logging.Entry) bool {
				return e.Payload == c.message && e.Severity == c.severity
			}))

			logger := blueLogging.Get(mockCloudLogger)

			logger.LogResponseMessage(w, r, c.statusCode, c.message)

			mockCloudLogger.AssertExpectations(t)

			s := string(w.buf)

			require.Contains(t, s, c.message)
			require.Equal(t, c.statusCode, w.statusCode)
		})
	}
}

func TestLogServiceError(t *testing.T) {
	testCases := []struct {
		name       string
		err        error
		statusCode int
		msg        string
		severity   logging.Severity
	}{
		{
			name:       "server error",
			severity:   logging.Error,
			statusCode: http.StatusInternalServerError,
			err:        errors.New("huge error"),
			msg:        "huge error",
		},
		{
			name:       "validation error",
			severity:   logging.Error,
			statusCode: http.StatusInternalServerError,
			err:        blueLogging.ValidationError{Root: errors.New("root error")},
			msg:        "root error",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			var (
				r               = &http.Request{}
				mockCloudLogger = new(MockCloudLogger)
				w               = &FakeWriter{}
			)

			mockCloudLogger.On("Log", mock.MatchedBy(func(e logging.Entry) bool {
				return e.Severity == c.severity
			}))

			logger := blueLogging.Get(mockCloudLogger)

			logger.LogServiceError(w, r, c.err)

			mockCloudLogger.AssertExpectations(t)

			s := string(w.buf)

			require.Contains(t, s, c.msg)
			require.Equal(t, c.statusCode, w.statusCode)
		})
	}
}
