package service

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/enums"
	"github.com/iamgafurov/journal/internal/models"
	"go.uber.org/zap"
	"log"
	"strconv"
	"time"
)

func (s *service) TopicGetAll(ctx context.Context, req dto.TopicAllRequest) (resp dto.Response) {
	if req.CourseId == 0 || req.UserUchprocCode == 0 {
		resp.ErrCode(enums.BadRequest)
		return
	}

	topics, err := s.mssqlDB.GetTopics(ctx, req.UserUchprocCode, req.CourseId)
	if err != nil {
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		sentry.CaptureException(err)
		s.log.Error("internal/service.topic.go, TopicGetAll,  s.mssqlDB.TopicGetAll", zap.Error(err), zap.Any("Request", req))
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
		s.log.Error("internal/service.topic.go, TopicDelete,  s.mssqlDB.DeleteTopic", zap.Error(err), zap.Any("Request", req))
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
		s.log.Error("internal/service.topic.go, TopicUpdate,  s.mssqlDB.UpdateTopic", zap.Error(err), zap.Any("Request", req))
		return
	}
	resp.ErrCode(enums.Success)
	return
}

func (s *service) TopicCreate(ctx context.Context, req dto.TopicCreateRequest) (resp dto.Response) {
	log.Println(req.UserUchprocCode)
	if req.UserUchprocCode == 0 {
		resp.ErrCode(enums.BadRequest)
		resp.ErrStr = "empty user code"
		return
	}

	if req.CourseId == 0 {
		resp.ErrCode(enums.BadRequest)
		resp.ErrStr = "empty courseId"
		return
	}

	if !req.Valid() {
		resp.ErrCode(enums.BadRequest)
		resp.ErrStr = "invalid topic"
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
		s.log.Error("internal/service.topic.go, TopicCreate,  s.mssqlDB.GetAttendanceStatement", zap.Error(err), zap.Any("Request", req))
		return
	}

	if statement.Kst != req.UserUchprocCode && statement.Kas != req.UserUchprocCode {
		resp.ErrCode(enums.BadRequest)
		resp.Message = "course does not belong to this user"
		return
	}

	cnzap, err := s.mssqlDB.GetCurrentCnzap(ctx, req.CourseId)
	if err != nil {
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		s.log.Error("internal/service.topic.go, TopicCreate,  s.mssqlDB.GetCurrentCnzap", zap.Error(err), zap.Any("Request", req))
		return
	}
	cn, err := strconv.Atoi(cnzap)
	if err != nil {
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		s.log.Error("internal/service.topic.go, TopicCreate,  cannot conver cnzap to int", zap.Error(err), zap.Any("Cnzap", cnzap))
		return
	}

	topic := models.Topic{
		Cnzap:     strconv.Itoa(cn + 1),
		Kvd:       req.CourseId,
		Tema:      req.Topic.Tema,
		Dtzap:     time.Now(),
		Kst:       req.UserUchprocCode,
		IsuSotId:  req.UserUchprocCode,
		DtActive:  time.Now(),
		ChActive:  1,
		VchAccess: statement.VchAccess,
		Chruchgod: statement.Chruchgod,
		KolLab:    req.Topic.KolLab,
		KolPrak:   req.Topic.KolPrak,
		KolKmd:    req.Topic.KolKmd,
		KolSem:    req.Topic.KolSem,
		KolLek:    req.Topic.KolLek,
	}
	topic.KolObsh = topic.KolPrak + topic.KolKmd + topic.KolLab + topic.KolSem + topic.KolLek

	id, err := s.mssqlDB.CreateTopic(ctx, topic)
	if err != nil {
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		s.log.Error("internal/service.topic.go, TopicCreate,  s.mssqlDB.TopicCreate", zap.Error(err), zap.Any("Topic", topic))
		return
	}

	resp.ErrCode(enums.Success)
	topic.Id = id
	topic.Editable = true
	resp.Payload = topic
	return
}
