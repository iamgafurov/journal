package dto

import "time"

type CheckAttendanceRequest struct {
	Students []CheckAttendanceItem `json:"students"`
}

type CheckAttendanceItem struct {
	RecordBook   string    `json:"recordBook"`
	AttendanceId int64     `json:"attendanceId"`
	TopicNumber  int       `json:"topicNumber"`
	LessonTime   time.Time `json:"lessonTime"`
	CourseId     int64     `json:"courseId"`
}

type CheckAttendanceResponse struct {
	AbsentStudents []CheckAttendanceItem `json:"absentStudents"`
}
