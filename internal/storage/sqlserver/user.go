package sqlserver

import (
	"context"
	"database/sql"
	"github.com/iamgafurov/journal/internal/dto"
	"strconv"
	"strings"
)

func (d *db) GetUserAuthParams(ctx context.Context, login string) (params dto.AuthParams, err error) {

	conn, err := d.pool.Conn(ctx)
	if err != nil {
		return
	}
	defer conn.Close()

	query := `SELECT kdn,tmk, prp FROM prl WHERE RTRIM(cLogin)=$1`
	err = conn.QueryRowContext(ctx, query, login).Scan(&params.UserId, &params.Tmk, &params.Password)
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
	params.Password = strings.Join(res, "|")

	return params, nil
}

func (d *db) UserGetLoginByUchprocId(ctx context.Context, uchprocId int64) (login string, err error) {
	conn, err := d.pool.Conn(ctx)
	if err != nil {
		return
	}
	defer conn.Close()

	query := `SELECT RTRIM(clogin) FROM adempiere.isu_prl WHERE isu_prl_id = $1`
	err = conn.QueryRowContext(ctx, query, uchprocId).Scan(&login)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", dto.ErrNoRows
		}
		return
	}
	return login, nil
}
