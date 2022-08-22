package rest

import (
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/enums"
	"github.com/iamgafurov/journal/internal/logger"
	"go.uber.org/zap"
	"net/http"
)

//Ping handler
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
