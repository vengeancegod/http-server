package entities

const (
	ErrConnectBeanstalk = "Failed to connnect to beanstalk server"
	ErrMarshalTask      = "Failed marshaling task"
	ErrPutTask          = "Failed put task to queue"
	ErrReserveTask      = "Failed reserve task"
	ErrUnmarshalTask    = "Failed unmarashaling task"
)

type SynchronizationTask struct {
	UnisenderKey string `json:"unisender_key"`
	AccountID    int64  `json:"account_id"`
}

