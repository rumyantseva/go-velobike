package velobike

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHistoryService_Get(t *testing.T) {
	// Test against a mocked API with fixtures
	{
		f, err := os.Open("../fixtures/history.json")
		require.NoError(t, err)
		ts := httptest.NewServer(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					io.Copy(w, f)
				}))
		defer ts.Close()
		client := NewClient(WithBaseURL(ts.URL))
		_, _, err = client.History.Get()
		require.NoError(t, err)
	}

	// Test against a mocked API with fixtures, empty history
	{
		f, err := os.Open("../fixtures/history_empty.json")
		require.NoError(t, err)
		ts := httptest.NewServer(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					io.Copy(w, f)
				}))
		defer ts.Close()
		client := NewClient(WithBaseURL(ts.URL))
		_, _, err = client.History.Get()
		require.NoError(t, err)
	}
}
