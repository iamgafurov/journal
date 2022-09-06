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

func (d *db) GetPointsJournal(ctx context.Context, courseId int64) (dto.PointJournal, error) {
	journal := dto.PointJournal{}
	points := make([]dto.StudentPoint, 0)

	query := `SELECT 
    			std.kdn, 
       			nst,
       			kzc,
       			tblvdstkr.oceblkr1, 
       			tblvdstkr.oceblkr2, 
       			tblvdstkr.oceblkr3, 
       			tblvdstkr.oceblkr4, 
       			tblvdstkr.oceblkr5, 
       			tblvdstkr.oceblkr6, 
       			tblvdstkr.oceblkr7, 
       			tblvdstkr.oceblkr8, 
       			tblvdstkr.oceblkr9, 
       			tblvdstkr.oceblkr10, 
       			tblvdstkr.oceblkr11, 
       			tblvdstkr.oceblkr12, 
       			tblvdstkr.oceblkr13, 
       			tblvdstkr.oceblkr14, 
       			tblvdstkr.oceblkr15, 
       			tblvdstkr.oceblkr16, 
       			tblvdstkr.oceblkr17, 
       			tblvdstkr.oceblkr18,
       			tblvdstkr.itoceblkr,
       			tblvdstkr.itocekr
			FROM std
			INNER JOIN tblvdstkr on std.kdn = tblvdstkr.kst
			WHERE tblvdstkr.kvd = $1
			ORDER BY pnn`

	rows, err := d.pool.QueryContext(ctx, query, courseId)
	if err != nil {
		return journal, err
	}

	for rows.Next() {
		pt := make([]float32, 18)
		st := dto.StudentPoint{}
		err = rows.Scan(
			&st.Id,
			&st.Name,
			&st.RecordBook,
			&pt[0], &pt[1], &pt[2], &pt[3], &pt[4], &pt[5], &pt[6], &pt[7], &pt[8], &pt[9], &pt[10], &pt[11], &pt[12], &pt[13], &pt[14], &pt[15], &pt[16], &pt[17],
			&st.PointsSum,
			&st.Grade,
		)
		if err != nil {
			return journal, err
		}

		st.Name = strings.TrimSpace(st.Name)
		for i := 0; i < 18; i++ {
			w := dto.WeekPoint{
				WeekNumber: i,
				Point:      pt[i],
			}
			st.WeekPoints = append(st.WeekPoints, w)
		}

		points = append(points, st)
	}

	journal.Points = points
	return journal, nil
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

func (d *db) GetPointUserCode(ctx context.Context, courseId int64) (code int64, err error) {
	err = d.pool.QueryRowContext(ctx, `SELECT kst FROM tblvdtkr WHERE kdn = $1`, courseId).Scan(&code)
	if err != nil {
		if isSqlNoRows(err) {
			return 0, dto.ErrNoRows
		}
	}
	return
}

func isSqlNoRows(err error) bool {
	return err == sql.ErrNoRows
}
