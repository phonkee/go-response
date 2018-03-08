// error helper package
package response

import (
	"sync"
	"github.com/pkg/errors"
)

var (
	em *errMap
)

func init() {
	// initialize global error map
	em = &errMap{
		mutex: &sync.RWMutex{},
		registry:map[error]int{},
	}
}

// RegisterError registers error to given http status
func RegisterError(err error, status int) {
	em.Register(err, status)
}

// GetErrorStatus returns appropriate http status for given error
func GetErrorStatus(err error) int {
	return em.GetStatus(err)
}

// errMap provides mapping from errors to http statuses
// then you can call response.Error(myerr) it will set appropriate status
type errMap struct {

	// mutex to secure our map
	mutex *sync.RWMutex

	// store map from error to int
	registry map[error]int
}

// Register registers error
func (e *errMap ) Register(err error, status int)  {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.registry[err] = status
	return
}

// GetStatus returns appropriate status for given error
func (e *errMap) GetStatus(err error) (status int) {
	var ok bool

	e.mutex.RLock()
	defer e.mutex.RUnlock()

	// get cause if available
	err = errors.Cause(err)

	if status, ok = e.registry[err]; !ok {
		status = 0
	}

	return
}
