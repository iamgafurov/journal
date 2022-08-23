package rest

import (
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/enums"
	"github.com/iamgafurov/journal/internal/logger"
	"github.com/iamgafurov/journal/internal/models"
	"go.uber.org/zap"
	"net/http"
)

func (s *Server) tokenize(w http.ResponseWriter, r *http.Request) {
	var (
		tokenizeRequest dto.TokenizeRequest
	)

	if !parseBody(r, &tokenizeRequest) {
		s.reply(w, dto.Response{Code: enums.BadRequest, Message: "bad request"})
		return
	}

	serviceResp := s.service.Tokenize(r.Context(), tokenizeRequest)

	if s.cfg.Debug {
		logger.Logger.Info("Tokenize", zap.Any("Request", tokenizeRequest), zap.Any("Response", serviceResp))
	}

	s.reply(w, serviceResp)
}

//Ping handler
func (s *Server) tokenDelete(w http.ResponseWriter, r *http.Request) {
	var (
		req dto.DeleteTokenRequest
	)

	if !parseBody(r, &req) {
		s.reply(w, dto.Response{Code: enums.BadRequest, Message: "bad request"})
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		s.reply(w, dto.Response{Code: enums.InternalError})
		logger.Logger.Error("handler/rest/user.go, tokenDelete cannot cast user from context", zap.Any("Request", req))
		return
	}
	req.Token = user.Token

	serviceResp := s.service.TokenDelete(r.Context(), req)

	if s.cfg.Debug {
		logger.Logger.Info("TokenDelete", zap.Any("Request", req), zap.Any("Response", serviceResp))
	}

	s.reply(w, serviceResp)
}
