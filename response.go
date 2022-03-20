package response

import (
	"encoding/json"
	"net/http"
)

// newResponse returns blank response
func newResponse() Response {
	return response{
		headers: http.Header{},
	}.Status(DefaultStatus)
}

// response implements Response interface
type response struct {
	status  int
	result  interface{}
	headers http.Header
}

// Error sets error and status
func (r response) Error(err error) Response {
	// if there is no error we do nothing
	if err == nil {
		return r
	}
	errText := err.Error()

	return r.Status(GetErrorStatus(err)).Result(errText)
}

// Header sets header value, if you need more functionality use Headers
func (r response) Header(kv ...string) Response {
	h := r.Headers()

	// add all given headers
	for i := 0; i <= (len(kv)/2)-1; i++ {
		h.Set(kv[2*i], kv[(2*i)+1])
	}

	return r
}

// Headers returns reference to headers
func (r response) Headers() http.Header {
	return r.headers
}

// Status sets status code and tries to resolve status text
func (r response) Status(status int) Response {
	r.status = status

	// if result was not set, and status is some meaningful value (non zero)
	if status != 0 && r.result == nil {
		r = r.StatusText(http.StatusText(status)).(response)
	}

	return r
}

// StatusText sets status text to given value
func (r response) StatusText(statusText string) Response {
	r.result = statusText
	return r
}

// Result sets json value
func (r response) Result(result interface{}) Response {
	r.result = result
	return r
}

// Write writes all data to response
func (r response) Write(_ *http.Request, w http.ResponseWriter) {
	// store headers pointer
	headers := w.Header()

	// now update all headers
	for key, _ := range r.Headers() {
		for _, value := range r.Headers().Values(key) {
			headers.Add(key, value)
		}
	}

	// add json content type (always)
	headers.Set("Content-Type", "application/json")

	// check if we have something to write first
	if r.result != nil {
		// encode to json
		if errMarshal := json.NewEncoder(w).Encode(r.result); errMarshal != nil {
			// cannot marshal json - internal server error
			w.WriteHeader(http.StatusInternalServerError)

			// I believe this will not fail
			_ = json.NewEncoder(w).Encode(errMarshal.Error())
			return
		}
	}

	// add status
	w.WriteHeader(r.status)
}
