package dto

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
	Points      []StudentPoint `json:"points"`
}

type Week struct {
	Number   int  `json:"number"`
	Editable bool `json:"editable"`
}

type StudentPoint struct {
	Id         int64       `json:"id"`
	Name       string      `json:"name"`
	RecordBook string      `json:"record_book"`
	PointsSum  float32     `json:"points_sum"`
	Grade      string      `json:"grade"`
	WeekPoints []WeekPoint `json:"week_points"`
}

type WeekPoint struct {
	WeekNumber int     `json:"week_number"`
	Point      float32 `json:"point"`
	Editable   bool    `json:"editable"`
}
