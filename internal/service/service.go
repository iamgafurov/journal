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
	Tokenize(ctx context.Context, request dto.TokenizeRequest) (resp dto.Response)
	TokenDelete(ctx context.Context, req dto.DeleteTokenRequest) (resp dto.Response)
	UserGetByToken(ctx context.Context, token string) (models.User, error)
	UserFaculties(ctx context.Context, se dto.ServiceNameExternalRef, uchprosId int64) (resp dto.Response)
	CheckUser(ctx context.Context, req dto.CheckUserRequest) (resp dto.Response)
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
