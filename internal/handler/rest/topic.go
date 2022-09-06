package rest

import (
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/enums"
	"github.com/iamgafurov/journal/internal/logger"
	"github.com/iamgafurov/journal/internal/models"
	"go.uber.org/zap"
	"net/http"
)

func (s *Server) topics(w http.ResponseWriter, r *http.Request) {
	var req dto.TopicsRequest

	if !parseBody(r, &req) {
		s.reply(w, dto.Response{Code: enums.BadRequest, Message: "bad request"})
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		s.reply(w, dto.Response{Code: enums.InternalError})
		logger.Logger.Error("handler/rest/user.go, topics cannot cast user from context", zap.Any("Request", req))
		return
	}

	req.UserUchprocCode = user.UchprocCode
	serviceResp := s.service.TopicGetAll(r.Context(), req)

	if s.cfg.Debug {
		logger.Logger.Info("topics", zap.Any("Request", req), zap.Any("Response", serviceResp))
	}

	s.reply(w, serviceResp)
}

func (s *Server) topicDelete(w http.ResponseWriter, r *http.Request) {
	var req dto.TopicDeleteRequest

	if !parseBody(r, &req) {
		s.reply(w, dto.Response{Code: enums.BadRequest, Message: "bad request"})
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		s.reply(w, dto.Response{Code: enums.InternalError})
		logger.Logger.Error("handler/rest/user.go, topicDelete cannot cast user from context", zap.Any("Request", req))
		return
	}

	req.UserUchprocCode = user.UchprocCode
	serviceResp := s.service.TopicDelete(r.Context(), req)

	if s.cfg.Debug {
		logger.Logger.Info("topicDelete", zap.Any("Request", req), zap.Any("Response", serviceResp))
	}

	s.reply(w, serviceResp)
}

func (s *Server) topicUpdate(w http.ResponseWriter, r *http.Request) {
	var req dto.TopicUpdateRequest

	if !parseBody(r, &req) {
		s.reply(w, dto.Response{Code: enums.BadRequest, Message: "bad request"})
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		s.reply(w, dto.Response{Code: enums.InternalError})
		logger.Logger.Error("handler/rest/user.go, topicUpdate cannot cast user from context", zap.Any("Request", req))
		return
	}

	req.UserUchprocCode = user.UchprocCode
	serviceResp := s.service.TopicUpdate(r.Context(), req)

	if s.cfg.Debug {
		logger.Logger.Info("topicUpdate", zap.Any("Request", req), zap.Any("Response", serviceResp))
	}

	s.reply(w, serviceResp)
}
