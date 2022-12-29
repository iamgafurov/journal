package sqlserver

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/models"
	"strings"
)

func (d *db) GetAttendanceJournal(ctx context.Context, courseId int64, limit int) ([]dto.StudentAttendance, error) {
	result := make([]dto.StudentAttendance, 0)
	rows, err := d.pool.QueryContext(ctx, attendanceGetQuery, courseId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		m := make([]string, 96)
		at := dto.StudentAttendance{}
		err = rows.Scan(&at.Id, &at.Name, &at.RecordBook, &at.Id,
			&m[0],
			&m[1],
			&m[2],
			&m[3],
			&m[4],
			&m[5],
			&m[6],
			&m[7],
			&m[8],
			&m[9],
			&m[10],
			&m[11],
			&m[12],
			&m[13],
			&m[14],
			&m[15],
			&m[16],
			&m[17],
			&m[18],
			&m[19],
			&m[20],
			&m[21],
			&m[22],
			&m[23],
			&m[24],
			&m[25],
			&m[26],
			&m[27],
			&m[28],
			&m[29],
			&m[30],
			&m[31],
			&m[32],
			&m[33],
			&m[34],
			&m[35],
			&m[36],
			&m[37],
			&m[38],
			&m[39],
			&m[40],
			&m[41],
			&m[42],
			&m[43],
			&m[44],
			&m[45],
			&m[46],
			&m[47],
			&m[48],
			&m[49],
			&m[50],
			&m[51],
			&m[52],
			&m[53],
			&m[54],
			&m[55],
			&m[56],
			&m[57],
			&m[58],
			&m[59],
			&m[60],
			&m[61],
			&m[62],
			&m[63],
			&m[64],
			&m[65],
			&m[66],
			&m[67],
			&m[68],
			&m[69],
			&m[70],
			&m[71],
			&m[72],
			&m[73],
			&m[74],
			&m[75],
			&m[76],
			&m[77],
			&m[78],
			&m[79],
			&m[80],
			&m[81],
			&m[82],
			&m[83],
			&m[84],
			&m[85],
			&m[86],
			&m[87],
			&m[88],
			&m[89],
			&m[90],
			&m[91],
			&m[92],
			&m[93],
			&m[94],
			&m[95],
		)
		if err != nil {
			return nil, err
		}
		for i, v := range m {
			at.Attendance = append(at.Attendance, dto.Attendance{Value: v, Number: i + 1})
		}
		at.Name = strings.TrimSpace(at.Name)
		at.RecordBook = strings.TrimSpace(at.RecordBook)
		at.Attendance = at.Attendance[:limit]
		result = append(result, at)
	}

	return result, nil
}

func (d *db) GetAttendanceStatement(ctx context.Context, id int64) (statement models.Statement, err error) {
	err = d.pool.QueryRowContext(ctx,
		`SELECT 
    				kdn, 
    				kst, 
    				kpr, 
    				kas, 
    				vchaccess, 
    				chruchgod 
				FROM tblvdtkr 
				WHERE kdn = $1`,
		id).Scan(
		&statement.Id,
		&statement.Kst,
		&statement.Kpr,
		&statement.Kas,
		&statement.VchAccess,
		&statement.Chruchgod,
	)
	if err != nil {
		if isSqlNoRows(err) {
			return models.Statement{}, dto.ErrNoRows
		}
	}

	return
}

func (d *db) UpdateAttendanceJournal(ctx context.Context, courseId int64, at []dto.StudentAttendance) (err error, atErr []dto.AttendanceJournalError) {
	for _, student := range at {
		cposesQuery := ``
		for _, attendance := range student.Attendance {
			attendance.Value = strings.TrimSpace(attendance.Value)
			if attendance.Value == "" {
				attendance.Value = " "
			}

			cposesQuery += fmt.Sprintf("cpos%v = '%s', ", attendance.Number, attendance.Value)
		}

		if cposesQuery == "" {
			continue
		}

		tx, err := d.pool.Begin()
		if err != nil {
			atErr = append(atErr, dto.AttendanceJournalError{StudentId: student.Id, Message: err.Error()})
			continue
		}

		//trim last comma
		cposesQuery = cposesQuery[:len(cposesQuery)-2]
		cmd, err := tx.ExecContext(ctx, `UPDATE tblvdtstkr SET `+cposesQuery+` WHERE kdn = $1 AND kvd = $2`, student.Id, courseId)
		if err != nil {
			tx.Rollback()
			atErr = append(atErr, dto.AttendanceJournalError{StudentId: student.Id, Message: student.Name + ", err:" + err.Error()})
			continue
		}

		ra, err := cmd.RowsAffected()
		if err != nil {
			tx.Rollback()
			atErr = append(atErr, dto.AttendanceJournalError{StudentId: student.Id, Message: err.Error()})
			continue
		}

		if ra != 1 {
			tx.Rollback()
			atErr = append(atErr, dto.AttendanceJournalError{StudentId: student.Id, Message: "student: " + student.Name + " not found"})
			continue
		}

		cposes, err := getCposesById(ctx, student.Id, tx)
		if err != nil {
			tx.Rollback()
			atErr = append(atErr, dto.AttendanceJournalError{StudentId: student.Id, Message: err.Error()})
			continue
		}

		totalMisses := 0
		for _, v := range cposes {
			if v == "н" {
				totalMisses++
			}
		}

		err = updateTotalMisses(ctx, totalMisses, student.Id, tx)
		if err != nil {
			tx.Rollback()
			atErr = append(atErr, dto.AttendanceJournalError{StudentId: student.Id, Message: err.Error()})
			continue
		}

		tx.Commit()
	}

	return
}

func updateTotalMisses(ctx context.Context, totalMisses int, kvd int64, tx *sql.Tx) error {
	cmd, err := tx.ExecContext(ctx, `UPDATE tblvdtstkr SET itjurkrs =$1 WHERE kdn = $2`, totalMisses, kvd)
	if err != nil {
		return err
	}

	ra, err := cmd.RowsAffected()
	if err != nil {
		return err
	}

	if ra != 1 {
		return dto.ErrNoRowsAffected
	}

	return nil
}
func getCposesById(ctx context.Context, kvd int64, tx *sql.Tx) ([]string, error) {
	rows, err := tx.QueryContext(ctx, cposesByKvdQuery, kvd)
	if err != nil {
		return nil, err
	}

	m := make([]string, 96)
	for rows.Next() {
		err = rows.Scan(
			&m[0],
			&m[1],
			&m[2],
			&m[3],
			&m[4],
			&m[5],
			&m[6],
			&m[7],
			&m[8],
			&m[9],
			&m[10],
			&m[11],
			&m[12],
			&m[13],
			&m[14],
			&m[15],
			&m[16],
			&m[17],
			&m[18],
			&m[19],
			&m[20],
			&m[21],
			&m[22],
			&m[23],
			&m[24],
			&m[25],
			&m[26],
			&m[27],
			&m[28],
			&m[29],
			&m[30],
			&m[31],
			&m[32],
			&m[33],
			&m[34],
			&m[35],
			&m[36],
			&m[37],
			&m[38],
			&m[39],
			&m[40],
			&m[41],
			&m[42],
			&m[43],
			&m[44],
			&m[45],
			&m[46],
			&m[47],
			&m[48],
			&m[49],
			&m[50],
			&m[51],
			&m[52],
			&m[53],
			&m[54],
			&m[55],
			&m[56],
			&m[57],
			&m[58],
			&m[59],
			&m[60],
			&m[61],
			&m[62],
			&m[63],
			&m[64],
			&m[65],
			&m[66],
			&m[67],
			&m[68],
			&m[69],
			&m[70],
			&m[71],
			&m[72],
			&m[73],
			&m[74],
			&m[75],
			&m[76],
			&m[77],
			&m[78],
			&m[79],
			&m[80],
			&m[81],
			&m[82],
			&m[83],
			&m[84],
			&m[85],
			&m[86],
			&m[87],
			&m[88],
			&m[89],
			&m[90],
			&m[91],
			&m[92],
			&m[93],
			&m[94],
			&m[95],
		)
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}

var attendanceGetQuery = `SELECT 
    			std.kdn, 
       			nst,
       			kzc,
       			tblvdtstkr.kdn,
       			tblvdtstkr.cpos1,
       			tblvdtstkr.cpos2,
       			tblvdtstkr.cpos3,
       			tblvdtstkr.cpos4,
       			tblvdtstkr.cpos5,
       			tblvdtstkr.cpos6,
       			tblvdtstkr.cpos7,
       			tblvdtstkr.cpos8,
       			tblvdtstkr.cpos9,
       			tblvdtstkr.cpos10,
       			tblvdtstkr.cpos11,
       			tblvdtstkr.cpos12,
       			tblvdtstkr.cpos13,
       			tblvdtstkr.cpos14,
       			tblvdtstkr.cpos15,
       			tblvdtstkr.cpos16,
       			tblvdtstkr.cpos17,
       			tblvdtstkr.cpos18,
       			tblvdtstkr.cpos19,
       			tblvdtstkr.cpos20,
       			tblvdtstkr.cpos21,
       			tblvdtstkr.cpos22,
       			tblvdtstkr.cpos23,
       			tblvdtstkr.cpos24,
       			tblvdtstkr.cpos25,
       			tblvdtstkr.cpos26,
       			tblvdtstkr.cpos27,
       			tblvdtstkr.cpos28,
       			tblvdtstkr.cpos29,
       			tblvdtstkr.cpos30,
       			tblvdtstkr.cpos31,
       			tblvdtstkr.cpos32,
       			tblvdtstkr.cpos33,
       			tblvdtstkr.cpos34,
       			tblvdtstkr.cpos35,
       			tblvdtstkr.cpos36,
       			tblvdtstkr.cpos37,
       			tblvdtstkr.cpos38,
       			tblvdtstkr.cpos39,
       			tblvdtstkr.cpos40,
       			tblvdtstkr.cpos41,
       			tblvdtstkr.cpos42,
       			tblvdtstkr.cpos43,
       			tblvdtstkr.cpos44,
       			tblvdtstkr.cpos45,
       			tblvdtstkr.cpos46,
       			tblvdtstkr.cpos47,
       			tblvdtstkr.cpos48,
       			tblvdtstkr.cpos49,
       			tblvdtstkr.cpos50,
       			tblvdtstkr.cpos51,
       			tblvdtstkr.cpos52,
       			tblvdtstkr.cpos53,
       			tblvdtstkr.cpos54,
       			tblvdtstkr.cpos55,
       			tblvdtstkr.cpos56,
       			tblvdtstkr.cpos57,
       			tblvdtstkr.cpos58,
       			tblvdtstkr.cpos59,
       			tblvdtstkr.cpos60,
       			tblvdtstkr.cpos61,
       			tblvdtstkr.cpos62,
       			tblvdtstkr.cpos63,
       			tblvdtstkr.cpos64,
       			tblvdtstkr.cpos65,
       			tblvdtstkr.cpos66,
       			tblvdtstkr.cpos67,
       			tblvdtstkr.cpos68,
       			tblvdtstkr.cpos69,
       			tblvdtstkr.cpos70,
       			tblvdtstkr.cpos71,
       			tblvdtstkr.cpos72,
       			tblvdtstkr.cpos73,
       			tblvdtstkr.cpos74,
       			tblvdtstkr.cpos75,
       			tblvdtstkr.cpos76,
       			tblvdtstkr.cpos77,
       			tblvdtstkr.cpos78,
       			tblvdtstkr.cpos79,
       			tblvdtstkr.cpos80,
       			tblvdtstkr.cpos81,
       			tblvdtstkr.cpos82,
       			tblvdtstkr.cpos83,
       			tblvdtstkr.cpos84,
       			tblvdtstkr.cpos85,
       			tblvdtstkr.cpos86,
       			tblvdtstkr.cpos87,
       			tblvdtstkr.cpos88,
       			tblvdtstkr.cpos89,
       			tblvdtstkr.cpos90,
       			tblvdtstkr.cpos91,
       			tblvdtstkr.cpos92,
       			tblvdtstkr.cpos93,
       			tblvdtstkr.cpos94,
       			tblvdtstkr.cpos95,
       			tblvdtstkr.cpos96
			FROM std
			INNER JOIN tblvdtstkr on std.kdn = tblvdtstkr.kst
			WHERE tblvdtstkr.kvd = $1
			ORDER BY pnn`

var cposesByKvdQuery = `SELECT cpos1,
					cpos2,
					cpos3,
					cpos4,
					cpos5,
					cpos6,
					cpos7,
					cpos8,
					cpos9,
					cpos10,
					cpos11,
					cpos12,
					cpos13,
					cpos14,
					cpos15,
					cpos16,
					cpos17,
					cpos18,
					cpos19,
					cpos20,
					cpos21,
					cpos22,
					cpos23,
					cpos24,
					cpos25,
					cpos26,
					cpos27,
					cpos28,
					cpos29,
					cpos30,
					cpos31,
					cpos32,
					cpos33,
					cpos34,
					cpos35,
					cpos36,
					cpos37,
					cpos38,
					cpos39,
					cpos40,
					cpos41,
					cpos42,
					cpos43,
					cpos44,
					cpos45,
					cpos46,
					cpos47,
					cpos48,
					cpos49,
					cpos50,
					cpos51,
					cpos52,
					cpos53,
					cpos54,
					cpos55,
					cpos56,
					cpos57,
					cpos58,
					cpos59,
					cpos60,
					cpos61,
					cpos62,
					cpos63,
					cpos64,
					cpos65,
					cpos66,
					cpos67,
					cpos68,
					cpos69,
					cpos70,
					cpos71,
					cpos72,
					cpos73,
					cpos74,
					cpos75,
					cpos76,
					cpos77,
					cpos78,
					cpos79,
					cpos80,
					cpos81,
					cpos82,
					cpos83,
					cpos84,
					cpos85,
					cpos86,
					cpos87,
					cpos88,
					cpos89,
					cpos90,
					cpos91,
					cpos92,
					cpos93,
					cpos94,
					cpos95,
					cpos96
				FROM tblvdtstkr
				WHERE kdn = $1`
