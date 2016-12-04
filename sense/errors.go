package sense

import "fmt"

type APIError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (err APIError) Error() string {
	return fmt.Sprintf("sense: %d %v", err.Code, err.Message)
}

func (err APIError) Empty() bool {
	return err.Code == 0
}

// relevantError returns any non-nil http-related error (creating the request,
// getting the response, decoding) if any. If the decoded apiError is non-zero
// the apiError is returned. Otherwise, no errors occurred, returns nil.
func relevantError(httpError error, apiError APIError) error {
	if httpError != nil {
		return httpError
	}
	if apiError.Empty() {
		return nil
	}

	return apiError
}
