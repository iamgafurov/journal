package sqlserver

import (
	"context"
	"database/sql"
	"github.com/iamgafurov/journal/internal/dto"
)

func (d *db) GetUserAuthParams(ctx context.Context, login string) (params dto.AuthParams, err error) {
	params = dto.AuthParams{
		UserId:   122,
		Tmk:      "123",
		Login:    "user1",
		Password: "myPass",
	}
	return
	conn, err := d.pool.Conn(ctx)
	if err != nil {
		return
	}
	defer conn.Close()

	query := `SELECT isu_prl_id,tmk, adempiere.getAsciiCode(prp)  FROM adempiere.isu_prl WHERE RTRIM(clogin)=$1`
	err = conn.QueryRowContext(ctx, query, login).Scan(&params.UserId, &params.Tmk, &params.Password, &params.AsciiCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return dto.AuthParams{}, dto.ErrNoRows
		}
		return
	}
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
