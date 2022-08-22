package storage

import (
	"context"
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/models"
)

type MSSQLDB interface {
	GetUserAuthParams(ctx context.Context, login string) (params dto.AuthParams, err error)
	Close()
}

type PostgresDB interface {
	UserInsert(ctx context.Context, user models.User) (id int64, err error)
	UserGetByToken(ctx context.Context, token string) (user models.User, err error)
	Close()
}
