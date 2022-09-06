package sqlserver

import (
	"context"
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/models"
	"github.com/iamgafurov/journal/internal/tools"
	"strings"
)

func (d *db) GetTopics(ctx context.Context, userCode, courseId int64) ([]models.Topic, error) {
	topics := make([]models.Topic, 0)

	query := `SELECT 
				kdn, 
				cnzap, 
				dtzap,
				ctema,
				nkolch,
				nkolchsem,
				nkolchprak,
				nkolchlab,
				nkolchkmdro,
				nkolchv
			FROM tblvdpstkr
			WHERE kst = $1 AND kvd = $2
			ORDER BY cnzap desc;`

	row, err := d.pool.QueryContext(ctx, query, userCode, courseId)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		t := models.Topic{}
		err = row.Scan(&t.Id, &t.Cnzap, &t.Dtzap, &t.Tema, &t.KolLek, &t.KolSem, &t.KolPrak, &t.KolLab, &t.KolKmd, &t.KolObsh)
		if err != nil {
			return nil, err
		}
		t.Editable = t.Dtzap.After(tools.TodayStart())
		t.Tema = strings.TrimSpace(t.Tema)
		t.Cnzap = strings.TrimSpace(t.Cnzap)
		topics = append(topics, t)
	}
	return topics, nil
}

func (d *db) DeleteTopic(ctx context.Context, topicId, userCode int64) error {
	cmd, err := d.pool.ExecContext(ctx, `DELETE FROM tblvdpstkr WHERE kdn = $1 AND kst = $2 AND dtzap > $3;`, topicId, userCode, tools.TodayStart())
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

func (d *db) UpdateTopic(ctx context.Context, topic models.Topic, userCode int64) error {
	query := `UPDATE tblvdpstkr
			SET ctema = $1,
				nkolch = $2,
				nkolchsem = $3,
				nkolchlab = $4,
				nkolchprak = $5,
				nkolchkmdro = $6,
				nkolchv = $7
			WHERE 
			    kdn = $8 
				AND kst = $9 
			  	AND dtzap > $10;`

	cmd, err := d.pool.ExecContext(ctx, query,
		topic.Tema,
		topic.KolLek,
		topic.KolSem,
		topic.KolLab,
		topic.KolPrak,
		topic.KolKmd,
		topic.KolObsh,
		topic.Id,
		userCode,
		tools.TodayStart(),
	)

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
