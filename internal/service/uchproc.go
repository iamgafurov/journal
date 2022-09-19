package service

import (
	"context"
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/enums"
	"go.uber.org/zap"
)

func (s *service) GetGroupCourses(ctx context.Context, req dto.GroupCoursesRequest) (resp dto.Response) {
	if req.GroupId < 1 {
		resp.ErrCode(enums.BadRequest)
		return
	}

	courses, err := s.mssqlDB.GetGroupCourses(ctx, req.GroupId, req.UserUchprocCode)
	if err != nil {
		if err == dto.ErrNoRows {
			resp.ErrCode(enums.NotFound)
			return
		}
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		s.log.Error("internal/service.uchproc.go, GetGroupCourses, GetGroupCourses", zap.Error(err), zap.Any("Request", req))
		return
	}

	resp.ErrCode(enums.Success)
	resp.Payload = dto.GroupCoursesPayload{Courses: courses}
	return
}
