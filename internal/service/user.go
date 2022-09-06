package service

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/enums"
	"github.com/iamgafurov/journal/internal/models"
	"github.com/iamgafurov/journal/internal/tools"
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
		s.log.Error("internal/service/user.go CheckUser/UserGetLoginByUchprocId", zap.Error(err), zap.Any("request", req))
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

func (s *service) UserFaculties(ctx context.Context, req dto.FacultiesRequest, uchprosId int64) (resp dto.Response) {
	if uchprosId == 0 {
		resp.ErrCode(enums.BadRequest)
		return
	}

	if tools.StrEmpty(req.AcademicYear) {
		resp.ErrCode(enums.BadRequest)
		return
	}

	faculties, err := s.mssqlDB.GetFaculties(ctx, uchprosId, req.AcademicYear)
	if err != nil {
		if err == dto.ErrNoRows {
			resp.ErrCode(enums.NotFound)
			return
		}
		resp.ErrCode(enums.InternalError)
		s.log.Error("internal/service/user.go, UserFaculties, s.mssqlDB.GetFaculties", zap.Error(err), zap.Int64("UchprocId", uchprosId), zap.Any("request", req))
		sentry.CaptureException(err)
		return
	}

	resp.ErrCode(enums.Success)
	resp.Payload = dto.MainFilterPayload{Faculties: faculties}
	return
}

func (s *service) AcademicYears(ctx context.Context, req dto.AcademicYearsRequest) (resp dto.Response) {
	if req.UserUchprocCode == 0 {
		resp.ErrCode(enums.BadRequest)
		return
	}

	years, err := s.mssqlDB.GetAcademicYears(ctx, req.UserUchprocCode)
	if err != nil {
		resp.ErrCode(enums.InternalError)
		s.log.Error("internal/service/user.go AcademicYears s.mssqlDB.GetAcademicYears", zap.Error(err), zap.Any("request", req))
		sentry.CaptureException(err)
		return
	}

	resp.ErrCode(enums.Success)
	resp.Payload = dto.AcademicYearsPayload{AcademicYears: years}
	return
}
