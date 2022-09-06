package service

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/enums"
	"go.uber.org/zap"
)

func (s *service) TopicGetAll(ctx context.Context, req dto.TopicsRequest) (resp dto.Response) {
	if req.CourseId == 0 || req.UserUchprocCode == 0 {
		resp.ErrCode(enums.BadRequest)
		return
	}
	//req.UserUchprocCode = 33
	//req.CourseId = 374

	topics, err := s.mssqlDB.GetTopics(ctx, req.UserUchprocCode, req.CourseId)
	if err != nil {
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		sentry.CaptureException(err)
		s.log.Error("internal/service.uchproc.go, TopicGetAll,  s.mssqlDB.TopicGetAll", zap.Error(err), zap.Any("Request", req))
		return
	}
	resp.ErrCode(enums.Success)
	resp.Payload = dto.GetTopicsPayload{Topics: topics}
	return
}

func (s *service) TopicDelete(ctx context.Context, req dto.TopicDeleteRequest) (resp dto.Response) {
	if req.TopicId == 0 || req.UserUchprocCode == 0 {
		resp.ErrCode(enums.BadRequest)
		return
	}

	err := s.mssqlDB.DeleteTopic(ctx, req.UserUchprocCode, req.TopicId)
	if err != nil {
		if err == dto.ErrNoRowsAffected {
			resp.ErrCode(enums.NotFound)
			resp.ErrStr = "topic with such id not exists"
			return
		}
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		sentry.CaptureException(err)
		s.log.Error("internal/service.uchproc.go, TopicDelete,  s.mssqlDB.DeleteTopic", zap.Error(err), zap.Any("Request", req))
		return
	}
	resp.ErrCode(enums.Success)
	return
}

func (s *service) TopicUpdate(ctx context.Context, req dto.TopicUpdateRequest) (resp dto.Response) {
	if req.Topic.Id == 0 || req.UserUchprocCode == 0 {
		resp.ErrCode(enums.BadRequest)
		return
	}

	req.Topic.KolObsh = req.Topic.KolLek + req.Topic.KolPrak + req.Topic.KolSem + req.Topic.KolLab + req.Topic.KolKmd
	err := s.mssqlDB.UpdateTopic(ctx, req.Topic, req.UserUchprocCode)
	if err != nil {
		if err == dto.ErrNoRowsAffected {
			resp.ErrCode(enums.NotFound)
			resp.ErrStr = "topic with such id not exists"
			return
		}
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		sentry.CaptureException(err)
		s.log.Error("internal/service.uchproc.go, TopicUpdate,  s.mssqlDB.UpdateTopic", zap.Error(err), zap.Any("Request", req))
		return
	}
	resp.ErrCode(enums.Success)
	return
}
