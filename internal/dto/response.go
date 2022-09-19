package dto

import "github.com/iamgafurov/journal/internal/enums"

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
	ErrStr  string      `json:"-"`
}

func (r *Response) ErrCode(code int) {
	switch code {
	case enums.Success:
		r.Code = enums.Success
		r.Status = "success"
		r.Message = "ok"
		return
	case enums.BadRequest:
		r.Code = enums.BadRequest
		r.Status = enums.StatusFailed
		r.Message = "Bad request"
		return
	case enums.InternalError:
		r.Code = enums.InternalError
		r.Status = "pending"
		r.Message = "Internal server error"
		return
	case enums.NotFound:
		r.Code = enums.NotFound
		r.Status = enums.StatusFailed
		r.Message = "Not found"
		return
	case enums.Unauthorized:
		r.Code = enums.Unauthorized
		r.Status = enums.StatusFailed
		r.Message = "Unauthorized"
		return
	case enums.TokenExpired:
		r.Code = enums.TokenExpired
		r.Status = enums.StatusFailed
		r.Message = "Token expired"
		return
	case enums.UserNotActive:
		r.Code = enums.UserNotActive
		r.Status = enums.StatusFailed
		r.Message = "The user is blocked"
		return
	case enums.SuccessPartially:
		r.Code = enums.SuccessPartially
		r.Message = "success partially"
		r.Status = "warning"
		return
	default:
		r.Code = code
		r.Status = ""
		r.Message = "undefined invalid code"
	}
}
