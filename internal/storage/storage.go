package storage

import (
	"context"
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/models"
	"time"
)

type MSSQLDB interface {
	GetUserAuthParams(ctx context.Context, login string) (params dto.AuthParams, err error)
	UserGetLoginByUchprocId(ctx context.Context, uchprocId int64) (login string, err error)
	GetFaculties(ctx context.Context, userUchprocId int64, studyYear string) (res []dto.Faculty, err error)
	GetGroupCourses(ctx context.Context, groupId, userId int64) (cs []dto.Course, err error)
	GetAcademicYears(ctx context.Context, userUchprocCode int64) ([]string, error)
	GetTopics(ctx context.Context, userCode, courseId int64) ([]models.Topic, error)
	DeleteTopic(ctx context.Context, topicId, userCode int64) error
	UpdateTopic(ctx context.Context, topic models.Topic, userCode int64) error
	GetPointsJournal(ctx context.Context, courseId int64) (dto.PointJournal, error)
	GetLastAYStartTime(ctx context.Context) (tm time.Time, err error)
	Close()
}

type PostgresDB interface {
	UserInsert(ctx context.Context, user models.User) (id int64, err error)
	UserGetByToken(ctx context.Context, token string) (user models.User, err error)
	UserDeleteByToken(ctx context.Context, token string) error
	Close()
}
