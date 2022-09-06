package service

import (
	"context"
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/enums"
	"go.uber.org/zap"
	"time"
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

func (s *service) GetPointsJournal(ctx context.Context, req dto.GetPointsJournalRequest) (resp dto.Response) {
	if req.CourseId < 1 || req.UserUchprocCode < 1 {
		resp.ErrCode(enums.BadRequest)
		return
	}

	journal, err := s.mssqlDB.GetPointsJournal(ctx, req.CourseId)
	if err != nil {
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		s.log.Error("internal/service.uchproc.go, GetPointsJournal, s.mssqlDB.GetPointsJournal", zap.Error(err), zap.Any("Request", req))
		return
	}

	tm, err := s.mssqlDB.GetLastAYStartTime(ctx)
	if err != nil {
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		s.log.Error("internal/service.uchproc.go, GetPointsJournal, s.mssqlDB.GetLastAYStartTime", zap.Error(err), zap.Any("Request", req))
		return
	}

	diff := time.Now().Sub(tm)

	//1 week = 168 hours
	journal.CurrentWeek = int(diff.Hours()/168) + 1

	resp.ErrCode(enums.Success)
	resp.Payload = journal

	return
}
