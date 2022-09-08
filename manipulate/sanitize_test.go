package manipulate_test

import (
	"testing"

	"github.com/blue-health/blue-go-toolbox/manipulate"
	"github.com/stretchr/testify/require"
)

func TestSanitizeFileName(t *testing.T) {
	testCases := []struct {
		in, out string
	}{
		{in: "Rechnung Schwarzwald-Apotheke Murg 138521-2022-6-A.pdf", out: "Rechnung-Schwarzwald-Apotheke-Murg-138521-2022-6-A.pdf"},
		{in: "Rechnung Blue Pharmacy 1 240137-2022-6-A.pdf", out: "Rechnung-Blue-Pharmacy-1-240137-2022-6-A.pdf"},
		{in: "Rechnung Stadt-Apotheke 159394-2022-6-A.pdf", out: "Rechnung-Stadt-Apotheke-159394-2022-6-A.pdf"},
		{in: "Rechnung Löwen-Apotheke 137099-2022-6-A.pdf", out: "Rechnung-Loewen-Apotheke-137099-2022-6-A.pdf"},
		{in: "Gutschrift Brücken-Apotheke 142864-2022-6-B-B.pdf", out: "Gutschrift-Bruecken-Apotheke-142864-2022-6-B-B.pdf"},
	}

	for _, c := range testCases {
		t.Run(c.in, func(t *testing.T) {
			require.Equal(t, c.out, manipulate.SanitizeFileName(c.in))
		})
	}
}
