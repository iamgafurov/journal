package dto

import "github.com/iamgafurov/journal/internal/models"

type TokenizeRequest struct {
	ServiceName string    `json:"serviceName"`
	ExternalRef string    `json:"externalRef"`
	LoginPass   LoginPass `json:"loginPass"`
}

type DeleteTokenRequest struct {
	Token       string `json:"token"`
	ServiceName string `json:"serviceName"`
	ExternalRef string `json:"externalRef"`
}

type Attendance struct {
}

type LoginPass struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type CheckUserRequest struct {
	Login     string `json:"login"`
	UchprocId int64  `json:"uchprocId"`
}

type FacultiesRequest struct {
	ServiceName  string `json:"serviceName"`
	ExternalRef  string `json:"externalRef"`
	AcademicYear string `json:"academicYear"`
}

type GroupCoursesRequest struct {
	ServiceName     string `json:"serviceName"`
	ExternalRef     string `json:"externalRef"`
	GroupId         int64  `json:"groupId"`
	UserUchprocCode int64
}

type AcademicYearsRequest struct {
	ServiceName     string `json:"serviceName"`
	ExternalRef     string `json:"externalRef"`
	UserUchprocCode int64
}

type TopicsRequest struct {
	ServiceName     string `json:"serviceName"`
	ExternalRef     string `json:"externalRef"`
	CourseId        int64  `json:"courseId"`
	UserUchprocCode int64
}

type TopicDeleteRequest struct {
	ServiceName     string `json:"serviceName"`
	ExternalRef     string `json:"externalRef"`
	TopicId         int64  `json:"topicId"`
	UserUchprocCode int64
}

type TopicUpdateRequest struct {
	ServiceName     string       `json:"serviceName"`
	ExternalRef     string       `json:"externalRef"`
	TopicId         int64        `json:"topicId"`
	Topic           models.Topic `json:"topic"`
	UserUchprocCode int64
}

type GetPointsJournalRequest struct {
	ServiceName     string `json:"serviceName"`
	ExternalRef     string `json:"externalRef"`
	CourseId        int64  `json:"course_id"`
	UserUchprocCode int64
}
