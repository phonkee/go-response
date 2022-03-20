package response

import (
	"net/http"
)

var (
	// DefaultStatus is set by default
	DefaultStatus = http.StatusOK
)

// Response
// this interface is helper for writing json methods
type Response interface {
	// Error sets error to response
	Error(error) Response

	// Header sets given headers and returns new response
	Header(kv ...string) Response

	// Headers returns reference to headers
	Headers() http.Header

	// Status sets http status
	Status(int) Response

	// StatusText sets status text (response)
	StatusText(string) Response

	// Result sets json result (body)
	Result(interface{}) Response

	// Write writes to given response writer
	Write(r *http.Request, w http.ResponseWriter)
}

// Error creates reponse with error
func Error(err error) Response {
	return newResponse().Error(err)
}

// New returns new response instance, if no status is given StatusOK is used
func New(statuses ...int) (result Response) {
	if len(statuses) > 0 && statuses[0] != 0 {
		return newResponse().Status(statuses[0])
	}
	return newResponse()
}

// OK is success response
func OK() Response {
	return New(http.StatusOK)
}

// Result returns instantiated response with json data and 200 - OK
func Result(result interface{}) Response {
	return New(http.StatusOK).Result(result)
}
