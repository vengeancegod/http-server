package entities

<<<<<<< HEAD
type Account struct {
	Id           int64  `json:"id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Expires      int64  `json:"expires"`
=======
const GrantType = "authorization_code"

type Account struct {
	ID           int64                `json:"id" gorm:"primaryKey"`
	AccessToken  string               `json:"access_token"`
	RefreshToken string               `json:"refresh_token"`
	Expires      int64                `json:"expires"`
	Integrations []AccountIntegration `gorm:"foreignKey:AccountID"`
	Contacts     []Contacts           `gorm:"foreignKey:AccountID"`
	UnisenderKey string               `gorm:"foreignKey:AccountID"`
}

type AuthRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	RedirectURI  string `json:"redirect_uri"`
}

type AuthResponse struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
>>>>>>> feature/SCHOOL-1312
}
