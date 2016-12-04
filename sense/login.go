package sense

import (
	"net/http"

	"github.com/dghubble/sling"
)

const clientID = "8d3c1664-05ae-47e4-bcdb-477489590aa4"
const clientSecret = "4f771f6f-5c10-4104-bbc6-3333f5b11bf9"

type LoginResponse struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	AccountID    string `json:"account_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginParams struct {
	GrantType    string `url:"grant_type"`
	ClientID     string `url:"client_id"`
	ClientSecret string `url:"client_secret"`
	Username     string `url:"username"`
	Password     string `url:"password"`
}

func Login(username string, password string) (*LoginResponse, *http.Response, error) {
	login := new(LoginResponse)
	apiError := new(APIError)
	params := &LoginParams{
		GrantType:    "password",
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Username:     username,
		Password:     password,
	}
	resp, err := sling.New().Base(senseAPI).Post("v1/oauth2/token").BodyForm(params).Receive(login, apiError)
	return login, resp, relevantError(err, *apiError)
}
