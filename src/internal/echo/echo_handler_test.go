package echo_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/erikcc02/echo-server/src/internal/echo"
	"github.com/stretchr/testify/assert"
)

func TestEchoHandler(t *testing.T) {

	body := `{ "message": "test" }`
	url := "/anyroute?foo=bar"
	headerName := "X-My-Custom-Header"

	req := httptest.NewRequest(http.MethodGet, url, strings.NewReader(body))
	w := httptest.NewRecorder()
	handler := echo.BuildHandler()

	req.Header.Add(headerName, "foo")

	handler(w, req)
	result := w.Result()
	defer result.Body.Close()

	data, err := ioutil.ReadAll(result.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	bodyStr := string(data)

	assert.Contains(t, bodyStr, "GET")
	assert.Contains(t, bodyStr, url)
	assert.Contains(t, bodyStr, headerName)
	assert.Contains(t, bodyStr, body)
}
