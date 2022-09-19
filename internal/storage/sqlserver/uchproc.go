package sqlserver

import (
	"context"
	"database/sql"
	"github.com/iamgafurov/journal/internal/dto"
	"strings"
	"time"
)

func (d *db) GetGroupCourses(ctx context.Context, groupId, userId int64) (cs []dto.Course, err error) {
	cs = []dto.Course{}
	var (
		assisTantId int64
		teacherId   int64
	)
	queryStr := `SELECT 
    				tblvdkr.kas,
					tblvdkr.kst,
					tblvdkr.kdn,
					tblvdkr.kgr,
					tblvdtkr.kdn,
					prk.npr, 
					sot.nst, 
					tblvdkr.kolkr
				FROM tblvdkr
				INNER JOIN prk ON tblvdkr.kpr = prk.kdn
				INNER JOIN sot ON (tblvdkr.kst = sot.kdn OR tblvdkr.kas = sot.kdn)
				INNER JOIN tblvdtkr ON (tblvdkr.kpr = tblvdtkr.kpr)
				WHERE tblvdkr.kgr = $1 AND (tblvdkr.kst = $2 OR tblvdkr.kas = $2)`
	rows, err := d.pool.QueryContext(ctx, queryStr, groupId, userId)
	if err != nil {
		return
	}

	for rows.Next() {
		c := dto.Course{}
		err = rows.Scan(&assisTantId, &teacherId, &c.AttendanceId, &c.GroupId, &c.PointId, &c.CourseName, &c.TeacherName, &c.CreditsCount)
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
