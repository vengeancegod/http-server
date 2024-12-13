package entities

const (
	ErrConnectBeanstalk = "Failed to connnect to beanstalk server"
	ErrMarshalTask      = "Failed marshaling task"
	ErrPutTask          = "Failed put task to queue"
	ErrReserveTask      = "Failed reserve task"
	ErrUnmarshalTask    = "Failed unmarashaling task"
	ErrAction           = "Unknown action"
)

type SynchronizationTask struct {
	UnisenderKey string `json:"unisender_key"`
	AccountID    int64  `json:"account_id"`
}

type ContactsTask struct {
	Action    string `json:"action"`
	ContactID int64  `json:"contact_id"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	AccountID int64  `json:"account_id,omitempty"`
}
