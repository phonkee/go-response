package response

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse_Header(t *testing.T) {
	ass := assert.New(t)

	// check header value
	r := New().Header("key", "value")
	ass.Equal(r.Headers().Get("key"), "value")

	// check multiple values
	r = New().Header("key", "new", "other", "yeah")
	ass.Equal(r.Headers().Get("key"), "new")
	ass.Equal(r.Headers().Get("other"), "yeah")
}

func TestResponse_Result(t *testing.T) {
	ass := assert.New(t)
	result := "hello"

	ass.Equal(New().Result(result).(response).result, result)
}

func TestResponse_Error(t *testing.T) {
	ass := assert.New(t)

	// first test for http.StatusInternalServerError
	e := fmt.Errorf("testing error")
	ass.Equal(Error(e).(response).result, e.Error())
	ass.Equal(Error(e).(response).status, http.StatusInternalServerError)

	// test registered error (
	ass.Equal(http.StatusTeapot, Error(fmt.Errorf("%w", customError)).(response).status)
}

func TestResponse_Write(t *testing.T) {
	ass := assert.New(t)

	data := []struct {
		given        Response
		expectStatus int
		expectBody   string
	}{
		{OK(), http.StatusOK, "\"OK\""},
	}

	for _, item := range data {
		recorder := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "", nil)
		ass.Nil(err)
		item.given.Write(req, recorder)
		x, _ := io.ReadAll(recorder.Result().Body)
		ass.Equal(item.expectBody, strings.TrimRight(string(x), "\n"))
	}
}
