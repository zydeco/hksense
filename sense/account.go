package sense

import "net/http"

type Account struct {
	Email         string    `json:"email"`
	Name          string    `json:"name"`
	FirstName     string    `json:"firstname"`
	LastName      string    `json:"lastname"`
	Gender        string    `json:"gender"`
	Height        int       `json:"height"`
	Weight        int       `json:"weight"`
	Created       Timestamp `json:"created"`
	LastModified  Timestamp `json:"last_modified"`
	EmailVerified bool      `json:"email_verified"`
	ProfilePhoto  string    `json:"profile_photo,omitempty"`
	ID            string    `json:"id"`
	DateOfBirth   string    `json:"dob"`
}

func (c *Client) GetAccount() (*Account, *http.Response, error) {
	account := new(Account)
	apiError := new(APIError)
	resp, err := c.sling.New().Get("v1/account").Receive(account, apiError)
	return account, resp, relevantError(err, *apiError)
}
