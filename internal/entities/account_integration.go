package entities

type AccountIntegration struct {
	Id                 int64  `json:"id"`
	AccountId          string `json:"account_id"`
	SecretKey          string `json:"secret_key"`
	ClientId           string `json:"client_id"`
	RedirectURL        string `json:"redirect_url"`
	AuthenticationCode string `json:"authentication_code"`
	AuthorizationCode  string `json:"authorization_code"`
}
