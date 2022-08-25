package storage

import (
	"context"
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/models"
)

type MSSQLDB interface {
	GetUserAuthParams(ctx context.Context, login string) (params dto.AuthParams, err error)
	UserGetLoginByUchprocId(ctx context.Context, uchprocId int64) (login string, err error)
	GetFaculties(ctx context.Context, userUchprocId int64) (res []dto.Faculty, err error)
	Close()
}

type PostgresDB interface {
	UserInsert(ctx context.Context, user models.User) (id int64, err error)
	UserGetByToken(ctx context.Context, token string) (user models.User, err error)
	UserDeleteByToken(ctx context.Context, token string) error
	Close()
}
