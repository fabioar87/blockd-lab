package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupAPI(t *testing.T) (string, func()) {
	t.Helper()
	ts := httptest.NewServer(customMux())

	return ts.URL, ts.Close
}

func TestGet(t *testing.T) {
	useCases := []struct {
		caseName      string
		path          string
		expStatusCode int
	}{
		{
			caseName:      "GetRoot",
			path:          "/",
			expStatusCode: http.StatusOK,
		},
		{
			caseName:      "NotFound",
			path:          "/wrong/path/status/500",
			expStatusCode: http.StatusNotFound,
		},
	}

	url, cleanup := setupAPI(t)
	defer cleanup()

	for _, tc := range useCases {
		t.Run(tc.caseName, func(t *testing.T) {
			var (
				err error
			)

			r, err := http.Get(url + tc.path)
			if err != nil {
				t.Error(err)
			}
			defer r.Body.Close()

			if r.StatusCode != tc.expStatusCode {
				t.Fatalf("Expected %q, got %q instead", http.StatusText(tc.expStatusCode), http.StatusText(r.StatusCode))
			}
		})
	}
}
