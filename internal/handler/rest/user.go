package rest

import (
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/enums"
	"github.com/iamgafurov/journal/internal/logger"
	"github.com/iamgafurov/journal/internal/models"
	"go.uber.org/zap"
	"net/http"
)

// Tokenization godoc
// @Summary Tokenize user
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param Service header string true "Service Name"
// @Param Request body dto.TokenizeRequest true "Request body"
// @Router /tokenize [post]
// @Success 200 {object} dto.TokenizePayload
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

// Delete token godoc
// @Summary Delete user token
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param Service header string true "Service Name"
// @Param Request body dto.DeleteTokenRequest true "Request body"
// @Router /untokenize [post]
// @Success 200 {object} dto.Response
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

// @Summary Get user faculties, specialities, years and groups
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param Service header string true "Service Name"
// @Param Request body dto.FacultiesRequest true "Request body"
// @Router /faculties [post]
// @Success 200 {object} dto.MainFilterPayload
func (s *Server) userFaculties(w http.ResponseWriter, r *http.Request) {
	var req dto.FacultiesRequest

	if !parseBody(r, &req) {
		s.reply(w, dto.Response{Code: enums.BadRequest, Message: "bad request"})
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		s.reply(w, dto.Response{Code: enums.InternalError})
		logger.Logger.Error("handler/rest/user.go, userFaculties cannot cast user from context", zap.Any("Request", req))
		return
	}

	serviceResp := s.service.UserFaculties(r.Context(), req, user.UchprocCode)

	if s.cfg.Debug {
		logger.Logger.Info("UserFaculties", zap.Any("Request", req), zap.Any("Response", serviceResp))
	}

	s.reply(w, serviceResp)
}

// @Summary Get academic years
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param Service header string true "Service Name"
// @Param Request body dto.AcademicYearsRequest true "Request body"
// @Router /academic_years [post]
// @Success 200 {object} dto.AcademicYearsPayload
func (s *Server) academicYears(w http.ResponseWriter, r *http.Request) {
	var req dto.AcademicYearsRequest

	if !parseBody(r, &req) {
		s.reply(w, dto.Response{Code: enums.BadRequest, Message: "bad request"})
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		s.reply(w, dto.Response{Code: enums.InternalError})
		logger.Logger.Error("handler/rest/user.go, academicYears cannot cast user from context", zap.Any("Request", req))
		return
	}

	req.UserUchprocCode = user.UchprocCode
	serviceResp := s.service.AcademicYears(r.Context(), req)

	if s.cfg.Debug {
		logger.Logger.Info("academicYears", zap.Any("Request", req), zap.Any("Response", serviceResp))
	}

	s.reply(w, serviceResp)
}

// @Summary Get poins journal
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param Service header string true "Service Name"
// @Param Request body dto.GetJournalRequest true "Request body"
// @Router /point_journal/get [get]
// @Success 200 {object} dto.PointJournal
func (s *Server) pointsJournal(w http.ResponseWriter, r *http.Request) {
	var req dto.GetJournalRequest

	if !parseBody(r, &req) {
		s.reply(w, dto.Response{Code: enums.BadRequest, Message: "bad request"})
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		s.reply(w, dto.Response{Code: enums.InternalError})
		logger.Logger.Error("handler/rest/user.go, pointsJournal cannot cast user from context", zap.Any("Request", req))
		return
	}

	req.UserUchprocCode = user.UchprocCode
	serviceResp := s.service.GetPointsJournal(r.Context(), req)

	if s.cfg.Debug {
		logger.Logger.Info("pointsJournal", zap.Any("Request", req), zap.Any("Response", serviceResp))
	}

	s.reply(w, serviceResp)
}

// @Summary Update poins journal
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param Service header string true "Service Name"
// @Param Request body dto.UpdatePointJournalRequest true "Request body"
// @Router /point_journal/update [post]
// @Success 200 {object} dto.Response
func (s *Server) pointsJournalUpdate(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdatePointJournalRequest

	if !parseBody(r, &req) {
		s.reply(w, dto.Response{Code: enums.BadRequest, Message: "bad request"})
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		s.reply(w, dto.Response{Code: enums.InternalError})
		logger.Logger.Error("handler/rest/user.go, pointsJournal cannot cast user from context", zap.Any("Request", req))
		return
	}

	req.UserUchprocCode = user.UchprocCode
	serviceResp := s.service.UpdatePointJournal(r.Context(), req)

	if s.cfg.Debug {
		logger.Logger.Info("pointsJournalUpdate", zap.Any("Request", req), zap.Any("Response", serviceResp))
	}

	s.reply(w, serviceResp)
}

// @Summary Get attendance journal
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param Service header string true "Service Name"
// @Param Request body dto.GetJournalRequest true "Request body"
// @Router /attendance_journal/get [get]
// @Success 200 {object} dto.GetAttendanceJournalPayload
func (s *Server) getAttendanceJournal(w http.ResponseWriter, r *http.Request) {
	var req dto.GetJournalRequest

	if !parseBody(r, &req) {
		s.reply(w, dto.Response{Code: enums.BadRequest, Message: "bad request"})
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		s.reply(w, dto.Response{Code: enums.InternalError})
		logger.Logger.Error("handler/rest/user.go, pointsJournal cannot cast user from context", zap.Any("Request", req))
		return
	}

	req.UserUchprocCode = user.UchprocCode
	serviceResp := s.service.GetAttendanceJournal(r.Context(), req)

	if s.cfg.Debug {
		logger.Logger.Info("getAttendanceJournal", zap.Any("Request", req), zap.Any("Response", serviceResp))
	}

	s.reply(w, serviceResp)
}

// @Summary Update attendance journal
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param Service header string true "Service Name"
// @Param Request body dto.UpdateAttendanceJournalRequest true "Request body"
// @Router /attendance_journal/update [post]
// @Success 200 {object} dto.AttendanceJournalError
func (s *Server) updateAttendanceJournal(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateAttendanceJournalRequest

	if !parseBody(r, &req) {
		s.reply(w, dto.Response{Code: enums.BadRequest, Message: "bad request"})
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		s.reply(w, dto.Response{Code: enums.InternalError})
		logger.Logger.Error("handler/rest/user.go, updateAttendanceJournal cannot cast user from context", zap.Any("Request", req))
		return
	}

	req.UserUchprocCode = user.UchprocCode
	serviceResp := s.service.UpdateAttendanceJournal(r.Context(), req)

	if s.cfg.Debug {
		logger.Logger.Info("updateAttendanceJournal", zap.Any("Request", req), zap.Any("Response", serviceResp))
	}

	s.reply(w, serviceResp)
}
