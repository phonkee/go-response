package response

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	customError = errors.New("test error")
)

func init() {
	// register custom error
	RegisterError(customError, http.StatusTeapot)
}

func TestError(t *testing.T) {
	ass := assert.New(t)

	// first test for http.StatusInternalServerError
	e := fmt.Errorf("testing error")
	ass.Equal(Error(e).(response).result, e.Error())
	ass.Equal(Error(e).(response).status, http.StatusInternalServerError)

	// test registered error (
	ass.Equal(http.StatusTeapot, Error(fmt.Errorf("%w", customError)).(response).status)
}

func TestNew(t *testing.T) {
	ass := assert.New(t)
	data := []struct {
		In     int
		Expect int
	}{
		{0, DefaultStatus},
		{http.StatusAccepted, http.StatusAccepted},
		{http.StatusInternalServerError, http.StatusInternalServerError},
	}

	for _, item := range data {
		ass.Equal(item.Expect, New(item.In).(response).status)
	}
}

func TestOK(t *testing.T) {
	ass := assert.New(t)
	ass.Equal(OK().(response).status, http.StatusOK)
}

func TestResult(t *testing.T) {
	ass := assert.New(t)
	result := "hello world"
	resp := Result(result)
	ass.Equal(resp.(response).status, http.StatusOK)
	ass.Equal(resp.(response).result, result)
}
