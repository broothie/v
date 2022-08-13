package v

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteTo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Render(w, http.StatusOK, Nodes{
			Doctype("html"),
			HTML(Attr{"lang": "en"},
				Head(nil,
					Title(nil, Text("test")),
				),
				Body(nil),
			),
		})
	}))
	defer server.Close()

	response, err := http.Get(server.URL)
	require.NoError(t, err)

	bytes, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	require.Equal(t, `<!doctype html><html lang="en"><head><title>test</title></head><body></body></html>`, string(bytes))
}
