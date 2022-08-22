package enums

const (
	Success = iota + 1
	BadRequest
	InternalError
	NotFound
	Unauthorized
	TokenExpired
	UserNotActive
)

const (
	StatusActive = "active"
)

const (
	ServiceMobi = "mobi"
	ServiceWeb  = "web"
)
