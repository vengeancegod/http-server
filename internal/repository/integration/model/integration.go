package model

type AccountIntegration struct {
	ID                 int64  `json:"id"`
	AccountID          int64  `json:"account_id"`
	SecretKey          string `json:"secret_key"`
	ClientID           string `json:"client_id"`
	RedirectURL        string `json:"redirect_url"`
	AuthenticationCode string `json:"authentication_code"`
	AuthorizationCode  string `json:"authorization_code"`
}
