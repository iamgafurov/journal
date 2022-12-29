package service

import (
	"context"
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/enums"
	"github.com/iamgafurov/journal/internal/tools"
	"go.uber.org/zap"
)

func (s *service) GetGroupCoursesAttendance(ctx context.Context, req dto.GroupCoursesRequest) (resp dto.Response) {
	if req.GroupId < 1 {
		resp.ErrCode(enums.BadRequest)
		return
	}
	if tools.StrEmpty(req.AcademicYear) {
		resp.ErrCode(enums.BadRequest)
		return
	}

	courses, err := s.mssqlDB.GetGroupCoursesAttendance(ctx, req.GroupId, req.UserUchprocCode, req.AcademicYear)
	if err != nil {
		if err == dto.ErrNoRows {
			resp.ErrCode(enums.NotFound)
			return
		}
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		s.log.Error("internal/service.uchproc.go, GetGroupCoursesAttendance, s.mssqlDB.GetGroupCoursesAttendance", zap.Error(err), zap.Any("Request", req))
		return
	}

	resp.ErrCode(enums.Success)
	resp.Payload = dto.GroupCoursesPayload{Courses: courses}
	return
}

func (s *service) GetGroupCoursesPoint(ctx context.Context, req dto.GroupCoursesRequest) (resp dto.Response) {
	if req.GroupId < 1 {
		resp.ErrCode(enums.BadRequest)
		return
	}
	if tools.StrEmpty(req.AcademicYear) {
		resp.ErrCode(enums.BadRequest)
		return
	}

	courses, err := s.mssqlDB.GetGroupCoursesPoint(ctx, req.GroupId, req.UserUchprocCode, req.AcademicYear)
	if err != nil {
		if err == dto.ErrNoRows {
			resp.ErrCode(enums.NotFound)
			return
		}
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		s.log.Error("internal/service.uchproc.go, GetGroupCoursesPoint, s.mssqlDB.GetGroupCoursesPoint", zap.Error(err), zap.Any("Request", req))
		return
	}

	resp.ErrCode(enums.Success)
	resp.Payload = dto.GroupCoursesPayload{Courses: courses}
	return
}
