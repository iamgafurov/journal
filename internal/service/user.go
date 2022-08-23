package service

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/enums"
	"github.com/iamgafurov/journal/internal/models"
	"go.uber.org/zap"
)

type ParamPass struct {
	TmK    string `json:"tmk"`
	PrP    string `json:"pr_p"`
	CLogin string `json:"clogin"`
}

func (s *service) UserGetByToken(ctx context.Context, token string) (user models.User, err error) {
	return s.postgresDB.UserGetByToken(ctx, token)
}

func (s *service) CheckUser(ctx context.Context, req dto.CheckUserRequest) (resp dto.Response) {
	login, err := s.mssqlDB.UserGetLoginByUchprocId(ctx, req.UchprocId)
	if err != nil {
		if err == dto.ErrNoRows {
			resp.ErrCode(enums.NotFound)
			resp.ErrStr = "user not found"
			return
		}
		resp.ErrCode(enums.InternalError)
		s.log.Error("internal/service/user.go CheckUser/UserGetLoginByUchprocId", zap.Error(err), zap.Any("Request", req))
		sentry.CaptureException(err)
		return
	}

	if login != req.Login {
		resp.ErrCode(enums.Unauthorized)
		resp.ErrStr = "login not correct"
		return
	}
	resp.ErrCode(enums.Success)
	return
}
