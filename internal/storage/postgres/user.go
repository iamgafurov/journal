package postgres

import (
	"context"
	"database/sql"
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/models"
)

func (d *db) UserInsert(ctx context.Context, user models.User) (id int64, err error) {
	sql := `INSERT INTO users(
				token,
				login,
				service,
				uchproc_id,
				expire_at,
				status
			)
			VALUES($1, $2, $3, $4, $5, $6)
			ON CONFLICT(login, service)
				DO UPDATE SET token = $1, updated_at = NOW(), expire_at = $5 
			RETURNING id;`
	err = d.pool.QueryRow(ctx, sql,
		user.Token,
		user.Login,
		user.Service,
		user.UchprocId,
		user.ExpireAt,
		user.Status,
	).Scan(&id)

	return
}

func (d *db) UserGetByToken(ctx context.Context, token string) (user models.User, err error) {
	sql := `SELECT
				id,
				token,
				login,
				service,
				uchproc_id,
				reg_date,
				updated_at,
				expire_at,
				status
			FROM users
			WHERE token = $1`
	err = d.pool.QueryRow(ctx, sql, token).
		Scan(&user.Id,
			&user.Token,
			&user.Login,
			&user.Service,
			&user.UchprocId,
			&user.RegDate,
			&user.UpdatedAt,
			&user.ExpireAt,
			&user.Status,
		)

	return
}

func (d *db) GetUserAuthParams(ctx context.Context, login string) (params dto.AuthParams, err error) {
	//params = dto.AuthParams{
	//	UserId:   122,
	//	Tmk:      "123",
	//	Login:    "user1",
	//	Password: "myPass",
	//}
	//return

	query := `SELECT isu_prl_id,tmk, adempiere.getAsciiCode(prp)  FROM adempiere.isu_prl WHERE RTRIM(clogin)=$1`
	err = d.pool.QueryRow(ctx, query, login).Scan(&params.UserId, &params.Tmk, &params.Password, &params.AsciiCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return dto.AuthParams{}, dto.ErrNoRows
		}
		return
	}
	return params, nil
}