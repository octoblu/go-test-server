package testserver

import "net/http"

// Transaction defines the response of an
// an http request, as well as store the request made
type Transaction struct {
	Request        *http.Request
	ResponseBody   string
	ResponseStatus int
}
