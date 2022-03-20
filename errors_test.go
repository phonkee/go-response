package response

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// given error
	given = errors.New("given error")
)

func TestErrMap_Register(t *testing.T) {
	ass := assert.New(t)
	m := newErrMap()
	m.Register(given, http.StatusTeapot)

	// get status
	ass.Equal(http.StatusTeapot, m.GetStatus(fmt.Errorf("%w", given)))
	ass.Equal(http.StatusTeapot, m.GetStatus(given))
}
