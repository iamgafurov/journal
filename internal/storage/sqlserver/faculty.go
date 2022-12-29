package sqlserver

import (
	"context"
	"github.com/iamgafurov/journal/internal/dto"
)

func (d *db) GetFaculties(ctx context.Context, userUchprocId int64, studyYear string) (res []dto.Faculty, err error) {
	query := `SELECT 
    			fak.kdn,
				fak.kfk,
				RTRIM(fak.nfk),
				spe.kdn,
				spe.ksp,
				RTRIM(spe.nsp),
				krs.kdn,
				krs.kkr,
				grp.kdn,
				grp.kgr
			FROM fak
			INNER JOIN spe ON fak.kdn = spe.kfk
			INNER JOIN krs ON spe.kdn = krs.ksp
			INNER JOIN grp ON krs.kdn = grp.kkr
			INNER JOIN tblvdkr ON grp.kdn = tblvdkr.kgr
			WHERE tblvdkr.kst = $1 AND tblvdkr.chruchgod = $2`
	rows, err := d.pool.QueryContext(ctx, query, userUchprocId, studyYear)
	if err != nil {
		return
	}

	for rows.Next() {
		var (
			fId   int64
			fCode string
			fName string
			sId   int64
			sCode string
			sName string
			cId   int64
			cCode string
			gId   int64
			gCode string
		)

		err = rows.Scan(&fId, &fCode, &fName, &sId, &sCode, &sName, &cId, &cCode, &gId, &gCode)
		if err != nil {
			return
		}

		res = set(res, dto.Faculty{Id: fId, Code: fCode, Name: fName,
			Specialties: []dto.Speciality{{Id: sId, Code: sCode, Name: sName,
				Years: []dto.Year{{Id: cId, Code: cCode,
					Groups: []dto.Group{{Id: gId, Code: gCode}},
				}},
			}},
		})

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
		return
	}

	for i, s := range fk[0].Specialties {
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
		if c.Id == faculty.Specialties[0].Years[0].Id {
			cInd = i
			break
		}
	}

	if cInd < 0 {
		fk[fInd].Specialties[sInd].Years = append(fk[fInd].Specialties[sInd].Years, faculty.Specialties[0].Years[0])
		return
	}

	for _, g := range fk[fInd].Specialties[sInd].Years[cInd].Groups {
		if g.Id == faculty.Specialties[0].Years[0].Groups[0].Id {
			return
		}
	}
	fk[fInd].Specialties[sInd].Years[cInd].Groups = append(fk[fInd].Specialties[sInd].Years[cInd].Groups, faculty.Specialties[0].Years[0].Groups[0])

	return
}
