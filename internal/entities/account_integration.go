package entities

type AccountIntegration struct {
<<<<<<< HEAD
	Id                 int64  `json:"id"`
	AccountId          string `json:"account_id"`
	SecretKey          string `json:"secret_key"`
	ClientId           string `json:"client_id"`
=======
	ID                 int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	AccountID          int64  `gorm:"index"`
	SecretKey          string `json:"secret_key"`
	ClientID           string `json:"client_id"`
>>>>>>> feature/SCHOOL-1312
	RedirectURL        string `json:"redirect_url"`
	AuthenticationCode string `json:"authentication_code"`
	AuthorizationCode  string `json:"authorization_code"`
}
