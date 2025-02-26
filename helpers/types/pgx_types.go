package types

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

func Timestamp(time time.Time) pgtype.Timestamp {

	return pgtype.Timestamp{Time: time, Valid: true}
}

func Text(str string) pgtype.Text {
	return pgtype.Text{String: str, Valid: true}

}
