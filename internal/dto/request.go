package dto

import "github.com/iamgafurov/journal/internal/models"

type TokenizeRequest struct {
	ServiceName string    `json:"service_name"`
	ExternalRef string    `json:"external_ref"`
	LoginPass   LoginPass `json:"login_pass"`
}

type DeleteTokenRequest struct {
	Token       string `json:"token"`
	ServiceName string `json:"service_name"`
	ExternalRef string `json:"external_ref"`
}

type LoginPass struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type CheckUserRequest struct {
	Login     string `json:"login"`
	UchprocId int64  `json:"uchproc_id"`
}

type FacultiesRequest struct {
	ServiceName  string `json:"service_name"`
	ExternalRef  string `json:"external_ref"`
	AcademicYear string `json:"academic_year"`
}

type GroupCoursesRequest struct {
	ServiceName     string `json:"service_name"`
	ExternalRef     string `json:"external_ref"`
	GroupId         int64  `json:"group_id"`
	AcademicYear    string `json:"academic_year"`
	UserUchprocCode int64
}

type AcademicYearsRequest struct {
	ServiceName     string `json:"service_name"`
	ExternalRef     string `json:"external_ref"`
	UserUchprocCode int64
}

type TopicAllRequest struct {
	ServiceName     string `json:"service_name"`
	ExternalRef     string `json:"external_ref"`
	CourseId        int64  `json:"course_id"`
	UserUchprocCode int64
}

type TopicDeleteRequest struct {
	ServiceName     string `json:"service_name"`
	ExternalRef     string `json:"external_ref"`
	TopicId         int64  `json:"topicId"`
	UserUchprocCode int64
}

type TopicUpdateRequest struct {
	ServiceName     string       `json:"service_name"`
	ExternalRef     string       `json:"external_ref"`
	TopicId         int64        `json:"topic_id"`
	Topic           models.Topic `json:"topic"`
	UserUchprocCode int64
}
type TopicCreateRequest struct {
	ServiceName     string       `json:"service_name"`
	ExternalRef     string       `json:"external_ref"`
	CourseId        int64        `json:"course_id"`
	Topic           models.Topic `json:"topic"`
	UserUchprocCode int64
}

type GetJournalRequest struct {
	ServiceName     string `json:"service_name"`
	ExternalRef     string `json:"external_ref"`
	Limit           int    `json:"limit"`
	CourseId        int64  `json:"course_id"`
	UserUchprocCode int64
}

type UpdateAttendanceJournalRequest struct {
	ServiceName     string `json:"service_name"`
	ExternalRef     string `json:"external_ref"`
	CourseId        int64  `json:"course_id"`
	UserUchprocCode int64
	Attendance      []StudentAttendance `json:"attendance"`
}

type UpdatePointJournalRequest struct {
	ServiceName     string        `json:"service_name"`
	ExternalRef     string        `json:"external_ref"`
	CourseId        int64         `json:"course_id"`
	Points          []PointUpdate `json:"points"`
	UserUchprocCode int64
}

func (r *TopicCreateRequest) Valid() bool {
	min := 0
	max := 10
	if !(r.Topic.KolSem >= min && r.Topic.KolSem <= max) {
		return false
	}
	if !(r.Topic.KolPrak >= min && r.Topic.KolPrak <= max) {
		return false
	}
	if !(r.Topic.KolKmd >= min && r.Topic.KolKmd <= max) {
		return false
	}
	if !(r.Topic.KolLab >= min && r.Topic.KolLab <= max) {
		return false
	}
	if !(r.Topic.KolLek >= min && r.Topic.KolLek <= max) {
		return false
	}

	return true
}
