package entities

type AccountIntegration struct {
	ID                 int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	AccountID          int64  `gorm:"index"`
	SecretKey          string `json:"secret_key"`
	ClientID           string `json:"client_id"`
	RedirectURL        string `json:"redirect_url"`
	AuthenticationCode string `json:"authentication_code"`
	AuthorizationCode  string `json:"authorization_code"`
}
