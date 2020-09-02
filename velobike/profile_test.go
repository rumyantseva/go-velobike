package velobike

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProfileService_Get(t *testing.T) {
	// Test against a mocked API with fixtures
	{
		f, err := os.Open("../fixtures/profile.json")
		require.NoError(t, err)
		ts := httptest.NewServer(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					io.Copy(w, f)
				}))
		defer ts.Close()
		client := NewClient(WithBaseURL(ts.URL))
		profile, _, err := client.Profile.Get()
		require.NoError(t, err)
		assert.Equal(t, "Hello", *profile.FirstName)
		assert.Equal(t, "World", *profile.LastName)
	}
}
