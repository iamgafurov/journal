package sqlserver

import (
	"context"
	"fmt"
	"github.com/iamgafurov/journal/internal/dto"
	"strings"
)

func (d *db) GetPointsJournal(ctx context.Context, courseId int64) (dto.PointJournal, error) {
	var (
		journal       = dto.PointJournal{}
		studentPoints = make([]dto.StudentPoint, 0)
	)
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
		var firsRatingSum float32
		var secondRatingSum float32

		for i := 0; i < 9; i++ {
			w := dto.WeekPoint{
				WeekNumber: i,
				Point:      pt[i],
			}
			firsRatingSum += pt[i]
			st.FirstRating = append(st.FirstRating, w)
		}

		for i := 9; i < 18; i++ {
			w := dto.WeekPoint{
				WeekNumber: i,
				Point:      pt[i],
			}
			secondRatingSum += pt[i]
			st.SecondRating = append(st.SecondRating, w)
		}

		st.FirstRatingSum = firsRatingSum
		st.SecondRatingSum = secondRatingSum
		studentPoints = append(studentPoints, st)
	}

	journal.Students = studentPoints
	return journal, nil
}

func (d *db) GetPointUserCode(ctx context.Context, courseId int64) (code int64, err error) {
	err = d.pool.QueryRowContext(ctx, `SELECT kst FROM tblvdkr WHERE kdn = $1`, courseId).Scan(&code)
	if err != nil {
		if isSqlNoRows(err) {
			return 0, dto.ErrNoRows
		}
	}
	return
}

func (d *db) UpdatePointsJournal(ctx context.Context, points []dto.PointUpdate, kvdId int64, currentWeek int) (err error) {
	txn, err := d.pool.Begin()
	if err != nil {
		return
	}
	currentWeek -= 20
	weekColumn := fmt.Sprintf("oceblkr%d", currentWeek)

	for _, p := range points {
		query := `UPDATE tblvdstkr SET ` + weekColumn + ` = $1 WHERE kdn = $2 AND kvd =$3`
		cmd, err := txn.ExecContext(ctx, query, p.Point, p.Id, kvdId)
		if err != nil {
			txn.Rollback()
			return err
		}

		ra, err := cmd.RowsAffected()
		if err != nil {
			txn.Rollback()
			return err
		}

		if ra != 1 {
			txn.Rollback()
			return dto.ErrNoRowsAffected
		}
	}
	return txn.Commit()
}
