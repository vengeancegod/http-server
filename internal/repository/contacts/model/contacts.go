package entities

type Contacts struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	AccountID int    `json:"account_id"`
	Email     string `json:"email"`
}
