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
	UserGetByToken(ctx context.Context, token string) (models.User, error)
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
