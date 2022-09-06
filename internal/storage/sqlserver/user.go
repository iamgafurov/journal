package sqlserver

import (
	"context"
	"database/sql"
	"github.com/iamgafurov/journal/internal/dto"
	"strconv"
	"strings"
)

func (d *db) GetUserAuthParams(ctx context.Context, login string) (params dto.AuthParams, err error) {
	query := `SELECT kdn,tmk, prp, kst, nmp FROM prl WHERE RTRIM(cLogin)=$1`
	err = d.pool.QueryRowContext(ctx, query, login).Scan(&params.UserId, &params.Tmk, &params.Password, &params.UserCode, &params.UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			return dto.AuthParams{}, dto.ErrNoRows
		}
		return
	}

	str := []rune(params.Password)
	res := []string{}
	for _, r := range str {
		ss := strconv.Itoa(int(r))
		res = append(res, ss)
	}
	params.Password = "|" + strings.Join(res, "|")

	return params, nil
}

func (d *db) UserGetLoginByUchprocId(ctx context.Context, uchprocId int64) (login string, err error) {
	query := `SELECT RTRIM(clogin) FROM prl WHERE kdn = $1`
	err = d.pool.QueryRowContext(ctx, query, uchprocId).Scan(&login)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", dto.ErrNoRows
		}
		return
	}
	return login, nil
}

func (d *db) GetAcademicYears(ctx context.Context, userUchprocCode int64) ([]string, error) {
	years := make([]string, 0)
	conn, err := d.pool.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := `SELECT chruchgod FROM tblvdkr WHERE kst = $1 GROUP BY chruchgod`
	rows, err := conn.QueryContext(ctx, query, userUchprocCode)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		y := ""
		err = rows.Scan(&y)
		if err != nil {
			return nil, err
		}
		years = append(years, y)
	}

	return years, nil
}
