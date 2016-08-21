package hubstorago

import "fmt"

// ErrorHttpBadStatus is returned when the API's response has an invalid status code
type ErrorHttpBadStatus struct {
	Code int
}

// ErrorJsonBadResponse is returned when the API's response has an invalid JSON
type ErrorJsonBadResponse struct {
	Body string
}

func (e *ErrorHttpBadStatus) Error() string {
	return fmt.Sprintf("Bad status code: %d", e.Code)
}

func (e *ErrorJsonBadResponse) Error() string {
	return fmt.Sprintf("Bad JSON response: %d", e.Body)
}
