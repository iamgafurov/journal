package service

import (
	"context"
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/enums"
	"go.uber.org/zap"
	"log"
	"time"
)

func (s *service) GetPointsJournal(ctx context.Context, req dto.GetJournalRequest) (resp dto.Response) {
	if req.CourseId < 1 || req.UserUchprocCode < 1 {
		resp.ErrCode(enums.BadRequest)
		return
	}

	journal, err := s.mssqlDB.GetPointsJournal(ctx, req.CourseId)
	if err != nil {
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		s.log.Error("internal/service.point_journal.go, GetPointsJournal, s.mssqlDB.GetPointsJournal", zap.Error(err), zap.Any("Request", req))
		return
	}

	cw, err := s.getCurrentWeek(ctx)
	if err != nil {
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		s.log.Error("internal/service.point_journal.go, GetPointsJournal, s.getCurrentWeek", zap.Error(err), zap.Any("Request", req))
	}
	journal.CurrentWeek = cw
	log.Println(len(journal.Students))
	resp.ErrCode(enums.Success)
	resp.Payload = journal

	return
}

func (s *service) UpdatePointJournal(ctx context.Context, req dto.UpdatePointJournalRequest) (resp dto.Response) {
	if req.CourseId == 0 {
		resp.ErrCode(enums.BadRequest)
		resp.ErrStr = "empty courseId"
		return
	}

	if req.UserUchprocCode == 0 {
		resp.ErrCode(enums.BadRequest)
		resp.ErrStr = "empty userCode"
		return
	}

	for _, p := range req.Points {
		if p.Id == 0 {
			resp.ErrCode(enums.BadRequest)
			resp.ErrStr = "null point.id"
			return
		}
	}

	userCode, err := s.mssqlDB.GetPointUserCode(ctx, req.CourseId)
	if err != nil {
		if err == dto.ErrNoRows {
			resp.ErrCode(enums.BadRequest)
			resp.Message = "course not exists"
			return
		}
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		s.log.Error("internal/service.point_journal.go, UpdatePointJournal, s.mssqlDB.GetPointUserCode", zap.Error(err), zap.Any("Request", req))
		return
	}

	if req.UserUchprocCode != userCode {
		resp.ErrCode(enums.BadRequest)
		resp.Message = "course does not belong to this user"
		return
	}

	cw, err := s.getCurrentWeek(ctx)
	if err != nil {
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		s.log.Error("internal/service.point_journal.go, UpdatePointJournal, s.getCurrentWeek", zap.Error(err), zap.Any("Request", req))
		return
	}

	err = s.mssqlDB.UpdatePointsJournal(ctx, req.Points, req.CourseId, cw)
	if err != nil {
		if err == dto.ErrNoRowsAffected {
			resp.ErrCode(enums.BadRequest)
			resp.Message = "some students don't exist"
			return
		}
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		s.log.Error("internal/service.point_journal.go, UpdatePointJournal,  s.mssqlDB.UpdatePointsJournal", zap.Error(err), zap.Any("Request", req))
		return
	}

	resp.ErrCode(enums.Success)
	return
}

func (s *service) getCurrentWeek(ctx context.Context) (int, error) {
	tm, err := s.mssqlDB.GetLastAYStartTime(ctx)
	if err != nil {
		return 0, err
	}

	diff := time.Now().Sub(tm)

	//1 week = 168 hours
	return int(diff.Hours()/168) + 1, nil
}
