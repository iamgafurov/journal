package service

import (
	"context"
	"fmt"
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/enums"
	"go.uber.org/zap"
	"log"
	"strings"
)

func (s *service) GetAttendanceJournal(ctx context.Context, req dto.GetJournalRequest) (resp dto.Response) {
	if req.CourseId == 0 {
		resp.ErrCode(enums.BadRequest)
		resp.ErrStr = "null course id"
		return
	}

	if req.Limit < 1 || req.Limit > 96 {
		resp.ErrCode(enums.BadRequest)
		resp.Message = "invalid limit"
		return
	}

	attendance, err := s.mssqlDB.GetAttendanceJournal(ctx, req.CourseId, req.Limit)
	if err != nil {
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		s.log.Error("internal/service.attendance_journal.go, GetAttendanceJournal, s.mssqlDB.GetAttendanceJournal", zap.Error(err), zap.Any("Request", req))
	}

	resp.ErrCode(enums.Success)
	resp.Payload = dto.GetAttendanceJournalPayload{Journal: attendance}
	return
}

func (s *service) UpdateAttendanceJournal(ctx context.Context, req dto.UpdateAttendanceJournalRequest) (resp dto.Response) {
	if req.CourseId == 0 {
		resp.ErrCode(enums.BadRequest)
		resp.ErrStr = "null course id"
		return
	}

	if req.UserUchprocCode == 0 {
		resp.ErrCode(enums.Unauthorized)
		resp.ErrStr = "null user code"
		return
	}

	statement, err := s.mssqlDB.GetAttendanceStatement(ctx, req.CourseId)
	if err != nil {
		if err == dto.ErrNoRows {
			resp.ErrCode(enums.BadRequest)
			resp.Message = "course statement not exist"
			resp.ErrStr = resp.Message
			return
		}
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		s.log.Error("internal/service.attendance_journal.go, UpdateAttendanceJournal,  s.mssqlDB.GetAttendanceStatement", zap.Error(err), zap.Any("Request", req))
		return
	}

	topics, err := s.mssqlDB.GetTopics(ctx, req.UserUchprocCode, req.CourseId)
	if err != nil {
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		s.log.Error("internal/service.attendance_journal.go, UpdateAttendanceJournal,  s.mssqlDB.GetTopics", zap.Error(err), zap.Any("Request", req))
		return
	}

	if statement.Kst != req.UserUchprocCode && statement.Kas != req.UserUchprocCode {
		resp.ErrCode(enums.BadRequest)
		resp.Message = "course does not belong to this user"
		return
	}

	if len(req.Attendance) < 1 {
		resp.ErrCode(enums.BadRequest)
		resp.Message = "empty request"
		return
	}

	for i, at := range req.Attendance {
		if at.Id == 0 {
			resp.ErrCode(enums.BadRequest)
			resp.Message = fmt.Sprintf("empty student id, in item:%v", i)
			resp.ErrStr = resp.Message
			return
		}

		if len(at.Attendance) == 0 {
			resp.ErrCode(enums.BadRequest)
			resp.Message = fmt.Sprintf("empty student attendance, in item:%v", i)
			resp.ErrStr = resp.Message
			return
		}

		for i, v := range at.Attendance {
			if v.Number < 1 {
				resp.ErrCode(enums.BadRequest)
				resp.Message = fmt.Sprintf("invalid topic_number:%v,  for item:%v ", v.Number, i)
				return
			}

			if v.Number > len(topics) {
				resp.ErrCode(enums.BadRequest)
				resp.Message = fmt.Sprintf("topic with number:%v not exist", v.Number)
				return
			}
			if !topics[v.Number-1].Editable {
				log.Println(topics[v.Number])
				resp.ErrCode(enums.BadRequest)
				resp.Message = fmt.Sprintf("topic with number:%v not editable", v.Number)
				return
			}

			if strings.TrimSpace(v.Value) != "" && v.Value != "н" {
				resp.ErrCode(enums.BadRequest)
				resp.Message = fmt.Sprintf("invalid attendance value, miss number:%v, value:%s", v.Number, v.Value)
				resp.ErrStr = resp.Message
				return
			}
		}
	}

	err, attErr := s.mssqlDB.UpdateAttendanceJournal(ctx, req.CourseId, req.Attendance)
	if err != nil {
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		s.log.Error("internal/service.attendance_journal.go, UpdateAttendanceJournal,  s.mssqlDB.UpdateAttendanceJournal", zap.Error(err), zap.Any("Request", req))
		return
	}

	if len(attErr) > 0 {
		resp.ErrCode(enums.SuccessPartially)
		resp.Payload = attErr
		resp.Message = "some records were not saved"
		return
	}

	resp.ErrCode(enums.Success)
	return
}
