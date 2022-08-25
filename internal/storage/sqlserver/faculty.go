package sqlserver

import (
	"context"
	"database/sql"
	"github.com/iamgafurov/journal/internal/dto"
)

func (d *db) GetFaculties(ctx context.Context, userUchprocId int64) (res []dto.Faculty, err error) {
	conn, err := d.pool.Conn(ctx)
	if err != nil {
		return
	}
	defer conn.Close()

	query := ``
	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dto.ErrNoRows
		}
		return
	}

	for rows.Next() {
		var (
			fId   int64
			fCode int64
			fName string
			sId   int64
			sCode int64
			sName string
			cId   int64
			cCode int64
			cName string
			gId   int64
			gCode int64
			gName string
		)
		err = rows.Scan(&fId, &fCode, &fName, &sId, &sCode, &sName, &cId, &cCode, &cName, &gId, &gCode, &gName)
		if err != nil {
			return
		}
	}

	return res, nil
}

func set(faculties []dto.Faculty, faculty dto.Faculty) (fk []dto.Faculty) {
	fk = faculties
	fInd, sInd, cInd := -1, -1, -1
	for i, f := range fk {
		if f.Id == faculty.Id {
			fInd = i
			break
		}
	}
	if fInd < 0 {
		fk = append(fk, faculty)
	}

	for i, s := range fk[fInd].Specialties {
		if s.Id == faculty.Specialties[0].Id {
			sInd = i
			break
		}
	}
	if sInd < 0 {
		fk[fInd].Specialties = append(fk[fInd].Specialties, faculty.Specialties[0])
		return
	}

	for i, c := range fk[fInd].Specialties[sInd].Years {
		if c.Id == faculty.Specialties[sInd].Years[0].Id {
			cInd = i
			break
		}
	}

	if cInd < 0 {
		fk[fInd].Specialties[sInd].Years = append(fk[fInd].Specialties[sInd].Years, faculty.Specialties[0].Years[0])
		return
	}

	fk[fInd].Specialties[sInd].Years[cInd].Groups = append(fk[fInd].Specialties[sInd].Years[cInd].Groups, faculty.Specialties[0].Years[0].Groups[0])

	return
}
