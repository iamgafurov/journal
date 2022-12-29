package rest

import (
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/enums"
	"github.com/iamgafurov/journal/internal/logger"
	"github.com/iamgafurov/journal/internal/models"
	"go.uber.org/zap"
	"net/http"
)

// @Summary Get group courses
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param Service header string true "Service Name"
// @Param Request body dto.GroupCoursesRequest true "Request body"
// @Router /courses/pt [post]
// @Success 200 {object} dto.GroupCoursesPayload
func (s *Server) groupCoursesPt(w http.ResponseWriter, r *http.Request) {
	var req dto.GroupCoursesRequest

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
	req.UserUchprocCode = user.UchprocCode
	serviceResp := s.service.GetGroupCoursesPoint(r.Context(), req)

	if s.cfg.Debug {
		logger.Logger.Info("UserFaculties", zap.Any("Request", req), zap.Any("Response", serviceResp))
	}

	s.reply(w, serviceResp)
}

// @Summary Get group courses
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param Service header string true "Service Name"
// @Param Request body dto.GroupCoursesRequest true "Request body"
// @Router /courses/at [post]
// @Success 200 {object} dto.GroupCoursesPayload
func (s *Server) groupCoursesAt(w http.ResponseWriter, r *http.Request) {
	var req dto.GroupCoursesRequest

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
	req.UserUchprocCode = user.UchprocCode
	serviceResp := s.service.GetGroupCoursesAttendance(r.Context(), req)

	if s.cfg.Debug {
		logger.Logger.Info("UserFaculties", zap.Any("Request", req), zap.Any("Response", serviceResp))
	}

	s.reply(w, serviceResp)
}
