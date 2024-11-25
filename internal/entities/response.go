package entities

const (
	AccountCreate         = "Account create successfully"
	AccountUpdate         = "Account update successfully"
	AccountDelete         = "Account delete successfully"
	FailedToDeleteAccount = "Failed delete account"
	IntegrationCreate     = "Integration create successfully"
	IntegrationUpdate     = "Integration update successfully"
	IntegrationDelete     = "Integrations delete successfully"
	IncorrectType         = "Incorrect type of iD"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
