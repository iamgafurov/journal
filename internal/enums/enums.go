package enums

const (
	Success = iota + 1
	BadRequest
	InternalError
	NotFound
	Unauthorized
	TokenExpired
	UserNotActive
	SuccessPartially
)

const (
	StatusActive  = "active"
	StatusPending = "pending"
	StatusFailed  = "failed"
)

const (
	ServiceMobi = "mobi"
	ServiceWeb  = "web"
)
