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
	GetGroupCoursesAttendance(ctx context.Context, groupId, userId int64, academicYear string) (cs []dto.Course, err error)
	GetGroupCoursesPoint(ctx context.Context, groupId, userId int64, academicYear string) (cs []dto.Course, err error)
	GetAcademicYears(ctx context.Context, userUchprocCode int64) ([]string, error)

	GetTopics(ctx context.Context, userCode, courseId int64) ([]models.Topic, error)
	CreateTopic(ctx context.Context, topic models.Topic) (id int64, err error)
	DeleteTopic(ctx context.Context, topicId, userCode int64) error
	UpdateTopic(ctx context.Context, topic models.Topic, userCode int64) error
	GetCurrentCnzap(ctx context.Context, courseId int64) (cnzap string, err error)

	GetPointsJournal(ctx context.Context, courseId int64) (dto.PointJournal, error)
	GetPointUserCode(ctx context.Context, courseId int64) (code int64, err error)
	UpdatePointsJournal(ctx context.Context, points []dto.PointUpdate, kvdId int64, currentWeek int) (err error)
	GetLastAYStartTime(ctx context.Context) (tm time.Time, err error)

	GetAttendanceJournal(ctx context.Context, courseId int64, limit int) ([]dto.StudentAttendance, error)

	GetAttendanceStatement(ctx context.Context, id int64) (statement models.Statement, err error)
	UpdateAttendanceJournal(ctx context.Context, courseId int64, at []dto.StudentAttendance) (err error, atErr []dto.AttendanceJournalError)

	GetTurnstileStudents(ctx context.Context) ([]dto.CheckAttendanceItem, error)
	Close()
}

type PostgresDB interface {
	UserInsert(ctx context.Context, user models.User) (id int64, err error)
	UserGetByToken(ctx context.Context, token string) (user models.User, err error)
	UserDeleteByToken(ctx context.Context, token string) error
	Close()
}
