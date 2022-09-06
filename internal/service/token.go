package service

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/enums"
	"github.com/iamgafurov/journal/internal/models"
	"github.com/iamgafurov/journal/internal/tools"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"log"
	"strings"
	"time"
)

func (s *service) Tokenize(ctx context.Context, request dto.TokenizeRequest) (resp dto.Response) {
	var (
		l = request.LoginPass.Login
		p = request.LoginPass.Password
	)
	//defer sentry.Recover()

	if request.ServiceName != enums.ServiceMobi && request.ServiceName != enums.ServiceWeb {
		resp.ErrCode(enums.BadRequest)
		resp.Message = "invalid service name"
		return
	}

	if tools.StrEmpty(request.LoginPass.Login) || tools.StrEmpty(request.LoginPass.Password) {
		resp.ErrCode(enums.BadRequest)
		resp.Message = "invalid login/password"
		return
	}

	params, err := s.mssqlDB.GetUserAuthParams(ctx, l)
	log.Println(params)
	if err != nil {
		if err == dto.ErrNoRows {
			resp.ErrCode(enums.Unauthorized)
			return
		}
		resp.ErrCode(enums.InternalError)
		s.log.Error("service/user.go Tokenize/GetUserAuthParams", zap.Error(err), zap.Any("ExternalRef", request.ExternalRef))
		sentry.CaptureException(err)
		return
	}
	paramPass := ParamPass{TmK: params.Tmk, CLogin: params.Login, PrP: params.Password}
	//paramPass := ParamPass{TmK: "18:19:18", PrP: "|8240|0100|1052|1026|0004|0016|1026|0064|0016|0004|1026|0008|0064|0001|0002|0002", CLogin: "Oper_05             "}
	hpass := DecodePassword(paramPass)
	if hpass != p {
		log.Println(hpass, p)
		resp.ErrCode(enums.Unauthorized)
		resp.Message = "Wrong password"
		return
	}

	th := dto.TokenHash{
		Login: l,
		Id:    params.UserId,
		Time:  time.Now().UnixNano(),
	}

	bt, err := jsoniter.Marshal(th)
	if err != nil {
		resp.ErrCode(enums.InternalError)
		sentry.CaptureException(err)
		s.log.Error("service/user.go Tokenize/Marshal", zap.Error(err))
		resp.ErrStr = err.Error()
		return
	}

	token := tools.HmacHash(bt, s.cfg.MasterKey)

	user := models.User{
		Login:       l,
		Token:       token,
		Service:     request.ServiceName,
		ExpireAt:    time.Now().Add(time.Duration(s.cfg.TokensDurationInHours) * time.Hour),
		Status:      enums.StatusActive,
		UchprocId:   params.UserId,
		UchprocCode: params.UserCode,
		Name:        strings.TrimSpace(params.UserName),
	}

	_, err = s.postgresDB.UserInsert(ctx, user)
	if err != nil {
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		sentry.CaptureException(err)
		s.log.Error("service/user Tokenize/UserInsert", zap.Error(err), zap.String("Service", request.ServiceName), zap.String("ExternalRef", request.ExternalRef))
		return
	}

	resp.ErrCode(enums.Success)
	resp.Payload = dto.TokenizePayload{
		Token:    token,
		ExpireAt: user.ExpireAt,
	}

	return
}

func (s *service) TokenDelete(ctx context.Context, req dto.DeleteTokenRequest) (resp dto.Response) {
	if tools.StrEmpty(req.Token) {
		resp.ErrCode(enums.BadRequest)
		return
	}

	err := s.postgresDB.UserDeleteByToken(ctx, req.Token)
	if err != nil {
		if err == dto.ErrNoRowsAffected {
			resp.ErrCode(enums.NotFound)
			resp.Message = "token not found"
			return
		}
		resp.ErrCode(enums.InternalError)
		s.log.Error("service/user.go, Untokenize/UserDeleteByToken", zap.Error(err), zap.String("ExternalRef", req.ExternalRef))
		return
	}

	resp.ErrCode(enums.Success)
	return
}
