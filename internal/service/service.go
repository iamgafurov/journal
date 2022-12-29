package service

import (
	"context"
	"github.com/iamgafurov/journal/internal/config"
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/models"
	"github.com/iamgafurov/journal/internal/storage"
	"go.uber.org/zap"
)

type Service interface {
	//token
	Tokenize(ctx context.Context, request dto.TokenizeRequest) (resp dto.Response)
	TokenDelete(ctx context.Context, req dto.DeleteTokenRequest) (resp dto.Response)
	UserGetByToken(ctx context.Context, token string) (models.User, error)

	UserFaculties(ctx context.Context, se dto.FacultiesRequest, uchprosId int64) (resp dto.Response)
	GetGroupCoursesAttendance(ctx context.Context, req dto.GroupCoursesRequest) (resp dto.Response)
	GetGroupCoursesPoint(ctx context.Context, req dto.GroupCoursesRequest) (resp dto.Response)
	CheckUser(ctx context.Context, req dto.CheckUserRequest) (resp dto.Response)
	AcademicYears(ctx context.Context, req dto.AcademicYearsRequest) (resp dto.Response)

	//topic
	TopicGetAll(ctx context.Context, req dto.TopicAllRequest) (resp dto.Response)
	TopicCreate(ctx context.Context, req dto.TopicCreateRequest) (resp dto.Response)
	TopicDelete(ctx context.Context, req dto.TopicDeleteRequest) (resp dto.Response)
	TopicUpdate(ctx context.Context, req dto.TopicUpdateRequest) (resp dto.Response)

	GetPointsJournal(ctx context.Context, req dto.GetJournalRequest) (resp dto.Response)
	UpdatePointJournal(ctx context.Context, req dto.UpdatePointJournalRequest) (resp dto.Response)

	GetAttendanceJournal(ctx context.Context, req dto.GetJournalRequest) (resp dto.Response)
	UpdateAttendanceJournal(ctx context.Context, req dto.UpdateAttendanceJournalRequest) (resp dto.Response)

	StartTurnstileWorker(ctx context.Context)
}

type service struct {
	postgresDB storage.PostgresDB
	mssqlDB    storage.MSSQLDB
	cfg        *config.Config
	log        *zap.Logger
}

func New(pdb storage.PostgresDB, mssqldb storage.MSSQLDB, cfg *config.Config, log *zap.Logger) Service {
	return &service{postgresDB: pdb, mssqlDB: mssqldb, cfg: cfg, log: log}
}
