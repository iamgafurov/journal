package dto

import (
	"github.com/iamgafurov/journal/internal/models"
	"time"
)

type TokenizePayload struct {
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expire_at"`
	FullName string    `json:"full_name"`
	UserCode int64     `json:"user_code"`
}

type TokenHash struct {
	Id    int64  `json:"id"`
	Login string `json:"login"`
	Time  int64  `json:"time"`
}

type MainFilterPayload struct {
	Faculties []Faculty `json:"faculties"`
}

type Faculty struct {
	Id          int64        `json:"id"`
	Code        string       `json:"code"`
	Name        string       `json:"name"`
	Specialties []Speciality `json:"specialties"`
}

type Speciality struct {
	Id    int64  `json:"id"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Years []Year `json:"years"`
}

type Year struct {
	Id     int64   `json:"id"`
	Code   string  `json:"code"`
	Groups []Group `json:"groups"`
}

type Group struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
}

type GroupCoursesPayload struct {
	Courses []Course `json:"courses"`
}

type AcademicYearsPayload struct {
	AcademicYears []string `json:"academic_years"`
}

type GetTopicsPayload struct {
	Topics []models.Topic `json:"topics"`
}

type GetAttendanceJournalPayload struct {
	Journal []StudentAttendance `json:"journal"`
}
