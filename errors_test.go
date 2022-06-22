package response

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	gerrors "github.com/pkg/errors"
)

var (
	ErrBase = errors.New("base error")
)

func TestError(t *testing.T) {
	em := newErrMap()
	em.Register(ErrBase, http.StatusTeapot)

	for _, item := range []struct {
		what error
	}{
		{fmt.Errorf("%w", ErrBase)},
		{fmt.Errorf("%w", fmt.Errorf("%w", ErrBase))},
		{gerrors.Wrap(ErrBase, "some")},
	} {
		if em.GetStatus(item.what) != http.StatusTeapot {
			t.Error("invalid status")
		}
	}

}
