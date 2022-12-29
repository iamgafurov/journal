package sqlserver

import (
	"context"
	"database/sql"
	"github.com/iamgafurov/journal/internal/dto"
	"strings"
	"time"
)

func (d *db) GetGroupCoursesAttendance(ctx context.Context, groupId, userId int64, academicYear string) (cs []dto.Course, err error) {
	cs = []dto.Course{}
	var (
		assisTantId int64
		teacherId   int64
	)
	queryStr := `SELECT 
    				t.kas,
					t.kst,
					t.kdn,
					t.kgr,
					prk.npr, 
					sot.nst, 
					t.kolkr
				FROM tblvdtkr t
				INNER JOIN prk ON t.kpr = prk.kdn
				INNER JOIN sot ON (t.kst = sot.kdn OR t.kas = sot.kdn)
				WHERE t.kgr = $1 AND (t.kst = $2 OR t.kas = $2) AND chruchgod = $3`
	rows, err := d.pool.QueryContext(ctx, queryStr, groupId, userId, academicYear)
	if err != nil {
		return
	}

	for rows.Next() {
		c := dto.Course{}
		err = rows.Scan(&assisTantId, &teacherId, &c.AttendanceId, &c.GroupId, &c.CourseName, &c.TeacherName, &c.CreditsCount)
		if err != nil {
			return
		}
		if assisTantId != 0 {
			c.IsAssistant = true
		}
		c.CourseName = strings.TrimSpace(c.CourseName)
		c.CreditsCount = strings.TrimSpace(c.CreditsCount)
		c.TeacherName = strings.TrimSpace(c.TeacherName)
		cs = append(cs, c)
	}

	return
}

func (d *db) GetGroupCoursesPoint(ctx context.Context, groupId, userId int64, academicYear string) (cs []dto.Course, err error) {
	cs = []dto.Course{}
	var (
		assisTantId int64
		teacherId   int64
	)
	queryStr := `SELECT 
    				t.kas,
					t.kst,
					t.kdn,
					t.kgr,
					prk.npr, 
					sot.nst, 
					t.kolkr
				FROM tblvdkr t
				INNER JOIN prk ON t.kpr = prk.kdn
				INNER JOIN sot ON (t.kst = sot.kdn OR t.kas = sot.kdn)
				WHERE t.kgr = $1 AND (t.kst = $2 OR t.kas = $2) AND chruchgod = $3`
	rows, err := d.pool.QueryContext(ctx, queryStr, groupId, userId, academicYear)
	if err != nil {
		return
	}

	for rows.Next() {
		c := dto.Course{}
		err = rows.Scan(&assisTantId, &teacherId, &c.PointId, &c.GroupId, &c.CourseName, &c.TeacherName, &c.CreditsCount)
		if err != nil {
			return
		}
		if assisTantId != 0 {
			c.IsAssistant = true
		}
		c.CourseName = strings.TrimSpace(c.CourseName)
		c.CreditsCount = strings.TrimSpace(c.CreditsCount)
		c.TeacherName = strings.TrimSpace(c.TeacherName)
		cs = append(cs, c)
	}

	return
}

func (d *db) GetLastAYStartTime(ctx context.Context) (tm time.Time, err error) {
	err = d.pool.QueryRowContext(ctx, `SELECT TOP 1  dto FROM ugd ORDER BY dto desc`).Scan(&tm)
	if err != nil {
		if isSqlNoRows(err) {
			return tm, dto.ErrNoRows
		}
	}
	return
}

func isSqlNoRows(err error) bool {
	return err == sql.ErrNoRows
}
