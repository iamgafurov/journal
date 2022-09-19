package dto

import "time"

type Course struct {
	PointId      int64  `json:"point_id"`
	AttendanceId int64  `json:"attendance_id"`
	CourseName   string `json:"course_name"`
	GroupId      int64  `json:"group_id"`
	TeacherName  string `json:"teacher_name"`
	CreditsCount string `json:"credits_count"`
	StartDate    string `json:"start_date"`
	IsAssistant  bool   `json:"is_assistant"`
}

type AuthParams struct {
	UserId   int64
	UserCode int64
	UserName string
	Tmk      string
	Password string
	Login    string
}

type PointJournal struct {
	WeeksNumber int            `json:"weeks_number"`
	CurrentWeek int            `json:"current_week"`
	MaxPoint    int            `json:"max_point"`
	Students    []StudentPoint `json:"students"`
}

type PointUpdate struct {
	Id    int64   `json:"id"`
	Point float32 `json:"point"`
}

type Week struct {
	Number   int  `json:"number"`
	Editable bool `json:"editable"`
}

type StudentPoint struct {
	Id              int64       `json:"id"`
	Name            string      `json:"name"`
	RecordBook      string      `json:"record_book"`
	PointsSum       float32     `json:"points_sum"`
	Grade           string      `json:"grade"`
	FirstRating     []WeekPoint `json:"first_rating"`
	SecondRating    []WeekPoint `json:"second_rating"`
	FirstRatingSum  float32     `json:"first_rating_sum"`
	SecondRatingSum float32     `json:"second_rating_sum"`
}

type WeekPoint struct {
	WeekNumber int     `json:"week_number"`
	Point      float32 `json:"point"`
	Editable   bool    `json:"-"`
}

type Attendance struct {
	Value  string `json:"value"`
	Number int    `json:"topic_number"`
}

type StudentAttendance struct {
	Id         int64        `json:"id"`
	Name       string       `json:"name"`
	RecordBook string       `json:"record_book"`
	Attendance []Attendance `json:"attendance"`
}

type Topic struct {
	Id       int64     `json:"id"`
	Cnzap    string    `json:"cnzap"`
	Dtzap    time.Time `json:"dtzap"`
	Tema     string    `json:"tema"`
	KolLek   int       `json:"kolLek"`
	KolSem   int       `json:"kolSem"`
	KolPrak  int       `json:"kolPrak"`
	KolLab   int       `json:"kolLab"`
	KolKmd   int       `json:"kolKmd"`
	KolObsh  int       `json:"kolObsh"`
	Editable bool      `json:"editable"`
}

type AttendanceJournalError struct {
	StudentId int64  `json:"student_id"`
	Message   string `json:"message"`
}
