package sqlserver

import (
	"context"
	"github.com/iamgafurov/journal/internal/dto"
)

func (d *db) GetTurnstileStudents(ctx context.Context) ([]dto.CheckAttendanceItem, error) {
	//todo get all students
	return []dto.CheckAttendanceItem{}, nil
}
