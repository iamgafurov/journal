package rest

import (
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/enums"
	"github.com/iamgafurov/journal/internal/logger"
	"github.com/iamgafurov/journal/internal/models"
	"go.uber.org/zap"
	"net/http"
)

// @Summary Get all topics
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param Service header string true "Service Name"
// @Param Request body dto.TopicAllRequest true "Request body"
// @Router /topic/all [get]
// @Success 200 {object} dto.GetTopicsPayload
func (s *Server) topics(w http.ResponseWriter, r *http.Request) {
	var req dto.TopicAllRequest

	if !parseBody(r, &req) {
		s.reply(w, dto.Response{Code: enums.BadRequest, Message: "bad request"})
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		s.reply(w, dto.Response{Code: enums.InternalError})
		logger.Logger.Error("handler/rest/topic.go, topics cannot cast user from context", zap.Any("Request", req))
		return
	}

	req.UserUchprocCode = user.UchprocCode
	serviceResp := s.service.TopicGetAll(r.Context(), req)

	if s.cfg.Debug {
		logger.Logger.Info("topics", zap.Any("Request", req), zap.Any("Response", serviceResp), zap.String("errStr", serviceResp.ErrStr))
	}

	s.reply(w, serviceResp)
}

// @Summary delete topic
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param Service header string true "Service Name"
// @Param Request body dto.TopicDeleteRequest true "Request body"
// @Router /topic/delete [post]
// @Success 200 {object} dto.Response
func (s *Server) topicDelete(w http.ResponseWriter, r *http.Request) {
	var req dto.TopicDeleteRequest

	if !parseBody(r, &req) {
		s.reply(w, dto.Response{Code: enums.BadRequest, Message: "bad request"})
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		s.reply(w, dto.Response{Code: enums.InternalError})
		logger.Logger.Error("handler/rest/topic.go, topicDelete cannot cast user from context", zap.Any("Request", req))
		return
	}

	req.UserUchprocCode = user.UchprocCode
	serviceResp := s.service.TopicDelete(r.Context(), req)

	if s.cfg.Debug {
		logger.Logger.Info("topicDelete", zap.Any("Request", req), zap.Any("Response", serviceResp), zap.String("errStr", serviceResp.ErrStr))
	}

	s.reply(w, serviceResp)
}

// @Summary update topic
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param Service header string true "Service Name"
// @Param Request body dto.TopicUpdateRequest true "Request body"
// @Router /topic/update [post]
// @Success 200 {object} dto.Response
func (s *Server) topicUpdate(w http.ResponseWriter, r *http.Request) {
	var req dto.TopicUpdateRequest

	if !parseBody(r, &req) {
		s.reply(w, dto.Response{Code: enums.BadRequest, Message: "bad request"})
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		s.reply(w, dto.Response{Code: enums.InternalError})
		logger.Logger.Error("handler/rest/topic.go, topicUpdate cannot cast user from context", zap.Any("Request", req))
		return
	}

	req.UserUchprocCode = user.UchprocCode
	serviceResp := s.service.TopicUpdate(r.Context(), req)

	if s.cfg.Debug {
		logger.Logger.Info("topicUpdate", zap.Any("Request", req), zap.Any("Response", serviceResp), zap.String("errStr", serviceResp.ErrStr))
	}

	s.reply(w, serviceResp)
}

// @Summary create topic
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param Service header string true "Service Name"
// @Param Request body dto.TopicUpdateRequest true "Request body"
// @Router /topic/create [post]
// @Success 200 {object} models.Topic
func (s *Server) topicCreate(w http.ResponseWriter, r *http.Request) {
	var req dto.TopicCreateRequest

	if !parseBody(r, &req) {
		s.reply(w, dto.Response{Code: enums.BadRequest, Message: "bad request"})
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		s.reply(w, dto.Response{Code: enums.InternalError})
		logger.Logger.Error("handler/rest/topic.go, topicCreate cannot cast user from context", zap.Any("Request", req))
		return
	}

	req.UserUchprocCode = user.UchprocCode
	serviceResp := s.service.TopicCreate(r.Context(), req)

	if s.cfg.Debug {
		logger.Logger.Info("topicCreate", zap.Any("Request", req), zap.Any("Response", serviceResp), zap.String("errStr", serviceResp.ErrStr))
	}

	s.reply(w, serviceResp)
}
