package velobike

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParkingsService_List(t *testing.T) {
	// Test against a mocked API with fixtures
	{
		f, err := os.Open("../fixtures/parkings.json")
		require.NoError(t, err)
		ts := httptest.NewServer(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					io.Copy(w, f)
				}))
		defer ts.Close()
		client := NewClient(WithBaseURL(ts.URL))
		parkings, _, err := client.Parkings.List()
		require.NoError(t, err)
		require.Greater(t, len(parkings.Items), 0)
	}

	// Test against a real API
	{
		client := NewClient()
		parkings, _, err := client.Parkings.List()
		require.NoError(t, err)
		require.Greater(t, len(parkings.Items), 0)
	}
}
