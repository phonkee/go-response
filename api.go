package response

import "net/http"

/*
New returns new response instance, if no status is given StatusOK is used
*/
func New(statuses ...int) (result Response) {
	result = &response{
		data:    map[string]interface{}{},
		headers: map[string]string{},
	}
	result.ContentType("application/json")
	if len(statuses) > 0 {
		result.Status(statuses[0])
	} else {
		result.Status(http.StatusOK)
	}
	return
}

/*
Body is helper to create response
*/
func Body(body interface{}) Response {
	return New().Body(body)
}

/*
Data is helper to create status ok response.
*/
func Data(key string, value interface{}) Response {
	return New().Data(key, value)
}

/*
Error is helper to create status ok response.
*/
func Error(err interface{}) Response {
	return New(http.StatusInternalServerError).Error(err)
}

/*
HTML returns response set to HTML
*/
func HTML(html string) Response {
	return New().HTML(html)
}

/*
Result is helper to create status ok response.
*/
func Result(result interface{}) Response {
	return New().Result(result)
}

/*
SliceResult is helper to create status ok response.
*/
func SliceResult(result interface{}) Response {
	return New().SliceResult(result)
}
